package controller

import (
	"testing"

	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/firestore"
)

func TestCreateDataController(t *testing.T) {
	fc, _ := firestore.NewFirestoreClient()
	dataController := NewDataController(fc)

	t.Log(dataController)
}
