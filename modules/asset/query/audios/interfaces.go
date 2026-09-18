package audios

import (
	"context"

	"nfxidentity/pkgs/httpx"

	"github.com/google/uuid"
)

type Single interface {
	ByID(ctx context.Context, id uuid.UUID) (*AudioVO, error)
}

type List interface {
	Generic(ctx context.Context, q ListQuery) (httpx.Page[AudioVO], error)
}

type Count interface {
	All(ctx context.Context) (int64, error)
	ByUploaderID(ctx context.Context, uploaderID uuid.UUID) (int64, error)
}
