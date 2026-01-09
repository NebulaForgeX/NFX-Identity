package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// AccountLockoutDomainToModel 将 Domain AccountLockout 转换为 Model AccountLockout
func AccountLockoutDomainToModel(al *account_lockouts.AccountLockout) *models.AccountLockout {
	if al == nil {
		return nil
	}

	return &models.AccountLockout{
		UserID:        al.UserID(),
		TenantID:      al.TenantID(),
		LockedUntil:   al.LockedUntil(),
		LockReason:    lockReasonDomainToEnum(al.LockReason()),
		LockedAt:      al.LockedAt(),
		LockedBy:      al.LockedBy(),
		ActorID:       al.ActorID(),
		UnlockedAt:    al.UnlockedAt(),
		UnlockedBy:    al.UnlockedBy(),
		UnlockActorID: al.UnlockActorID(),
		CreatedAt:     al.CreatedAt(),
		UpdatedAt:     al.UpdatedAt(),
	}
}

// AccountLockoutModelToDomain 将 Model AccountLockout 转换为 Domain AccountLockout
func AccountLockoutModelToDomain(m *models.AccountLockout) *account_lockouts.AccountLockout {
	if m == nil {
		return nil
	}

	state := account_lockouts.AccountLockoutState{
		UserID:        m.UserID,
		TenantID:      m.TenantID,
		LockedUntil:   m.LockedUntil,
		LockReason:    lockReasonEnumToDomain(m.LockReason),
		LockedAt:      m.LockedAt,
		LockedBy:      m.LockedBy,
		ActorID:       m.ActorID,
		UnlockedAt:    m.UnlockedAt,
		UnlockedBy:    m.UnlockedBy,
		UnlockActorID: m.UnlockActorID,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}

	return account_lockouts.NewAccountLockoutFromState(state)
}

// AccountLockoutModelToUpdates 将 Model AccountLockout 转换为更新字段映射
func AccountLockoutModelToUpdates(m *models.AccountLockout) map[string]any {
	return map[string]any{
		models.AccountLockoutCols.LockedUntil:   m.LockedUntil,
		models.AccountLockoutCols.LockReason:   m.LockReason,
		models.AccountLockoutCols.LockedAt:     m.LockedAt,
		models.AccountLockoutCols.LockedBy:     m.LockedBy,
		models.AccountLockoutCols.ActorID:      m.ActorID,
		models.AccountLockoutCols.UnlockedAt:   m.UnlockedAt,
		models.AccountLockoutCols.UnlockedBy:   m.UnlockedBy,
		models.AccountLockoutCols.UnlockActorID: m.UnlockActorID,
		models.AccountLockoutCols.UpdatedAt:     m.UpdatedAt,
	}
}

// 枚举转换辅助函数

func lockReasonDomainToEnum(lr account_lockouts.LockReason) enums.AuthLockReason {
	switch lr {
	case account_lockouts.LockReasonTooManyAttempts:
		return enums.AuthLockReasonTooManyAttempts
	case account_lockouts.LockReasonAdminLock:
		return enums.AuthLockReasonAdminLock
	case account_lockouts.LockReasonRiskDetected:
		return enums.AuthLockReasonRiskDetected
	case account_lockouts.LockReasonSuspiciousActivity:
		return enums.AuthLockReasonSuspiciousActivity
	case account_lockouts.LockReasonCompliance:
		return enums.AuthLockReasonCompliance
	case account_lockouts.LockReasonOther:
		return enums.AuthLockReasonOther
	default:
		return enums.AuthLockReasonTooManyAttempts
	}
}

func lockReasonEnumToDomain(lr enums.AuthLockReason) account_lockouts.LockReason {
	switch lr {
	case enums.AuthLockReasonTooManyAttempts:
		return account_lockouts.LockReasonTooManyAttempts
	case enums.AuthLockReasonAdminLock:
		return account_lockouts.LockReasonAdminLock
	case enums.AuthLockReasonRiskDetected:
		return account_lockouts.LockReasonRiskDetected
	case enums.AuthLockReasonSuspiciousActivity:
		return account_lockouts.LockReasonSuspiciousActivity
	case enums.AuthLockReasonCompliance:
		return account_lockouts.LockReasonCompliance
	case enums.AuthLockReasonOther:
		return account_lockouts.LockReasonOther
	default:
		return account_lockouts.LockReasonTooManyAttempts
	}
}
