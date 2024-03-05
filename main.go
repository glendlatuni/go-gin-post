package main

import(
	"github.com/glendlatuni/go-gin-post/src/utils"
	"github.com/username/go-gin-postgresql-backend/src/middlewares"
	"github.com/glendlatuni/go-gin-post/src/models"
	"github.com/glendlatuni/go-gin-post/src/routes"
	
	
)

func main() {
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8080")
}