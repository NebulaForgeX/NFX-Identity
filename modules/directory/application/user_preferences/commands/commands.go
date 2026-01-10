package commands

import (
	"github.com/google/uuid"
)

// CreateUserPreferenceCmd 创建用户偏好命令
type CreateUserPreferenceCmd struct {
	UserID        uuid.UUID
	Theme         string
	Language      string
	Timezone      string
	Notifications map[string]interface{}
	Privacy       map[string]interface{}
	Display       map[string]interface{}
	Other         map[string]interface{}
}

// UpdateUserPreferenceCmd 更新用户偏好命令
type UpdateUserPreferenceCmd struct {
	UserPreferenceID uuid.UUID
	Theme            *string
	Language         *string
	Timezone         *string
	Notifications    map[string]interface{}
	Privacy          map[string]interface{}
	Display         map[string]interface{}
	Other            map[string]interface{}
}

// DeleteUserPreferenceCmd 删除用户偏好命令
type DeleteUserPreferenceCmd struct {
	UserPreferenceID uuid.UUID
}
