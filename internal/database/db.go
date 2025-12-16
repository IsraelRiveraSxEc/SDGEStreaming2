package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
    Conn *sql.DB
}

//NewSQLiteDB crea e inicializa la conexión a la base de datos.
func NewSQLiteDB(path string) (*DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
// verificación de conectividad
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}
// Close cierra la conexión a la base de datos
func (db *DB) Close() error {
	return db.Conn.Close()
}