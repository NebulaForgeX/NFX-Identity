package single

import (
	"context"
	"errors"
	"nfxidentity/errors/src/asset"

	"nfxidentity/modules/asset/infrastructure/rdb/views"
	imagesQuery "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/utils/ptr"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imagesQuery.Single { return &Handler{db: db} }

func toVO(r views.Imagesactiveview) imagesQuery.ImageVO {
	return imagesQuery.ImageVO{
		ID: ptr.Deref(r.ID), FilePath: ptr.Deref(r.FilePath), FileName: ptr.Deref(r.FileName),
		FileSize: ptr.Deref(r.FileSize), MimeType: ptr.Deref(r.MimeType), UploaderID: ptr.Deref(r.UploaderID),
		CreatedAt: ptr.Deref(r.CreatedAt), UpdatedAt: ptr.Deref(r.UpdatedAt),
	}
}

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*imagesQuery.ImageVO, error) {
	var row views.Imagesactiveview
	if err := h.db.WithContext(ctx).Table(views.Imagesactiveview{}.TableName()).
		Where("id = ?", id).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asset.ErrAssetNotFound
		}
		return nil, err
	}
	vo := toVO(row)
	return &vo, nil
}

func (h *Handler) ListByUploader(ctx context.Context, uploaderID uuid.UUID) ([]imagesQuery.ImageVO, error) {
	var rows []views.Imagesactiveview
	if err := h.db.WithContext(ctx).Table(views.Imagesactiveview{}.TableName()).
		Where("uploader_id = ?", uploaderID).Order("created_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]imagesQuery.ImageVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, nil
}
