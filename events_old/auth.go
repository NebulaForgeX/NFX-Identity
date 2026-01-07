package events

// =============== Auth -> Auth Events (Internal) ===============

// AuthToAuth_SuccessEvent Auth 服务内部成功事件
type AuthToAuth_SuccessEvent struct {
	Operation string                 `json:"operation"` // "user.registered", "user.logged_in", "user.updated", "profile.updated" etc.
	EntityID  string                 `json:"entity_id"`
	UserID    string                 `json:"user_id"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

func (AuthToAuth_SuccessEvent) EventType() EventType { return ETAuthToAuth_Success }
func (AuthToAuth_SuccessEvent) TopicKey() TopicKey   { return TKAuth }

// AuthToAuth_TestEvent Auth 服务内部测试事件
type AuthToAuth_TestEvent struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (AuthToAuth_TestEvent) EventType() EventType { return ETAuthToAuth_Test }
func (AuthToAuth_TestEvent) TopicKey() TopicKey   { return TKAuth }

// AuthToAuth_UserCreatedEvent 用户创建事件（Auth 内部，用于通知其他服务）
type AuthToAuth_UserCreatedEvent struct {
	UserID   string                 `json:"user_id"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Phone    string                 `json:"phone"`
	Status   string                 `json:"status"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

func (AuthToAuth_UserCreatedEvent) EventType() EventType {
	return ETAuthToAuth_UserCreated
}
func (AuthToAuth_UserCreatedEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_UserDeletedEvent 用户删除事件（Auth 内部，用于通知其他服务）
type AuthToAuth_UserDeletedEvent struct {
	UserID   string                 `json:"user_id"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

func (AuthToAuth_UserDeletedEvent) EventType() EventType {
	return ETAuthToAuth_UserDeleted
}
func (AuthToAuth_UserDeletedEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_UserUpdatedEvent 用户更新事件（Auth 内部，用于通知其他服务）
type AuthToAuth_UserUpdatedEvent struct {
	UserID   string                 `json:"user_id"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Phone    string                 `json:"phone"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

func (AuthToAuth_UserUpdatedEvent) EventType() EventType {
	return ETAuthToAuth_UserUpdated
}
func (AuthToAuth_UserUpdatedEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_UserInvalidateCacheEvent 用户缓存清除事件（Auth 内部）
type AuthToAuth_UserInvalidateCacheEvent struct {
	UserID    string `json:"user_id"`   // 要清除的用户ID
	Operation string `json:"operation"` // 操作类型：updated, deleted, role_changed
}

func (AuthToAuth_UserInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_UserInvalidateCache
}
func (AuthToAuth_UserInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_ProfileInvalidateCacheEvent 用户资料缓存清除事件（Auth 内部）
type AuthToAuth_ProfileInvalidateCacheEvent struct {
	ProfileID string `json:"profile_id"` // 要清除的资料ID
	UserID    string `json:"user_id"`    // 关联的用户ID
	Operation string `json:"operation"`  // 操作类型：updated, deleted, avatar_updated, background_updated
}

func (AuthToAuth_ProfileInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_ProfileInvalidateCache
}
func (AuthToAuth_ProfileInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_RoleInvalidateCacheEvent 角色缓存清除事件（Auth 内部）
type AuthToAuth_RoleInvalidateCacheEvent struct {
	RoleID    string `json:"role_id"`   // 要清除的角色ID
	Operation string `json:"operation"` // 操作类型：updated, deleted, permissions_changed
}

func (AuthToAuth_RoleInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_RoleInvalidateCache
}
func (AuthToAuth_RoleInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_BadgeInvalidateCacheEvent 徽章缓存清除事件（Auth 内部）
type AuthToAuth_BadgeInvalidateCacheEvent struct {
	BadgeID   string `json:"badge_id"`  // 要清除的徽章ID
	Operation string `json:"operation"` // 操作类型：created, updated, deleted
}

func (AuthToAuth_BadgeInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_BadgeInvalidateCache
}
func (AuthToAuth_BadgeInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_EducationInvalidateCacheEvent 教育经历缓存清除事件（Auth 内部）
type AuthToAuth_EducationInvalidateCacheEvent struct {
	EducationID string `json:"education_id"` // 要清除的教育经历ID
	ProfileID   string `json:"profile_id"`   // 关联的资料ID
	Operation   string `json:"operation"`    // 操作类型：created, updated, deleted
}

func (AuthToAuth_EducationInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_EducationInvalidateCache
}
func (AuthToAuth_EducationInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_OccupationInvalidateCacheEvent 职业信息缓存清除事件（Auth 内部）
type AuthToAuth_OccupationInvalidateCacheEvent struct {
	OccupationID string `json:"occupation_id"` // 要清除的职业信息ID
	ProfileID    string `json:"profile_id"`    // 关联的资料ID
	Operation    string `json:"operation"`     // 操作类型：created, updated, deleted
}

func (AuthToAuth_OccupationInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_OccupationInvalidateCache
}
func (AuthToAuth_OccupationInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// AuthToAuth_ProfileBadgeInvalidateCacheEvent 用户徽章关联缓存清除事件（Auth 内部）
type AuthToAuth_ProfileBadgeInvalidateCacheEvent struct {
	ProfileBadgeID string `json:"profile_badge_id"` // 要清除的关联ID
	ProfileID      string `json:"profile_id"`       // 关联的资料ID
	BadgeID        string `json:"badge_id"`         // 关联的徽章ID
	Operation      string `json:"operation"`        // 操作类型：created, updated, deleted
}

func (AuthToAuth_ProfileBadgeInvalidateCacheEvent) EventType() EventType {
	return ETAuthToAuth_ProfileBadgeInvalidateCache
}
func (AuthToAuth_ProfileBadgeInvalidateCacheEvent) TopicKey() TopicKey { return TKAuth }

// =============== Auth -> Image Events ===============

// AuthToImage_ImageDeleteEvent Auth 服务请求 Image 服务删除图片
type AuthToImage_ImageDeleteEvent struct {
	ImagePath  string `json:"image_path"`  // 要删除的图片路径（相对路径或逗号分隔的多个路径）
	EntityID   string `json:"entity_id"`   // 关联的实体ID（user/profile/image ID）
	EntityType string `json:"entity_type"` // 实体类型：user, profile, image
	UserID     string `json:"user_id"`     // 用户ID（可选）
}

func (AuthToImage_ImageDeleteEvent) EventType() EventType { return ETAuthToImage_ImageDelete }
func (AuthToImage_ImageDeleteEvent) TopicKey() TopicKey   { return TKImage }

// AuthToImage_ImageSuccessEvent Auth 服务通知 Image 服务操作成功
type AuthToImage_ImageSuccessEvent struct {
	Operation string                 `json:"operation"`
	EntityID  string                 `json:"entity_id"`
	UserID    string                 `json:"user_id"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

func (AuthToImage_ImageSuccessEvent) EventType() EventType { return ETAuthToImage_ImageSuccess }
func (AuthToImage_ImageSuccessEvent) TopicKey() TopicKey   { return TKImage }

// AuthToImage_ImageTestEvent Auth 服务发送给 Image 服务的测试事件
type AuthToImage_ImageTestEvent struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (AuthToImage_ImageTestEvent) EventType() EventType { return ETAuthToImage_ImageTest }
func (AuthToImage_ImageTestEvent) TopicKey() TopicKey   { return TKImage }
