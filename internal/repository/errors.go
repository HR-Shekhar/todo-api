package repository

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrTodoNotFound = errors.New("todo not found")