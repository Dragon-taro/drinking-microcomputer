package datastore

import "github.com/Dragon-taro/drinking-microcomputer/server/entity"

type DataRepository struct {
	FirestoreHandler *FirestoreHandler
}

func (repo *DataRepository) Store(d entity.Data) error {
	_, _, err := repo.FirestoreHandler.Client.Collection("Data").Add(repo.FirestoreHandler.Ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DataRepository) FindAll() (entity.Data, error) {
	return entity.Data{}, nil
}
