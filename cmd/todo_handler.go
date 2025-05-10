package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Auxesia23/todo_list/internal/models"
	"github.com/go-chi/chi/v5"
)


func (app *application) CreateTodo (w http.ResponseWriter, r *http.Request){
	//ambil email user yang melakukan request dari context yang dikirim middleware
	userEmail := r.Context().Value("userEmail").(string)
	
	log.Println(r.Body)
	
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
	completedParam := r.URL.Query().Get("completed")
	var todosResponse []models.TodoResponse
	
	
	if completedParam == "" {
	    // Jika parameter 'completed' tidak disediakan, ambil semua todo
		var err error
	    todosResponse, err = app.Todo.GetAll(context.Background(), userEmail)
		if err != nil {
	        http.Error(w, err.Error(), http.StatusNotFound)
	        return
	    }
	} else {
	    // Konversi parameter 'completed' menjadi integer
	    completedInt, err := strconv.Atoi(completedParam)
	    if err != nil || (completedInt != 0 && completedInt != 1) {
	        http.Error(w, "Invalid 'completed' parameter. Must be 0 or 1.", http.StatusBadRequest)
	        return
	    }
	
	    // Ambil todo berdasarkan status 'completed'
	    todosResponse, err = app.Todo.GetByCompleted(context.Background(), userEmail, completedInt)
	    if err != nil {
	        http.Error(w, err.Error(), http.StatusNotFound)
	        return
	    }
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todosResponse)
}

func (app *application) UpdateTodo(w http.ResponseWriter, r *http.Request){
	//ambil email user yang melakukan request dari context yang dikirim middleware
	userEmail := r.Context().Value("userEmail").(string)
	
	todoId := chi.URLParam(r, "id")
	todoIdInt, _ := strconv.ParseInt(todoId, 10, 64)
	todoIdUint := uint(todoIdInt)
	
	response, err := app.Todo.Update(context.Background(), userEmail, todoIdUint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}