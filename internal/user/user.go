package user

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Repository interface {
	Create(u *User) error
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
}

type Service interface {
	Create(u *UserCreate) error
	FindByID(id uint) (*User, error)
	Login(u *UserLogin) (string, error)
}

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

type UserClaim struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
