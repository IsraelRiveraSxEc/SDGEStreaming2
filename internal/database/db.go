package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
    Conn *sql.DB
}

//NewSQLiteDB crea e inicializa la conexión a la base de datos.
func NewPostgresDB() (*DB, error) {
	conn, err := sql.Open("postgres", dsn)
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