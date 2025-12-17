package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"sdgestreaming/api"
	"sdgestreaming/internal/database"
)

func main() {
	// Carga las variables desde .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("No se encontro archivo .env, usando variables de entorno del sistema")
	}
	// Lee variables del .env
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	
	// Constuye DSN estilo key=value para PostgreSQL
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, dbname, sslmode,
	)
	// Inicializa la base de datos (PostgreSQL)
	db, err := database.NewPostgresDB(dsn)
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