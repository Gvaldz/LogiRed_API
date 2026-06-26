package controllers

import (
	"logired/src/internal/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(delete *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{
		deleteUser: delete,
	}
}
// Delete godoc
// @Summary      Eliminar un usuario
// @Description  Elimina la cuenta (solo admin o el propio usuario)
// @Tags         users
// @Produce      json
// @Param        id path int true "ID del usuario"
// @Security     Bearer
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      403 {object} map[string]string "no autorizado"
// @Failure      404 {object} map[string]string "usuario no encontrado"
// @Router       /users/{id} [delete]
func (h *DeleteUserController) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.deleteUser.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
