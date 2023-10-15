package server

type contextKey struct {
	string
}

var (
	databaseContextKey = contextKey{"databaseContextKey"}
)
