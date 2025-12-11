package queries

import (
	"context"
	imageDomainViews "nfxid/modules/image/domain/image/views"

	"github.com/google/uuid"
)

type ImageQuery interface {
	GetByID(ctx context.Context, id uuid.UUID) (imageDomainViews.ImageView, error)
	GetList(ctx context.Context, q ImageListQuery) ([]imageDomainViews.ImageView, int64, error)
}
