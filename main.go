package main

import(
	
)

func main() {
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Run(":8080")
}