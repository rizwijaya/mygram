package main

import (
	_ "mygram/docs"
	"mygram/infrastructures/config"
	database "mygram/infrastructures/databases"
	routesCommentsV1 "mygram/modules/v1/comments/routes"
	routesUsersV1 "mygram/modules/v1/users/routes"
	error "mygram/pkg/http-error"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			MyGram API
// @description	This is a sample server for a MyGram.
// @version		1.0.0
// @termsOfService	http://swagger.io/terms/
// @contact.name	Swagger API Team
// @contact.email	admin@rizwijaya.com
// @licence.name	MIT
// @licence.url	http://opensource.org/licenses/MIT
// @host			localhost:8080
// @BasePath		/
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	config := config.New()

	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router = routesUsersV1.NewRouter(router, db)
	router = routesCommentsV1.NewRouter(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(error.NotFound())
	router.NoMethod(error.NoMethod())

	router.Run(":" + config.App.Port)
}
