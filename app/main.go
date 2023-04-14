package main

import (
	"mygram/infrastructures/config"
	database "mygram/infrastructures/databases"
	routesCommentsV1 "mygram/modules/v1/comments/routes"
	routesUsersV1 "mygram/modules/v1/users/routes"
	error "mygram/pkg/http-error"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//	@title			Book Library API
//	@description	This is a sample server for a book library.
//	@version		1.0.0
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Swagger API Team
//	@contact.email	admin@rizwijaya.com
//	@licence.name	MIT
//	@licence.url	http://opensource.org/licenses/MIT
//	@host			localhost:8080
//	@BasePath		/

func main() {
	config := config.New()

	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router = routesUsersV1.NewRouter(router, db)
	router = routesCommentsV1.NewRouter(router, db)

	router.NoRoute(error.NotFound())
	router.NoMethod(error.NoMethod())

	router.Run(":" + config.App.Port)
}
