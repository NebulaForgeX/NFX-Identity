package commands

import (
	"nfxid/modules/tenants/domain/groups"

	"github.com/google/uuid"
)

// CreateGroupCmd 创建组命令
type CreateGroupCmd struct {
	GroupID       string
	TenantID      uuid.UUID
	Name          string
	Type          groups.GroupType
	ParentGroupID *uuid.UUID
	Description   *string
	CreatedBy     *uuid.UUID
	Metadata      map[string]interface{}
}

// UpdateGroupCmd 更新组命令
type UpdateGroupCmd struct {
	GroupID       uuid.UUID
	Name          string
	Type          groups.GroupType
	ParentGroupID *uuid.UUID
	Description   *string
	Metadata      map[string]interface{}
}

// DeleteGroupCmd 删除组命令
type DeleteGroupCmd struct {
	GroupID uuid.UUID
}
