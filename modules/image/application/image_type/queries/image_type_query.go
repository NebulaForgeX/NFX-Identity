package queries

import (
	"context"
	imageTypeDomainViews "nfxid/modules/image/domain/image_type/views"

	"github.com/google/uuid"
)

type ImageTypeQuery interface {
	GetByID(ctx context.Context, id uuid.UUID) (imageTypeDomainViews.ImageTypeView, error)
	GetByKey(ctx context.Context, key string) (imageTypeDomainViews.ImageTypeView, error)
	GetList(ctx context.Context, q ImageTypeListQuery) ([]imageTypeDomainViews.ImageTypeView, int64, error)
}
