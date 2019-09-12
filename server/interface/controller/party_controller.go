package controller

import (
	"net/http"

	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"github.com/Dragon-taro/drinking-microcomputer/server/usecase"
)

type PartyController struct {
	Interactor usecase.PartyInteractor
}

func NewPartyController(fc *datastore.FirestoreClient) *PartyController {
	return &PartyController{
		Interactor: usecase.PartyInteractor{
			PartyRepository: &datastore.PartyRepository{
				FirestoreClient: fc,
			},
		},
	}
}

func (controller *PartyController) Create(c Context) error {
	if err := controller.Interactor.Add(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (controller *PartyController) Index(c Context) error {
	p, err := controller.Interactor.LatestParty()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, p)
}

func (controller *PartyController) FinishLatest(c Context) error {
	if err := controller.Interactor.FinishLatest(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "success")
}
