package firestore

import (
	"os"
	"testing"
)

func TestNewFirestoreClient(t *testing.T) {
	os.Setenv("JSON_PATH", "../../credential/firestore.json")
	fh, err := NewFirestoreClient()
	if err != nil {
		t.Error(err)
	}

	t.Log(fh)
}
