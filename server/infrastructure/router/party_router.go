package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/controller"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func PartyRouter(e *echo.Echo, fc *datastore.FirestoreClient) {
	partyController := controller.NewPartyController(fc)
	e.GET("/api/party", func(c echo.Context) error { return partyController.Index(c) })
	e.POST("/api/party", func(c echo.Context) error { return partyController.Create(c) })
	e.PATCH("/api/party/finish", func(c echo.Context) error { return partyController.FinishLatest(c) })
}
