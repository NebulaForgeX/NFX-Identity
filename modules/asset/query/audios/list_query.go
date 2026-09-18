package audios

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	IDs         []uuid.UUID `json:"ids,omitempty"`
	UploaderIDs []uuid.UUID `json:"uploader_ids,omitempty"`
	MimeTypes   []string    `json:"mime_types,omitempty"`
	Search      *string     `json:"search,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedAt, Order: "desc"},
	)
}

func (q ListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
