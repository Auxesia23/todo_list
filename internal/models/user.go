package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     		string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Username      	string `json:"username" gorm:"type:varchar(100);not null"`
	Password  		string `json:"password" gorm:"type:varchar(100);not null"`
}

type UserInput struct {
	Username     string `json:"username"`
	Email      	 string `json:"email"`
	Password 	 string `json:"password"`
}

type UserResponse struct {
	Username  	*string `json:"username"`
	Email 		*string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
