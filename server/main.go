package main

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/firestore"
	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fh, err := firestore.NewFirestoreClient()
	if err != nil {
		panic(err)
	}
	router.Router(e, fh)

	e.Logger.Fatal(e.Start(":3000"))
}
