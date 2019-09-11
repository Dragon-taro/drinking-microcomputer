package firestore

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"google.golang.org/api/option"
)

func NewFirestoreClient() (*datastore.FirestoreClient, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./credential/drinking-microcomputer-test-3a6d76de6be6.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	fc := &datastore.FirestoreClient{
		Ctx:    ctx,
		Client: client,
	}

	return fc, nil
}
