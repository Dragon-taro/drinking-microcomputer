package controller

import (
	"testing"

	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/firestore"
)

func TestCreateDataController(t *testing.T) {
	fh, _ := firestore.NewFirestoreClient()
	dataController := NewDataController(fh)

	t.Log(dataController)
}
