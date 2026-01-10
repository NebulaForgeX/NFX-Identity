package image

import (
	"nfxid/events"
)

// ImageTagsInvalidateCacheEvent ImageTags 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.ImageTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type ImageTagsInvalidateCacheEvent struct {
	events.ImageTopic
	ID        string `json:"id"`         // 要清除的 Image Tag ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "image_tag"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ImageTypesInvalidateCacheEvent ImageTypes 缓存清除事件
type ImageTypesInvalidateCacheEvent struct {
	events.ImageTopic
	ID        string `json:"id"`         // 要清除的 Image Type ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "image_type"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ImageVariantsInvalidateCacheEvent ImageVariants 缓存清除事件
type ImageVariantsInvalidateCacheEvent struct {
	events.ImageTopic
	ID        string `json:"id"`         // 要清除的 Image Variant ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "image_variant"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ImagesInvalidateCacheEvent Images 缓存清除事件
type ImagesInvalidateCacheEvent struct {
	events.ImageTopic
	ID        string `json:"id"`         // 要清除的 Image ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "image"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
