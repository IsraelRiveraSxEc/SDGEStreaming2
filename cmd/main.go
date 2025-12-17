package main

import (
	"log"
	"sdgestreaming/api"
	"sdgestreaming/internal/database"
)

func main() {
	// Inicializa la base de datos (SQLite)
	db, err := database.NewSQLiteDB("sdgestreaming.db")
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	defer db.Close()
	// Ejecuta migraciones (crea o actualiza tablas)
	if err := database.RunMigrations(db, "internal/database/migrations"); err != nil {
	    log.Fatalf("Error al ejecutar las migraciones: %v", err)
	}
	// Crea el servidor
	server := api.NewServer(db)
	// Inicia el servidor en el puerto 8080
	log.Println("Iniciando servidor en http://localhost:8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}