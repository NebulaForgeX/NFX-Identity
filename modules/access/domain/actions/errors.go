package actions

import "errors"

var (
	ErrActionNotFound        = errors.New("action not found")
	ErrActionKeyRequired     = errors.New("action key is required")
	ErrActionNameRequired    = errors.New("action name is required")
	ErrActionServiceRequired = errors.New("action service is required")
	ErrActionKeyExists       = errors.New("action key already exists")
	ErrSystemActionDelete    = errors.New("cannot delete system action")
)
