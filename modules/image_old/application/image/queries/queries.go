package queries

import (
	"github.com/google/uuid"
)

type ImageListQuery struct {
	Page         int       `json:"page"`
	PageSize     int       `json:"page_size"`
	Search       string    `json:"search"`
	TypeID       *uuid.UUID `json:"type_id"`
	UserID       *uuid.UUID `json:"user_id"`
	SourceDomain *string   `json:"source_domain"`
	IsPublic     *bool     `json:"is_public"`
	OrderBy      string    `json:"order_by"` // "created_at", "updated_at", "size"
	Order        string    `json:"order"`    // "asc", "desc"
}

func (q *ImageListQuery) Normalize() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 20
	}
	if q.PageSize > 100 {
		q.PageSize = 100
	}
	if q.OrderBy == "" {
		q.OrderBy = "created_at"
	}
	if q.Order == "" {
		q.Order = "desc"
	}
}

func (q ImageListQuery) CacheKey() string {
	return "image:list"
}

type CommonQuery struct {
	Page     int
	PageSize int
	Search   string
	OrderBy  string
	Order    string
}

