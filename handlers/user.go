package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserData struct {
	UserID      string
	AccountType string
	FirstName   string
	MiddleName  string
	LastName    string
	DOB         string
	Location    string
	Pronouns    string
	Orientation string
	Race1       string
	Race2       string
	Ethnicity   string
	Religion    string
	Politics    string
	Family      string
	FamilyPlans string
	LookingFor  string
	AboutMe     string
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Thruplr back end is successfully responding to REST requests.")
}

func getUserDataFromLogin(userid string) string {
	userQuery := `select u.user_id, u.first_name, u.middle_name, u.last_name, u.account_type,
	fd.pronouns, fd.orientation, fd.dob, ifnull(fd.race1, '') as race1, ifnull(fd.race2, '') as race2, 
	ifnull(fd.ethnicity, '') as ethnicity, fd.location, ifnull(usd.about_me, '') as about_me, 
	ifnull(usd.religion, '') as religion, ifnull(usd.politics, '') as politics, ifnull(usd.family, '') as family,
	ifnull(usd.family_plans, '') as family_plans, ifnull(usd.looking_for, '') as looking_for
	from users u
	left join user_demographics fd on fd.user_id = u.user_id
	left join user_details usd on u.user_id = fd.user_id
	where u.user_id = ?`

	db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Query(userQuery, userid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	user_values := populateUserData(stmt)

	res, err := json.Marshal(user_values)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)

}

func populateUserData(stmt *sql.Rows) UserData {
	var (
		user_values UserData
	)

	if stmt.Next() {
		err := stmt.Scan(&user_values.UserID, &user_values.FirstName, &user_values.MiddleName, &user_values.LastName,
			&user_values.AccountType, &user_values.Pronouns, &user_values.Orientation, &user_values.DOB, &user_values.Race1,
			&user_values.Race2, &user_values.Ethnicity, &user_values.Location, &user_values.AboutMe, &user_values.Religion,
			&user_values.Politics, &user_values.Family, &user_values.FamilyPlans, &user_values.LookingFor)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user_values

}

func GetUserData(c echo.Context) error {
	userQuery := `select u.user_id, u.first_name, u.middle_name, u.last_name, u.account_type,
	fd.pronouns, fd.orientation, fd.dob, ifnull(fd.race1, '') as race1, ifnull(fd.race2, '') as race2, 
	ifnull(fd.ethnicity, '') as ethnicity, fd.location, ifnull(usd.about_me, '') as about_me, 
	ifnull(usd.religion, '') as religion, ifnull(usd.politics, '') as politics, ifnull(usd.family, '') as family,
	ifnull(usd.family_plans, '') as family_plans, ifnull(usd.looking_for, '') as looking_for
	from users u
	left join user_demographics fd on fd.user_id = u.user_id
	left join user_details usd on u.user_id = fd.user_id
	where u.user_id = ?`

	userid := c.Param("id")

	db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Query(userQuery, userid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	user_values := populateUserData(stmt)

	return c.JSON(http.StatusOK, user_values)
}
