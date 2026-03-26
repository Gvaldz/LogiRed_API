package controllers

import (
    "net/http"
    "os"
    "fmt"
    "path/filepath"
    "strings"
    "strconv"

    "logired/src/internal/users/application"
    userEntities "logired/src/internal/users/domain/entities"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

const (
    UserTypeCliente   = 1
    UserTypeConductor = 2
)

type CreateUserController struct {
    createUser     *application.CreateUser      
    registerDriver *application.RegisterDriver 
}

func NewCreateUserController(
    createUser *application.CreateUser,
    registerDriver *application.RegisterDriver,
) *CreateUserController {
    return &CreateUserController{
        createUser:     createUser,
        registerDriver: registerDriver,
    }
}

func (ctrl *CreateUserController) Create(c *gin.Context) {
    if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser multipart/form-data"})
        return
    }

    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
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

    var imageURL string
    file, header, err := c.Request.FormFile("image")
    if err != nil && err != http.ErrMissingFile {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener imagen: " + err.Error()})
        return
    }
    if file != nil {
        defer file.Close()
        ext := strings.ToLower(filepath.Ext(header.Filename))
        if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de imagen no permitido"})
            return
        }
        newFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
        if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear directorio"})
            return
        }
        if err := c.SaveUploadedFile(header, filepath.Join("./uploads", newFilename)); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar imagen"})
            return
        }
        imageURL = fmt.Sprintf("https://liveshop.myddns.me/uploads/%s", newFilename)
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
        citywork := c.Request.FormValue("citywork")
        if citywork == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "El campo citywork es obligatorio para conductores"})
            return
        }
        created, err := ctrl.registerDriver.Execute(userEntities.RegisterDriverInput{
            User:     user,
            Citywork: citywork,
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