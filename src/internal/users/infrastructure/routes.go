package infrastructure

import (
	"logired/src/internal/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	CreateUserController      *controllers.CreateUserController
	GetAllUsersController     *controllers.GetAllUsersController
	GetUserController         *controllers.GetByUserIDController
	GetUserProfileController  *controllers.GetUserProfileController
	GetMyProfileController    *controllers.GetMyProfileController
	UpdateUserController      *controllers.UpdateUserController
	UpdatePasswordController  *controllers.UpdatePasswordController
	RessetPasswordController  *controllers.RessetPasswordController
	DeleteUserController      *controllers.DeleteUserController
	AuthMiddleware            gin.HandlerFunc
}

func NewUserRoutes(
	createUserController    *controllers.CreateUserController,
	getAllUsersController    *controllers.GetAllUsersController,
	getUserController       *controllers.GetByUserIDController,
	getUserProfileController *controllers.GetUserProfileController,
	getMyProfileController  *controllers.GetMyProfileController,
	updateUserController    *controllers.UpdateUserController,
	updatePasswordController *controllers.UpdatePasswordController,
	ressetPasswordController *controllers.RessetPasswordController,
	deleteUserController    *controllers.DeleteUserController,
	authMiddleware          gin.HandlerFunc,
) *UserRoutes {
	return &UserRoutes{
		CreateUserController:     createUserController,
		GetAllUsersController:    getAllUsersController,
		GetUserController:        getUserController,
		GetUserProfileController: getUserProfileController,
		GetMyProfileController:   getMyProfileController,
		UpdateUserController:     updateUserController,
		UpdatePasswordController: updatePasswordController,
		RessetPasswordController: ressetPasswordController,
		DeleteUserController:     deleteUserController,
		AuthMiddleware:           authMiddleware,
	}
}

func (r *UserRoutes) AttachRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", r.CreateUserController.Create)
		userGroup.PUT("/password-reset", r.RessetPasswordController.RessetPassword)

		protected := userGroup.Group("")
		protected.Use(r.AuthMiddleware)
		{
			protected.GET("", r.GetAllUsersController.GetAll)
			protected.GET("/me", r.GetMyProfileController.GetMyProfile)
			protected.GET("/:id", r.GetUserController.GetByUserID)
			protected.GET("/profile-driver/:id", r.GetUserProfileController.GetProfile)
			protected.PUT("/:id", r.UpdateUserController.UpdateUser)
			protected.PUT("/update-password", r.UpdatePasswordController.UpdatePassword)
			protected.DELETE("/:id", r.DeleteUserController.Delete)
		}
	}
}
