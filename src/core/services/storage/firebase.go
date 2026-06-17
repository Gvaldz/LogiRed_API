package storage

import (
	"context"
	"io"
	"mime/multipart"
	"cloud.google.com/go/storage"
	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type StorageService interface {
	Upload(file multipart.File, destination string) (string, error)
}

type firebaseStorage struct {
	bucket *storage.BucketHandle
}

func NewFirebaseStorage(credentialsFile, bucketName string) (StorageService, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil { return nil, err }
	
	client, err := app.Storage(context.Background())
	if err != nil { return nil, err }
	
	bucket, err := client.Bucket(bucketName)
	return &firebaseStorage{bucket: bucket}, err
}

func (fs *firebaseStorage) Upload(file multipart.File, destination string) (string, error) {
	wc := fs.bucket.Object(destination).NewWriter(context.Background())
	if _, err := io.Copy(wc, file); err != nil { return "", err }
	if err := wc.Close(); err != nil { return "", err }
	return "https://storage.googleapis.com/TU_BUCKET/" + destination, nil
}