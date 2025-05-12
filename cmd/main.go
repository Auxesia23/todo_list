package main

import (
	"log"
	"os"

	"github.com/Auxesia23/todo_list/internal/database"
	"github.com/Auxesia23/todo_list/internal/repository"
	"github.com/Auxesia23/todo_list/internal/utils"
)

func main() {
	
	db := database.InitDB()
	UserRepo := repository.NewUserRepository(db)
	TodoRepo := repository.NewTodoRepository(db)
	LogsRepo := repository.NewLogsRepository(db)
	
	utils.SetupGoogleOAuth()
	
	cfg := config{
			addr: os.Getenv("PORT"),
		}
		
	app := &application{
		Config : cfg,
		User: UserRepo,
		Todo: TodoRepo,
		Logs: LogsRepo,
	}
	

	
	mux := app.mount()
	log.Fatal(app.run(mux))
}	