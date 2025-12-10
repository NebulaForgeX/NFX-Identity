package events

// =============== Image -> Auth Events ===============

// ImageToAuth_ImageDeleteEvent Image 服务通知 Auth 服务图片已删除
type ImageToAuth_ImageDeleteEvent struct {
	ImagePath  string `json:"image_path"`  // 已删除的图片路径（相对路径或逗号分隔的多个路径）
	EntityID   string `json:"entity_id"`   // 关联的实体ID（user/profile/image ID）
	EntityType string `json:"entity_type"` // 实体类型：user, profile, image
	UserID     string `json:"user_id"`     // 用户ID（可选）
}

func (ImageToAuth_ImageDeleteEvent) EventType() EventType { return ETImageToAuth_ImageDelete }
func (ImageToAuth_ImageDeleteEvent) TopicKey() TopicKey   { return TKAuth }

// ImageToAuth_ImageSuccessEvent Image 服务通知 Auth 服务操作成功
type ImageToAuth_ImageSuccessEvent struct {
	Operation string                 `json:"operation"`
	EntityID  string                 `json:"entity_id"`
	UserID    string                 `json:"user_id"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

func (ImageToAuth_ImageSuccessEvent) EventType() EventType { return ETImageToAuth_ImageSuccess }
func (ImageToAuth_ImageSuccessEvent) TopicKey() TopicKey   { return TKAuth }

// ImageToAuth_ImageTestEvent Image 服务发送给 Auth 服务的测试事件
type ImageToAuth_ImageTestEvent struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (ImageToAuth_ImageTestEvent) EventType() EventType { return ETImageToAuth_ImageTest }
func (ImageToAuth_ImageTestEvent) TopicKey() TopicKey   { return TKAuth }
