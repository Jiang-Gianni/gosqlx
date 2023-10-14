package server

import (
	"net/http"
	"strconv"

	"github.com/Jiang-Gianni/gosqlx/rdms"
	"github.com/Jiang-Gianni/gosqlx/views"
	"github.com/go-chi/chi/v5"
)

var (
	databaseGroup = func(r chi.Router) {
		r.Get("/database/{id}", getDatabase)
	}
)

func getDatabase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	connectionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.WriteIndexPage(w, errorPage, views.Error(err))
		return
	}
	conn, err := query.GetConnectionById(ctx, int64(connectionId))
	if err != nil {
		views.WriteIndexPage(w, errorPage, views.Error(err))
		return
	}
	database, err := rdms.DatabaseFromConnection(conn)
	if err != nil {
		views.WriteIndexPage(w, errorPage, views.Error(err))
		return
	}
	tables, err := database.GetTables()
	if err != nil {
		views.WriteIndexPage(w, errorPage, views.Error(err))
		return
	}
	databasePage := &views.Page{
		Title:       conn.Name,
		Description: conn.Datasource,
	}
	views.WriteIndexPage(w, databasePage, views.DatabaseBody(conn, tables))
}
