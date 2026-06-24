package storage

import (
	"context"
	"fmt"
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
    ctx := context.Background()
    
    obj := fs.bucket.Object(destination)
    wc := obj.NewWriter(ctx)

    // 2. Copiamos el archivo
    if _, err := io.Copy(wc, file); err != nil { 
        return "", fmt.Errorf("error al copiar archivo: %w", err) 
    }
    if err := wc.Close(); err != nil { 
        return "", fmt.Errorf("error al cerrar el writer: %w", err) 
    }

    if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return "", fmt.Errorf("error al hacer público el archivo: %w", err)
    }

    return "https://storage.googleapis.com/" + fs.bucket.BucketName() + "/" + destination, nil
}