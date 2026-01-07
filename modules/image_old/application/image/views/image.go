package views

import (
	imageDomainViews "nfxid/modules/image/domain/image/views"
	"time"
)

func ImageViewMapper(domainView imageDomainViews.ImageView) ImageView {
	typeID := ""
	if domainView.TypeID != nil {
		typeID = domainView.TypeID.String()
	}
	userID := ""
	if domainView.UserID != nil {
		userID = domainView.UserID.String()
	}
	sourceDomain := ""
	if domainView.SourceDomain != nil {
		sourceDomain = *domainView.SourceDomain
	}
	url := ""
	if domainView.URL != nil {
		url = *domainView.URL
	}

	return ImageView{
		ID:               domainView.ID.String(),
		TypeID:           typeID,
		UserID:           userID,
		SourceDomain:     sourceDomain,
		Filename:         domainView.Filename,
		OriginalFilename: domainView.OriginalFilename,
		MimeType:         domainView.MimeType,
		Size:             domainView.Size,
		Width:            domainView.Width,
		Height:           domainView.Height,
		StoragePath:      domainView.StoragePath,
		URL:              url,
		IsPublic:         domainView.IsPublic,
		Metadata:         domainView.Metadata,
		CreatedAt:        domainView.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        domainView.UpdatedAt.Format(time.RFC3339),
	}
}

type ImageView struct {
	ID               string                 `json:"id"`
	TypeID           string                 `json:"type_id,omitempty"`
	UserID           string                 `json:"user_id,omitempty"`
	SourceDomain     string                 `json:"source_domain,omitempty"`
	Filename         string                 `json:"filename"`
	OriginalFilename string                 `json:"original_filename"`
	MimeType         string                 `json:"mime_type"`
	Size             int64                  `json:"size"`
	Width            *int                   `json:"width,omitempty"`
	Height           *int                   `json:"height,omitempty"`
	StoragePath      string                 `json:"storage_path"`
	URL              string                 `json:"url,omitempty"`
	IsPublic         bool                   `json:"is_public"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt        string                 `json:"created_at"`
	UpdatedAt        string                 `json:"updated_at"`
}
