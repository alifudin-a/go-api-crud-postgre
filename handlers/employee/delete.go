package handlers

import (
	"net/http"

	"github.com/alifudin-a/go-api-crud-postgre/db"
	"github.com/labstack/echo/v4"
)

// DeleteEmployee : delete employee by id
func DeleteEmployee(c echo.Context) error {
	db := db.DBConn()

	id := c.Param("id")

	sqlStatement := "DELETE FROM employee WHERE id=$1"
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	} else {
		count, err := res.RowsAffected()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to delete employee!")
		}

		if count == 0 {
			return c.JSON(http.StatusInternalServerError, "Employee id not found!")
		}
	}

	return c.JSON(http.StatusOK, "Employee Deleted")
}
