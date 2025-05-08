package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Auxesia23/todo_list/internal/models"
)


func (app *application) CreateTodo (w http.ResponseWriter, r *http.Request){
	//ambil email user yang melakukan request dari context yang dikirim middleware
	userEmail := r.Context().Value("userEmail").(string)
	
	//Decode request body dan bind ke variabel
	var todoInput models.TodoInput
	err := json.NewDecoder(r.Body).Decode(&todoInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//parsing string yang dikirim dari input menjadi date
	parsedDate,err := time.Parse("2006-01-02", todoInput.DueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//merubah input menjadi format yang diterima fungsi create
	todo := models.Todo{
		Title: todoInput.Title,
		Description: todoInput.Description,
		DueDate: parsedDate,
		UserEmail: userEmail,
	}
	
	//membuat todo baru
	todoResponse, err := app.Todo.Create(context.Background(), todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todoResponse)
}

func (app *application) GetAllTodos(w http.ResponseWriter, r *http.Request){
	//ambil email user yang melakukan request dari context yang dikirim middleware
	userEmail := r.Context().Value("userEmail").(string)
	
	todosResponse, err := app.Todo.GetAll(context.Background(), userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todosResponse)
}