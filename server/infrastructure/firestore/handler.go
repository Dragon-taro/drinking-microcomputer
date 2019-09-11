package firestore

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"github.com/Dragon-taro/drinking-microcomputer/server/interface/datastore"
	"google.golang.org/api/option"
)

func NewFirestoreClient() (*datastore.FirestoreClient, error) {
	ctx := context.Background()

	jsonPath := os.Getenv("JSON_PATH")
	sa := option.WithCredentialsFile(jsonPath)
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
