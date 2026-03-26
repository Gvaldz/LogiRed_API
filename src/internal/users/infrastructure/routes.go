package infrastructure

import (
	"logired/src/internal/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	CreateUserController       *controllers.CreateUserController
	GetAllUsersController      *controllers.GetAllUsersController
	GetUserController          *controllers.GetByUserIDController
	GetUserProfileController   *controllers.GetUserProfileController   // ← nuevo
	UpdateUserController       *controllers.UpdateUserController
	UpdatePasswordController   *controllers.UpdatePasswordController
	DeleteUserController       *controllers.DeleteUserController
	AuthMiddleware             gin.HandlerFunc
}

func NewUserRoutes(
	createUserController *controllers.CreateUserController,
	getAllUsersController *controllers.GetAllUsersController,
	getUserController *controllers.GetByUserIDController,
	getUserProfileController *controllers.GetUserProfileController,  
	updateUserController *controllers.UpdateUserController,
	updatePasswordController *controllers.UpdatePasswordController,
	deleteUserController *controllers.DeleteUserController,
	authMiddleware gin.HandlerFunc,
) *UserRoutes {
	return &UserRoutes{
		CreateUserController:     createUserController,
		GetAllUsersController:    getAllUsersController,
		GetUserController:        getUserController,
		GetUserProfileController: getUserProfileController,   
		UpdateUserController:     updateUserController,
		UpdatePasswordController: updatePasswordController,
		DeleteUserController:     deleteUserController,
		AuthMiddleware:           authMiddleware,
	}
}

func (r *UserRoutes) AttachRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", r.CreateUserController.Create)
		userGroup.GET("", r.GetAllUsersController.GetAll)
		userGroup.GET("/:id", r.GetUserController.GetByUserID)
		userGroup.GET("/profile/:id", r.GetUserProfileController.GetProfile)

		protected := userGroup.Group("")
		protected.Use(r.AuthMiddleware)
		{
			protected.PUT("/:id", r.UpdateUserController.UpdateUser)
			protected.PUT("/password/:id", r.UpdatePasswordController.UpdatePassword)
			protected.DELETE("/:id", r.DeleteUserController.Delete)
		}
	}
}