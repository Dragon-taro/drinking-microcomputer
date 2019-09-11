package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/controller"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func DataRouter(e *echo.Echo, fh *datastore.FirestoreHandler) {
	dataController := controller.NewDataController(fh)
	e.GET("/hc", func(c echo.Context) error { return dataController.HealthCheck(c) })
	e.POST("/data", func(c echo.Context) error { return dataController.Create(c) })
}
