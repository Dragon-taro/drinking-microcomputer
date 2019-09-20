package datastore

import (
	"time"

	firestore "cloud.google.com/go/firestore"
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
	"github.com/mitchellh/mapstructure"
)

type PartyRepository struct {
	FirestoreClient *FirestoreClient
}

func (repo *PartyRepository) Store(p entity.Party) error {
	ref, _, err := repo.FirestoreClient.Client.Collection("Party").Add(repo.FirestoreClient.Ctx, p)
	if err != nil {
		return err
	}

	d := entity.Data{
		CreatedAt: time.Now(),
	}

	if _, _, err := ref.Collection("Data").Add(repo.FirestoreClient.Ctx, d); err != nil {
		return err
	}

	return nil
}

func (repo *PartyRepository) FindLatest() (*entity.Party, error) {
	doc, err := repo.latestPartySnap()
	if err != nil {
		return nil, err
	}

	d := &entity.Party{}
	mapstructure.Decode(doc.Data(), &d)
	return d, nil
}

func (repo *PartyRepository) FinishAll() error {
	docs, err := repo.FirestoreClient.Client.Collection("Party").Where("IsFinished", "==", false).Documents(repo.FirestoreClient.Ctx).GetAll()
	if err != nil {
		return err
	}

	isFinished := createFinishData()
	for _, doc := range docs {
		_, err := doc.Ref.Set(repo.FirestoreClient.Ctx, isFinished, firestore.MergeAll)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *PartyRepository) FinishLatest() error {
	doc, err := repo.latestPartySnap()
	if err != nil {
		return err
	}

	isFinished := createFinishData()
	if _, err := doc.Ref.Set(repo.FirestoreClient.Ctx, isFinished, firestore.MergeAll); err != nil {
		return err
	}

	return nil
}

func (repo *PartyRepository) latestPartySnap() (*firestore.DocumentSnapshot, error) {
	return repo.FirestoreClient.Client.Collection("Party").OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(repo.FirestoreClient.Ctx).Next()
}

func createFinishData() map[string]interface{} {
	return map[string]interface{}{
		"IsFinished": true,
		"FinishedAt": time.Now(),
	}
}
