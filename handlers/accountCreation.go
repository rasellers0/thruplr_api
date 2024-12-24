package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostAccountCreation(c echo.Context) error {
	fmt.Println("got to PostAccountCreation!!")

	db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")

	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(res)
	}
	jsonResponse := "{'value' : 'success'}"
	defer db.Close()
	return c.JSON(http.StatusOK, jsonResponse)
	// return c.String(http.StatusOK, "Account Created!!")
}
