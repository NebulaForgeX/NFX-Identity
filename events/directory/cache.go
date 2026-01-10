package directory

import (
	"nfxid/events"
)

// BadgesInvalidateCacheEvent Badges 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.DirectoryTopic 自动提供
type BadgesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 Badge ID
}

// UserBadgesInvalidateCacheEvent UserBadges 缓存清除事件
type UserBadgesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Badge ID
}

// UserEducationsInvalidateCacheEvent UserEducations 缓存清除事件
type UserEducationsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Education ID
}

// UserEmailsInvalidateCacheEvent UserEmails 缓存清除事件
type UserEmailsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Email ID
}

// UserOccupationsInvalidateCacheEvent UserOccupations 缓存清除事件
type UserOccupationsInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Occupation ID
}

// UserPhonesInvalidateCacheEvent UserPhones 缓存清除事件
type UserPhonesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Phone ID
}

// UserPreferencesInvalidateCacheEvent UserPreferences 缓存清除事件
type UserPreferencesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Preference ID
}

// UserProfilesInvalidateCacheEvent UserProfiles 缓存清除事件
type UserProfilesInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User Profile ID
}

// UsersInvalidateCacheEvent Users 缓存清除事件
type UsersInvalidateCacheEvent struct {
	events.DirectoryTopic
	ID string `json:"id"` // 要清除的 User ID
}
