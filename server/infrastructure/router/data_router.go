package router

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/controller"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/labstack/echo/v4"
)

func DataRouter(e *echo.Echo, fc *datastore.FirestoreClient) {
	dataController := controller.NewDataController(fc)
	e.GET("/api/hc", func(c echo.Context) error { return dataController.HealthCheck(c) })
	e.GET("/api/data", func(c echo.Context) error { return dataController.Index(c) })
	e.POST("/api/data", func(c echo.Context) error { return dataController.Create(c) })
}
