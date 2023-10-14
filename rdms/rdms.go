package rdms

import (
	"errors"

	"github.com/Jiang-Gianni/gosqlx/db"
)

type Database interface {
	GetTables() ([]string, error)
}

func DatabaseFromConnection(conn db.Connection) (Database, error) {
	switch conn.IDRdms {
	case 1:
		return NewPostgreSql(conn)
	case 2:
		return NewMySql(conn)
	default:
		return nil, errors.New("database not found")
	}
}
