package mapper

import (
	"encoding/json"

	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

func ToModel(p *authorityprofile.Profile) *models.AuthorityProfile {
	st := p.State()
	pref := st.Preference
	if len(pref) == 0 {
		pref = json.RawMessage(`{}`)
	}
	return &models.AuthorityProfile{
		ID: st.ID, AccountID: st.AccountID, AuthorityRoles: st.Roles, ProfileLanguage: st.ProfileLanguage,
		Preference: datatypes.JSON(pref), DisplayName: st.DisplayName, FirstName: st.FirstName, LastName: st.LastName,
		Country: st.Country, City: st.City, Gender: st.Gender, Birthday: st.Birthday, Website: st.Website,
		Timezone: st.Timezone, Bio: st.Bio, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
}
func ToDomain(m *models.AuthorityProfile) *authorityprofile.Profile {
	return authorityprofile.NewFromState(authorityprofile.State{
		ID: m.ID, AccountID: m.AccountID, Roles: m.AuthorityRoles, ProfileLanguage: m.ProfileLanguage,
		Preference: json.RawMessage(m.Preference), DisplayName: m.DisplayName, FirstName: m.FirstName, LastName: m.LastName,
		Country: m.Country, City: m.City, Gender: m.Gender, Birthday: m.Birthday, Website: m.Website,
		Timezone: m.Timezone, Bio: m.Bio, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt,
	})
}
