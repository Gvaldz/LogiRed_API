package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"logired/src/internal/users/application"
	"logired/src/internal/users/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserController struct {
	updateUserUC *application.UpdateUser
}

func NewUpdateUserController(updateUserUC *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUserUC: updateUserUC}
}

func (c *UpdateUserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var idInt int32
	if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	contentType := ctx.GetHeader("Content-Type")

	var user entities.User

	if strings.HasPrefix(contentType, "multipart/form-data") {
		if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
			return
		}

		file, header, err := ctx.Request.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener la imagen: " + err.Error()})
			return
		}

		if file != nil {
			defer file.Close()

			ext := strings.ToLower(filepath.Ext(header.Filename))
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato no permitido. Use jpg, jpeg, png o gif"})
				return
			}

			newFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
			uploadDir := "./uploads"
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear directorio"})
				return
			}

			filePath := filepath.Join(uploadDir, newFilename)
			if err := ctx.SaveUploadedFile(header, filePath); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen: " + err.Error()})
				return
			}

			user.ImageURL = fmt.Sprintf("https://liveshop.myddns.me/uploads/%s", newFilename)
		}

		user.Name        = ctx.Request.FormValue("name")
		user.Lastname    = ctx.Request.FormValue("lastname")
		user.Email       = ctx.Request.FormValue("email")
		user.NumberPhone = ctx.Request.FormValue("phone_number")
		user.Birthdate   = ctx.Request.FormValue("birthdate")

	} else {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := c.updateUserUC.Execute(idInt, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}