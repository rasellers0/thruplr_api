package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DoError(c echo.Context) error {
	fmt.Println("got to DoError!")
	return c.String(http.StatusOK, "Unknown Route")
}
