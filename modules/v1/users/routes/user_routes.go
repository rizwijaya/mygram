package routes

import (
	"mygram/infrastructures/middlewares"
	userControllerV1 "mygram/modules/v1/users/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userControllerV1 := userControllerV1.NewUserController(db)
	mid := middlewares.NewMiddleware(db)

	//User
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userControllerV1.Register)
			users.POST("/login", userControllerV1.Login)
		}

		//Social media
		social := api.Group("/media", mid.Auth())
		{
			social.GET("", userControllerV1.GetAllSocialMedia)
			social.GET("/:id", userControllerV1.GetOneSocialMedia)
			social.POST("", userControllerV1.CreateSocialMedia)
			social.PUT("/:id", userControllerV1.UpdateSocialMedia)
			social.DELETE("/:id", userControllerV1.DeleteSocialMedia)
		}
	}
	return router
}
