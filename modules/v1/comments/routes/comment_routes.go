package routes

import (
	"mygram/infrastructures/middlewares"
	commentControllerV1 "mygram/modules/v1/comments/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	commentControllerV1 := commentControllerV1.NewCommentController(db)
	mid := middlewares.NewMiddleware(db)

	api := router.Group("/api/v1", mid.Auth())
	{
		comments := api.Group("/comments")
		{
			comments.GET("/:id_photos", commentControllerV1.GetAllComments)
			comments.GET("/id/:id", commentControllerV1.GetCommentById)
			comments.POST("", commentControllerV1.CreateComment)
			comments.PUT("/:id", commentControllerV1.UpdateComment)
			comments.DELETE("/:id", commentControllerV1.DeleteComment)
		}

		photos := api.Group("/photos")
		{
			photos.GET("", commentControllerV1.GetAllPhotos)
			photos.GET("/:id", commentControllerV1.GetPhotoById)
		}
	}
	return router
}
