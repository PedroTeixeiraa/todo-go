package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente existentes.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Valida se todas as variáveis necessárias estão definidas
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Erro: Variáveis de ambiente do banco de dados não estão definidas corretamente.")
	}

	// Monta o DSN para conectar ao banco de dados
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	fmt.Printf("DSN (completo): %s\n", dsn)

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Banco de dados conectado com sucesso!")
}
