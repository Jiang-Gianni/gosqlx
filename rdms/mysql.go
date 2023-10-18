package rdms

import (
	"fmt"

	"github.com/Jiang-Gianni/gosqlx/db"
	_ "github.com/go-sql-driver/mysql"
)

// root:password@tcp(localhost:3306)/test
func MySqlConnectionString(conn db.Connection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&multiStatements=true&interpolateParams=true", conn.User, conn.Password, conn.Host, conn.Port, conn.Datasource)
}

func MySqlSelectAllTables(conn db.Connection) string {
	return fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%s';", conn.Datasource)
}
