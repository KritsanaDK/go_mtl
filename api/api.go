package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"strconv"

	module "github.com/KaDingMeaw/godb/modules"
	"github.com/labstack/echo/v4"
)

func GetItem(c echo.Context, db *sql.DB) error {
	item, err := module.GetAllItem(db)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, item)
}

func GetItemID(c echo.Context, db *sql.DB) error {

	// Get the ID from the URL parameter
	id := c.Param("id")
	num, err := strconv.Atoi(id)

	fmt.Println(id)

	item, err := module.GetItem(db, num)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, item)
}

func DeleteItemID(c echo.Context, db *sql.DB) error {

	// Get the ID from the URL parameter
	id := c.Param("id")
	num, err := strconv.Atoi(id)

	fmt.Println(id)

	err = module.DelItem(db, num)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	item := map[string]string{"message": "OK"}

	return c.JSON(http.StatusOK, item)
}
