package controllers

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
    "logired/src/core/services/storage"
    "logired/src/internal/users/application"
    userEntities "logired/src/internal/users/domain/entities"
    carEntities "logired/src/internal/cars/domain/entities"
    documentEntities "logired/src/internal/documents/domain/entities"
)

const (
    UserTypeCliente   = 1
    UserTypeConductor = 2
)

type CreateUserController struct {
    createUser     *application.CreateUser      
    registerDriver *application.RegisterDriver 
    storage        storage.StorageService 
}

func NewCreateUserController(
    createUser *application.CreateUser,
    registerDriver *application.RegisterDriver,
    storage storage.StorageService, 
) *CreateUserController {
    return &CreateUserController{
        createUser:     createUser,
        registerDriver: registerDriver,
        storage:        storage, 
    }
}

func uploadHelper(c *gin.Context, fieldName string, storage storage.StorageService, folder string, isPrivate bool) (string, error) {
    file, header, err := c.Request.FormFile(fieldName)
    if err != nil {
        return "", fmt.Errorf("error obteniendo el archivo %s: %w", fieldName, err)
    }
    defer file.Close()

    timestamp := time.Now().Unix()
    destination := fmt.Sprintf("%s/%d_%s", folder, timestamp, header.Filename)

    var url string
    const maxRetries = 3
    for attempt := 1; attempt <= maxRetries; attempt++ {
        if isPrivate {
            url, err = storage.UploadPrivate(file, destination)
        } else {
            url, err = storage.Upload(file, destination)
        }
        if err == nil {
            return url, nil
        }

        if attempt < maxRetries {
            time.Sleep(time.Duration(attempt*100) * time.Millisecond)
            file.Seek(0, 0)
        }
    }

    return "", fmt.Errorf("error subiendo %s después de %d intentos: %w", fieldName, maxRetries, err)
}

// Create godoc
// @Summary      Registrar un nuevo usuario (Cliente o Conductor)
// @Description  Crea una cuenta. Si user_type=1 (Cliente), solo requiere datos básicos. Si user_type=2 (Conductor), requiere además datos del vehículo y 2 documentos obligatorios (Identificación Oficial y Licencia).
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
//
// @Param        user_type               formData int    true  "Tipo de usuario: 1=Cliente, 2=Conductor"
// @Param        name                    formData string true  "Nombre"
// @Param        lastname                formData string true  "Apellido"
// @Param        email                   formData string true  "Correo electrónico"
// @Param        numberphone             formData string false "Teléfono"
// @Param        birthdate               formData string false "Fecha de nacimiento (YYYY-MM-DD)"
// @Param        password                formData string true  "Contraseña"
// @Param        image                   formData file   false "Foto de perfil (Opcional)"
//
// @Param        car_registration        formData string false "Placa del vehículo (Requerido para conductor)"
// @Param        brand                   formData string false "Marca del vehículo (Requerido para conductor)"
// @Param        model                   formData string false "Modelo del vehículo (Requerido para conductor)"
// @Param        color                   formData string false "Color del vehículo (Requerido para conductor)"
// @Param        max_capacity            formData int    false "Capacidad máxima (Requerido para conductor)"
//
// @Param        frontview_image         formData file   false "Imagen frontal del carro (Requerido para conductor)"
// @Param        backview_image          formData file   false "Imagen trasera del carro (Requerido para conductor)"
// @Param        leftview_image          formData file   false "Imagen lado izquierdo (Requerido para conductor)"
// @Param        rightview_image         formData file   false "Imagen lado derecho (Requerido para conductor)"
// @Param        space_image             formData file   false "Imagen del espacio de carga (Requerido para conductor)"
// @Param        plates_image            formData file   false "Imagen de las placas (Requerido para conductor)"
//
// @Param        document_identificacion_adelante formData file false "Identificación oficial (frente)"
// @Param        document_identificacion_atras    formData file false "Identificación oficial (reverso)"
// @Param        document_licencia                formData file false "Licencia de conducir"
// @Param        document_comprobante_domicilio   formData file false "Comprobante de domicilio"
//
// @Success      201 {object} map[string]interface{} "Usuario creado exitosamente"
// @Failure      400 {object} map[string]interface{} "Error en los datos (Faltan campos, archivos incorrectos, etc.)"
// @Failure      500 {object} map[string]interface{} "Error interno del servidor (BD o Storage)"
// @Router       /users [post]
func (ctrl *CreateUserController) Create(c *gin.Context) {
    if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser multipart/form-data"})
        return
    }

    if err := c.Request.ParseMultipartForm(30 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
        return
    }

    userTypeStr := c.Request.FormValue("user_type")
    userType, err := strconv.Atoi(userTypeStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuario inválido"})
        return
    }

    name := c.Request.FormValue("name")
    lastname := c.Request.FormValue("lastname")
    email := c.Request.FormValue("email")
    phone := c.Request.FormValue("numberphone")
    birthdate := c.Request.FormValue("birthdate")
    password := c.Request.FormValue("password")

    if name == "" || lastname == "" || email == "" || password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos obligatorios del usuario"})
        return
    }

    imageURL, _ := uploadHelper(c, "image", ctrl.storage, "profiles", false)

    user := userEntities.User{
        Name: name, Lastname: lastname, Email: email, NumberPhone: phone,
        Birthdate: birthdate, Password: password, UserType: userType, ImageURL: imageURL,
    }

    switch userType {
    case UserTypeCliente:
        created, err := ctrl.createUser.Execute(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"message": "Cliente registrado", "user": created})

    case UserTypeConductor:
        carRegistration := c.Request.FormValue("car_registration")
        brand := c.Request.FormValue("brand")
        model := c.Request.FormValue("model")
        color := c.Request.FormValue("color")
        maxCapacityStr := c.Request.FormValue("max_capacity")

        if carRegistration == "" || brand == "" || maxCapacityStr == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos obligatorios del carro"})
            return
        }
        maxCapacity, _ := strconv.ParseInt(maxCapacityStr, 10, 32)

        var wg sync.WaitGroup
        var mu sync.Mutex
        results := make(map[string]string)
        var errUpload error

        uploadFiles := []struct {
            field     string
            folder    string
            name      string
            isPrivate bool
        }{
            {"frontview_image", "cars", "frontView", false},
            {"backview_image", "cars", "backView", false},
            {"leftview_image", "cars", "leftView", false},
            {"rightview_image", "cars", "rightView", false},
            {"space_image", "cars", "spaces", false},
            {"plates_image", "cars", "plates", false},
            {"document_identificacion_adelante", "documents", "docIdAdelante", true},
            {"document_identificacion_atras", "documents", "docIdAtras", true},
            {"document_licencia", "documents", "docLicencia", true},
            {"document_comprobante_domicilio", "documents", "docComprobante", true},
        }

        semaphore := make(chan struct{}, 3)
        for _, f := range uploadFiles {
            wg.Add(1)
            go func(field, folder, name string, isPrivate bool) {
                defer wg.Done()
                semaphore <- struct{}{}
                defer func() { <-semaphore }()

                url, err := uploadHelper(c, field, ctrl.storage, folder, isPrivate)
                mu.Lock()
                defer mu.Unlock()
                if err != nil && errUpload == nil {
                    errUpload = fmt.Errorf("%s: %v", field, err)
                } else {
                    results[name] = url
                }
            }(f.field, f.folder, f.name, f.isPrivate)
        }
        wg.Wait()

        if errUpload != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": errUpload.Error()})
            return
        }

        car := carEntities.Car{
            CarRegistration: carRegistration, Brand: brand, Model: model, Color: color,
            MaxCapacity: int32(maxCapacity), FrontViewImage: results["frontView"],
            BackViewImage: results["backView"], LeftViewImage: results["leftView"],
            RightViewImage: results["rightView"], PlatesImage: results["plates"],
            SpacesImage: results["spaces"],
        }

        documents := []documentEntities.Document{
            {IdDocumentType: 1, URL: results["docIdAdelante"]},
            {IdDocumentType: 2, URL: results["docIdAtras"]},
            {IdDocumentType: 3, URL: results["docLicencia"]},
            {IdDocumentType: 4, URL: results["docComprobante"]},
        }

        created, err := ctrl.registerDriver.Execute(application.RegisterDriverInput{
            User:      user,
            Car:       car,
            Documents: documents,
        })
        if err != nil {
     
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message":  "Conductor registrado. Pendiente de aprobación por un administrador.",
            "approved": false,
            "user":     created,
        })

    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuario no reconocido"})
    }
}