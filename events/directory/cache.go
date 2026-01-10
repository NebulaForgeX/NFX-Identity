package directory

import (
	"nfxid/events"
)

// BadgesInvalidateCacheEvent Badges 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.DirectoryTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type BadgesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 Badge ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "badge"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserBadgesInvalidateCacheEvent UserBadges 缓存清除事件
type UserBadgesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Badge ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_badge"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserEducationsInvalidateCacheEvent UserEducations 缓存清除事件
type UserEducationsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Education ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_education"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserEmailsInvalidateCacheEvent UserEmails 缓存清除事件
type UserEmailsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Email ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_email"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserOccupationsInvalidateCacheEvent UserOccupations 缓存清除事件
type UserOccupationsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Occupation ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_occupation"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserPhonesInvalidateCacheEvent UserPhones 缓存清除事件
type UserPhonesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Phone ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_phone"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserPreferencesInvalidateCacheEvent UserPreferences 缓存清除事件
type UserPreferencesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Preference ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_preference"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserProfilesInvalidateCacheEvent UserProfiles 缓存清除事件
type UserProfilesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User Profile ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_profile"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UsersInvalidateCacheEvent Users 缓存清除事件
type UsersInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID        string `json:"id"`         // 要清除的 User ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
