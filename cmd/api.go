package main

import (
	"log"
	"net/http"

	"github.com/Auxesia23/todo_list/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type application struct{
	Config config
	User repository.UserRepository
	Todo repository.TodoRepository
	Logs repository.LogsRepository
}

type config struct {
	addr string
}

func (app *application) mount () http.Handler {
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	
	r.Use(app.LoggingMiddleware)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, 
		}))
	
	r.Route("/v1", func(r chi.Router){
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
		
		r.Route("/auth", func(r chi.Router){
			r.Post("/register", app.RegisterUser)
			r.Post("/login", app.UserLogin)
		})
		
		r.Route("/user", func(r chi.Router){
			r.Use(UserAuth)
			r.Get("/info", app.UserInfo)
		})
		
		r.Route("/todo",  func(r chi.Router) {
			r.Use(UserAuth)
			r.Post("/", app.CreateTodo)
			r.Get("/", app.GetAllTodos)
			r.Put("/{id}", app.UpdateTodo)
			r.Delete("/{id}", app.DeleteTodo)
		})
		
		r.Route("/logs", func(r chi.Router){
			r.Use(AdminAuth)
			r.Get("/", app.GetAllLogs)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
			Addr:    app.Config.addr,
			Handler: mux,
		}

		log.Println("Server running on port" + app.Config.addr)

		return srv.ListenAndServe()
}