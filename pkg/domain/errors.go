package domain

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("item not found")
	ErrAlreadyExist   = errors.New("item alreade exist")
	ErrBadRequst      = errors.New("invalid request")
	ErrInvalidStatus  = errors.New("invalid status")
)
