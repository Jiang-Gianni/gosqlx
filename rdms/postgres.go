package rdms

import (
	"fmt"

	"github.com/Jiang-Gianni/gosqlx/db"
	_ "github.com/lib/pq"
)

// postgresql://root:password@localhost:5432/test
func PostgreSqlConnectionString(conn db.Connection) string {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", conn.User, conn.Password, conn.Host, conn.Port, conn.Datasource)
	if !conn.Ssl.Bool {
		connString += "?sslmode=disable"
	}
	return connString
}

func PostgreSqlSelectAllTables(conn db.Connection) string {
	return fmt.Sprintf("select table_name from information_schema.tables where table_catalog = '%s' and table_schema = 'public';", conn.Datasource)
}
