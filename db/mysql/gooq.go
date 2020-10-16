package mysql

import (
	"database/sql"
	"errors"
)

// Gooq instance
type Gooq struct{}

// Exec execute sql
func (g *Gooq) Exec(sql string, args ...interface{}) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	return tx.Exec(sql, args...)
}

// Query query rows
func (g *Gooq) Query(sql string, rowMapper func(*sql.Rows) error, args ...interface{}) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	rows, err := tx.Query(sql, args...)
	if err != nil {
		return err
	}

	defer rows.Close()
	return rowMapper(rows)
}

// QueryRow query one row
func (g *Gooq) QueryRow(sql string, rowMapper func(row *sql.Row) error, args ...interface{}) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	row := tx.QueryRow(sql, args...)

	return rowMapper(row)
}
