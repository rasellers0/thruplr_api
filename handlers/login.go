package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DoLogin(c echo.Context) error {
	fmt.Println("got to PostAccountCreation!!")

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	} else {
		name := json_map["username"]
		pass := json_map["password"]

		var (
			pw_value string
		)

		db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")

		stmt, err := db.Query("select AES_DECRYPT(hash_creds, 'some_passkey') from users where username = ?", name)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		if stmt.Next() {
			err := stmt.Scan(&pw_value)
			if err != nil {
				log.Fatal(err)
			}

		}

		var (
			jsonResponse string
		)

		if pw_value == pass {
			jsonResponse = "{\"message\": \"Login Successful!!\", \"value\": \"success\"}"
		} else {
			jsonResponse = "{\"message\": \"Password Incorect\", \"value\": \"failure\"}"
		}

		return c.JSON(http.StatusOK, jsonResponse)
	}

}
