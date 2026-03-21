package controllers

import (
	"net/http"

	"logired/src/internal/users/application"
	"logired/src/internal/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUser *application.CreateUser
}

func NewCreateUserController(createUser *application.CreateUser) *CreateUserController {
	return &CreateUserController{createUser: createUser}
}

func (h *CreateUserController) CreateUser(c *gin.Context) {
	var userRequest entities.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos: " + err.Error()})
		return
	}

	createdUser, err := h.createUser.Execute(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado correctamente",
		"user": gin.H{
			"id":       createdUser.IdUser,
			"name":     createdUser.Name,
			"lastname": createdUser.Lastname,
			"email":    createdUser.Email,
			"user_type": createdUser.UserType,
		},
	})
}
