package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo, fc *datastore.FirestoreClient) {
	e.Use(middleware.CORS())

	DataRouter(e, fc)
	PartyRouter(e, fc)
}
