package reqdto

import (
	imageAppCommands "nfxid/modules/image/application/images/commands"

	"github.com/google/uuid"
)

type ImageCreateRequestDTO struct {
	TypeID           *uuid.UUID             `json:"type_id,omitempty"`
	UserID           *uuid.UUID             `json:"user_id,omitempty"`
	TenantID         *uuid.UUID             `json:"tenant_id,omitempty"`
	AppID            *uuid.UUID             `json:"app_id,omitempty"`
	SourceDomain     *string                `json:"source_domain,omitempty"`
	Filename         string                 `json:"filename" validate:"required"`
	OriginalFilename string                 `json:"original_filename" validate:"required"`
	MimeType         string                 `json:"mime_type" validate:"required"`
	Size             int64                  `json:"size" validate:"required,min=1"`
	Width            *int                   `json:"width,omitempty"`
	Height           *int                   `json:"height,omitempty"`
	StoragePath      string                 `json:"storage_path" validate:"required"`
	URL              *string                `json:"url,omitempty"`
	IsPublic         bool                   `json:"is_public"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

type ImageByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type ImageUpdateRequestDTO struct {
	TypeID           *uuid.UUID             `json:"type_id,omitempty"`
	UserID           *uuid.UUID             `json:"user_id,omitempty"`
	TenantID         *uuid.UUID             `json:"tenant_id,omitempty"`
	AppID            *uuid.UUID             `json:"app_id,omitempty"`
	SourceDomain     *string                `json:"source_domain,omitempty"`
	Filename         *string                `json:"filename,omitempty"`
	OriginalFilename *string                `json:"original_filename,omitempty"`
	MimeType         *string                `json:"mime_type,omitempty"`
	Size             *int64                 `json:"size,omitempty"`
	Width            *int                   `json:"width,omitempty"`
	Height           *int                   `json:"height,omitempty"`
	StoragePath      *string                `json:"storage_path,omitempty"`
	URL              *string                `json:"url,omitempty"`
	IsPublic         *bool                  `json:"is_public,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

type ImageUpdateURLRequestDTO struct {
	URL string `json:"url" validate:"required"`
}

type ImageUpdatePublicRequestDTO struct {
	IsPublic bool `json:"is_public" validate:"required"`
}

func (r *ImageCreateRequestDTO) ToCreateCmd() imageAppCommands.CreateImageCmd {
	return imageAppCommands.CreateImageCmd{
		TypeID:           r.TypeID,
		UserID:           r.UserID,
		TenantID:         r.TenantID,
		AppID:            r.AppID,
		SourceDomain:     r.SourceDomain,
		Filename:         r.Filename,
		OriginalFilename: r.OriginalFilename,
		MimeType:         r.MimeType,
		Size:             r.Size,
		Width:            r.Width,
		Height:           r.Height,
		StoragePath:      r.StoragePath,
		URL:              r.URL,
		IsPublic:         r.IsPublic,
		Metadata:         r.Metadata,
	}
}

func (r *ImageUpdateRequestDTO) ToUpdateCmd(imageID uuid.UUID) imageAppCommands.UpdateImageCmd {
	return imageAppCommands.UpdateImageCmd{
		ImageID:          imageID,
		TypeID:           r.TypeID,
		UserID:           r.UserID,
		TenantID:         r.TenantID,
		AppID:            r.AppID,
		SourceDomain:     r.SourceDomain,
		Filename:         r.Filename,
		OriginalFilename: r.OriginalFilename,
		MimeType:         r.MimeType,
		Size:             r.Size,
		Width:            r.Width,
		Height:           r.Height,
		StoragePath:      r.StoragePath,
		URL:              r.URL,
		IsPublic:         r.IsPublic,
		Metadata:         r.Metadata,
	}
}
