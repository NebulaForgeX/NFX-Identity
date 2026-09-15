package mapper

import (
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func ToModel(t *refreshtoken.RefreshToken) *models.RefreshToken {
	st := t.State()
	return &models.RefreshToken{ID: st.ID, AccountID: st.AccountID, IdentityID: st.IdentityID, ProfileID: st.ProfileID, ProfileScope: st.ProfileScope, DeviceID: st.DeviceID, TokenHash: st.TokenHash, ExpiresAt: st.ExpiresAt, RevokedAt: st.RevokedAt, CreatedAt: st.CreatedAt, DeletedAt: st.DeletedAt}
}
func ToDomain(m *models.RefreshToken) *refreshtoken.RefreshToken {
	return refreshtoken.NewFromState(refreshtoken.RefreshTokenState{ID: m.ID, AccountID: m.AccountID, IdentityID: m.IdentityID, ProfileID: m.ProfileID, ProfileScope: m.ProfileScope, DeviceID: m.DeviceID, TokenHash: m.TokenHash, ExpiresAt: m.ExpiresAt, RevokedAt: m.RevokedAt, CreatedAt: m.CreatedAt, DeletedAt: m.DeletedAt})
}
