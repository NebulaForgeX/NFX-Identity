package commands

import (
	"github.com/google/uuid"
)

// CreateImageCmd 创建图片命令
type CreateImageCmd struct {
	TypeID          *uuid.UUID
	UserID          *uuid.UUID
	TenantID        *uuid.UUID
	AppID           *uuid.UUID
	SourceDomain    *string
	Filename        string
	OriginalFilename string
	MimeType        string
	Size            int64
	Width           *int
	Height          *int
	StoragePath     string
	URL             *string
	IsPublic        bool
	Metadata        map[string]interface{}
}

// UpdateImageCmd 更新图片命令
type UpdateImageCmd struct {
	ImageID         uuid.UUID
	TypeID          *uuid.UUID
	UserID          *uuid.UUID
	TenantID        *uuid.UUID
	AppID           *uuid.UUID
	SourceDomain    *string
	Filename        *string
	OriginalFilename *string
	MimeType        *string
	Size            *int64
	Width           *int
	Height          *int
	StoragePath     *string
	URL             *string
	IsPublic        *bool
	Metadata        map[string]interface{}
}

// UpdateImageURLCmd 更新图片URL命令
type UpdateImageURLCmd struct {
	ImageID uuid.UUID
	URL     string
}

// UpdateImagePublicCmd 更新图片公开状态命令
type UpdateImagePublicCmd struct {
	ImageID  uuid.UUID
	IsPublic bool
}

// DeleteImageCmd 删除图片命令
type DeleteImageCmd struct {
	ImageID uuid.UUID
}
