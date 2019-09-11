package usecase

import (
	"time"

	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type DataInteractor struct {
	DataRepository DataRepository
}

func (i *DataInteractor) Add(dr entity.DataReqest) error {
	d, err := i.createData(dr)
	if err != nil {
		return err
	}

	err = i.DataRepository.Store(d)
	if err != nil {
		return err
	}

	return nil
}

func (i *DataInteractor) Data() (entity.Data, error) {
	d, err := i.DataRepository.FindAll()
	if err != nil {
		return entity.Data{}, err
	}

	return d, nil
}

func (i *DataInteractor) createData(dr entity.DataReqest) (entity.Data, error) {
	latestData, err := i.DataRepository.FindLatest()
	if err != nil {
		return entity.Data{}, err
	}

	diff := 0
	totalAmount := latestData.TotalAmount
	if latestData.Amount > dr.Amount {
		diff = latestData.Amount - dr.Amount
		totalAmount = totalAmount + diff
	}

	d := entity.Data{
		Amount:      dr.Amount,
		TotalAmount: totalAmount,
		Diff:        diff,
		CreatedAt:   time.Now(),
	}

	return d, nil
}
