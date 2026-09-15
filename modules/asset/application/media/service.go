package media

import (
	"context"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	assetrdb "nfxidentity/modules/asset/infrastructure/rdb"
	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

const (
	KindImages = "images"
	KindFiles  = "files"
	KindVideos = "videos"
	KindAudios = "audios"
)

type Service struct {
	db     *gorm.DB
	minio  *minio.Client
	bucket string
}

func NewService(db *gorm.DB, client *minio.Client, bucket string) *Service {
	return &Service{db: db, minio: client, bucket: bucket}
}

func (s *Service) EnsureBucket(ctx context.Context) error {
	ok, err := s.minio.BucketExists(ctx, s.bucket)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return s.minio.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{})
}

type PrepareResult struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
}

func (s *Service) Prepare(ctx context.Context, uploaderID uuid.UUID, kind, fileName, mimeType string) (*PrepareResult, error) {
	kind = NormalizeKind(kind)
	id := uuid.New()
	ext := path.Ext(fileName)
	if ext == "" {
		ext = ".bin"
	}
	key := fmt.Sprintf("%s/%s/%s%s", uploaderID.String(), kind, id.String(), ext)
	u, err := s.minio.PresignedPutObject(ctx, s.bucket, key, 15*time.Minute)
	if err != nil {
		return nil, errx.Internal("PRESIGN_FAILED", err.Error())
	}
	now := time.Now()
	if err := s.insertRow(ctx, kind, id.String(), key, fileName, mimeType, uploaderID.String(), now); err != nil {
		return nil, err
	}
	return &PrepareResult{ID: id.String(), UploadURL: u.String(), FilePath: key}, nil
}

func (s *Service) Confirm(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	row, err := s.getOwned(ctx, kind, id, uploaderID.String())
	if err != nil {
		return err
	}
	info, err := s.minio.StatObject(ctx, s.bucket, row.FilePath, minio.StatObjectOptions{})
	if err != nil {
		return errx.FailedPrecond("OBJECT_MISSING", "upload not found in object storage")
	}
	updates := map[string]any{"file_size": info.Size, "updated_at": time.Now()}
	if info.ContentType != "" {
		updates["mime_type"] = info.ContentType
	}
	return s.db.WithContext(ctx).Table(tableFor(kind)).Where("id = ?", id).Updates(updates).Error
}

func (s *Service) Delete(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	row, err := s.getOwned(ctx, kind, id, uploaderID.String())
	if err != nil {
		return err
	}
	_ = s.minio.RemoveObject(ctx, s.bucket, row.FilePath, minio.RemoveObjectOptions{})
	return s.db.WithContext(ctx).Table(tableFor(kind)).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (s *Service) Open(ctx context.Context, kind, id string) (io.ReadCloser, string, string, error) {
	kind = NormalizeKind(kind)
	row, err := s.getAny(ctx, kind, id)
	if err != nil {
		return nil, "", "", err
	}
	obj, err := s.minio.GetObject(ctx, s.bucket, row.FilePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", "", err
	}
	return obj, row.MimeType, row.FileName, nil
}

type ListItem struct {
	ID         string     `json:"id"`
	FilePath   string     `json:"file_path"`
	FileName   string     `json:"file_name"`
	FileSize   int64      `json:"file_size"`
	MimeType   string     `json:"mime_type"`
	UploaderID string     `json:"uploader_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (s *Service) Get(ctx context.Context, kind, id string) (*ListItem, error) {
	kind = NormalizeKind(kind)
	var row ListItem
	err := s.db.WithContext(ctx).Table(tableFor(kind)).
		Where("id = ? AND deleted_at IS NULL", id).
		Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func (s *Service) GetMany(ctx context.Context, kind string, ids []string) ([]ListItem, error) {
	kind = NormalizeKind(kind)
	if len(ids) == 0 {
		return nil, nil
	}
	var rows []ListItem
	err := s.db.WithContext(ctx).Table(tableFor(kind)).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Find(&rows).Error
	return rows, err
}

func (s *Service) List(ctx context.Context, uploaderID uuid.UUID, kind string) ([]ListItem, error) {
	kind = NormalizeKind(kind)
	var rows []ListItem
	err := s.db.WithContext(ctx).Table(tableFor(kind)).
		Where("uploader_id = ? AND deleted_at IS NULL", uploaderID.String()).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}

type rowView struct {
	FilePath string
	FileName string
	MimeType string
}

func (s *Service) insertRow(ctx context.Context, kind, id, filePath, fileName, mimeType, uploaderID string, now time.Time) error {
	switch kind {
	case KindImages:
		return s.db.WithContext(ctx).Create(&assetrdb.Image{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case KindFiles:
		return s.db.WithContext(ctx).Create(&assetrdb.File{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case KindVideos:
		return s.db.WithContext(ctx).Create(&assetrdb.Video{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	case KindAudios:
		return s.db.WithContext(ctx).Create(&assetrdb.Audio{
			ID: id, FilePath: filePath, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}).Error
	default:
		return errx.InvalidArg("INVALID_ASSET_KIND", "kind must be images, files, videos, or audios")
	}
}

func (s *Service) getOwned(ctx context.Context, kind, id, uploaderID string) (*rowView, error) {
	var row rowView
	err := s.db.WithContext(ctx).Table(tableFor(kind)).
		Select("file_path", "file_name", "mime_type").
		Where("id = ? AND uploader_id = ? AND deleted_at IS NULL", id, uploaderID).
		Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func (s *Service) getAny(ctx context.Context, kind, id string) (*rowView, error) {
	var row rowView
	err := s.db.WithContext(ctx).Table(tableFor(kind)).
		Select("file_path", "file_name", "mime_type").
		Where("id = ? AND deleted_at IS NULL", id).
		Take(&row).Error
	if err != nil {
		return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	}
	return &row, nil
}

func tableFor(kind string) string {
	switch kind {
	case KindImages:
		return "asset.Images"
	case KindFiles:
		return "asset.Files"
	case KindVideos:
		return "asset.Videos"
	case KindAudios:
		return "asset.Audios"
	default:
		return "asset.Images"
	}
}

func NormalizeKind(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	switch v {
	case KindImages, KindFiles, KindVideos, KindAudios:
		return v
	default:
		return KindImages
	}
}
