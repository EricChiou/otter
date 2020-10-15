package mysql

import (
	"database/sql"
	"errors"
)

// Gooq instance
type Gooq struct{}

// Exec execute sql
func (g *Gooq) Exec(sql string) (sql.Result, error) {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return nil, errors.New("db not initialized")
	}

	return tx.Exec(sql)
}

// Query query rows
func (g *Gooq) Query(sql string, rowMapper func(result Rows) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	rows, err := tx.Query(sql)
	defer rows.Close()
	if err != nil {
		return err
	}

	return rowMapper(Rows{Rows: rows})
}

// QueryRow query one row
func (g *Gooq) QueryRow(sql string, rowMapper func(result Row) error) error {
	tx, err := DB.Begin()
	defer tx.Commit()
	if err != nil {
		return errors.New("db not initialized")
	}

	row := tx.QueryRow(sql)

	return rowMapper(Row{Row: row})
}
