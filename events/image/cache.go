package image

import (
	"nfxid/events"
)

// ImageTagsInvalidateCacheEvent ImageTags 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.ImageTopic 自动提供
type ImageTagsInvalidateCacheEvent struct {
	events.ImageTopic
	ID string `json:"id"` // 要清除的 Image Tag ID
}

// ImageTypesInvalidateCacheEvent ImageTypes 缓存清除事件
type ImageTypesInvalidateCacheEvent struct {
	events.ImageTopic
	ID string `json:"id"` // 要清除的 Image Type ID
}

// ImageVariantsInvalidateCacheEvent ImageVariants 缓存清除事件
type ImageVariantsInvalidateCacheEvent struct {
	events.ImageTopic
	ID string `json:"id"` // 要清除的 Image Variant ID
}

// ImagesInvalidateCacheEvent Images 缓存清除事件
type ImagesInvalidateCacheEvent struct {
	events.ImageTopic
	ID string `json:"id"` // 要清除的 Image ID
}
