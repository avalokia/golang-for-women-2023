package database

import (
	"fmt"
	"gorm/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
		return err
	}

	host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed connecting DB:", err)
		return err
	}

	// Auto migration
	err = db.Debug().AutoMigrate(models.Product{}, models.Variants{})
	if err != nil {
		log.Fatalln("Failed to auto migrate:", err)
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		StartDB()
	}
	return db
}
