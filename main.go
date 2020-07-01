package main

import (
	"fmt"

	"github.com/alifudin-a/go-api-crud-postgre/db"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("GO")
	db.DBConn()
}
