package db

import (
	"database/sql"
	"strings"

	cons "otter/constants"
)

// MySQL mysql connecting
var MySQL *sql.DB

// InitMySQL connect MySQL
func InitMySQL(addr, port, userName, password, dbName string) (err error) {
	MySQL, err = sql.Open("mysql", userName+":"+password+"@tcp("+addr+":"+port+")/"+dbName)
	return err
}

// CloseMySQL close mysql connecting
func CloseMySQL() {
	if MySQL != nil {
		MySQL.Close()
	}
}

// MySQLErrMsgHandler error message handler
func MySQLErrMsgHandler(err error) string {
	if strings.Contains(err.Error(), "Duplicate") {
		return cons.APIResult.Duplicate
	} else {
		return cons.APIResult.DBError
	}
}
