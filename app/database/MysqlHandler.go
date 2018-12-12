package database

import (
	"database/sql"
	"goRest/app/interfaces"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

type MySQL struct {
	Conn *sql.DB
}

func (db MySQL) Execute(statement string) {
	db.Conn.Exec(statement)
}

func (db MySQL) Query(statement string) (interfaces.IRow, error) {
	rows, err := db.Conn.Query(statement)
	if err != nil {
		log.Fatal(err)
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

type SQLRow struct {
	Rows *sql.Rows
}

func (r SQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (r SQLRow) Next() bool {
	return r.Rows.Next()
}
