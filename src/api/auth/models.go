package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex"`
	Email string `gorm:"uniqueIndex"`
	PasswordHash string
	Name string
}

func NewUser(username, email, passwordHash, name string) User {
	return User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
