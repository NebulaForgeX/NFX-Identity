package commands

import (
	"github.com/google/uuid"
)

// CreateUserBadgeCmd 创建用户徽章命令
type CreateUserBadgeCmd struct {
	UserID      uuid.UUID
	BadgeID     uuid.UUID
	Description string
	Level       int
	EarnedAt    *string
}

// DeleteUserBadgeCmd 删除用户徽章命令
type DeleteUserBadgeCmd struct {
	UserBadgeID uuid.UUID
}
