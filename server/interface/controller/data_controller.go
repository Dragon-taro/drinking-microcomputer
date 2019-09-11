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

func NewDataController(fc *datastore.FirestoreClient) *DataController {
	return &DataController{
		Interactor: usecase.DataInteractor{
			DataRepository: &datastore.DataRepository{
				FirestoreClient: fc,
			},
		},
	}
}

func (controller *DataController) Create(c Context) error {
	dr := entity.DataReqest{}
	c.Bind(&dr)

	if err := controller.Interactor.Add(dr); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (controller *DataController) HealthCheck(c Context) error {
	return c.JSON(http.StatusOK, "ok")
}
