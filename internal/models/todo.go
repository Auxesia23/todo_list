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
	Status      bool      `json:"status" gorm:"default:false"`
	UserID      uint      `json:"user_id" gorm:"not null"`

	User User `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
}

type TodoInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type TodoResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      bool      `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
}
