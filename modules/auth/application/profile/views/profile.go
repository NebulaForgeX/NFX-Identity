package views

import (
	"time"

	profileDomainViews "nebulaid/modules/auth/domain/profile/views"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ImageInfo 图片信息（通过 gRPC 获取）
type ImageInfo struct {
	ID       string  `json:"id"`
	TypeID   string  `json:"type_id,omitempty"`
	URL      *string `json:"url,omitempty"`
	IsPublic bool    `json:"is_public"`
}

type ProfileView struct {
	ID              uuid.UUID       `json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	Username        string          `json:"username"`
	Email           string          `json:"email"`
	UserPhone       *string         `json:"user_phone"`
	UserStatus      string          `json:"user_status"`
	IsVerified      bool            `json:"is_verified"`
	FirstName       *string         `json:"first_name"`
	LastName        *string         `json:"last_name"`
	Nickname        *string         `json:"nickname"`
	DisplayName     *string         `json:"display_name"`
	AvatarID        *uuid.UUID      `json:"avatar_id"`
	BackgroundID    *uuid.UUID      `json:"background_id"`
	BackgroundIds   []uuid.UUID     `json:"background_ids"`
	Images          []ImageInfo     `json:"images,omitempty"` // 通过 gRPC 获取的用户所有图片
	Bio             *string         `json:"bio"`
	Phone           *string         `json:"phone"`
	Birthday        *time.Time      `json:"birthday"`
	Age             *int            `json:"age"`
	Gender          *string         `json:"gender"`
	Location        *string         `json:"location"`
	Website         *string         `json:"website"`
	Github          *string         `json:"github"`
	SocialLinks     *datatypes.JSON `json:"social_links"`
	Preferences     *datatypes.JSON `json:"preferences"`
	Skills          *datatypes.JSON `json:"skills"`
	PrivacySettings *datatypes.JSON `json:"privacy_settings"`
	Occupations     *datatypes.JSON `json:"occupations"`
	Educations      *datatypes.JSON `json:"educations"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// ProfileViewMapper 将 Domain ProfileView 转换为 Application ProfileView
func ProfileViewMapper(v profileDomainViews.ProfileView) ProfileView {
	return ProfileView{
		ID:            v.ID,
		UserID:        v.UserID,
		Username:      v.Username,
		Email:         v.Email,
		UserPhone:     v.UserPhone,
		UserStatus:    v.UserStatus,
		IsVerified:    v.IsVerified,
		FirstName:     v.FirstName,
		LastName:      v.LastName,
		Nickname:      v.Nickname,
		DisplayName:   v.DisplayName,
		AvatarID:      v.AvatarID,
		BackgroundID:  v.BackgroundID,
		BackgroundIds: v.BackgroundIds,
		// Images 字段会在 application layer 中通过 gRPC 填充
		Bio:             v.Bio,
		Phone:           v.Phone,
		Birthday:        v.Birthday,
		Age:             v.Age,
		Gender:          v.Gender,
		Location:        v.Location,
		Website:         v.Website,
		Github:          v.Github,
		SocialLinks:     v.SocialLinks,
		Preferences:     v.Preferences,
		Skills:          v.Skills,
		PrivacySettings: v.PrivacySettings,
		Occupations:     v.Occupations,
		Educations:      v.Educations,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}
}
