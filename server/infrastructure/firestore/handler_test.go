package firestore

import "testing"

func TestNewFirestoreClient(t *testing.T) {
	fh, err := NewFirestoreClient()
	if err != nil {
		t.Error(err)
	}

	t.Log(fh)
}
