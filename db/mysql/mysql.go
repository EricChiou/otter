package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"otter/config"
	"otter/constants/api"
)

// DB mysql connecting
var DB *sql.DB

var specificCharStr string = "\"':.,;(){}[]&|=+-*%/\\<>^"
var specificChar [128]bool

// RowResult db QueryRow result
type RowResult struct {
	Row *sql.Row
}

// RowsResult db Query result
type RowsResult struct {
	Rows *sql.Rows
}

// Init connect MySQL
func Init() (err error) {
	for _, c := range specificCharStr {
		specificChar[int(c)] = true
	}

	cfg := config.Get()
	userName := cfg.MySQLUserName
	password := cfg.MySQLPassword
	addr := cfg.MySQLAddr
	port := cfg.MySQLPort
	dbName := cfg.MySQLDBNAME

	DB, err = sql.Open("mysql", userName+":"+password+"@tcp("+addr+":"+port+")/"+dbName)
	return err
}

// Close close mysql connecting
func Close() {
	if DB != nil {
		DB.Close()
	}
}

// ErrMsgHandler error message handler
func ErrMsgHandler(err error) api.RespStatus {
	if strings.Contains(err.Error(), "Duplicate") {
		return api.Duplicate
	} else {
		return api.DBError
	}
}

// Insert insert data
func Insert(table string, columnValues sqlParams) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	columnSQL := ""
	valueSQL := ""
	var args []interface{}
	for k, v := range columnValues.kv {
		columnSQL += ", " + k
		valueSQL += ", ?"
		args = append(args, v)
	}
	if len(columnSQL) > 2 {
		columnSQL = columnSQL[2:]
	}
	if len(valueSQL) > 2 {
		valueSQL = valueSQL[2:]
	}

	return tx.Exec("INSERT INTO "+table+"( "+columnSQL+" ) values( "+valueSQL+" )", args...)
}

// Update upadte data
func Update(table string, set sqlParams, where sqlParams) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	var args []interface{}
	setSQL, args := getSetSQL(set.kv, args)
	whereSQL, args := getWhereSQL(where.kv, args)

	return tx.Exec("UPDATE "+table+" SET "+setSQL+whereSQL, args...)
}

// Delete delete data
func Delete(table string, where sqlParams) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}
	whereSQL, args := getWhereSQL(where.kv, []interface{}{})

	return tx.Exec("DELETE FROM "+table+whereSQL, args...)
}

// Query query data
func Query(sql string, params sqlParams, rowMapper func(RowsResult) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	convertSQL, args := execSQL(sql, params.kv)
	rows, err := tx.Query(convertSQL, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	return rowMapper(RowsResult{Rows: rows})
}

// QueryRow query one row
func QueryRow(sql string, params sqlParams, rowMapper func(RowResult) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	convertSQL, args := execSQL(sql, params.kv)
	row := tx.QueryRow(convertSQL, args...)

	return rowMapper(RowResult{Row: row})
}

// Page paging data
func Page(table, pk string, column []string, whereKV map[string]interface{}, orderBy string, page, limit int, rowMapper func(RowsResult) error) (int, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return 0, errors.New("db not initialized")
	}

	where, args := getWhereSQL(whereKV, []interface{}{})
	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM "+table+where, args...).Scan(&total)
	if err != nil {
		return total, err
	}

	columns := getColumnSQL(column)
	args = append(args, (page-1)*limit, limit)
	rows, err := tx.Query(
		"SELECT "+columns+
			" FROM "+table+
			" JOIN "+"( SELECT "+pk+" FROM "+table+where+" ORDER BY "+orderBy+" LIMIT ?, ? ) t"+
			" USING ("+pk+")",
		args...,
	)
	defer rows.Close()
	if err != nil {
		return total, err
	}

	return total, rowMapper(RowsResult{Rows: rows})
}

func getSetSQL(kv map[string]interface{}, args []interface{}) (string, []interface{}) {
	setSQL := ""
	for k, v := range kv {
		setSQL += ", " + k + "=?"
		args = append(args, v)
	}
	if len(setSQL) > 2 {
		setSQL = setSQL[2:]
	}

	return setSQL, args
}

func getWhereSQL(kv map[string]interface{}, args []interface{}) (string, []interface{}) {
	whereSQL := ""
	for k, v := range kv {
		whereSQL += " AND " + k + "=?"
		args = append(args, v)
	}
	if len(whereSQL) > 5 {
		whereSQL = " WHERE " + whereSQL[5:]
	}

	return whereSQL, args
}

func getColumnSQL(column []string) string {
	columns := ""
	for _, key := range column {
		columns += ", " + key
	}
	if len(columns) > 2 {
		columns = columns[2:]
	} else {
		columns = "*"
	}

	return columns
}

func execSQL(originalSql string, params map[string]interface{}) (string, []interface{}) {
	convertSql := ""
	args := []interface{}{}

	preIndex := 0
	for i := 0; i < len(originalSql)-1; i++ {
		switch originalSql[i : i+1] {
		case "#":
			key := getKey(originalSql, i+1)
			value := fmt.Sprintf("%v", params[key])
			if len(value) > 0 {
				convertSql += originalSql[preIndex:i] + value
				i += len(key)
				preIndex = i + 1
			}

		case ":":
			key := getKey(originalSql, i+1)
			value := params[key]
			if value != nil {
				convertSql += originalSql[preIndex:i] + "?"
				args = append(args, value)
				i += len(key)
				preIndex = i + 1
			}
		}
	}
	convertSql += originalSql[preIndex:]

	return convertSql, args
}

func getKey(original string, startIndex int) string {
	for j := startIndex; j < len(original); j++ {
		if isSpecificChar([]rune(original[j : j+1])[0]) {
			key := original[startIndex:j]
			return key
		}
	}

	return original[startIndex:]
}

func isSpecificChar(c rune) bool {
	return (c < 128 && specificChar[c]) || c == ' '
}
