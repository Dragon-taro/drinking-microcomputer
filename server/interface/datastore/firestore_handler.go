package datastore

import (
	"context"

	firestore "cloud.google.com/go/firestore"
)

type FirestoreClient struct {
	// この層にしては具体的なのを持ちすぎ？
	// でも、抽象化の方法がいまいちわからん。。
	Client *firestore.Client
	Ctx    context.Context
}
