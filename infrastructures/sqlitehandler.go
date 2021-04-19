package infrastructures

import (
	"database/sql"
	"fmt"

	"github.com/rizkypascal/go-clean-rest-api/interfaces"
)

type SQLiteHandler struct {
	Conn *sql.DB
}

func (handler *SQLiteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SQLiteHandler) Query(statement string, args ...interface{}) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement, args)

	if err != nil {
		fmt.Println(err)
		return new(SqliteRow), err
	}
	row := new(SqliteRow)
	row.Rows = rows

	return row, nil
}

func (handler *SQLiteHandler) PrepareAndExec(statement string, args ...interface{}) (sql.Result, error) {
	stmt, err := handler.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(args)

	if err != nil {
		return nil, err
	}

	return result, nil
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
