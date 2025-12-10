package reqdto

import (
	imageTypeApp "nebulaid/modules/image/application/image_type"
	imageTypeAppQueries "nebulaid/modules/image/application/image_type/queries"

	"github.com/google/uuid"
)

type ImageTypeCreateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Description *string `json:"description,omitempty"`
	MaxWidth    *int    `json:"max_width,omitempty"`
	MaxHeight   *int    `json:"max_height,omitempty"`
	AspectRatio *string `json:"aspect_ratio,omitempty"`
	IsSystem    bool    `json:"is_system"`
}

type ImageTypeUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Key         *string   `json:"key,omitempty"`
	Description *string   `json:"description,omitempty"`
	MaxWidth    *int      `json:"max_width,omitempty"`
	MaxHeight   *int      `json:"max_height,omitempty"`
	AspectRatio *string   `json:"aspect_ratio,omitempty"`
	IsSystem    *bool     `json:"is_system,omitempty"`
}

type ImageTypeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ImageTypeByKeyRequestDTO struct {
	Key string `params:"key" validate:"required"`
}

type ImageTypeQueryParamsDTO struct {
	Page     *int    `query:"page"`
	PageSize *int    `query:"page_size"`
	Search   *string `query:"search"`
	IsSystem *bool   `query:"is_system"`
	OrderBy  *string `query:"order_by"`
	Order    *string `query:"order"`
}

func (r *ImageTypeCreateRequestDTO) ToCreateCmd() imageTypeApp.CreateImageTypeCmd {
	return imageTypeApp.CreateImageTypeCmd{
		Key:         r.Key,
		Description: r.Description,
		MaxWidth:    r.MaxWidth,
		MaxHeight:   r.MaxHeight,
		AspectRatio: r.AspectRatio,
		IsSystem:    r.IsSystem,
	}
}

func (r *ImageTypeUpdateRequestDTO) ToUpdateCmd() imageTypeApp.UpdateImageTypeCmd {
	cmd := imageTypeApp.UpdateImageTypeCmd{
		ID:          r.ID,
		Description: r.Description,
		MaxWidth:    r.MaxWidth,
		MaxHeight:   r.MaxHeight,
		AspectRatio: r.AspectRatio,
	}

	if r.Key != nil {
		cmd.Key = *r.Key
	}
	if r.IsSystem != nil {
		cmd.IsSystem = *r.IsSystem
	}

	return cmd
}

func (q *ImageTypeQueryParamsDTO) ToListQuery() imageTypeAppQueries.ImageTypeListQuery {
	query := imageTypeAppQueries.ImageTypeListQuery{}

	if q.Page != nil {
		query.Page = *q.Page
	}
	if q.PageSize != nil {
		query.PageSize = *q.PageSize
	}
	if q.Search != nil {
		query.Search = *q.Search
	}
	if q.IsSystem != nil {
		query.IsSystem = q.IsSystem
	}
	if q.OrderBy != nil {
		query.OrderBy = *q.OrderBy
	}
	if q.Order != nil {
		query.Order = *q.Order
	}

	return query
}
