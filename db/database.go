package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importe o driver PostgreSQL
)

func Connect() error {
	if os.Getenv("CONNECTION_STRING") == "" {
		panic("connection string vazia, carregue as envs")
	}
	connStr := os.Getenv("CONNECTION_STRING")

	// Abre a conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// boas vindas
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return nil
}
