package rdms

import (
	"database/sql"
	"fmt"

	"github.com/Jiang-Gianni/gosqlx/db"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	*sql.DB
	db.Connection
}

// postgresql://root:password@localhost:5432/test
func NewPostgreSql(conn db.Connection) (*PostgreSql, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", conn.User, conn.Password, conn.Host, conn.Port, conn.Datasource)
	if !conn.Ssl.Bool {
		connString += "?sslmode=disable"
	}
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	postgreSql := &PostgreSql{
		db,
		conn,
	}
	if err != nil {
		return postgreSql, err
	}
	return postgreSql, nil
}

func (p *PostgreSql) GetTables() ([]string, error) {
	tables := []string{}
	selectTable := fmt.Sprintf("select table_name from information_schema.tables where table_catalog = '%s' and table_schema = 'public';", p.Connection.Datasource)
	rows, err := p.Query(selectTable)
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
