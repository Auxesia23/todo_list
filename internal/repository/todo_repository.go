package repository

import (
	"context"

	"github.com/Auxesia23/todo_list/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface{
	Create(ctx context.Context,todo models.Todo)(models.TodoResponse, error)
	GetAll(ctx context.Context, email string)([]models.TodoResponse, error)
	GetByCompleted(ctx context.Context, email string, completed int)([]models.TodoResponse, error)
	Update(ctx context.Context, email string, id uint)(models.TodoResponse, error)
	Delete(ctx context.Context, email string, id uint)( error)
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
		ID: &newTodo.ID,
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
	err := repo.DB.WithContext(ctx).Where("user_email = ?", email).Order("created_at DESC").Find(&todos).Error
	if err != nil {
		return []models.TodoResponse{}, err
	}
	
	//loop setiap todo agar sesuai format response dan di appeand ke variabel baru
	var todosResponse []models.TodoResponse
	for _,todo := range(todos){
		todosResponse = append(todosResponse, models.TodoResponse{
			ID: &todo.ID,
			Title: &todo.Title,
			Description: &todo.Description,
			DueDate: &todo.DueDate,
			Completed: &todo.Completed,
			UpdatedAt: &todo.UpdatedAt,
		})
	}

	return todosResponse, nil
}

func (repo *TodoRepo) GetByCompleted(ctx context.Context, email string, completed int)([]models.TodoResponse, error){
	//mengambil semua todo berdasarkan email dan status completed
	var todos []models.Todo
	if completed == 1 {
		err := repo.DB.WithContext(ctx).Where("completed = true AND user_email = ?", email).Order("created_at DESC").Find(&todos).Error
		if err != nil {
			return []models.TodoResponse{}, err
		}
	}else {
		err := repo.DB.WithContext(ctx).Where("completed = false AND user_email = ?", email).Order("created_at DESC").Find(&todos).Error
		if err != nil {
			return []models.TodoResponse{}, err
		}
	}
	
	//loop setiap todo agar sesuai format response dan di appeand ke variabel baru
	var todosResponse []models.TodoResponse
	for _,todo := range(todos){
		todosResponse = append(todosResponse, models.TodoResponse{
			ID: &todo.ID,
			Title: &todo.Title,
			Description: &todo.Description,
			DueDate: &todo.DueDate,
			Completed: &todo.Completed,
			UpdatedAt: &todo.UpdatedAt,
		})
	}

	return todosResponse, nil
}

func (repo *TodoRepo) Update(ctx context.Context, email string, id uint) (models.TodoResponse, error) {
	var todo models.Todo
	err := repo.DB.WithContext(ctx).Where("id = ? AND user_email = ?", id, email).First(&todo).Error
	if err != nil {
		return models.TodoResponse{}, err
	}

	// Toggle completed
	newCompleted := !todo.Completed
	err = repo.DB.WithContext(ctx).Model(&todo).Update("completed", newCompleted).Error
	if err != nil {
		return models.TodoResponse{}, err
	}

	// Ambil ulang todo untuk memastikan UpdatedAt dan Completed terbaru
	err = repo.DB.WithContext(ctx).First(&todo, todo.ID).Error
	if err != nil {
		return models.TodoResponse{}, err
	}

	response := models.TodoResponse{
		ID: &todo.ID,
		Title:       &todo.Title,
		Description: &todo.Description,
		DueDate:     &todo.DueDate,
		Completed:   &todo.Completed,
		UpdatedAt:   &todo.UpdatedAt,
	}

	return response, nil
}

func (repo *TodoRepo) Delete(ctx context.Context, email string, id uint)(error){
	var todo models.Todo
	err := repo.DB.WithContext(ctx).Where("id = ? AND user_email = ?", id,email).First(&todo).Error
	if err != nil {
		return err
	}
	
	err = repo.DB.Delete(&todo, id).Error
	if err != nil {
		return err
	}
	
	return nil
}
