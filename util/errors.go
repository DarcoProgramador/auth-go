package util

import "errors"

var (
	ErrNotFoundUser             = errors.New("user not found")
	ErrCreateUser               = errors.New("error create user")
	ErrBadRequest               = errors.New("error parsing data")
	ErrUnauthorized             = errors.New("error unauthorized")
	ErrTokenInternalServerError = errors.New("internal server error parsing token data")
	ErrLogin                    = errors.New("error login")
	ErrToken                    = errors.New("error token")
	ErrUserExists               = errors.New("user already exist")
)
