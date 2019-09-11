package controller

import (
	"os"
	"testing"

	"github.com/Dragon-taro/drinking-microcomputer/server/infrastructure/firestore"
)

func TestCreateDataController(t *testing.T) {
	os.Setenv("JSON_PATH", "../../credential/firestore.json")
	fc, _ := firestore.NewFirestoreClient()
	dataController := NewDataController(fc)

	t.Log(dataController)
}
