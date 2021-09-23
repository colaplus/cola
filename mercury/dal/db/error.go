package db

import "errors"

var (
	ErrUserExists        = errors.New("username is exists")
	ErrUserNotExists     = errors.New("username is not exists")
	ErrUserPasswordWrong = errors.New("username or password wrong")
)
