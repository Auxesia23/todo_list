package repository

import (
	"context"
	"errors"

	"github.com/Auxesia23/todo_list/internal/models"
	"github.com/Auxesia23/todo_list/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.UserResponse, error)
	Login(ctx context.Context, email, password string)(string, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Create (ctx context.Context, user models.User) (models.UserResponse, error) {
	//Check apakah user dengan email ini sudah ada atau belum
	var existingUser models.User
	err := repo.DB.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return models.UserResponse{},errors.New("user with this email already exist")
	}
	
	//Membuat user baru
	err = repo.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return models.UserResponse{}, err
	}
	
	//Mengambil data user yang baru dibuat
	var newUser models.User
	err = repo.DB.WithContext(ctx).First(&newUser, user.ID).Error
	if err != nil {
		return models.UserResponse{}, err
	}
	
	response := models.UserResponse{
		Username: &newUser.Username,
		Email: &newUser.Email,
	}
	
	return response, nil
}

func(repo *UserRepo) Login(ctx context.Context, email, password string) (string, error){
	//Ambil data user dari db jika ada
	var user models.User
	err := repo.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		//jika tidak ada return error
		return "", errors.New("invalid credencial")
	}
	
	//check password dan hashed password di db
	if !utils.CheckPasswordHash(password, user.Password){
		return "", errors.New("invalid credencial")
	}
	
	//Membuat token sesuai data user
	token, err := utils.GenerateToken(&user)
	if err != nil {
		return "", err
	}
	
	return token, nil
}
