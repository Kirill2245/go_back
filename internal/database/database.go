package database

import (
	"log"
	"os"

	"backToDo/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	log.Println("Успешное подключение к БД")

	if err := db.AutoMigrate(model.AllModels()...); err != nil {
		log.Fatal("Не удалось выполнить миграцию:", err)
	}
	log.Println("Миграция выполнена")

	return db
}
