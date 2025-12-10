package reqdto

import (
	"strings"
	"time"

	profileAppCommands "nebulaid/modules/auth/application/profile/commands"
	profileAppQueries "nebulaid/modules/auth/application/profile/queries"
	profileDomain "nebulaid/modules/auth/domain/profile"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"

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

type ProfileCreateRequestDTO struct {
	UserID          uuid.UUID              `json:"user_id" validate:"required"`
	FirstName       *string                `json:"first_name,omitempty"`
	LastName        *string                `json:"last_name,omitempty"`
	Nickname        *string                `json:"nickname,omitempty"`     // 昵称（唯一）
	DisplayName     *string                `json:"display_name,omitempty"` // 显示名称（不唯一）
	AvatarID        *uuid.UUID             `json:"avatar_id,omitempty"`
	BackgroundID    *uuid.UUID             `json:"background_id,omitempty"`
	BackgroundIds   []uuid.UUID            `json:"background_ids,omitempty"` // 背景图片ID数组
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
}

type ProfileUpdateRequestDTO struct {
	ID              uuid.UUID              `params:"id" validate:"required,uuid"`
	FirstName       *string                `json:"first_name,omitempty"`
	LastName        *string                `json:"last_name,omitempty"`
	Nickname        *string                `json:"nickname,omitempty"`     // 昵称（唯一）
	DisplayName     *string                `json:"display_name,omitempty"` // 显示名称（不唯一）
	AvatarID        *uuid.UUID             `json:"avatar_id,omitempty"`
	BackgroundID    *uuid.UUID             `json:"background_id,omitempty"`
	BackgroundIds   []uuid.UUID            `json:"background_ids,omitempty"` // 背景图片ID数组
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
}

type ProfileByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ProfileByUserIDRequestDTO struct {
	UserID uuid.UUID `params:"user_id" validate:"required,uuid"`
}

type ProfileQueryParamsDTO struct {
	Offset *int       `query:"offset"`
	Limit  *int       `query:"limit"`
	UserID *uuid.UUID `query:"user_id"`
	Search *string    `query:"search"`
	Sort   []string   `query:"sort"`
}

func (r *ProfileQueryParamsDTO) ToListQuery() profileAppQueries.ProfileListQuery {
	var userIDs []uuid.UUID
	if r.UserID != nil {
		userIDs = []uuid.UUID{*r.UserID}
	}

	return profileAppQueries.ProfileListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]profileAppQueries.SortField{
			"created_time": profileAppQueries.SortByCreatedTime,
			"display_name": profileAppQueries.SortByDisplayName,
			"nickname":     profileAppQueries.SortByNickname,
		}),
		Search:  r.Search,
		UserIDs: userIDs,
	}
}

func (r *ProfileCreateRequestDTO) ToCreateCmd() profileAppCommands.CreateProfileCmd {
	// 转换 Location DTO 为 "C P C T" 格式的字符串（空格分隔）
	var locationStr *string
	if r.Location != nil {
		var parts []string
		if r.Location.Country != nil {
			parts = append(parts, *r.Location.Country)
		} else {
			parts = append(parts, "")
		}
		if r.Location.Province != nil {
			parts = append(parts, *r.Location.Province)
		} else {
			parts = append(parts, "")
		}
		if r.Location.City != nil {
			parts = append(parts, *r.Location.City)
		} else {
			parts = append(parts, "")
		}
		if r.Location.Timezone != nil {
			parts = append(parts, *r.Location.Timezone)
		} else {
			parts = append(parts, "")
		}
		locationStrVal := strings.Join(parts, " ")
		locationStr = &locationStrVal
	}

	// 转换 Social DTO 为 map[string]interface{}
	var socialLinksMap map[string]interface{}
	if r.Social != nil {
		socialLinksMap = make(map[string]interface{})
		if r.Social.Twitter != nil {
			socialLinksMap["twitter"] = *r.Social.Twitter
		}
		if r.Social.Linkedin != nil {
			socialLinksMap["linkedin"] = *r.Social.Linkedin
		}
		if r.Social.Instagram != nil {
			socialLinksMap["instagram"] = *r.Social.Instagram
		}
		if r.Social.Youtube != nil {
			socialLinksMap["youtube"] = *r.Social.Youtube
		}
	}

	return profileAppCommands.CreateProfileCmd{
		UserID: r.UserID,
		Editable: profileDomain.ProfileEditable{
			FirstName:       r.FirstName,
			LastName:        r.LastName,
			Nickname:        r.Nickname,
			DisplayName:     r.DisplayName,
			AvatarID:        r.AvatarID,
			BackgroundID:    r.BackgroundID,
			BackgroundIds:   r.BackgroundIds,
			Bio:             r.Bio,
			Phone:           r.Phone,
			Birthday:        r.Birthday,
			Age:             r.Age,
			Gender:          r.Gender,
			Location:        locationStr,
			Website:         r.Website,
			Github:          r.Github,
			SocialLinks:     socialLinksMap,
			Preferences:     r.Preferences,
			Skills:          r.Skills,
			PrivacySettings: r.PrivacySettings,
		},
	}
}

func (r *ProfileUpdateRequestDTO) ToUpdateCmd() profileAppCommands.UpdateProfileCmd {
	// 转换 Location DTO 为 "C P C T" 格式的字符串（空格分隔）
	var locationStr *string
	if r.Location != nil {
		var parts []string
		if r.Location.Country != nil {
			parts = append(parts, *r.Location.Country)
		} else {
			parts = append(parts, "")
		}
		if r.Location.Province != nil {
			parts = append(parts, *r.Location.Province)
		} else {
			parts = append(parts, "")
		}
		if r.Location.City != nil {
			parts = append(parts, *r.Location.City)
		} else {
			parts = append(parts, "")
		}
		if r.Location.Timezone != nil {
			parts = append(parts, *r.Location.Timezone)
		} else {
			parts = append(parts, "")
		}
		locationStrVal := strings.Join(parts, " ")
		locationStr = &locationStrVal
	}

	// 转换 Social DTO 为 map[string]interface{}
	var socialLinksMap map[string]interface{}
	if r.Social != nil {
		socialLinksMap = make(map[string]interface{})
		if r.Social.Twitter != nil {
			socialLinksMap["twitter"] = *r.Social.Twitter
		}
		if r.Social.Linkedin != nil {
			socialLinksMap["linkedin"] = *r.Social.Linkedin
		}
		if r.Social.Instagram != nil {
			socialLinksMap["instagram"] = *r.Social.Instagram
		}
		if r.Social.Youtube != nil {
			socialLinksMap["youtube"] = *r.Social.Youtube
		}
	}

	editable := profileDomain.ProfileEditable{}
	if r.FirstName != nil {
		editable.FirstName = r.FirstName
	}
	if r.LastName != nil {
		editable.LastName = r.LastName
	}
	if r.Nickname != nil {
		editable.Nickname = r.Nickname
	}
	if r.DisplayName != nil {
		editable.DisplayName = r.DisplayName
	}
	if r.AvatarID != nil {
		editable.AvatarID = r.AvatarID
	}
	if r.BackgroundID != nil {
		editable.BackgroundID = r.BackgroundID
	}
	if r.BackgroundIds != nil {
		editable.BackgroundIds = r.BackgroundIds
	}
	if r.Bio != nil {
		editable.Bio = r.Bio
	}
	if r.Phone != nil {
		editable.Phone = r.Phone
	}
	if r.Birthday != nil {
		editable.Birthday = r.Birthday
	}
	if r.Age != nil {
		editable.Age = r.Age
	}
	if r.Gender != nil {
		editable.Gender = r.Gender
	}
	if locationStr != nil {
		editable.Location = locationStr
	}
	if r.Website != nil {
		editable.Website = r.Website
	}
	if r.Github != nil {
		editable.Github = r.Github
	}
	if socialLinksMap != nil {
		editable.SocialLinks = socialLinksMap
	}
	if r.Preferences != nil {
		editable.Preferences = r.Preferences
	}
	if r.Skills != nil {
		editable.Skills = r.Skills
	}
	if r.PrivacySettings != nil {
		editable.PrivacySettings = r.PrivacySettings
	}

	return profileAppCommands.UpdateProfileCmd{
		ProfileID: r.ID,
		Editable:  editable,
	}
}
