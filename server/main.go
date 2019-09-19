package main

import (
	"log"

	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/firestore"
	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	fc, err := firestore.NewFirestoreClient()
	if err != nil {
		panic(err)
	}
	router.Router(e, fc)

	e.Logger.Fatal(e.Start(":3000"))
}
