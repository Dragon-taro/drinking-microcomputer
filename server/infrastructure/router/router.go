package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, fc *datastore.FirestoreClient) {
	DataRouter(e, fc)
	PartyRouter(e, fc)
}
