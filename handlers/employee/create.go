package handlers

import (
	"net/http"

	"github.com/alifudin-a/go-api-crud-postgre/db"
	"github.com/alifudin-a/go-api-crud-postgre/models"
	"github.com/labstack/echo/v4"
)

// CreateEmployee : create employee
func CreateEmployee(c echo.Context) error {
	db := db.DBConn()

	employee := new(models.Employee)
	if err := c.Bind(employee); err != nil {
		return err
	}

	err := db.QueryRowx("SELECT * FROM employee  WHERE id=$1", employee.ID).StructScan(&employee)
	if err == nil {
		return nil
	}

	sqlStatement := `INSERT INTO employee (name, city) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(sqlStatement, employee.Name, employee.City).Scan(&employee.ID)
	if err != nil {
		return err
	}

	defer db.Close()

	return c.JSON(http.StatusOK, employee)
}
