package datastore

import "github.com/Dragon-taro/drinking-microcomputer/server/entity"

type DataRepository struct {
	FirestoreClient *FirestoreClient
}

func (repo *DataRepository) Store(d entity.Data) error {
	_, _, err := repo.FirestoreClient.Client.Collection("Data").Add(repo.FirestoreClient.Ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DataRepository) FindAll() (entity.Data, error) {
	return entity.Data{}, nil
}
