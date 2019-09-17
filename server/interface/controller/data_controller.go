package controller

import (
	"net/http"

	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/Dragon-taro/drinking-microcomputer/server/usecase"
	"gopkg.in/go-playground/validator.v9"
)

type DataController struct {
	Interactor usecase.DataInteractor
}

var validate *validator.Validate

func NewDataController(fc *datastore.FirestoreClient) *DataController {
	validate = validator.New()

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

	if err := validate.Struct(dr); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	if err := controller.Interactor.Add(dr); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "success")
}

func (controller *DataController) HealthCheck(c Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func (controller *DataController) Index(c Context) error {
	d, err := controller.Interactor.Data()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, d)
}
