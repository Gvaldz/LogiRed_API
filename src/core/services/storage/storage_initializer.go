package storage

import (
	"log"
	"os"
)

func InitStorage() StorageService {
	credsFile := os.Getenv("FCM_CREDENTIALS_FILE")
	bucketName := os.Getenv("FIREBASE_BUCKET_NAME") 

	if credsFile == "" || bucketName == "" {
		log.Fatal("ERROR: FCM_CREDENTIALS_FILE o FIREBASE_BUCKET_NAME no están configuradas en el archivo .env")
	}

	storageService, err := NewFirebaseStorage(credsFile, bucketName)
	if err != nil {
		log.Fatalf("Error crítico al inicializar Firebase Storage: %v", err)
	}

	log.Println("¡Firebase Storage inicializado correctamente!")
	return storageService
}