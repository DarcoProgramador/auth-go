package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/DarcoProgramador/auth-go/util"
	"github.com/golang-jwt/jwt/v5"
)

const (
	SecretKey = "secret"
)

type serv struct {
	Repository
}

func NewService(r Repository) Service {
	return &serv{r}
}

func (s *serv) Create(u *UserCreate) error {
	password, err := util.HashPassword(strings.TrimSpace(u.Password))
	if err != nil {
		return err
	}

	user := &User{
		Username: strings.TrimSpace(u.Username),
		Password: password,
		Email:    strings.TrimSpace(u.Email),
	}
	return s.Repository.Create(user)
}

func (s *serv) Login(u *UserLogin) (string, error) {
	var user *User
	var err error

	if u.Username != "" {
		user, err = s.Repository.FindByUsername(u.Username)
		if err != nil {
			return "", err
		}
	}

	if u.Email != "" {
		user, err = s.Repository.FindByEmail(u.Email)
		if err != nil {
			return "", err
		}
	}

	if user == nil {
		return "", util.ErrLogin
	}

	if util.CheckPassword(u.Password, user.Password) != nil {
		return "", util.ErrLogin
	}

	claims := UserClaim{
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    strconv.Itoa(int(user.ID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", util.ErrToken
	}

	return tokenString, err
}

func (s *serv) FindByID(id uint) (*User, error) {
	return s.Repository.FindByID(id)
}
