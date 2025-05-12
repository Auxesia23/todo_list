package database

import (
	"log"
	"os"

	"github.com/Auxesia23/todo_list/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//Load environment variabel untuk db
	dsn := os.Getenv("DB_DSN")

	//membuka koneksi ke db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//menerapkan migrate models ke db
	err = db.AutoMigrate(&models.User{},models.Todo{}, models.LogEntry{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}