package database

import (
	"gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Ocorreu um erro ao conectar com banco de dados -", err)
	}

	errMigrate := DB.AutoMigrate(&models.Student{})
	if errMigrate != nil {
		log.Panic("Ocorreu um erro ao realizar migration -", errMigrate)
	}
}
