package rdms

import (
	"database/sql"

	"github.com/Jiang-Gianni/gosqlx/db"
)

type Database struct {
	Connection db.Connection
	*sql.DB
	connectionString string
	driver           string
	selectAllTable   string
}

func DatabaseFromConnection(conn db.Connection) (*Database, error) {
	database := &Database{
		Connection: conn,
	}
	switch conn.IDRdms {
	case 1:
		database.connectionString = PostgreSqlConnectionString(conn)
		database.driver = "postgres"
		database.selectAllTable = PostgreSqlSelectAllTables(conn)
	case 2:
		database.connectionString = MySqlConnectionString(conn)
		database.driver = "mysql"
		database.selectAllTable = MySqlSelectAllTables(conn)
	default:
	}
	db, err := sql.Open(database.driver, database.connectionString)
	if err != nil {
		return nil, err
	}
	database.DB = db
	err = database.DB.Ping()
	if err != nil {
		return database, err
	}
	return database, nil
}

func (d *Database) GetTables() ([]string, error) {
	tables := []string{}
	rows, err := d.DB.Query(d.selectAllTable)
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
