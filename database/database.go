package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DoomGuy1818/gofiber-test/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getEnvVariables() (string, string, string, string, string){
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
	}
	return os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT")
}

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	Db_host, Db_username, Db_password, Db_name, Db_port := getEnvVariables()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", Db_host, Db_username, Db_password, Db_name, Db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("connected Successully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{
		Db: db,
	}
}
