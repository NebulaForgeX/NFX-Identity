package mapper

import (
	"encoding/json"
	"strings"
	"time"

	profileAppViews "nfxid/modules/auth/application/profile/views"
	profilepb "nfxid/protos/gen/auth/profile"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ProfileViewToProto 将 ProfileView 转换为 proto Profile 消息
// includeBadges: 是否包含徽章信息（通过 profile_badges 关联）
func ProfileViewToProto(v *profileAppViews.ProfileView, includeBadges bool) *profilepb.Profile {
	if v == nil {
		return nil
	}

	profile := &profilepb.Profile{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.FirstName != nil {
		profile.FirstName = v.FirstName
	}
	if v.LastName != nil {
		profile.LastName = v.LastName
	}
	if v.Nickname != nil {
		profile.Nickname = v.Nickname
	}
	if v.DisplayName != nil {
		profile.DisplayName = v.DisplayName
	}
	if v.AvatarID != nil {
		avatarIDStr := v.AvatarID.String()
		profile.AvatarId = &avatarIDStr
	}
	if v.BackgroundID != nil {
		backgroundIDStr := v.BackgroundID.String()
		profile.BackgroundId = &backgroundIDStr
	}
	if len(v.BackgroundIds) > 0 {
		profile.BackgroundIds = make([]string, len(v.BackgroundIds))
		for i, id := range v.BackgroundIds {
			profile.BackgroundIds[i] = id.String()
		}
	}
	if v.Bio != nil {
		profile.Bio = v.Bio
	}
	if v.Phone != nil {
		profile.Phone = v.Phone
	}
	if v.Birthday != nil {
		birthdayStr := formatDate(*v.Birthday)
		profile.Birthday = &birthdayStr
	}
	if v.Age != nil {
		ageInt32 := int32(*v.Age)
		profile.Age = &ageInt32
	}
	if v.Gender != nil {
		profile.Gender = v.Gender
	}
	// 转换 Location string ("C P C T" 格式) 为 proto Location message
	if v.Location != nil && *v.Location != "" {
		locationStr := *v.Location
		profile.Location = &locationStr

		parts := strings.Fields(locationStr) // 按空格分割
		locationStruct := &profilepb.Location{}
		if len(parts) > 0 && parts[0] != "" {
			locationStruct.Country = &parts[0]
		}
		if len(parts) > 1 && parts[1] != "" {
			locationStruct.Province = &parts[1]
		}
		if len(parts) > 2 && parts[2] != "" {
			locationStruct.City = &parts[2]
		}
		if len(parts) > 3 && parts[3] != "" {
			locationStruct.Timezone = &parts[3]
		}
		profile.LocationStruct = locationStruct
	}
	if v.Website != nil {
		profile.Website = v.Website
	}
	if v.Github != nil {
		profile.Github = v.Github
	}

	// 转换 SocialLinks (*datatypes.JSON) 为 proto Social message
	if v.SocialLinks != nil {
		socialMap := make(map[string]interface{})
		if err := json.Unmarshal(*v.SocialLinks, &socialMap); err == nil {
			social := &profilepb.Social{}
			if twitter, ok := socialMap["twitter"].(string); ok && twitter != "" {
				social.Twitter = &twitter
			}
			if linkedin, ok := socialMap["linkedin"].(string); ok && linkedin != "" {
				social.Linkedin = &linkedin
			}
			if instagram, ok := socialMap["instagram"].(string); ok && instagram != "" {
				social.Instagram = &instagram
			}
			if youtube, ok := socialMap["youtube"].(string); ok && youtube != "" {
				social.Youtube = &youtube
			}
			profile.Social = social
		}
	}

	// 转换 Preferences (*datatypes.JSON) 为 map[string]string
	if v.Preferences != nil {
		prefsMap := make(map[string]interface{})
		if err := json.Unmarshal(*v.Preferences, &prefsMap); err == nil {
			profile.Preferences = convertMapToStringMap(prefsMap)
		}
	}

	// 转换 Skills (*datatypes.JSON) 为 map[string]int32
	if v.Skills != nil {
		skillsMap := make(map[string]interface{})
		if err := json.Unmarshal(*v.Skills, &skillsMap); err == nil {
			skillsProto := make(map[string]int32)
			for k, v := range skillsMap {
				if num, ok := v.(float64); ok {
					skillsProto[k] = int32(num)
				} else if num, ok := v.(int); ok {
					skillsProto[k] = int32(num)
				}
			}
			if len(skillsProto) > 0 {
				profile.Skills = skillsProto
			}
		}
	}

	// 转换 PrivacySettings (*datatypes.JSON) 为 map[string]string
	if v.PrivacySettings != nil {
		privacyMap := make(map[string]interface{})
		if err := json.Unmarshal(*v.PrivacySettings, &privacyMap); err == nil {
			profile.PrivacySettings = convertMapToStringMap(privacyMap)
		}
	}

	// Badges 字段会在 handler 中填充（如果需要）
	// 这里不填充，由 handler 负责调用 profileBadgeAppService 获取

	return profile
}

// formatDate 将 time.Time 格式化为日期字符串 (YYYY-MM-DD)
func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// convertMapToStringMap 将 map[string]interface{} 转换为 map[string]string
func convertMapToStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		if str, ok := v.(string); ok {
			result[k] = str
		} else {
			// 如果不是字符串，转换为字符串
			result[k] = ""
		}
	}
	return result
}
