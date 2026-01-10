package commands

import (
	"nfxid/modules/auth/domain/mfa_factors"

	"github.com/google/uuid"
)

// CreateMFAFactorCmd 创建MFA因子命令
type CreateMFAFactorCmd struct {
	FactorID         string
	TenantID         uuid.UUID
	UserID           uuid.UUID
	Type             mfa_factors.MFAType
	SecretEncrypted  *string
	Phone            *string
	Email            *string
	Name             *string
	Enabled          bool
	RecoveryCodesHash *string
}

// UpdateMFAFactorCmd 更新MFA因子命令
type UpdateMFAFactorCmd struct {
	FactorID         string
	Name             *string
	Enabled          bool
}

// EnableMFAFactorCmd 启用MFA因子命令
type EnableMFAFactorCmd struct {
	FactorID string
}

// DisableMFAFactorCmd 禁用MFA因子命令
type DisableMFAFactorCmd struct {
	FactorID string
}

// UpdateLastUsedCmd 更新最后使用时间命令
type UpdateLastUsedCmd struct {
	FactorID string
}

// DeleteMFAFactorCmd 删除MFA因子命令
type DeleteMFAFactorCmd struct {
	FactorID string
}
