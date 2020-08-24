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

// SQLRow db QueryRow result
type Row struct {
	Row *sql.Row
}

// SQLRows db Query result
type Rows struct {
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

	return tx.Exec("DELETE FROM "+table+" "+whereSQL, args...)
}

// Query query data
func Query(sql string, params sqlParams, rowMapper func(Rows) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	execSQL, args := convertSQL(sql, params.kv)

	rows, err := tx.Query(execSQL, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	return rowMapper(Rows{Rows: rows})
}

// QueryRow query one row
func QueryRow(sql string, params sqlParams, rowMapper func(Row) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	execSQL, args := convertSQL(sql, params.kv)
	row := tx.QueryRow(execSQL, args...)

	return rowMapper(Row{Row: row})
}

// WhereSQL get where sql
func WhereSQL(params whereParams) string {
	whereSQL := ""
	for k, v := range params.kv {
		whereSQL += "AND " + k + "=" + v
	}
	if len(whereSQL) > 5 {
		whereSQL = "WHERE " + whereSQL[4:] + " "
	}

	return whereSQL
}

func getSetSQL(kv map[string]interface{}, args []interface{}) (string, []interface{}) {
	setSQL := ""
	for k, v := range kv {
		setSQL += ", " + k + "=?"
		args = append(args, v)
	}
	if len(setSQL) > 2 {
		setSQL = setSQL[2:] + " "
	}

	return setSQL, args
}

func getWhereSQL(kv map[string]interface{}, args []interface{}) (string, []interface{}) {
	whereSQL := ""
	for k, v := range kv {
		whereSQL += "AND " + k + "=?"
		args = append(args, v)
	}
	if len(whereSQL) > 5 {
		whereSQL = "WHERE " + whereSQL[4:]
	}

	return whereSQL, args
}

func convertSQL(originalSql string, params map[string]interface{}) (string, []interface{}) {
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
