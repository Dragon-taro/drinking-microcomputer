package controller

import (
	"net/http"

	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/Dragon-taro/drinking-microcomputer/server/usecase"
)

type DataController struct {
	Interactor usecase.DataInteractor
}

func NewDataController(fh *datastore.FirestoreHandler) *DataController {
	return &DataController{
		Interactor: usecase.DataInteractor{
			DataRepository: &datastore.DataRepository{
				FirestoreHandler: fh,
			},
		},
	}
}

func (controller *DataController) Create(c Context) error {
	d := entity.Data{}
	c.Bind(&d)

	if err := controller.Interactor.Add(d); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (controller *DataController) HealthCheck(c Context) error {
	return c.JSON(http.StatusOK, "ok")
}
