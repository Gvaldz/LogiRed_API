package controllers

import (
	"fmt"
	"logired/src/core"
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

	savedPath := filepath.Join(uploadDir, newFilename)

	if err := c.SaveUploadedFile(header, savedPath); err != nil {
		return "", fmt.Errorf("error al guardar %s: %w", field, err)
	}

	if err := core.FixImageOrientation(savedPath); err != nil {
		fmt.Printf("advertencia: no se pudo corregir orientación de %s: %v\n", newFilename, err)
	}

	if err := os.Chmod(savedPath, 0644); err != nil {
		fmt.Printf("advertencia: no se pudo cambiar permisos: %v\n", err)

	}
	return fmt.Sprintf("https://logiredapi.redirectme.net/uploads/cars/%s", newFilename), nil
}