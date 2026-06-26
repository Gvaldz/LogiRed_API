package controllers

import (
	"logired/src/internal/rides/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetRidesByClientController struct {
	getRidesByClient *application.GetRidesByClient
}

func NewGetRidesByClientController(get *application.GetRidesByClient) *GetRidesByClientController {
	return &GetRidesByClientController{getRidesByClient: get}
}

// GetByClient godoc
// @Summary      Obtener viajes del cliente autenticado
// @Description  Devuelve todos los viajes solicitados por el cliente
// @Tags         rides
// @Produce      json
// @Security     Bearer
// @Success      200 {object} map[string]interface{} "lista de viajes"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/client [get]
func (ctrl *GetRidesByClientController) GetByClient(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	clientID := userIDInterface.(int32)

	rides, err := ctrl.getRidesByClient.Execute(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rides": rides})
}