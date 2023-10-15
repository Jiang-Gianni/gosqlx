package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Jiang-Gianni/gosqlx/rdms"
	"github.com/Jiang-Gianni/gosqlx/views"
	"github.com/go-chi/chi/v5"
)

var (
	databaseGroup = func(r chi.Router) {
		r.Route("/database/{id}", func(r chi.Router) {
			r.Use(databaseCtx)
			r.Post("/exec", postExec)
			r.Get("/", getDatabase)
		})
	}
)

func databaseCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		ctx = context.WithValue(ctx, databaseContextKey, database)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func databaseFromContext(r *http.Request) *rdms.Database {
	return r.Context().Value(databaseContextKey).(*rdms.Database)
}

func getDatabase(w http.ResponseWriter, r *http.Request) {
	database := databaseFromContext(r)
	tables, err := database.GetTables()
	if err != nil {
		views.WriteIndexPage(w, errorPage, views.Error(err))
		return
	}
	databasePage := &views.Page{
		Title:       database.Connection.Name,
		Description: database.Connection.Datasource,
	}
	views.WriteIndexPage(w, databasePage, views.DatabaseBody(database.Connection, tables))
}

func postExec(w http.ResponseWriter, r *http.Request) {
	database := databaseFromContext(r)
	query := r.FormValue("query")
	rows, err := database.ExecQuery(query)
	if err != nil {
		views.WriteError(w, err)
	}
	views.WriteQueryResult(w, rows)
}
