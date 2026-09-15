package kinds

import (
	"context"
	"time"

	"nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/errx"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func New(db *gorm.DB) *Handler { return &Handler{db: db} }

type Row struct {
	FilePath string
	FileName string
	MimeType string
}

func (h *Handler) Insert(ctx context.Context, kind, id, filePath, fileName, mimeType, uploaderID string, now time.Time) error {
	switch kind {
	case "images":
		return h.db.WithContext(ctx).Create(&models.Image{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case "files":
		return h.db.WithContext(ctx).Create(&models.File{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case "videos":
		return h.db.WithContext(ctx).Create(&models.Video{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case "audios":
		return h.db.WithContext(ctx).Create(&models.Audio{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	default:
		return errx.InvalidArg("INVALID_ASSET_KIND", "kind must be images, files, videos, or audios")
	}
}

func tableFor(kind string) string {
	switch kind {
	case "files":
		return "asset.Files"
	case "videos":
		return "asset.Videos"
	case "audios":
		return "asset.Audios"
	default:
		return "asset.Images"
	}
}

func viewFor(kind string) string {
	switch kind {
	case "files":
		return "asset.FilesActiveView"
	case "videos":
		return "asset.VideosActiveView"
	case "audios":
		return "asset.AudiosActiveView"
	default:
		return "asset.ImagesActiveView"
	}
}

func (h *Handler) GetOwned(ctx context.Context, kind, id, uploaderID string) (*Row, error) {
	var row Row
	err := h.db.WithContext(ctx).Table(tableFor(kind)).
		Select("file_path", "file_name", "mime_type").
		Where("id = ? AND uploader_id = ? AND deleted_at IS NULL", id, uploaderID).
		Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func (h *Handler) GetAny(ctx context.Context, kind, id string) (*Row, error) {
	var row Row
	err := h.db.WithContext(ctx).Table(tableFor(kind)).
		Select("file_path", "file_name", "mime_type").
		Where("id = ? AND deleted_at IS NULL", id).
		Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func (h *Handler) UpdateSize(ctx context.Context, kind, id string, size int64, mime string) error {
	updates := map[string]any{"file_size": size, "updated_at": time.Now()}
	if mime != "" {
		updates["mime_type"] = mime
	}
	return h.db.WithContext(ctx).Table(tableFor(kind)).Where("id = ?", id).Updates(updates).Error
}

func (h *Handler) SoftDelete(ctx context.Context, kind, id string) error {
	return h.db.WithContext(ctx).Table(tableFor(kind)).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

type ListItem struct {
	ID         string    `json:"id"`
	FilePath   string    `json:"file_path"`
	FileName   string    `json:"file_name"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
	UploaderID string    `json:"uploader_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (h *Handler) GetFromView(ctx context.Context, kind, id string) (*ListItem, error) {
	var row ListItem
	err := h.db.WithContext(ctx).Table(viewFor(kind)).Where("id = ?", id).Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func (h *Handler) GetManyFromView(ctx context.Context, kind string, ids []string) ([]ListItem, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var rows []ListItem
	err := h.db.WithContext(ctx).Table(viewFor(kind)).Where("id IN ?", ids).Find(&rows).Error
	return rows, err
}

func (h *Handler) ListFromView(ctx context.Context, kind, uploaderID string) ([]ListItem, error) {
	var rows []ListItem
	err := h.db.WithContext(ctx).Table(viewFor(kind)).
		Where("uploader_id = ?", uploaderID).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}
