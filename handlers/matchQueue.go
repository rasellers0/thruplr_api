package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetUserMatchCandidates(c echo.Context) error {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := "neo4j+s://0fc6f0f6.databases.neo4j.io"
	dbUser := "neo4j"
	dbPassword := "0Dr_M-2CcIbAX1l4BWOccOap4hDKB3MubVIa9NATbf4"
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPassword, ""))
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established.")
	return c.JSON(http.StatusOK, "some information goes here")
}
