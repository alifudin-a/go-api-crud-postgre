package routes

import (
	"log"
	"os"

	handlers "github.com/alifudin-a/go-api-crud-postgre/handlers/employee"
	"github.com/alifudin-a/go-api-crud-postgre/utils"
	"github.com/labstack/echo/v4"
)

func Routes() {
	utils.Init()
	e := echo.New()

	v1 := e.Group("/v1")
	v1.GET("/employee", handlers.ListEmployee)

	log.Fatal(e.Start(os.Getenv("API_PORT")))
}
