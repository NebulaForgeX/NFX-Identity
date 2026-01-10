package commands

import (
	"nfxid/modules/clients/domain/rate_limits"

	"github.com/google/uuid"
)

// CreateRateLimitCmd 创建速率限制命令
type CreateRateLimitCmd struct {
	AppID         uuid.UUID
	LimitType     rate_limits.RateLimitType
	LimitValue    int
	WindowSeconds int
	Description   *string
	Status        string
	CreatedBy     *uuid.UUID
}

// UpdateRateLimitCmd 更新速率限制命令
type UpdateRateLimitCmd struct {
	RateLimitID   uuid.UUID
	LimitValue    int
	WindowSeconds int
	Description   *string
	UpdatedBy     *uuid.UUID
}

// UpdateRateLimitStatusCmd 更新速率限制状态命令
type UpdateRateLimitStatusCmd struct {
	RateLimitID uuid.UUID
	Status      string
}

// DeleteRateLimitCmd 删除速率限制命令
type DeleteRateLimitCmd struct {
	RateLimitID uuid.UUID
}
