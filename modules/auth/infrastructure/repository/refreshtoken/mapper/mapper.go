package mapper

import (
	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"
)

func ToModel(t *refreshtoken.RefreshToken) *models.Refreshtoken {
	st := t.State()
	var scope *enums.AuthProfileScope
	if st.ProfileScope != nil {
		s := enums.AuthProfileScope(*st.ProfileScope)
		scope = &s
	}
	return &models.Refreshtoken{
		ID:           st.ID,
		AccountID:    st.AccountID,
		IdentityID:   st.IdentityID,
		ProfileID:    st.ProfileID,
		ProfileScope: scope,
		DeviceID:     st.DeviceID,
		TokenHash:    st.TokenHash,
		ExpiresAt:    st.ExpiresAt,
		RevokedAt:    st.RevokedAt,
		CreatedAt:    st.CreatedAt,
		DeletedAt:    timex.TimeToGormDeletedAt(st.DeletedAt),
	}
}

func ToDomain(m *models.Refreshtoken) *refreshtoken.RefreshToken {
	var scope *string
	if m.ProfileScope != nil {
		s := string(*m.ProfileScope)
		scope = &s
	}
	return refreshtoken.NewFromState(refreshtoken.RefreshTokenState{
		ID:           m.ID,
		AccountID:    m.AccountID,
		IdentityID:   m.IdentityID,
		ProfileID:    m.ProfileID,
		ProfileScope: scope,
		DeviceID:     m.DeviceID,
		TokenHash:    m.TokenHash,
		ExpiresAt:    m.ExpiresAt,
		RevokedAt:    m.RevokedAt,
		CreatedAt:    m.CreatedAt,
		DeletedAt:    timex.GormDeletedAtToTime(m.DeletedAt),
	})
}
