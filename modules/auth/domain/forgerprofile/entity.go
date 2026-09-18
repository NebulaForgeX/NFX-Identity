package forgerprofile

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Profile struct{ state State }

type State struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	Roles           pq.StringArray
	ProfileLanguage string
	Preference      json.RawMessage
	DisplayName     *string
	FirstName       *string
	LastName        *string
	Country         *string
	City            *string
	Gender          *string
	Birthday        *time.Time
	Website         *string
	Timezone        *string
	Bio             *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func NewFromState(st State) *Profile           { return &Profile{state: st} }
func (p *Profile) ID() uuid.UUID               { return p.state.ID }
func (p *Profile) AccountID() uuid.UUID        { return p.state.AccountID }
func (p *Profile) Roles() []string             { return []string(p.state.Roles) }
func (p *Profile) ProfileLanguage() string     { return p.state.ProfileLanguage }
func (p *Profile) Preference() json.RawMessage { return p.state.Preference }
func (p *Profile) DisplayName() *string        { return p.state.DisplayName }
func (p *Profile) FirstName() *string          { return p.state.FirstName }
func (p *Profile) LastName() *string           { return p.state.LastName }
func (p *Profile) Country() *string            { return p.state.Country }
func (p *Profile) City() *string               { return p.state.City }
func (p *Profile) Gender() *string             { return p.state.Gender }
func (p *Profile) Birthday() *time.Time        { return p.state.Birthday }
func (p *Profile) Website() *string            { return p.state.Website }
func (p *Profile) Timezone() *string           { return p.state.Timezone }
func (p *Profile) Bio() *string                { return p.state.Bio }
func (p *Profile) CreatedAt() time.Time        { return p.state.CreatedAt }
func (p *Profile) UpdatedAt() time.Time        { return p.state.UpdatedAt }
func (p *Profile) DeletedAt() *time.Time       { return p.state.DeletedAt }
func (p *Profile) State() State                { return p.state }

func (p *Profile) SoftDelete(at time.Time) {
	p.state.DeletedAt = &at
	p.state.UpdatedAt = at
}
func (p *Profile) SetRoles(roles []string) {
	p.state.Roles = pq.StringArray(roles)
	p.state.UpdatedAt = time.Now()
}
func (p *Profile) SetPreference(raw json.RawMessage) {
	p.state.Preference = raw
	p.state.UpdatedAt = time.Now()
}
func (p *Profile) ApplyPatch(patch map[string]any) {
	now := time.Now()
	p.state.UpdatedAt = now
	str := func(key string, dest **string) {
		v, ok := patch[key]
		if !ok {
			return
		}
		if v == nil {
			*dest = nil
			return
		}
		if s, ok := v.(string); ok {
			*dest = &s
		}
	}
	str("display_name", &p.state.DisplayName)
	str("first_name", &p.state.FirstName)
	str("last_name", &p.state.LastName)
	str("country", &p.state.Country)
	str("city", &p.state.City)
	str("gender", &p.state.Gender)
	str("website", &p.state.Website)
	str("timezone", &p.state.Timezone)
	str("bio", &p.state.Bio)
	if v, ok := patch["profile_language"].(string); ok {
		p.state.ProfileLanguage = v
	}
	if v, ok := patch["birthday"]; ok {
		if v == nil {
			p.state.Birthday = nil
		} else if s, ok := v.(string); ok {
			if t, err := time.Parse("2006-01-02", s); err == nil {
				p.state.Birthday = &t
			}
		}
	}
}
func (p *Profile) HasRole(role string) bool {
	for _, r := range p.state.Roles {
		if r == role {
			return true
		}
	}
	return false
}
