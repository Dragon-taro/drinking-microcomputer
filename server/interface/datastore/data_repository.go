package datastore

import (
	firestore "cloud.google.com/go/firestore"
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
	"github.com/mitchellh/mapstructure"
)

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

func (repo *DataRepository) FindAll() ([]entity.Data, error) {
	docs, err := repo.FirestoreClient.Client.Collection("Data").OrderBy("CreatedAt", firestore.Desc).Documents(repo.FirestoreClient.Ctx).GetAll()
	if err != nil {
		return make([]entity.Data, 0), err
	}
	data := make([]entity.Data, len(docs))
	for i, doc := range docs {
		d := entity.Data{}
		mapstructure.Decode(doc.Data(), &d)
		data[i] = d
	}
	return data, nil
}

func (repo *DataRepository) FindLatest() (entity.Data, error) {
	doc, err := repo.FirestoreClient.Client.Collection("Data").OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(repo.FirestoreClient.Ctx).Next()
	if err != nil {
		return entity.Data{}, err
	}

	d := entity.Data{}
	mapstructure.Decode(doc.Data(), &d)
	return d, nil
}
