package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-api-crud-postgre/db"
	"github.com/alifudin-a/go-api-crud-postgre/models"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func ListEmployee(c echo.Context) (err error) {
	db := db.DBConn()

	pageLimit, _ := strconv.Atoi(c.QueryParam("pageLimit"))
	if pageLimit == 0 {
		pageLimit = 10
	}

	sqlStatement := "SELECT * FROM employee ORDER BY id ASC limit $1"
	rows, err := db.Queryx(sqlStatement, pageLimit)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, rows)
	}

	defer db.Close()

	employees := models.Employees{}

	for rows.Next() {
		employee := models.Employee{}
		err = rows.Scan(&employee.ID, &employee.Name, &employee.City)
		if err != nil {
			log.Println(err)
		}

		employees.Employees = append(employees.Employees, employee)
	}

	return c.JSON(http.StatusOK, employees)
}
