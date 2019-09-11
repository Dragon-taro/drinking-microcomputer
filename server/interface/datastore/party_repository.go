package datastore

import (
	firestore "cloud.google.com/go/firestore"
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
	"github.com/mitchellh/mapstructure"
)

type PartyRepository struct {
	FirestoreClient *FirestoreClient
}

func (repo *PartyRepository) Store(p entity.Party) error {
	_, _, err := repo.FirestoreClient.Client.Collection("Party").Add(repo.FirestoreClient.Ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PartyRepository) FindLatest() (entity.Party, error) {
	doc, err := repo.FirestoreClient.Client.Collection("Party").OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(repo.FirestoreClient.Ctx).Next()
	if err != nil {
		return entity.Party{}, err
	}

	d := entity.Party{}
	mapstructure.Decode(doc.Data(), &d)
	return d, nil
}
