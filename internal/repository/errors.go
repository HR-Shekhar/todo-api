package repository

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrTodoNotFound = errors.New("todo not found")
var ErrDeletionFailed = errors.New("failed to delete todo")