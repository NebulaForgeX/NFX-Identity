package results

import (
	"time"

	"nfxid/modules/image/domain/images"

	"github.com/google/uuid"
)

type ImageRO struct {
	ID              uuid.UUID
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
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

// ImageMapper 将 Domain Image 转换为 Application ImageRO
func ImageMapper(i *images.Image) ImageRO {
	if i == nil {
		return ImageRO{}
	}

	return ImageRO{
		ID:              i.ID(),
		TypeID:          i.TypeID(),
		UserID:          i.UserID(),
		TenantID:        i.TenantID(),
		AppID:           i.AppID(),
		SourceDomain:    i.SourceDomain(),
		Filename:        i.Filename(),
		OriginalFilename: i.OriginalFilename(),
		MimeType:        i.MimeType(),
		Size:            i.Size(),
		Width:           i.Width(),
		Height:          i.Height(),
		StoragePath:     i.StoragePath(),
		URL:             i.URL(),
		IsPublic:        i.IsPublic(),
		Metadata:        i.Metadata(),
		CreatedAt:       i.CreatedAt(),
		UpdatedAt:       i.UpdatedAt(),
		DeletedAt:       i.DeletedAt(),
	}
}
