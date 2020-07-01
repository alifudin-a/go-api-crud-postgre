package db

import (
	"fmt"
	"log"
	"os"

	"github.com/alifudin-a/go-api-crud-postgre/utils"
	"github.com/jmoiron/sqlx"
)

func DBConn() *sqlx.DB {
	utils.Init()
	var (
		dbDriver = os.Getenv("DB_DRIVER")
		dbHost   = os.Getenv("DB_HOST")
		dbUser   = os.Getenv("DB_USER")
		dbPass   = os.Getenv("DB_PASS")
		dbName   = os.Getenv("DB_NAME")
		dbPort   = os.Getenv("DB_PORT")
		sslmode  = os.Getenv("SSLMODE")
	)

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUser, dbPass, dbName, dbHost, dbPort, sslmode)
	db, err := sqlx.Open(dbDriver, psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("Database succesfully connected!")

	return db
}
