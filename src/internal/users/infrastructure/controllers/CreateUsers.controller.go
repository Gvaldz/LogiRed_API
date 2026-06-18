package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"logired/src/internal/users/application"
	userEntities "logired/src/internal/users/domain/entities"
	storageService "logired/src/core/services/storage"
	"github.com/gin-gonic/gin"
)

const (
	UserTypeCliente   = 1
	UserTypeConductor = 2
)

type CreateUserController struct {
	createUser     *application.CreateUser      
	registerDriver *application.RegisterDriver 
	storage        storageService.StorageService 
}

func NewCreateUserController(
	createUser *application.CreateUser,
	registerDriver *application.RegisterDriver,
	storage storageService.StorageService, 
) *CreateUserController {
	return &CreateUserController{
		createUser:     createUser,
		registerDriver: registerDriver,
		storage:        storage, 
	}
}
// Create godoc
// @Summary      Registrar un nuevo usuario
// @Description  Crea una cuenta (cliente o conductor) con imagen opcional
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
// @Param        name         formData string true  "Nombre"
// @Param        lastname     formData string true  "Apellido"
// @Param        email        formData string true  "Correo electrónico"
// @Param        numberphone  formData string true  "Teléfono"
// @Param        birthdate    formData string true  "Fecha de nacimiento (YYYY-MM-DD)"
// @Param        password     formData string true  "Contraseña"
// @Param        user_type    formData int    true  "1=cliente, 2=conductor"
// @Param        image        formData file   false "Foto de perfil (jpg, png, gif)"
// @Success      201 {object} map[string]interface{} "usuario creado"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /users [post]
func (ctrl *CreateUserController) Create(c *gin.Context) {
	if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser multipart/form-data"})
		return
	}

	if err := c.Request.ParseMultipartForm(15 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
		return
	}

	name        := c.Request.FormValue("name")
	lastname    := c.Request.FormValue("lastname")
	email       := c.Request.FormValue("email")
	phone       := c.Request.FormValue("numberphone")
	birthdate   := c.Request.FormValue("birthdate")
	password    := c.Request.FormValue("password")
	userTypeStr := c.Request.FormValue("user_type")

	if name == "" || lastname == "" || email == "" || phone == "" || birthdate == "" || password == "" || userTypeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos obligatorios"})
		return
	}

	userType, err := strconv.Atoi(userTypeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuario inválido"})
		return
	}

	// Llamada adaptada a Firebase
	imageURL, err := saveUserImage(c, ctrl.storage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userEntities.User{
		Name:        name,
		Lastname:    lastname,
		Email:       email,
		NumberPhone: phone,
		Birthdate:   birthdate,
		Password:    password,
		UserType:    userType,
		ImageURL:    imageURL,
	}

	switch userType {
	case UserTypeCliente:
		created, err := ctrl.createUser.Execute(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Cliente registrado correctamente", "user": created})

	case UserTypeConductor:
		created, err := ctrl.registerDriver.Execute(userEntities.RegisterDriverInput{
			User:     user,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Conductor registrado correctamente", "user": created})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuario no reconocido"})
	}
}