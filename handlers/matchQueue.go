package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserMatchCandidates(c echo.Context) error {

	return c.JSON(http.StatusOK, "some information goes here")
}
