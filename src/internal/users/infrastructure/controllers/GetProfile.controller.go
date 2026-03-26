package controllers

import (
    "fmt"
    "logired/src/internal/users/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetUserProfileController struct {
    getProfile *application.GetUserProfile
}

func NewGetUserProfileController(uc *application.GetUserProfile) *GetUserProfileController {
    return &GetUserProfileController{getProfile: uc}
}

func (ctrl *GetUserProfileController) GetProfile(c *gin.Context) {
    var idInt int32
    if _, err := fmt.Sscanf(c.Param("id"), "%d", &idInt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    profile, err := ctrl.getProfile.Execute(idInt)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, profile)
}