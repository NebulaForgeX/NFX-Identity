package commands

import (
	"github.com/google/uuid"
)

// CreateImageTagCmd 创建图片标签命令
type CreateImageTagCmd struct {
	ImageID    uuid.UUID
	Tag        string
	Confidence *float64
}

// UpdateImageTagCmd 更新图片标签命令
type UpdateImageTagCmd struct {
	ImageTagID uuid.UUID
	Tag        string
	Confidence *float64
}

// DeleteImageTagCmd 删除图片标签命令
type DeleteImageTagCmd struct {
	ImageTagID uuid.UUID
}

// DeleteImageTagByImageIDAndTagCmd 根据图片ID和标签删除图片标签命令
type DeleteImageTagByImageIDAndTagCmd struct {
	ImageID uuid.UUID
	Tag     string
}
