package password_history

import "errors"

var (
	ErrPasswordHistoryNotFound = errors.New("password history not found")
	ErrUserIDRequired          = errors.New("user id is required")
	ErrTenantIDRequired        = errors.New("tenant id is required")
	ErrPasswordHashRequired    = errors.New("password hash is required")
)
