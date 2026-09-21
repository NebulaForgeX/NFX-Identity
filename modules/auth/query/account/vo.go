package account

import (
	"time"

	"nfxidentity/enums"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AccountVO struct {
	ID             uuid.UUID                `json:"id"`
	AccountStatus  enums.AuthAccountStatus  `json:"account_status"`
	SignupPlatform enums.AuthSignupPlatform `json:"signup_platform"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
}

type EmailVO struct {
	ID         uuid.UUID  `json:"id"`
	AccountID  uuid.UUID  `json:"account_id"`
	Email      string     `json:"email"`
	IsPrimary  bool       `json:"is_primary"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// PrimaryEmailVO is the resolved primary email of a community profile's account,
// with the profile language for localized outbound email.
type PrimaryEmailVO struct {
	Email           string  `json:"email"`
	ProfileLanguage string  `json:"profile_language"`
	DisplayName     *string `json:"display_name,omitempty"`
}

type PhoneVO struct {
	ID         uuid.UUID  `json:"id"`
	AccountID  uuid.UUID  `json:"account_id"`
	Phone      string     `json:"phone"`
	IsPrimary  bool       `json:"is_primary"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type CommunityProfileAvatarVO struct {
	ID        uuid.UUID `json:"id"`
	ProfileID uuid.UUID `json:"profile_id"`
	ImageID   uuid.UUID `json:"image_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommunityProfileBackgroundVO struct {
	ID        uuid.UUID `json:"id"`
	ProfileID uuid.UUID `json:"profile_id"`
	ImageID   uuid.UUID `json:"image_id"`
	SortOrder int32     `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommunityProfileSettingsVO struct {
	LoginNotification bool      `json:"login_notification"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CommunityProfileVO struct {
	ProfileID       uuid.UUID                      `json:"profile_id"`
	AccountID       uuid.UUID                      `json:"account_id"`
	ForgerRoles     []enums.AuthForgerRole         `json:"forger_roles"`
	ProfileLanguage enums.AuthProfileLanguage      `json:"profile_language"`
	Preference      *datatypes.JSON                `json:"preference"`
	DisplayName     *string                        `json:"display_name"`
	FirstName       *string                        `json:"first_name"`
	LastName        *string                        `json:"last_name"`
	Country         *string                        `json:"country"`
	City            *string                        `json:"city"`
	Gender          *string                        `json:"gender"`
	Birthday        *time.Time                     `json:"birthday"`
	Website         *string                        `json:"website"`
	Timezone        *string                        `json:"timezone"`
	Bio             *string                        `json:"bio"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	Avatars         []CommunityProfileAvatarVO     `json:"avatars"`
	Backgrounds     []CommunityProfileBackgroundVO `json:"backgrounds"`
	Settings        *CommunityProfileSettingsVO    `json:"settings"`
}

type AuthorityProfileAvatarVO struct {
	ID        uuid.UUID `json:"id"`
	ProfileID uuid.UUID `json:"profile_id"`
	ImageID   uuid.UUID `json:"image_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthorityProfileBackgroundVO struct {
	ID        uuid.UUID `json:"id"`
	ProfileID uuid.UUID `json:"profile_id"`
	ImageID   uuid.UUID `json:"image_id"`
	SortOrder int32     `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthorityProfileSettingsVO struct {
	LoginNotification bool      `json:"login_notification"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type AuthorityProfileVO struct {
	ProfileID       uuid.UUID                      `json:"profile_id"`
	AccountID       uuid.UUID                      `json:"account_id"`
	AuthorityRoles  []enums.AuthAuthorityRole      `json:"authority_roles"`
	ProfileLanguage enums.AuthProfileLanguage      `json:"profile_language"`
	Preference      *datatypes.JSON                `json:"preference"`
	DisplayName     *string                        `json:"display_name"`
	FirstName       *string                        `json:"first_name"`
	LastName        *string                        `json:"last_name"`
	Country         *string                        `json:"country"`
	City            *string                        `json:"city"`
	Gender          *string                        `json:"gender"`
	Birthday        *time.Time                     `json:"birthday"`
	Website         *string                        `json:"website"`
	Timezone        *string                        `json:"timezone"`
	Bio             *string                        `json:"bio"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	Avatars         []AuthorityProfileAvatarVO     `json:"avatars"`
	Backgrounds     []AuthorityProfileBackgroundVO `json:"backgrounds"`
	Settings        *AuthorityProfileSettingsVO    `json:"settings"`
}

// FullAccountInformationVO 是账号聚合（含账号下全部 profile），用于管理端列表等场景。
type FullAccountInformationVO struct {
	Account           AccountVO            `json:"account"`
	Emails            []EmailVO            `json:"emails"`
	Phones            []PhoneVO            `json:"phones"`
	CommunityProfiles []CommunityProfileVO `json:"community_profiles"`
	AuthorityProfiles []AuthorityProfileVO `json:"authority_profiles"`
}

// FullAccountInformationWithCommunityProfileVO 是「当前登录态」视角的账号信息：
// 只携带当前选中的单个 profile（由 token 中的 profile_id 决定），前端无需再自行 resolve。
// 账号 id 在 Account.ID，不再单独冗余一个顶层 AccountID。
type FullAccountInformationWithCommunityProfileVO struct {
	Account          AccountVO           `json:"account"`
	Emails           []EmailVO           `json:"emails"`
	Phones           []PhoneVO           `json:"phones"`
	CommunityProfile *CommunityProfileVO `json:"community_profile"`
}

// FullAccountInformationWithAuthorityProfileVO 是「当前登录态（权限档）」视角的账号信息：
// 只携带当前选中的单个 authority profile（由 token 中的 profile_id 决定）。
type FullAccountInformationWithAuthorityProfileVO struct {
	Account          AccountVO           `json:"account"`
	Emails           []EmailVO           `json:"emails"`
	Phones           []PhoneVO           `json:"phones"`
	AuthorityProfile *AuthorityProfileVO `json:"authority_profile"`
}
