package profile

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"nfxidentity/enums"
	"time"
)

// * =============================== Authority profile entity =============================== !//
type AuthorityProfile struct {
	state AuthorityProfileState
}

type AuthorityProfileState struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	AuthorityRoles  []enums.AuthAuthorityRole
	ProfileLanguage enums.AuthProfileLanguage
	Preference      *datatypes.JSON
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

func (u *AuthorityProfile) ID() uuid.UUID        { return u.state.ID }
func (u *AuthorityProfile) AccountID() uuid.UUID { return u.state.AccountID }
func (u *AuthorityProfile) AuthorityRoles() []enums.AuthAuthorityRole {
	return append([]enums.AuthAuthorityRole(nil), u.state.AuthorityRoles...)
}
func (u *AuthorityProfile) ProfileLanguage() enums.AuthProfileLanguage {
	return u.state.ProfileLanguage
}
func (u *AuthorityProfile) Preference() *datatypes.JSON { return u.state.Preference }
func (u *AuthorityProfile) DisplayName() *string        { return u.state.DisplayName }
func (u *AuthorityProfile) FirstName() *string          { return u.state.FirstName }
func (u *AuthorityProfile) LastName() *string           { return u.state.LastName }
func (u *AuthorityProfile) Country() *string            { return u.state.Country }
func (u *AuthorityProfile) City() *string               { return u.state.City }
func (u *AuthorityProfile) Gender() *string             { return u.state.Gender }
func (u *AuthorityProfile) Birthday() *time.Time        { return u.state.Birthday }
func (u *AuthorityProfile) Website() *string            { return u.state.Website }
func (u *AuthorityProfile) Timezone() *string           { return u.state.Timezone }
func (u *AuthorityProfile) Bio() *string                { return u.state.Bio }
func (u *AuthorityProfile) CreatedAt() time.Time        { return u.state.CreatedAt }
func (u *AuthorityProfile) UpdatedAt() time.Time        { return u.state.UpdatedAt }
func (u *AuthorityProfile) DeletedAt() *time.Time       { return u.state.DeletedAt }

func NewAuthorityProfileFromState(st AuthorityProfileState) *AuthorityProfile {
	return &AuthorityProfile{state: st}
}

//! ====================================================================================== !//

// AuthorityProfileAvatar 表示用户资料头像关联行（auth.AuthorityProfileAvatars）。
type AuthorityProfileAvatar struct {
	state AuthorityProfileAvatarState
}

// AuthorityProfileAvatarState 与表 AuthorityProfileAvatars 对齐。
type AuthorityProfileAvatarState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (a *AuthorityProfileAvatar) State() AuthorityProfileAvatarState { return a.state }

func (a *AuthorityProfileAvatar) ID() uuid.UUID { return a.state.ID }

func (a *AuthorityProfileAvatar) ProfileID() uuid.UUID { return a.state.ProfileID }

func (a *AuthorityProfileAvatar) ImageID() uuid.UUID { return a.state.ImageID }

func (a *AuthorityProfileAvatar) IsActive() bool { return a.state.IsActive }

func (a *AuthorityProfileAvatar) CreatedAt() time.Time { return a.state.CreatedAt }

func (a *AuthorityProfileAvatar) UpdatedAt() time.Time { return a.state.UpdatedAt }

func (a *AuthorityProfileAvatar) DeletedAt() *time.Time { return a.state.DeletedAt }

func NewAuthorityProfileAvatarFromState(st AuthorityProfileAvatarState) *AuthorityProfileAvatar {
	return &AuthorityProfileAvatar{state: st}
}

//! ====================================================================================== !//

// AuthorityProfileBackground 表示用户资料背景图关联行（auth.AuthorityProfileBackgrounds）。
type AuthorityProfileBackground struct {
	state AuthorityProfileBackgroundState
}

// AuthorityProfileBackgroundState 与表 AuthorityProfileBackgrounds 对齐。
type AuthorityProfileBackgroundState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (b *AuthorityProfileBackground) State() AuthorityProfileBackgroundState { return b.state }
func (b *AuthorityProfileBackground) ID() uuid.UUID                          { return b.state.ID }
func (b *AuthorityProfileBackground) ProfileID() uuid.UUID                   { return b.state.ProfileID }
func (b *AuthorityProfileBackground) ImageID() uuid.UUID                     { return b.state.ImageID }
func (b *AuthorityProfileBackground) SortOrder() int                         { return b.state.SortOrder }
func (b *AuthorityProfileBackground) CreatedAt() time.Time                   { return b.state.CreatedAt }
func (b *AuthorityProfileBackground) UpdatedAt() time.Time                   { return b.state.UpdatedAt }
func (b *AuthorityProfileBackground) DeletedAt() *time.Time                  { return b.state.DeletedAt }
func NewAuthorityProfileBackgroundFromState(st AuthorityProfileBackgroundState) *AuthorityProfileBackground {
	return &AuthorityProfileBackground{state: st}
}

//! ====================================================================================== !//

// AuthorityProfileSettings 表示用户资料系统设置行（auth.AuthorityProfileSettings，1:1，id = profile id）。
type AuthorityProfileSettings struct {
	state AuthorityProfileSettingsState
}

type AuthorityProfileSettingsState struct {
	ID                uuid.UUID
	LoginNotification bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func (s *AuthorityProfileSettings) ID() uuid.UUID           { return s.state.ID }
func (s *AuthorityProfileSettings) LoginNotification() bool { return s.state.LoginNotification }
func (s *AuthorityProfileSettings) CreatedAt() time.Time    { return s.state.CreatedAt }
func (s *AuthorityProfileSettings) UpdatedAt() time.Time    { return s.state.UpdatedAt }
func (s *AuthorityProfileSettings) DeletedAt() *time.Time   { return s.state.DeletedAt }

func NewAuthorityProfileSettingsFromState(st AuthorityProfileSettingsState) *AuthorityProfileSettings {
	return &AuthorityProfileSettings{state: st}
}

// * =============================== Forger profile entity =============================== !//
type ForgerProfile struct {
	state ForgerProfileState
}

type ForgerProfileState struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	ForgerRoles     []enums.AuthForgerRole
	ProfileLanguage enums.AuthProfileLanguage
	Preference      *datatypes.JSON
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

func (u *ForgerProfile) ID() uuid.UUID        { return u.state.ID }
func (u *ForgerProfile) AccountID() uuid.UUID { return u.state.AccountID }
func (u *ForgerProfile) ForgerRoles() []enums.AuthForgerRole {
	return append([]enums.AuthForgerRole(nil), u.state.ForgerRoles...)
}
func (u *ForgerProfile) ProfileLanguage() enums.AuthProfileLanguage {
	return u.state.ProfileLanguage
}
func (u *ForgerProfile) Preference() *datatypes.JSON { return u.state.Preference }
func (u *ForgerProfile) DisplayName() *string        { return u.state.DisplayName }
func (u *ForgerProfile) FirstName() *string          { return u.state.FirstName }
func (u *ForgerProfile) LastName() *string           { return u.state.LastName }
func (u *ForgerProfile) Country() *string            { return u.state.Country }
func (u *ForgerProfile) City() *string               { return u.state.City }
func (u *ForgerProfile) Gender() *string             { return u.state.Gender }
func (u *ForgerProfile) Birthday() *time.Time        { return u.state.Birthday }
func (u *ForgerProfile) Website() *string            { return u.state.Website }
func (u *ForgerProfile) Timezone() *string           { return u.state.Timezone }
func (u *ForgerProfile) Bio() *string                { return u.state.Bio }
func (u *ForgerProfile) CreatedAt() time.Time        { return u.state.CreatedAt }
func (u *ForgerProfile) UpdatedAt() time.Time        { return u.state.UpdatedAt }
func (u *ForgerProfile) DeletedAt() *time.Time       { return u.state.DeletedAt }

func NewForgerProfileFromState(st ForgerProfileState) *ForgerProfile {
	return &ForgerProfile{state: st}
}

//! ====================================================================================== !//

// ForgerProfileAvatar 表示用户资料头像关联行（auth.ForgerProfileAvatars）。
type ForgerProfileAvatar struct {
	state ForgerProfileAvatarState
}

// ForgerProfileAvatarState 与表 ForgerProfileAvatars 对齐。
type ForgerProfileAvatarState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (a *ForgerProfileAvatar) State() ForgerProfileAvatarState { return a.state }

func (a *ForgerProfileAvatar) ID() uuid.UUID { return a.state.ID }

func (a *ForgerProfileAvatar) ProfileID() uuid.UUID { return a.state.ProfileID }

func (a *ForgerProfileAvatar) ImageID() uuid.UUID { return a.state.ImageID }

func (a *ForgerProfileAvatar) IsActive() bool { return a.state.IsActive }

func (a *ForgerProfileAvatar) CreatedAt() time.Time { return a.state.CreatedAt }

func (a *ForgerProfileAvatar) UpdatedAt() time.Time { return a.state.UpdatedAt }

func (a *ForgerProfileAvatar) DeletedAt() *time.Time { return a.state.DeletedAt }

func NewForgerProfileAvatarFromState(st ForgerProfileAvatarState) *ForgerProfileAvatar {
	return &ForgerProfileAvatar{state: st}
}

//! ====================================================================================== !//

// ForgerProfileBackground 表示用户资料背景图关联行（auth.ForgerProfileBackgrounds）。
type ForgerProfileBackground struct {
	state ForgerProfileBackgroundState
}

// ForgerProfileBackgroundState 与表 ForgerProfileBackgrounds 对齐。
type ForgerProfileBackgroundState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (b *ForgerProfileBackground) State() ForgerProfileBackgroundState { return b.state }
func (b *ForgerProfileBackground) ID() uuid.UUID                       { return b.state.ID }
func (b *ForgerProfileBackground) ProfileID() uuid.UUID                { return b.state.ProfileID }
func (b *ForgerProfileBackground) ImageID() uuid.UUID                  { return b.state.ImageID }
func (b *ForgerProfileBackground) SortOrder() int                      { return b.state.SortOrder }
func (b *ForgerProfileBackground) CreatedAt() time.Time                { return b.state.CreatedAt }
func (b *ForgerProfileBackground) UpdatedAt() time.Time                { return b.state.UpdatedAt }
func (b *ForgerProfileBackground) DeletedAt() *time.Time               { return b.state.DeletedAt }
func NewForgerProfileBackgroundFromState(st ForgerProfileBackgroundState) *ForgerProfileBackground {
	return &ForgerProfileBackground{state: st}
}

//! ====================================================================================== !//

// ForgerProfileSettings 表示用户资料系统设置行（auth.ForgerProfileSettings，1:1，id = profile id）。
type ForgerProfileSettings struct {
	state ForgerProfileSettingsState
}

type ForgerProfileSettingsState struct {
	ID                uuid.UUID
	LoginNotification bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func (s *ForgerProfileSettings) ID() uuid.UUID           { return s.state.ID }
func (s *ForgerProfileSettings) LoginNotification() bool { return s.state.LoginNotification }
func (s *ForgerProfileSettings) CreatedAt() time.Time    { return s.state.CreatedAt }
func (s *ForgerProfileSettings) UpdatedAt() time.Time    { return s.state.UpdatedAt }
func (s *ForgerProfileSettings) DeletedAt() *time.Time   { return s.state.DeletedAt }

func NewForgerProfileSettingsFromState(st ForgerProfileSettingsState) *ForgerProfileSettings {
	return &ForgerProfileSettings{state: st}
}
