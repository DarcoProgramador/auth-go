package user

import (
	"strings"

	"github.com/DarcoProgramador/auth-go/db"
	"github.com/DarcoProgramador/auth-go/util"
)

type repo struct {
	Database *db.DB
}

func NewRepository(db *db.DB) Repository {
	return &repo{Database: db}
}

func (r *repo) Create(u *User) error {
	if err := r.Database.Create(u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return util.ErrUserExists
		}
		return util.ErrCreateUser
	}
	return nil
}

func (r *repo) FindByUsername(username string) (*User, error) {
	var user User
	err := r.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, util.ErrNotFoundUser
	}
	return &user, err
}

func (r *repo) FindByEmail(email string) (*User, error) {
	var user User
	err := r.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, util.ErrNotFoundUser
	}
	return &user, err
}

func (r *repo) FindByID(id uint) (*User, error) {
	var user User
	err := r.Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, util.ErrNotFoundUser
	}
	return &user, err
}
