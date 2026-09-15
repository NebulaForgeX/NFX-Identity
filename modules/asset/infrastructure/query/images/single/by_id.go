package single

import (
	"context"
	"errors"

	"nfxidentity/modules/asset/infrastructure/rdb/views"
	imagesQuery "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imagesQuery.Single { return &Handler{db: db} }

func toVO(r views.ImagesActiveView) imagesQuery.ImageVO {
	return imagesQuery.ImageVO{
		ID: r.ID, FilePath: r.FilePath, FileName: r.FileName, FileSize: r.FileSize,
		MimeType: r.MimeType, UploaderID: r.UploaderID, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt,
	}
}

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*imagesQuery.ImageVO, error) {
	var row views.ImagesActiveView
	if err := h.db.WithContext(ctx).Table(views.ImagesActiveView{}.TableName()).
		Where("id = ?", id).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
		}
		return nil, err
	}
	vo := toVO(row)
	return &vo, nil
}

func (h *Handler) ListByUploader(ctx context.Context, uploaderID uuid.UUID) ([]imagesQuery.ImageVO, error) {
	var rows []views.ImagesActiveView
	if err := h.db.WithContext(ctx).Table(views.ImagesActiveView{}.TableName()).
		Where("uploader_id = ?", uploaderID).Order("created_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]imagesQuery.ImageVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, nil
}
