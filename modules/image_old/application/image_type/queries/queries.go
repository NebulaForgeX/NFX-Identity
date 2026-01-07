package queries

type ImageTypeListQuery struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Search   string `json:"search"`
	IsSystem *bool  `json:"is_system"`
	OrderBy  string `json:"order_by"` // "created_at", "key"
	Order    string `json:"order"`   // "asc", "desc"
}

func (q *ImageTypeListQuery) Normalize() {
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

func (q ImageTypeListQuery) CacheKey() string {
	return "image_type:list"
}

