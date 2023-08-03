package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"player_backend/models"
)


func LoadEnvVariables(){
	if err := godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}
}

var DB *gorm.DB 

func ConnectToDB(){
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database")
	}
}

func Migrate(){
	DB.AutoMigrate(&models.Player{})
}

