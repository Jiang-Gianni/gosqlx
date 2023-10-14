package main

import (
	"log"
	"os"

	"github.com/Jiang-Gianni/gosqlx/server"
)

func main() {

	db := server.InitSqlite()

	sqlList := []string{
		"sql/connection.sql",
		"sql/rdms.sql",
	}

	for _, sqlFile := range sqlList {
		b, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Fatal(err)
		}
		r, err := db.Exec(string(b))
		if err != nil {
			log.Println(err)
		}
		log.Println(r.RowsAffected())
	}
}
