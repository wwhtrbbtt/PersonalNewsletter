package main

// https://pkg.go.dev/cloud.google.com/go/firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Firestore struct {
	Client     *firestore.Client
	Collection string
	Ctx        *context.Context
}

func (fs *Firestore) Connect(keyFilePath, projectID string) error {
	ctx := context.Background()
	var options []option.ClientOption
	options = append(options, option.WithCredentialsFile(keyFilePath))

	client, err := firestore.NewClient(ctx, projectID, options...)
	if err != nil {
		return err
	}
	fs.Client = client
	fs.Ctx = &ctx
	return nil
}

func (fs *Firestore) GetDoc(id string) (map[string]interface{}, error) {
	col := fs.Client.Collection(fs.Collection)
	doc := col.Doc(id)

	docsnap, err := doc.Get(*fs.Ctx)

	if err != nil {
		return make(map[string]interface{}), nil
	}
	dataMap := docsnap.Data()
	return dataMap, nil
}

func (fs *Firestore) SetDoc(id string, data interface{}) error {
	col := fs.Client.Collection(fs.Collection)
	doc := col.Doc(id)
	_, err := doc.Set(*fs.Ctx, data)
	return err
}

func (fs *Firestore) WhereDoc(q1, q2, q3 string) ([]map[string]interface{}, error) {
	// Supported operators include '<', '<=', '>', '>=', '==', 'in', 'array-contains', and 'array-contains-any'.
	// eg. "foo", "!=", "bar" (returns all)
	col := fs.Client.Collection(fs.Collection)
	q := col.Where(q1, q2, q3)

	iter := q.Documents(*fs.Ctx)
	defer iter.Stop()
	var docs []map[string]interface{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return docs, err
		}
		docs = append(docs, doc.Data())
	}
	return docs, nil
}
