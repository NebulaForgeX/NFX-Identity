package mapper

import (
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func RefreshTokenDomainToModel(t *refreshtokenDomain.RefreshToken) *rdbmodels.RefreshToken {
	if t == nil {
		return nil
	}
	return &rdbmodels.RefreshToken{
		ID:           t.ID(),
		AccountID:    t.AccountID(),
		IdentityID:   t.IdentityID(),
		ProfileID:    t.ProfileID(),
		ProfileScope: t.ProfileScope(),
		DeviceID:     t.DeviceID(),
		TokenHash:    t.TokenHash(),
		ExpiresAt:    t.ExpiresAt(),
		RevokedAt:    t.RevokedAt(),
		CreatedAt:    t.CreatedAt(),
		DeletedAt:    ptrx.TimePtrToDeletedAt(t.DeletedAt()),
	}
}

func RefreshTokenModelToDomain(m *rdbmodels.RefreshToken) *refreshtokenDomain.RefreshToken {
	if m == nil {
		return nil
	}
	return refreshtokenDomain.NewRefreshTokenFromState(refreshtokenDomain.RefreshTokenState{
		ID:           m.ID,
		AccountID:    m.AccountID,
		IdentityID:   m.IdentityID,
		ProfileID:    m.ProfileID,
		ProfileScope: m.ProfileScope,
		DeviceID:     m.DeviceID,
		TokenHash:    m.TokenHash,
		ExpiresAt:    m.ExpiresAt,
		RevokedAt:    m.RevokedAt,
		CreatedAt:    m.CreatedAt,
		DeletedAt:    ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func RefreshTokenDomainToUpdates(t *refreshtokenDomain.RefreshToken) map[string]any {
	m := RefreshTokenDomainToModel(t)
	return map[string]any{
		rdbmodels.RefreshTokenCols.AccountID:    m.AccountID,
		rdbmodels.RefreshTokenCols.IdentityID:   m.IdentityID,
		rdbmodels.RefreshTokenCols.ProfileID:    m.ProfileID,
		rdbmodels.RefreshTokenCols.ProfileScope: m.ProfileScope,
		rdbmodels.RefreshTokenCols.DeviceID:     m.DeviceID,
		rdbmodels.RefreshTokenCols.TokenHash:  m.TokenHash,
		rdbmodels.RefreshTokenCols.ExpiresAt:  m.ExpiresAt,
		rdbmodels.RefreshTokenCols.RevokedAt:  m.RevokedAt,
		rdbmodels.RefreshTokenCols.DeletedAt:  m.DeletedAt,
	}
}
