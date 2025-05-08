package repository

import (
	"context"

	"github.com/Auxesia23/todo_list/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface{
	Create(ctx context.Context,todo models.Todo)(models.TodoResponse, error)
	GetAll(ctx context.Context, email string)([]models.TodoResponse, error)
}

type TodoRepo struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository{
	return &TodoRepo{
		DB : db,
	}
}

func (repo *TodoRepo) Create(ctx context.Context,todo models.Todo)(models.TodoResponse, error){
	//membuat todo baru
	err := repo.DB.WithContext(ctx).Create(&todo).Error
	if err != nil {
		return models.TodoResponse{}, err
	}
	
	//mengambil data todo baru
	var newTodo models.Todo
	err = repo.DB.WithContext(ctx).First(&newTodo, todo.ID).Error
	if err != nil {
		return models.TodoResponse{}, err
	}
	
	//mengubah data todo agar sesuai untuk response
	response := models.TodoResponse {
		Title: &newTodo.Title,
		Description: &newTodo.Description,
		DueDate: &newTodo.DueDate,
		Completed: &newTodo.Completed,
		UpdatedAt: &newTodo.UpdatedAt,
	}
	
	return response, nil
}

func (repo *TodoRepo) GetAll(ctx context.Context, email string)([]models.TodoResponse, error){
	//mengambil semua todo berdasarkan email
	var todos []models.Todo
	err := repo.DB.WithContext(ctx).Where("user_email = ?", email).Find(&todos).Error
	if err != nil {
		return []models.TodoResponse{}, err
	}
	
	//loop setiap todo agar sesuai format response dan di appeand ke variabel baru
	var todosResponse []models.TodoResponse
	for _,todo := range(todos){
		todosResponse = append(todosResponse, models.TodoResponse{
			Title: &todo.Title,
			Description: &todo.Description,
			DueDate: &todo.DueDate,
			Completed: &todo.Completed,
			UpdatedAt: &todo.UpdatedAt,
		})
	}

	return todosResponse, nil
}