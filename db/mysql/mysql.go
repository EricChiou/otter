package mysql

import (
	"database/sql"
	"strings"

	"otter/config"
	cons "otter/constants"
)

// DB mysql connecting
var DB *sql.DB

// Init connect MySQL
func Init() (err error) {
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
func ErrMsgHandler(err error) cons.ApiResult {
	if strings.Contains(err.Error(), "Duplicate") {
		return cons.RSDuplicate
	} else {
		return cons.RSDBError
	}
}

// Insert insert data
func Insert(tx *sql.Tx, table string, kv map[string]interface{}) (sql.Result, error) {
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
func Update(tx *sql.Tx, table string, setKV map[string]interface{}, whereKV map[string]interface{}) (sql.Result, error) {
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
func Delete(tx *sql.Tx, table string, whereKV map[string]interface{}) (sql.Result, error) {
	var args []interface{}
	where, args := WhereString(whereKV, args)

	return tx.Exec("DELETE FROM "+table+where, args...)
}

// Query query data
func Query(tx *sql.Tx, table string, column []string, whereKV map[string]interface{}, orderBy string) (*sql.Rows, error) {
	var args []interface{}
	columns := ColumnString(column)
	where, args := WhereString(whereKV, args)
	if len(orderBy) > 0 {
		orderBy = " ORDER BY " + orderBy
	}

	return tx.Query("SELECT "+columns+" FROM "+table+where+orderBy, args...)
}

// QueryRow query one data
func QueryRow(tx *sql.Tx, table string, column []string, whereKV map[string]interface{}) *sql.Row {
	var args []interface{}
	columns := ColumnString(column)
	where, args := WhereString(whereKV, args)

	return tx.QueryRow("SELECT "+columns+" FROM "+table+where, args...)
}

// Page paging data
func Page(tx *sql.Tx, table, pk string, column []string, whereKV map[string]interface{}, orderBy string, page, limit int) (*sql.Rows, error) {
	var args []interface{}
	columns := ColumnString(column)
	where, args := WhereString(whereKV, args)
	args = append(args, (page-1)*limit, limit)

	return tx.Query(
		"SELECT "+columns+
			" FROM "+table+
			" JOIN "+
			"( SELECT "+pk+" FROM "+table+where+" ORDER BY "+orderBy+" LIMIT ?, ? ) t"+
			" USING ("+pk+")",
		args...,
	)
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

// ExecString get exec string, "?" will replaced by params string and "!" will replaced by "?"
func ExecString(original string, params ...string) string {
	for _, param := range params {
		original = strings.Replace(original, "?", param, 1)
	}
	original = strings.Replace(original, "!", "?", -1)
	return original
}
