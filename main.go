package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	handler "thruplr_api/handlers"

	// "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// db connection
	dbc := dbConnect()

	var (
		id   int
		name string
	)

	rows, err := dbc.Query("select user_id, CONCAT(first_name, ' ', last_name) as name from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		fmt.Println(strconv.Itoa(id) + ": " + name)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(id, name)
		}
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.Hello)
	e.GET("/user/:id", handler.GetUserData)
	e.POST("createAccount", handler.PostAccountCreation)
	e.POST("/login", handler.DoLogin)
	e.POST("*", handler.DoError)
	e.GET("*", handler.DoError)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return db
}
