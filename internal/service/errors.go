package service

import "errors"

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUsernameAlreadyExists = errors.New("username already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrTodoNotFound = errors.New("todo not found")
var ErrTodoCreationFailed = errors.New("Todo Creation Failed")