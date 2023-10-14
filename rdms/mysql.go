package rdms

import (
	"database/sql"
	"fmt"

	"github.com/Jiang-Gianni/gosqlx/db"
	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	*sql.DB
	db.Connection
}

// root:password@tcp(localhost:3306)/test
func NewMySql(conn db.Connection) (*MySql, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conn.User, conn.Password, conn.Host, conn.Port, conn.Datasource)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	mySql := &MySql{
		db,
		conn,
	}
	if err != nil {
		return mySql, err
	}
	return mySql, nil
}

func (m *MySql) GetTables() ([]string, error) {
	tables := []string{}
	selectTable := fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%s';", m.Connection.Datasource)
	rows, err := m.Query(selectTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
