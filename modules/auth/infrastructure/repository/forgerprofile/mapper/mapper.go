package mapper

import (
	"encoding/json"

	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

func ToModel(p *forgerprofile.Profile) *models.ForgerProfile {
	st := p.State()
	pref := st.Preference
	if len(pref) == 0 {
		pref = json.RawMessage(`{}`)
	}
	return &models.ForgerProfile{
		ID: st.ID, AccountID: st.AccountID, ForgerRoles: st.Roles, ProfileLanguage: st.ProfileLanguage,
		Preference: datatypes.JSON(pref), DisplayName: st.DisplayName, FirstName: st.FirstName, LastName: st.LastName,
		Country: st.Country, City: st.City, Gender: st.Gender, Birthday: st.Birthday, Website: st.Website,
		Timezone: st.Timezone, Bio: st.Bio, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
}
func ToDomain(m *models.ForgerProfile) *forgerprofile.Profile {
	return forgerprofile.NewFromState(forgerprofile.State{
		ID: m.ID, AccountID: m.AccountID, Roles: m.ForgerRoles, ProfileLanguage: m.ProfileLanguage,
		Preference: json.RawMessage(m.Preference), DisplayName: m.DisplayName, FirstName: m.FirstName, LastName: m.LastName,
		Country: m.Country, City: m.City, Gender: m.Gender, Birthday: m.Birthday, Website: m.Website,
		Timezone: m.Timezone, Bio: m.Bio, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt,
	})
}
