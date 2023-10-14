package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jiang-Gianni/gosqlx/db"
	"github.com/Jiang-Gianni/gosqlx/views"
	"github.com/go-chi/chi/v5"
)

var (
	connectionPage = &views.Page{
		Title:       "GoSQLX",
		Description: "GoSQLXperience: A Seamless SQL Client Powered by Go and HTMX",
	}
	connectionGroup = func(r chi.Router) {
		r.Get("/", getIndex)
		r.Post("/connection", postConnection)
		r.Delete("/connection/{id}", deleteConnection)
		r.Post("/connection/{id}", postConnectionId)
	}
)

func insertConnectionParams(conn *db.Connection) db.InsertConnectionParams {
	return db.InsertConnectionParams{
		IDRdms:     conn.IDRdms,
		Ssl:        conn.Ssl,
		Name:       conn.Name,
		Host:       conn.Host,
		Port:       conn.Port,
		Datasource: conn.Datasource,
		User:       conn.User,
		Password:   conn.Password,
	}
}

func updateConnectionParams(conn *db.Connection) db.UpdateConnectionParams {
	return db.UpdateConnectionParams{
		ID:         conn.ID,
		IDRdms:     conn.IDRdms,
		Ssl:        conn.Ssl,
		Name:       conn.Name,
		Host:       conn.Host,
		Port:       conn.Port,
		Datasource: conn.Datasource,
		User:       conn.User,
		Password:   conn.Password,
	}
}

func connectionFromRequest(r *http.Request) (*db.Connection, error) {
	port, err := strconv.Atoi(r.FormValue("port"))
	if err != nil {
		return nil, err
	}
	idRdms, err := strconv.Atoi(r.FormValue("id_rdms"))
	if err != nil {
		return nil, err
	}
	var conn = &db.Connection{
		IDRdms: int64(idRdms),
		Ssl: sql.NullBool{
			Bool:  r.FormValue("ssl") == "on",
			Valid: true,
		},
		Name:       r.FormValue("name"),
		Host:       r.FormValue("host"),
		Port:       int64(port),
		Datasource: r.FormValue("datasource"),
		User:       r.FormValue("user"),
		Password:   r.FormValue("password"),
	}
	return conn, nil
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conns, err := query.GetAllConnections(ctx)
	if err != nil {
		views.WriteError(w, err)
		fmt.Println(err)
	}
	rdms, err := query.GetAllRdms(ctx)
	if err != nil {
		views.WriteError(w, err)
	}
	views.WriteIndexPage(w, connectionPage, views.ConnectionBody(conns, rdms))
}

func postConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conn, err := connectionFromRequest(r)
	if err != nil {
		views.WriteError(w, err)
	}
	new, err := query.InsertConnection(ctx, insertConnectionParams(conn))
	if err != nil {
		views.WriteError(w, err)
	}
	rdms, err := query.GetAllRdms(ctx)
	if err != nil {
		views.WriteError(w, err)
	}
	views.WriteModifyConnection(w, new, rdms)
}

func postConnectionId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conn, err := connectionFromRequest(r)
	if err != nil {
		views.WriteError(w, err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.WriteError(w, err)
	}
	conn.ID = int64(id)
	updated, err := query.UpdateConnection(ctx, updateConnectionParams(conn))
	if err != nil {
		views.WriteError(w, err)
	}
	rdms, err := query.GetAllRdms(ctx)
	if err != nil {
		views.WriteError(w, err)
	}
	views.WriteModifyConnection(w, updated, rdms)
}

func deleteConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.WriteError(w, err)
	}
	err = query.DeleteConnection(ctx, int64(id))
	if err != nil {
		views.WriteError(w, err)
	}
}
