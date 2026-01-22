package respdto

import (
	"time"

	imageAppResult "nfxid/modules/image/application/images/results"

	"github.com/google/uuid"
)

type ImageDTO struct {
	ID              uuid.UUID              `json:"id"`
	TypeID          *uuid.UUID             `json:"type_id,omitempty"`
	UserID          *uuid.UUID             `json:"user_id,omitempty"`
	TenantID        *uuid.UUID             `json:"tenant_id,omitempty"`
	AppID           *uuid.UUID             `json:"app_id,omitempty"`
	SourceDomain    *string                `json:"source_domain,omitempty"`
	Filename        string                 `json:"filename"`
	OriginalFilename string                `json:"original_filename"`
	MimeType        string                 `json:"mime_type"`
	Size            int64                  `json:"size"`
	Width           *int                   `json:"width,omitempty"`
	Height          *int                   `json:"height,omitempty"`
	StoragePath     string                 `json:"storage_path"`
	URL             *string                `json:"url,omitempty"`
	IsPublic        bool                   `json:"is_public"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	DeletedAt       *time.Time             `json:"deleted_at,omitempty"`
}

// ImageROToDTO converts application ImageRO to response DTO
func ImageROToDTO(v *imageAppResult.ImageRO) *ImageDTO {
	if v == nil {
		return nil
	}

	return &ImageDTO{
		ID:              v.ID,
		TypeID:          v.TypeID,
		UserID:          v.UserID,
		TenantID:        v.TenantID,
		AppID:           v.AppID,
		SourceDomain:    v.SourceDomain,
		Filename:        v.Filename,
		OriginalFilename: v.OriginalFilename,
		MimeType:        v.MimeType,
		Size:            v.Size,
		Width:           v.Width,
		Height:          v.Height,
		StoragePath:     v.StoragePath,
		URL:             v.URL,
		IsPublic:        v.IsPublic,
		Metadata:        v.Metadata,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
		DeletedAt:       v.DeletedAt,
	}
}
