package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"logired/src/internal/users/application"
	"logired/src/internal/users/domain/entities"
	storageService "logired/src/core/services/storage"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserUC   *application.UpdateUser
	storage        storageService.StorageService 
}

func NewUpdateUserController(
	updateUserUC *application.UpdateUser,
	storage storageService.StorageService, 
) *UpdateUserController {
	return &UpdateUserController{
		updateUserUC:   updateUserUC,
		storage:        storage, 
	}
}
// UpdateUser godoc
// @Summary      Actualizar perfil de usuario (con imagen opcional)
// @Description  Actualiza los datos del usuario autenticado (solo el propio usuario)
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "ID del usuario"
// @Param        name        formData string false "Nombre"
// @Param        lastname    formData string false "Apellido"
// @Param        email       formData string false "Email"
// @Param        numberphone formData string false "Teléfono"
// @Param        birthdate   formData string false "Fecha de nacimiento"
// @Param        image       formData file   false "Nueva foto de perfil"
// @Security     Bearer
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      403 {object} map[string]string "no autorizado"
// @Failure      404 {object} map[string]string "usuario no encontrado"
// @Router       /users/{id} [put]
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

	if strings.HasPrefix(contentType, "multipart/form-data") {
		if err := ctx.Request.ParseMultipartForm(15 << 20); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
			return
		}

		// Llamada adaptada a Firebase utilizando la nueva lógica
		imageURL, err := saveUserImage(ctx, c.storage)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.ImageURL = imageURL

		user.Name        = ctx.Request.FormValue("name")
		user.Lastname    = ctx.Request.FormValue("lastname")
		user.Email       = ctx.Request.FormValue("email")
		user.NumberPhone = ctx.Request.FormValue("numberphone")
		user.Birthdate   = ctx.Request.FormValue("birthdate")

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

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}