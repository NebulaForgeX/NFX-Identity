package commands

import (
	"github.com/google/uuid"
)

// CreateImageVariantCmd 创建图片变体命令
type CreateImageVariantCmd struct {
	ImageID     uuid.UUID
	VariantKey  string
	Width       *int
	Height      *int
	Size        *int64
	MimeType    *string
	StoragePath string
	URL         *string
}

// UpdateImageVariantCmd 更新图片变体命令
type UpdateImageVariantCmd struct {
	ImageVariantID uuid.UUID
	Width          *int
	Height         *int
	Size           *int64
	MimeType       *string
	StoragePath    *string
	URL            *string
}

// UpdateImageVariantURLCmd 更新图片变体URL命令
type UpdateImageVariantURLCmd struct {
	ImageVariantID uuid.UUID
	URL            string
}

// DeleteImageVariantCmd 删除图片变体命令
type DeleteImageVariantCmd struct {
	ImageVariantID uuid.UUID
}

// DeleteImageVariantByImageIDAndVariantKeyCmd 根据图片ID和变体Key删除图片变体命令
type DeleteImageVariantByImageIDAndVariantKeyCmd struct {
	ImageID    uuid.UUID
	VariantKey string
}
