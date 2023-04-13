package routes

import (
	userControllerV1 "mygram/modules/v1/users/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userControllerV1 := userControllerV1.NewUserController(db)
	api := router.Group("/api/v1")
	users := api.Group("/users")
	{
		users.POST("/register", userControllerV1.Register)
		users.POST("/login", userControllerV1.Login)
	}

	// //Social media
	// api := router.Group("/api/v1")
	// {

	// }
	return router
}
