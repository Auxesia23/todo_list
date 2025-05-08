package utils

import (
	"errors"
	"net/mail"
)

func VerifyPasswordAndEmail(password, email string) error {
	_, err := mail.ParseAddress(email)
    if err != nil {
        return errors.New("invalid email format")
    }
    
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }
    return nil
}

