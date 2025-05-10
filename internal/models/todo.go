package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	UserEmail   string    `json:"user_email" gorm:"not null"`

	User User `json:"user" gorm:"foreignKey:UserEmail;references:Email;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
}

type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type TodoResponse struct {
	ID          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Completed   *bool      `json:"completed"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
