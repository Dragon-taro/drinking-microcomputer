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
	latestParty, err := repo.getLatestpartyRef()
	if err != nil {
		return err
	}

	if _, _, err := latestParty.Ref.Collection("Data").Add(repo.FirestoreClient.Ctx, d); err != nil {
		return err
	}

	return nil
}

func (repo *DataRepository) FindAll() ([]entity.Data, error) {
	latestParty, err := repo.getLatestpartyRef()
	if err != nil {
		return nil, err
	}

	docs, err := latestParty.Ref.Collection("Data").OrderBy("CreatedAt", firestore.Desc).Documents(repo.FirestoreClient.Ctx).GetAll()
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
	latestParty, err := repo.getLatestpartyRef()
	if err != nil {
		return entity.Data{}, err
	}

	doc, err := latestParty.Ref.Collection("Data").OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(repo.FirestoreClient.Ctx).Next()
	if err != nil {
		return entity.Data{}, err
	}

	d := entity.Data{}
	mapstructure.Decode(doc.Data(), &d)
	return d, nil
}

func (repo *DataRepository) getLatestpartyRef() (*firestore.DocumentSnapshot, error) {
	return repo.FirestoreClient.Client.Collection("Party").Where("IsFinished", "==", false).Limit(1).Documents(repo.FirestoreClient.Ctx).Next()
}
