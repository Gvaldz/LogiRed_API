package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func saveCarImage(c *gin.Context, field string) (string, error) {
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

	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	uploadDir := "./uploads/cars"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error al crear directorio: %w", err)
	}

	if err := c.SaveUploadedFile(header, filepath.Join(uploadDir, newFilename)); err != nil {
		return "", fmt.Errorf("error al guardar %s: %w", field, err)
	}

	return fmt.Sprintf("https://liveshop.myddns.me/uploads/cars/%s", newFilename), nil
}