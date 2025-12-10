package reqdto

import (
	imageApp "nebulaid/modules/image/application/image"
	imageAppQueries "nebulaid/modules/image/application/image/queries"

	"github.com/google/uuid"
)

type ImageCreateRequestDTO struct {
	TypeID           *string                `json:"type_id,omitempty"`
	UserID           *string                `json:"user_id,omitempty"`
	SourceDomain     *string                `json:"source_domain,omitempty"`
	Filename         string                 `json:"filename" validate:"required"`
	OriginalFilename string                 `json:"original_filename" validate:"required"`
	MimeType         string                 `json:"mime_type" validate:"required"`
	Size             int64                  `json:"size" validate:"required,gt=0"`
	Width            *int                   `json:"width,omitempty"`
	Height           *int                   `json:"height,omitempty"`
	StoragePath      string                 `json:"storage_path" validate:"required"`
	URL              *string                `json:"url,omitempty"`
	IsPublic         bool                   `json:"is_public"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

type ImageUpdateRequestDTO struct {
	ID               uuid.UUID              `params:"id" validate:"required,uuid"`
	TypeID           *string                `json:"type_id,omitempty"`
	UserID           *string                `json:"user_id,omitempty"`
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

type ImageByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ImageQueryParamsDTO struct {
	Page     *int    `query:"page"`
	PageSize *int    `query:"page_size"`
	Search   *string `query:"search"`
	TypeID   *string `query:"type_id"`
	UserID   *string `query:"user_id"`
	IsPublic *bool   `query:"is_public"`
	OrderBy  *string `query:"order_by"`
	Order    *string `query:"order"`
}

func (r *ImageCreateRequestDTO) ToCreateCmd() imageApp.CreateImageCmd {
	var typeID, userID *uuid.UUID
	if r.TypeID != nil {
		if id, err := uuid.Parse(*r.TypeID); err == nil {
			typeID = &id
		}
	}
	if r.UserID != nil {
		if id, err := uuid.Parse(*r.UserID); err == nil {
			userID = &id
		}
	}

	return imageApp.CreateImageCmd{
		TypeID:           typeID,
		UserID:           userID,
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

func (r *ImageUpdateRequestDTO) ToUpdateCmd() imageApp.UpdateImageCmd {
	var typeID, userID *uuid.UUID
	if r.TypeID != nil {
		if id, err := uuid.Parse(*r.TypeID); err == nil {
			typeID = &id
		}
	}
	if r.UserID != nil {
		if id, err := uuid.Parse(*r.UserID); err == nil {
			userID = &id
		}
	}

	cmd := imageApp.UpdateImageCmd{
		ID:           r.ID,
		TypeID:       typeID,
		UserID:       userID,
		SourceDomain: r.SourceDomain,
		Width:        r.Width,
		Height:       r.Height,
		URL:          r.URL,
		Metadata:     r.Metadata,
	}

	if r.MimeType != nil {
		cmd.MimeType = *r.MimeType
	}
	if r.IsPublic != nil {
		cmd.IsPublic = *r.IsPublic
	}

	if r.Filename != nil {
		cmd.Filename = *r.Filename
	}
	if r.OriginalFilename != nil {
		cmd.OriginalFilename = *r.OriginalFilename
	}
	if r.Size != nil {
		cmd.Size = *r.Size
	}
	if r.StoragePath != nil {
		cmd.StoragePath = *r.StoragePath
	}

	return cmd
}

func (q *ImageQueryParamsDTO) ToListQuery() imageAppQueries.ImageListQuery {
	query := imageAppQueries.ImageListQuery{}

	if q.Page != nil {
		query.Page = *q.Page
	}
	if q.PageSize != nil {
		query.PageSize = *q.PageSize
	}
	if q.Search != nil {
		query.Search = *q.Search
	}
	if q.TypeID != nil {
		if id, err := uuid.Parse(*q.TypeID); err == nil {
			query.TypeID = &id
		}
	}
	if q.UserID != nil {
		if id, err := uuid.Parse(*q.UserID); err == nil {
			query.UserID = &id
		}
	}
	if q.IsPublic != nil {
		query.IsPublic = q.IsPublic
	}
	if q.OrderBy != nil {
		query.OrderBy = *q.OrderBy
	}
	if q.Order != nil {
		query.Order = *q.Order
	}

	return query
}
