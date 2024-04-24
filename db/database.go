package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, err error) {
	if os.Getenv("CONNECTION_STRING") == "" {
		panic("connection string vazia, carregue as envs")
	}
	connStr := os.Getenv("CONNECTION_STRING")

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return db, nil
}
