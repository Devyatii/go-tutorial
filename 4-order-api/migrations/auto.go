package main

import (
	"4-order-api/intrenal/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Данный метод нужен для настройки миграции БД при изменении полей
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbErr := db.AutoMigrate(&product.Product{})
	if dbErr != nil {
		return
	}
}
