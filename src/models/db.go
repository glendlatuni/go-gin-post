package models

import (
	"fmt"
	
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Database  *gorm.DB

func OpenDatabaseConnection() {
	var err error

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	Database , err = grom.Open(postgres.Open(dsn), &gorm.Config{})
	
	if eff != nil {
		panic(err)
	} else {
		fmt.Println("Connection Opened to Database")
	}
}


func AutoMigrateModels(){
	DB.AutoMigrate(&Stratup{})
}