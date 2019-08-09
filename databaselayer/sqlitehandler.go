package databaselayer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //comment
)

//SQLiteHandler ...
type SQLiteHandler struct {
	*SQLHandler
}

//NewSQLiteHandler ...
func NewSQLiteHandler(connection string) (*SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", connection)
	return &SQLiteHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err

}
