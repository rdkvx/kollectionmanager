package db

import (
	"fmt"
	"kollectionmanager/m/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, err error) {
	if os.Getenv(utils.ConnectionString) == "" {
		panic(utils.ConnStrErr)
	}
	connStr := os.Getenv(utils.ConnectionString)

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println(utils.DbConnectionSuccess)

	return db, nil
}
