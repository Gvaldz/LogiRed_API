package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	storageService "logired/src/core/services/storage"
)

func saveUserImage(c *gin.Context, storage storageService.StorageService) (string, error) {
	file, header, err := c.Request.FormFile("image")
	if err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("error al obtener imagen: %w", err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return "", fmt.Errorf("formato de imagen no permitido")
	}

	// Destino ordenado en Firebase
	newFilename := fmt.Sprintf("users/%s%s", uuid.New().String(), ext)

	url, err := storage.Upload(file, newFilename)
	if err != nil {
		return "", fmt.Errorf("error al subir a Firebase: %w", err)
	}

	return url, nil
}