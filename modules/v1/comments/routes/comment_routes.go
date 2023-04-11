package routes

import (
	//commentControllerV1 "mygram/modules/v1/comments/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// commentControllerV1 := commentControllerV1.NewCommentController(db)
	// api := router.Group("/api/v1")
	// {
	// 	comments := api.Group("/comments")
	// 	{

	// 	}
	// }

	return router
}
