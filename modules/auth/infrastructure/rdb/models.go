package rdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Account struct {
	ID             uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	AccountStatus  string     `gorm:"column:account_status"`
	SignupPlatform string     `gorm:"column:signup_platform"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
}

func (Account) TableName() string { return "auth.Accounts" }

type Identity struct {
	ID               uuid.UUID      `gorm:"column:id;type:uuid;primaryKey"`
	AccountID        uuid.UUID      `gorm:"column:account_id;type:uuid"`
	IdentityProvider string         `gorm:"column:identity_provider"`
	ProviderSubject  string         `gorm:"column:provider_subject"`
	PasswordHash     *string        `gorm:"column:password_hash"`
	Metadata         datatypes.JSON `gorm:"column:metadata"`
	LastLoginAt      *time.Time     `gorm:"column:last_login_at"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        *time.Time     `gorm:"column:deleted_at"`
}

func (Identity) TableName() string { return "auth.Identities" }

type Email struct {
	ID         uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	AccountID  uuid.UUID  `gorm:"column:account_id;type:uuid"`
	Email      string     `gorm:"column:email"`
	IsPrimary  bool       `gorm:"column:is_primary"`
	VerifiedAt *time.Time `gorm:"column:verified_at"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (Email) TableName() string { return "auth.Emails" }

type Phone struct {
	ID         uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	AccountID  uuid.UUID  `gorm:"column:account_id;type:uuid"`
	Phone      string     `gorm:"column:phone"`
	IsPrimary  bool       `gorm:"column:is_primary"`
	VerifiedAt *time.Time `gorm:"column:verified_at"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (Phone) TableName() string { return "auth.Phones" }

type RefreshToken struct {
	ID           uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	AccountID    uuid.UUID  `gorm:"column:account_id;type:uuid"`
	IdentityID   *uuid.UUID `gorm:"column:identity_id;type:uuid"`
	ProfileID    *uuid.UUID `gorm:"column:profile_id;type:uuid"`
	ProfileScope *string    `gorm:"column:profile_scope"`
	DeviceID     *string    `gorm:"column:device_id"`
	TokenHash    string     `gorm:"column:token_hash"`
	ExpiresAt    time.Time  `gorm:"column:expires_at"`
	RevokedAt    *time.Time `gorm:"column:revoked_at"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

func (RefreshToken) TableName() string { return "auth.RefreshTokens" }

type ForgerProfile struct {
	ID              uuid.UUID      `gorm:"column:id;type:uuid;primaryKey"`
	AccountID       uuid.UUID      `gorm:"column:account_id;type:uuid"`
	ForgerRoles     pq.StringArray `gorm:"column:forger_roles;type:auth.forger_role[]"`
	ProfileLanguage string         `gorm:"column:profile_language"`
	Preference      datatypes.JSON `gorm:"column:preference"`
	DisplayName     *string        `gorm:"column:display_name"`
	FirstName       *string        `gorm:"column:first_name"`
	LastName        *string        `gorm:"column:last_name"`
	Country         *string        `gorm:"column:country"`
	City            *string        `gorm:"column:city"`
	Gender          *string        `gorm:"column:gender"`
	Birthday        *time.Time     `gorm:"column:birthday"`
	Website         *string        `gorm:"column:website"`
	Timezone        *string        `gorm:"column:timezone"`
	Bio             *string        `gorm:"column:bio"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at"`
	DeletedAt       *time.Time     `gorm:"column:deleted_at"`
}

func (ForgerProfile) TableName() string { return "auth.ForgerProfiles" }

type AuthorityProfile struct {
	ID              uuid.UUID      `gorm:"column:id;type:uuid;primaryKey"`
	AccountID       uuid.UUID      `gorm:"column:account_id;type:uuid"`
	AuthorityRoles  pq.StringArray `gorm:"column:authority_roles;type:auth.authority_role[]"`
	ProfileLanguage string         `gorm:"column:profile_language"`
	Preference      datatypes.JSON `gorm:"column:preference"`
	DisplayName     *string        `gorm:"column:display_name"`
	FirstName       *string        `gorm:"column:first_name"`
	LastName        *string        `gorm:"column:last_name"`
	Country         *string        `gorm:"column:country"`
	City            *string        `gorm:"column:city"`
	Gender          *string        `gorm:"column:gender"`
	Birthday        *time.Time     `gorm:"column:birthday"`
	Website         *string        `gorm:"column:website"`
	Timezone        *string        `gorm:"column:timezone"`
	Bio             *string        `gorm:"column:bio"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at"`
	DeletedAt       *time.Time     `gorm:"column:deleted_at"`
}

func (AuthorityProfile) TableName() string { return "auth.AuthorityProfiles" }

type ProfileAvatar struct {
	ID        uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	ProfileID uuid.UUID  `gorm:"column:profile_id;type:uuid"`
	ImageID   uuid.UUID  `gorm:"column:image_id;type:uuid"`
	IsActive  bool       `gorm:"column:is_active"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (ProfileAvatar) TableName() string { return "auth.ForgerProfileAvatars" }

type ForgerProfileAvatar struct {
	ProfileAvatar
}

func (ForgerProfileAvatar) TableName() string { return "auth.ForgerProfileAvatars" }

type AuthorityProfileAvatar struct {
	ProfileAvatar
}

func (AuthorityProfileAvatar) TableName() string { return "auth.AuthorityProfileAvatars" }

type ProfileBackground struct {
	ID        uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	ProfileID uuid.UUID  `gorm:"column:profile_id;type:uuid"`
	ImageID   uuid.UUID  `gorm:"column:image_id;type:uuid"`
	SortOrder int        `gorm:"column:sort_order"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

type ForgerProfileBackground struct {
	ProfileBackground
}

func (ForgerProfileBackground) TableName() string { return "auth.ForgerProfileBackgrounds" }

type AuthorityProfileBackground struct {
	ProfileBackground
}

func (AuthorityProfileBackground) TableName() string { return "auth.AuthorityProfileBackgrounds" }

type ProfileSettings struct {
	ID                uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	LoginNotification bool       `gorm:"column:login_notification"`
	CreatedAt         time.Time  `gorm:"column:created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at"`
}

type ForgerProfileSettings struct {
	ProfileSettings
}

func (ForgerProfileSettings) TableName() string { return "auth.ForgerProfileSettings" }

type AuthorityProfileSettings struct {
	ProfileSettings
}

func (AuthorityProfileSettings) TableName() string { return "auth.AuthorityProfileSettings" }

