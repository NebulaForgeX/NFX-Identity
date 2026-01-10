package results

import (
	"time"

	"nfxid/modules/auth/domain/mfa_factors"

	"github.com/google/uuid"
)

type MFAFactorRO struct {
	ID                uuid.UUID
	FactorID          string
	TenantID          uuid.UUID
	UserID            uuid.UUID
	Type              mfa_factors.MFAType
	SecretEncrypted   *string
	Phone             *string
	Email             *string
	Name              *string
	Enabled           bool
	CreatedAt         time.Time
	LastUsedAt        *time.Time
	RecoveryCodesHash *string
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

// MFAFactorMapper 将 Domain MFAFactor 转换为 Application MFAFactorRO
func MFAFactorMapper(mf *mfa_factors.MFAFactor) MFAFactorRO {
	if mf == nil {
		return MFAFactorRO{}
	}

	return MFAFactorRO{
		ID:                mf.ID(),
		FactorID:          mf.FactorID(),
		TenantID:          mf.TenantID(),
		UserID:            mf.UserID(),
		Type:              mf.Type(),
		SecretEncrypted:   mf.SecretEncrypted(),
		Phone:             mf.Phone(),
		Email:             mf.Email(),
		Name:              mf.Name(),
		Enabled:           mf.Enabled(),
		CreatedAt:         mf.CreatedAt(),
		LastUsedAt:        mf.LastUsedAt(),
		RecoveryCodesHash: mf.RecoveryCodesHash(),
		UpdatedAt:         mf.UpdatedAt(),
		DeletedAt:         mf.DeletedAt(),
	}
}
