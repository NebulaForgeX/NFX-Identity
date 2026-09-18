package mapper

import (
	"encoding/json"

	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

func ToModel(p *forgerprofile.Profile) *models.Forgerprofile {
	st := p.State()
	pref := datatypes.JSON(st.Preference)
	if len(pref) == 0 {
		pref = datatypes.JSON(`{}`)
	}
	return &models.Forgerprofile{
		ID:              st.ID,
		AccountID:       st.AccountID,
		ForgerRoles:     pgArrayLiteral(st.Roles),
		ProfileLanguage: enums.AuthProfileLanguage(st.ProfileLanguage),
		Preference:      &pref,
		DisplayName:     st.DisplayName,
		FirstName:       st.FirstName,
		LastName:        st.LastName,
		Country:         st.Country,
		City:            st.City,
		Gender:          st.Gender,
		Birthday:        st.Birthday,
		Website:         st.Website,
		Timezone:        st.Timezone,
		Bio:             st.Bio,
		CreatedAt:       st.CreatedAt,
		UpdatedAt:       st.UpdatedAt,
		DeletedAt:       timex.TimeToGormDeletedAt(st.DeletedAt),
	}
}

func ToDomain(m *models.Forgerprofile) *forgerprofile.Profile {
	var pref json.RawMessage
	if m.Preference != nil {
		pref = json.RawMessage(*m.Preference)
	}
	return forgerprofile.NewFromState(forgerprofile.State{
		ID:              m.ID,
		AccountID:       m.AccountID,
		Roles:           pgArrayScan(m.ForgerRoles),
		ProfileLanguage: string(m.ProfileLanguage),
		Preference:      pref,
		DisplayName:     m.DisplayName,
		FirstName:       m.FirstName,
		LastName:        m.LastName,
		Country:         m.Country,
		City:            m.City,
		Gender:          m.Gender,
		Birthday:        m.Birthday,
		Website:         m.Website,
		Timezone:        m.Timezone,
		Bio:             m.Bio,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       timex.GormDeletedAtToTime(m.DeletedAt),
	})
}

func pgArrayLiteral(roles pq.StringArray) string {
	if len(roles) == 0 {
		return "{}"
	}
	v, err := roles.Value()
	if err != nil || v == nil {
		return "{}"
	}
	switch t := v.(type) {
	case string:
		return t
	case []byte:
		return string(t)
	default:
		return "{}"
	}
}

func pgArrayScan(s string) pq.StringArray {
	var a pq.StringArray
	if s == "" {
		return a
	}
	_ = a.Scan(s)
	return a
}
