package views

import (
	imageTypeDomainViews "nebulaid/modules/image/domain/image_type/views"
	"time"
)

func ImageTypeViewMapper(domainView imageTypeDomainViews.ImageTypeView) ImageTypeView {
	description := ""
	if domainView.Description != nil {
		description = *domainView.Description
	}
	aspectRatio := ""
	if domainView.AspectRatio != nil {
		aspectRatio = *domainView.AspectRatio
	}

	return ImageTypeView{
		ID:          domainView.ID.String(),
		Key:         domainView.Key,
		Description: description,
		MaxWidth:    domainView.MaxWidth,
		MaxHeight:   domainView.MaxHeight,
		AspectRatio: aspectRatio,
		IsSystem:    domainView.IsSystem,
		CreatedAt:   domainView.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   domainView.UpdatedAt.Format(time.RFC3339),
	}
}

type ImageTypeView struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	MaxWidth    *int   `json:"max_width,omitempty"`
	MaxHeight   *int   `json:"max_height,omitempty"`
	AspectRatio string `json:"aspect_ratio,omitempty"`
	IsSystem    bool   `json:"is_system"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
