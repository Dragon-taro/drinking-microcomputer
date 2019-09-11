package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, fh *datastore.FirestoreHandler) {
	DataRouter(e, fh)
}
