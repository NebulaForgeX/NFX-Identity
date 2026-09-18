package views

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type ListEmailItem struct {
	ID         uuid.UUID  `gorm:"column:id"`
	AccountID  uuid.UUID  `gorm:"column:account_id"`
	Email      string     `gorm:"column:email"`
	IsPrimary  bool       `gorm:"column:is_primary"`
	VerifiedAt *time.Time `gorm:"column:verified_at"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
}

func (ListEmailItem) TableName() string { return "auth.ListEmailItems" }

type ListPhoneItem struct {
	ID         uuid.UUID  `gorm:"column:id"`
	AccountID  uuid.UUID  `gorm:"column:account_id"`
	Phone      string     `gorm:"column:phone"`
	IsPrimary  bool       `gorm:"column:is_primary"`
	VerifiedAt *time.Time `gorm:"column:verified_at"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
}

func (ListPhoneItem) TableName() string { return "auth.ListPhoneItems" }

type ListForgerProfileItem struct {
	ProfileID         uuid.UUID      `gorm:"column:profile_id"`
	AccountID         uuid.UUID      `gorm:"column:account_id"`
	ForgerRoles       pq.StringArray `gorm:"column:forger_roles;type:auth.forger_role[]"`
	DisplayName       *string        `gorm:"column:display_name"`
	ProfileLanguage   string         `gorm:"column:profile_language"`
	City              *string        `gorm:"column:city"`
	Country           *string        `gorm:"column:country"`
	Website           *string        `gorm:"column:website"`
	Timezone          *string        `gorm:"column:timezone"`
	Birthday          *time.Time     `gorm:"column:birthday"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	AvatarImageID     *uuid.UUID     `gorm:"column:avatar_image_id"`
	BackgroundImageID *uuid.UUID     `gorm:"column:background_image_id"`
}

func (ListForgerProfileItem) TableName() string { return "auth.ListForgerProfileItems" }

type ListAuthorityProfileItem struct {
	ProfileID         uuid.UUID      `gorm:"column:profile_id"`
	AccountID         uuid.UUID      `gorm:"column:account_id"`
	AuthorityRoles    pq.StringArray `gorm:"column:authority_roles;type:auth.authority_role[]"`
	DisplayName       *string        `gorm:"column:display_name"`
	ProfileLanguage   string         `gorm:"column:profile_language"`
	City              *string        `gorm:"column:city"`
	Country           *string        `gorm:"column:country"`
	Website           *string        `gorm:"column:website"`
	Timezone          *string        `gorm:"column:timezone"`
	Birthday          *time.Time     `gorm:"column:birthday"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	AvatarImageID     *uuid.UUID     `gorm:"column:avatar_image_id"`
	BackgroundImageID *uuid.UUID     `gorm:"column:background_image_id"`
}

func (ListAuthorityProfileItem) TableName() string { return "auth.ListAuthorityProfileItems" }

type FullAccountInformation struct {
	AccountID      uuid.UUID      `gorm:"column:account_id"`
	AccountStatus  string         `gorm:"column:account_status"`
	SignupPlatform string         `gorm:"column:signup_platform"`
	PrimaryEmail   *string        `gorm:"column:primary_email"`
	PrimaryPhone   *string        `gorm:"column:primary_phone"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	Preference     datatypes.JSON `gorm:"column:preference"`
}

func (FullAccountInformation) TableName() string { return "auth.FullAccountInformation" }
