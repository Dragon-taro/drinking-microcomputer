package usecase

import (
	"time"

	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type PartyInteractor struct {
	PartyRepository PartyRepository
}

func (i *PartyInteractor) Add() (*entity.Party, error) {
	// 終わってないPartyは一個しか存在しないようにする
	if err := i.PartyRepository.FinishAll(); err != nil {
		return nil, err
	}

	p := entity.Party{
		CreatedAt:  time.Now(),
		IsFinished: false,
	}
	if err := i.PartyRepository.Store(p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (i *PartyInteractor) LatestParty() (*entity.Party, error) {
	p, err := i.PartyRepository.FindLatest()

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (i *PartyInteractor) FinishLatest() error {
	if err := i.PartyRepository.FinishLatest(); err != nil {
		return err
	}

	return nil
}
