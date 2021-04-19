package interfaces

import "database/sql"

type IDbHandler interface {
	Execute(statement string)
	Query(statement string, args ...interface{}) (IRow, error)
	PrepareAndExec(statement string, args ...interface{}) (sql.Result, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
