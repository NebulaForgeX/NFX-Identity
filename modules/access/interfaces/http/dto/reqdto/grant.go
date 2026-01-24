package reqdto

import (
	grantAppCommands "nfxid/modules/access/application/grants/commands"
	grantDomain "nfxid/modules/access/domain/grants"

	"github.com/google/uuid"
)

type GrantCreateRequestDTO struct {
	SubjectType  string     `json:"subject_type" validate:"required"`
	SubjectID    uuid.UUID  `json:"subject_id" validate:"required"`
	GrantType    string     `json:"grant_type" validate:"required"`
	GrantRefID   uuid.UUID  `json:"grant_ref_id" validate:"required"`
	TenantID     *uuid.UUID `json:"tenant_id,omitempty"`
	AppID        *uuid.UUID `json:"app_id,omitempty"`
	ResourceType *string    `json:"resource_type,omitempty"`
	ResourceID   *uuid.UUID `json:"resource_id,omitempty"`
	Effect       string     `json:"effect" validate:"required"`
	ExpiresAt    *string    `json:"expires_at,omitempty"`
}

type GrantUpdateRequestDTO struct {
	GrantID   uuid.UUID `params:"id" validate:"required,uuid"`
	ExpiresAt *string   `json:"expires_at,omitempty"`
}

type GrantRevokeRequestDTO struct {
	GrantID      uuid.UUID `params:"id" validate:"required,uuid"`
	RevokedBy    uuid.UUID `json:"revoked_by" validate:"required"`
	RevokeReason *string   `json:"revoke_reason,omitempty"`
}

type GrantByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type GrantBySubjectRequestDTO struct {
	SubjectType string    `query:"subject_type" validate:"required"`
	SubjectID   uuid.UUID `query:"subject_id" validate:"required"`
	TenantID    *uuid.UUID `query:"tenant_id,omitempty"`
}

// ToQueryCmd 转换为查询命令（用于 GetGrantsBySubject）
// 返回 subjectType 字符串和 subjectID，由 application 层处理类型转换
func (r *GrantBySubjectRequestDTO) ToQueryParams() (string, uuid.UUID, *uuid.UUID) {
	return r.SubjectType, r.SubjectID, r.TenantID
}

func (r *GrantCreateRequestDTO) ToCreateCmd() grantAppCommands.CreateGrantCmd {
	cmd := grantAppCommands.CreateGrantCmd{
		SubjectID:    r.SubjectID,
		GrantRefID:   r.GrantRefID,
		TenantID:     r.TenantID,
		AppID:        r.AppID,
		ResourceType: r.ResourceType,
		ResourceID:   r.ResourceID,
		ExpiresAt:    r.ExpiresAt,
	}
	
	// Parse enums
	if r.SubjectType != "" {
		cmd.SubjectType = grantDomain.SubjectType(r.SubjectType)
	}
	if r.GrantType != "" {
		cmd.GrantType = grantDomain.GrantType(r.GrantType)
	}
	if r.Effect != "" {
		cmd.Effect = grantDomain.GrantEffect(r.Effect)
	}
	
	return cmd
}

func (r *GrantUpdateRequestDTO) ToUpdateCmd() grantAppCommands.UpdateGrantCmd {
	return grantAppCommands.UpdateGrantCmd{
		GrantID:   r.GrantID,
		ExpiresAt: r.ExpiresAt,
	}
}

func (r *GrantRevokeRequestDTO) ToRevokeCmd() grantAppCommands.RevokeGrantCmd {
	return grantAppCommands.RevokeGrantCmd{
		GrantID:      r.GrantID,
		RevokedBy:    r.RevokedBy,
		RevokeReason: r.RevokeReason,
	}
}
