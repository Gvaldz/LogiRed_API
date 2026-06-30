package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"logired/src/core"
)

type StorageService interface {
	Upload(file multipart.File, destination string) (string, error)
	UploadPrivate(file multipart.File, destination string) (string, error)
}

type firebaseStorage struct {
	bucket *storage.BucketHandle
}

func NewFirebaseStorage(credentialsFile, bucketName string) (StorageService, error) {
	app, err := core.GetFirebaseApp(credentialsFile)
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

func (fs *firebaseStorage) UploadPrivate(file multipart.File, destination string) (string, error) {
    ctx := context.Background()
    
    obj := fs.bucket.Object(destination)
    wc := obj.NewWriter(ctx)

    // Copiamos el archivo
    if _, err := io.Copy(wc, file); err != nil { 
        return "", fmt.Errorf("error al copiar archivo privado: %w", err) 
    }
    if err := wc.Close(); err != nil { 
        return "", fmt.Errorf("error al cerrar el writer: %w", err) 
    }


    return "https://storage.googleapis.com/" + fs.bucket.BucketName() + "/" + destination, nil
}