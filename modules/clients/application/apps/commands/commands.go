package commands

import (
	"nfxid/modules/clients/domain/apps"

	"github.com/google/uuid"
)

// CreateAppCmd 创建应用命令
type CreateAppCmd struct {
	AppID       string
	TenantID    uuid.UUID
	Name        string
	Description *string
	Type        apps.AppType
	Status      apps.AppStatus
	Environment apps.Environment
	CreatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

// UpdateAppCmd 更新应用命令
type UpdateAppCmd struct {
	AppID       uuid.UUID
	Name        string
	Description *string
	Type        apps.AppType
	Environment apps.Environment
	UpdatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

// UpdateAppStatusCmd 更新应用状态命令
type UpdateAppStatusCmd struct {
	AppID   uuid.UUID
	Status  apps.AppStatus
}

// DeleteAppCmd 删除应用命令
type DeleteAppCmd struct {
	AppID uuid.UUID
}
