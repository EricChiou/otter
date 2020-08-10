package mysql

import (
	"database/sql"
	"errors"
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
func Insert(table string, kv map[string]interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	keys := ""
	values := ""
	var args []interface{}
	for k, v := range kv {
		keys += ", " + k
		values += ", ?"
		args = append(args, v)
	}
	if len(keys) > 2 {
		keys = keys[2:]
	}
	if len(values) > 2 {
		values = values[2:]
	}

	return tx.Exec("INSERT INTO "+table+"( "+keys+" ) values( "+values+" )", args...)
}

// Update upadte data
func Update(table string, setKV map[string]interface{}, whereKV map[string]interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	var args []interface{}
	set := ""
	for k, v := range setKV {
		set += ", " + k + "=?"
		args = append(args, v)
	}
	if len(set) > 2 {
		set = set[2:]
	}
	where, args := WhereString(whereKV, args)

	return tx.Exec("UPDATE "+table+" SET "+set+where, args...)
}

// Delete delete data
func Delete(table string, whereKV map[string]interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}
	where, args := WhereString(whereKV, []interface{}{})

	return tx.Exec("DELETE FROM "+table+where, args...)
}

// Query query data
func Query(sql string, params map[string]string, args []interface{}, rowMapper func(RowsResult) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	convertSql := execSQL(sql, params)
	rows, err := tx.Query(convertSql, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	return rowMapper(RowsResult{Rows: rows})
}

// QueryRow query one data
func QueryRow(sql string, params map[string]string, args []interface{}, rowMapper func(RowResult) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	convertSql := execSQL(sql, params)
	row := tx.QueryRow(convertSql, args...)
	return rowMapper(RowResult{Row: row})
}

// Page paging data
func Page(table, pk string, column []string, whereKV map[string]interface{}, orderBy string, page, limit int, rowMapper func(RowsResult) error) (int, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return 0, errors.New("db not initialized")
	}

	where, args := WhereString(whereKV, []interface{}{})
	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM "+table+where, args...).Scan(&total)
	if err != nil {
		return total, err
	}

	columns := ColumnString(column)
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

// WhereString get db where string
func WhereString(whereKV map[string]interface{}, args []interface{}) (string, []interface{}) {
	where := ""
	for k, v := range whereKV {
		where += " AND " + k + "=?"
		args = append(args, v)
	}
	if len(where) > 5 {
		where = " WHERE " + where[5:]
	}

	return where, args
}

// ColumnString get db column string
func ColumnString(column []string) string {
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

func execSQL(originalSql string, params map[string]string) string {
	convertSql := ""
	preIndex := 0
	for i := 0; i < len(originalSql)-1; i++ {
		if originalSql[i:i+1] == "#" {
			key := getKey(originalSql, i+1)
			value := params[key]
			if len(value) > 0 {
				convertSql += originalSql[preIndex:i] + value
				i += len(key)
				preIndex = i + 1
			}
		}
	}
	convertSql += originalSql[preIndex:]

	return convertSql
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
