-- name: GetAllRdms :many
select * from rdms;

-- name: GetAllConnections :many
select * from connection;

-- name: GetConnectionById :one
select * from connection where id = ?;

-- name: DeleteConnection :exec
delete from connection where id = ?;

-- name: UpdateConnection :one
update connection
set id_rdms = ?,
    ssl = ?,
    name = ?,
    host = ?,
    port = ?,
    datasource = ?,
    user = ?,
    password = ?
where id = ?
returning *;

-- name: InsertConnection :one
insert into connection(id_rdms, ssl, name, host, port, datasource, user, password)
values (?, ?, ?, ?, ?, ?, ?, ?) returning *;