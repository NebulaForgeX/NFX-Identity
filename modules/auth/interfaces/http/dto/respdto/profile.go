package respdto

import (
	"strings"
	"time"

	profileAppViews "nfxid/modules/auth/application/profile/views"

	"github.com/google/uuid"
)

// LocationDTO 位置信息（结构化）
type LocationDTO struct {
	Country  *string `json:"country,omitempty"`
	Province *string `json:"province,omitempty"`
	City     *string `json:"city,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// SocialDTO 社交链接（结构化，不包含 github/website）
type SocialDTO struct {
	Twitter   *string `json:"twitter,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	Youtube   *string `json:"youtube,omitempty"`
}

// ImageInfoDTO 图片信息 DTO
type ImageInfoDTO struct {
	ID       string  `json:"id"`
	TypeID   string  `json:"type_id,omitempty"`
	URL      *string `json:"url,omitempty"`
	IsPublic bool    `json:"is_public"`
}

type ProfileDTO struct {
	ID              uuid.UUID              `json:"id"`
	UserID          uuid.UUID              `json:"user_id"`
	Username        string                 `json:"username"`
	Email           string                 `json:"email"`
	UserPhone       *string                `json:"user_phone,omitempty"`
	UserStatus      string                 `json:"user_status"`
	IsVerified      bool                   `json:"is_verified"`
	FirstName       *string                `json:"first_name,omitempty"`
	LastName        *string                `json:"last_name,omitempty"`
	Nickname        *string                `json:"nickname,omitempty"`     // 昵称（唯一）
	DisplayName     *string                `json:"display_name,omitempty"` // 显示名称（不唯一）
	AvatarID        *uuid.UUID             `json:"avatar_id,omitempty"`
	BackgroundID    *uuid.UUID             `json:"background_id,omitempty"`
	BackgroundIds   []uuid.UUID            `json:"background_ids,omitempty"` // 背景图片ID数组
	Images          []ImageInfoDTO         `json:"images,omitempty"`         // 用户的所有图片（通过 gRPC 获取）
	Bio             *string                `json:"bio,omitempty"`
	Phone           *string                `json:"phone,omitempty"`
	Birthday        *time.Time             `json:"birthday,omitempty"`
	Age             *int                   `json:"age,omitempty"` // 年龄
	Gender          *string                `json:"gender,omitempty"`
	Location        *LocationDTO           `json:"location,omitempty"` // 结构化位置信息
	Website         *string                `json:"website,omitempty"`
	Github          *string                `json:"github,omitempty"` // GitHub 用户名或 URL
	Social          *SocialDTO             `json:"social,omitempty"` // 结构化社交链接（不包含 github/website）
	Preferences     map[string]interface{} `json:"preferences,omitempty"`
	Skills          map[string]int         `json:"skills,omitempty"`           // 技能：{"golang": 10, "python": 8, ...}
	PrivacySettings map[string]interface{} `json:"privacy_settings,omitempty"` // 隐私设置
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// ProfileViewToDTO converts application ProfileView to response DTO
func ProfileViewToDTO(v *profileAppViews.ProfileView) *ProfileDTO {
	if v == nil {
		return nil
	}

	// 转换 Location string ("C P C T" 格式) 为 LocationDTO
	var locationDTO *LocationDTO
	if v.Location != nil && *v.Location != "" {
		parts := strings.Fields(*v.Location) // 按空格分割
		locationDTO = &LocationDTO{}
		if len(parts) > 0 && parts[0] != "" {
			locationDTO.Country = &parts[0]
		}
		if len(parts) > 1 && parts[1] != "" {
			locationDTO.Province = &parts[1]
		}
		if len(parts) > 2 && parts[2] != "" {
			locationDTO.City = &parts[2]
		}
		if len(parts) > 3 && parts[3] != "" {
			locationDTO.Timezone = &parts[3]
		}
	}

	// 转换 SocialLinks datatypes.JSON 为 SocialDTO
	var socialDTO *SocialDTO
	if v.SocialLinks != nil {
		// Note: Need to unmarshal datatypes.JSON to map[string]interface{}
		// For now, leave it empty or implement JSON unmarshaling if needed
		socialDTO = &SocialDTO{}
	}

	// 转换 Preferences, Skills, PrivacySettings from datatypes.JSON
	var preferences map[string]interface{}
	var skills map[string]int
	var privacySettings map[string]interface{}
	// Note: Need to unmarshal datatypes.JSON to map[string]interface{}
	// For now, leave them empty or implement JSON unmarshaling if needed

	return &ProfileDTO{
		ID:              v.ID,
		UserID:          v.UserID,
		Username:        v.Username,
		Email:           v.Email,
		UserPhone:       v.UserPhone,
		UserStatus:      v.UserStatus,
		IsVerified:      v.IsVerified,
		FirstName:       v.FirstName,
		LastName:        v.LastName,
		Nickname:        v.Nickname,
		DisplayName:     v.DisplayName,
		AvatarID:        v.AvatarID,
		BackgroundID:    v.BackgroundID,
		BackgroundIds:   v.BackgroundIds,
		Images:          ImageInfoListToDTO(v.Images),
		Bio:             v.Bio,
		Phone:           v.Phone,
		Birthday:        v.Birthday,
		Age:             v.Age,
		Gender:          v.Gender,
		Location:        locationDTO,
		Website:         v.Website,
		Github:          v.Github,
		Social:          socialDTO,
		Preferences:     preferences,
		Skills:          skills,
		PrivacySettings: privacySettings,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}
}

// ProfileListViewToDTO converts list of ProfileView to DTOs
func ProfileListViewToDTO(views []profileAppViews.ProfileView) []ProfileDTO {
	dtos := make([]ProfileDTO, len(views))
	for i, v := range views {
		if dto := ProfileViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}

// ImageInfoListToDTO converts list of ImageInfo to DTOs
func ImageInfoListToDTO(images []profileAppViews.ImageInfo) []ImageInfoDTO {
	if images == nil {
		return nil
	}
	dtos := make([]ImageInfoDTO, len(images))
	for i, img := range images {
		dtos[i] = ImageInfoDTO{
			ID:       img.ID,
			TypeID:   img.TypeID,
			URL:      img.URL,
			IsPublic: img.IsPublic,
		}
	}
	return dtos
}
