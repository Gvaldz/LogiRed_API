package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	driver "logired/src/internal/drivers/application"
	"logired/src/internal/users/application"
	"logired/src/internal/users/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserController struct {
	updateUserUC   *application.UpdateUser
	updateDriverUC *driver.UpdateDriverProfile
}

func NewUpdateUserController(
	updateUserUC *application.UpdateUser,
	updateDriverUC *driver.UpdateDriverProfile,
) *UpdateUserController {
	return &UpdateUserController{
		updateUserUC:   updateUserUC,
		updateDriverUC: updateDriverUC, 
	}
}

func (c *UpdateUserController) UpdateUser(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	tokenUserID := userIDInterface.(int32)

	var idInt int32
	if _, err := fmt.Sscanf(ctx.Param("id"), "%d", &idInt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
		if tokenUserID != idInt {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para editar este perfil"})
		return
	}

	contentType := ctx.GetHeader("Content-Type")
	var user entities.User
	var citywork string

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
			if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear directorio"})
				return
			}
			if err := ctx.SaveUploadedFile(header, filepath.Join("./uploads", newFilename)); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen: " + err.Error()})
				return
			}
			user.ImageURL = fmt.Sprintf("https://liveshop.myddns.me/uploads/%s", newFilename)
		}

		user.Name        = ctx.Request.FormValue("name")
		user.Lastname    = ctx.Request.FormValue("lastname")
		user.Email       = ctx.Request.FormValue("email")
		user.NumberPhone = ctx.Request.FormValue("numberphone")
		user.Birthdate   = ctx.Request.FormValue("birthdate")
		citywork         = ctx.Request.FormValue("citywork") 

	} else {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := c.updateUserUC.Execute(idInt, user); err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if citywork != "" {
		if err := c.updateDriverUC.Execute(idInt, citywork); err != nil {
			if strings.Contains(err.Error(), "no encontrado") {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Perfil de conductor no encontrado"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar conductor: " + err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}