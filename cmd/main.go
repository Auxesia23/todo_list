package main

import (
	"log"
	"os"

	"github.com/Auxesia23/todo_list/internal/database"
	"github.com/Auxesia23/todo_list/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db := database.InitDB()
	UserRepo := repository.NewUserRepository(db)
	
	cfg := config{
			addr: os.Getenv("PORT"),
		}
		
	app := &application{
		Config : cfg,
		User: UserRepo,
	}
	
	mux := app.mount()
	log.Fatal(app.run(mux))
}	