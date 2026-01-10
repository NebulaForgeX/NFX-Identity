package commands

import (
	"github.com/google/uuid"
)

// CreateBadgeCmd 创建徽章命令
type CreateBadgeCmd struct {
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string
	IsSystem    bool
}

// UpdateBadgeCmd 更新徽章命令
type UpdateBadgeCmd struct {
	BadgeID     uuid.UUID
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string
}

// DeleteBadgeCmd 删除徽章命令
type DeleteBadgeCmd struct {
	BadgeID uuid.UUID
}
