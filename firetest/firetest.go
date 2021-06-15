package firetest

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// StoreClient tries to create firestore client for testing.
// It also delete the documents in the collection after the tests.
func StoreClient(t *testing.T, projectID, collection string) (
	client *firestore.Client,
) {
	t.Helper()

	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		t.Fatalf("error creating firestore client: %v", err)
	}

	t.Cleanup(func() {
		ctx := context.Background()
		iter := client.Collection(collection).DocumentRefs(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				t.Fatalf("error iterating to next document: %v", err)
			}
			_, err = doc.Delete(ctx)
			if err != nil {
				t.Fatalf("error deleting document: %v", err)
			}
		}

		client.Close()
	})

	return client
}
