package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// build the DSN from .env
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )
	dsn := os.Getenv("DB_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&LandlordDetails{}, &Listing{}, &Utilities{}, &Features{}, &Location{}, &PropertyMedia{})

	log.Println("Database migrated successfully!")
}
