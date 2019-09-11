package usecase

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type DataInteractor struct {
	DataRepository DataRepository
}

func (i *DataInteractor) Add(d entity.Data) error {
	if err := i.DataRepository.Store(d); err != nil {
		return err
	}

	return nil
}
