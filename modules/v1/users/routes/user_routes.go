package routes

import (
	userControllerV1 "mygram/modules/v1/users/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userControllerV1 := userControllerV1.NewUserController(db)
	users := router.Group("")
	{
		users.POST("/register", userControllerV1.Register)
	}

	// //Social media
	// api := router.Group("/api/v1")
	// {

	// }
	return router
}
