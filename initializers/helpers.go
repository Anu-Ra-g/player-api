package initializers

import (
	"fmt"
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
	fmt.Println(`  ___  _____  _  _  _  _  ____   ___  ____  ____  ____     ____  _____    ____     __    ____    __    ____    __    ___  ____ 
	/ __)(  _  )( \( )( \( )( ___) / __)(_  _)( ___)(  _ \   (_  _)(  _  )  (  _ \   /__\  (_  _)  /__\  (  _ \  /__\  / __)( ___)
   ( (__  )(_)(  )  (  )  (  )__) ( (__   )(   )__)  )(_) )    )(   )(_)(    )(_) ) /(__)\   )(   /(__)\  ) _ < /(__)\ \__ \ )__) 
	\___)(_____)(_)\_)(_)\_)(____) \___) (__) (____)(____/    (__) (_____)  (____/ (__)(__) (__) (__)(__)(____/(__)(__)(___/(____)
   `)
}

