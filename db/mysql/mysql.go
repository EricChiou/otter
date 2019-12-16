package mysql

import (
	"database/sql"
	"strings"

	cons "otter/constants"
)

// DB mysql connecting
var DB *sql.DB

// Init connect MySQL
func Init(addr, port, userName, password, dbName string) (err error) {
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
func ErrMsgHandler(err error) string {
	if strings.Contains(err.Error(), "Duplicate") {
		return cons.APIResult.Duplicate
	} else {
		return cons.APIResult.DBError
	}
}
