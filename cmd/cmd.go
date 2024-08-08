package cmd

import (
	"cloud/app/service/auth/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return nil
}

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", os.Getenv("db_host"), os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_name"), os.Getenv("db_port"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrateErr := db.AutoMigrate(&models.User{})
	if migrateErr != nil {
		panic(migrateErr)
	}
	return db, nil
}

func Test() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("db_host"), os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_name"), os.Getenv("db_port"))
	return dsn
}
