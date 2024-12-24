package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	util "thruplr_api/util"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	fmt.Println("so it'll hit here, ok")
	return c.String(http.StatusOK, "Fuck you!")
}

func GetUserData(c echo.Context) error {
	jsonFile, err := os.Open("./json/user.json")
	if err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "No User Data")
	}
	log.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	user := new(util.User)
	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, "No User Data")
	} else {
		log.Println("no errors found")
	}
	log.Println("User id: " + strconv.Itoa(int(user.ID)))
	log.Println("User First Name: " + user.FirstName)
	log.Println("User Last Name: " + user.LastName)

	return c.JSON(http.StatusOK, user)
}
