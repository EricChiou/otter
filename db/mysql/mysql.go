package mysql

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"otter/config"
	"otter/constants/api"
)

// DB mysql connecting
var DB *sql.DB

var specificCharStr string = `"':.,;(){}[]&|=+-*%/\<>^`
var specificChar [128]bool

// SQLRow db QueryRow result
type Row struct {
	Row *sql.Row
}

// SQLRows db Query result
type Rows struct {
	Rows *sql.Rows
}

// Page QueryPage input struct
type Page struct {
	TableName   string
	ColumnNames []string
	UniqueKey   string
	OrderBy     string
	Page        int
	Limit       int
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
func Insert(table string, columnValues map[string]interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	columnSQL := ""
	valueSQL := ""
	var args []interface{}
	for k, v := range columnValues {
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

// Exec execute sql
func Exec(sql string, params sqlParams, args []interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	execSQL := convertSQL(sql, params.kv)
	return tx.Exec(execSQL, args...)
}

// Query query data
func Query(sql string, params sqlParams, args []interface{}, rowMapper func(result Rows) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	execSQL := convertSQL(sql, params.kv)

	rows, err := tx.Query(execSQL, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	return rowMapper(Rows{Rows: rows})
}

// QueryRow query one row
func QueryRow(sql string, params sqlParams, args []interface{}, rowMapper func(result Row) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	execSQL := convertSQL(sql, params.kv)
	row := tx.QueryRow(execSQL, args...)

	return rowMapper(Row{Row: row})
}

// QueryPage query page format
func QueryPage(page Page, whereSQL string, args []interface{}, rowMapper func(result Rows, total int) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	var columnSQL string
	for _, columnName := range page.ColumnNames {
		columnSQL += ", " + columnName
	}
	if len(columnSQL) > 0 {
		columnSQL = columnSQL[2:]
	} else {
		columnSQL = "*"
	}

	sql := "SELECT " + columnSQL + " "
	sql += "FROM " + page.TableName + " "
	sql += "    JOIN ( "
	sql += "    SELECT " + page.UniqueKey + " FROM " + page.TableName + " "
	sql += "    ORDER BY " + page.OrderBy + " "
	sql += "    LIMIT " + strconv.Itoa((page.Page-1)*page.Limit) + ", " + strconv.Itoa(page.Limit) + " "
	sql += ") t "
	sql += "USING ( " + page.UniqueKey + " ) "
	sql += whereSQL

	rows, err := tx.Query(sql, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	var total int
	sql = "SELECT COUNT(*) FROM " + page.TableName + " " + whereSQL
	row := tx.QueryRow(sql, args...)
	row.Scan(&total)

	rowMapper(Rows{Rows: rows}, total)
	return nil
}

func convertSQL(originalSql string, params map[string]string) string {
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
