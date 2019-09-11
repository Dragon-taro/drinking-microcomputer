package usecase

import (
	"time"

	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type PartyInteractor struct {
	PartyRepository PartyRepository
}

func (i *PartyInteractor) Add() error {
	p := entity.Party{
		CreatedAt:  time.Now(),
		IsFinished: false,
		Data: []entity.Data{
			entity.Data{
				CreatedAt: time.Now(),
			},
		},
	}

	if err := i.PartyRepository.Store(p); err != nil {
		return err
	}

	return nil
}

func (i *PartyInteractor) LatestParty() (entity.Party, error) {
	p, err := i.PartyRepository.FindLatest()

	if err != nil {
		return entity.Party{}, err
	}

	return p, nil
}
