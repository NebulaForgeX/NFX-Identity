package commands

import (
	"github.com/google/uuid"
)

// CreateImageTypeCmd 创建图片类型命令
type CreateImageTypeCmd struct {
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
}

// UpdateImageTypeCmd 更新图片类型命令
type UpdateImageTypeCmd struct {
	ImageTypeID uuid.UUID
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
}

// DeleteImageTypeCmd 删除图片类型命令
type DeleteImageTypeCmd struct {
	ImageTypeID uuid.UUID
}
