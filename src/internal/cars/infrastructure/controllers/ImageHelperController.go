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

func saveCarImage(c *gin.Context, field string, storage storageService.StorageService) (string, error) {
	file, header, err := c.Request.FormFile(field)
	if err == http.ErrMissingFile {
		return "", nil 
	}
	if err != nil {
		return "", fmt.Errorf("error al obtener %s: %w", field, err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("formato no permitido en %s, use jpg, jpeg o png", field)
	}

	newFilename := fmt.Sprintf("cars/%s%s", uuid.New().String(), ext)

	url, err := storage.Upload(file, newFilename)
	if err != nil {
		return "", fmt.Errorf("error al subir %s a Firebase: %w", field, err)
	}

	return url, nil
}