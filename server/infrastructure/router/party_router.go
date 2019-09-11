package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/controller"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func PartyRouter(e *echo.Echo, fc *datastore.FirestoreClient) {
	partyController := controller.NewPartyController(fc)
	e.GET("/party", func(c echo.Context) error { return partyController.Index(c) })
	e.POST("/party", func(c echo.Context) error { return partyController.Create(c) })
}
