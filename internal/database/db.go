package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Auxesia23/todo_list/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	sslmode := "disable"


	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}


	err = db.AutoMigrate(&models.User{},models.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}