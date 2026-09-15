package media

import (
	"context"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"nfxidentity/modules/asset/infrastructure/objectstore"
	"nfxidentity/modules/asset/infrastructure/repository/kinds"
	imagesQuery "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

const (
	KindImages = "images"
	KindFiles  = "files"
	KindVideos = "videos"
	KindAudios = "audios"
)

type Service struct {
	tx     transaction.TxManager
	store  *objectstore.Store
	rows   *kinds.Handler
	images *imagesQuery.Query
}

func NewService(tx transaction.TxManager, store *objectstore.Store, rows *kinds.Handler, images *imagesQuery.Query) *Service {
	return &Service{tx: tx, store: store, rows: rows, images: images}
}

func (s *Service) EnsureBucket(ctx context.Context) error {
	return s.store.EnsureBucket(ctx)
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
	u, err := s.store.PresignPut(ctx, key, 15*time.Minute)
	if err != nil {
		return nil, errx.Internal("PRESIGN_FAILED", err.Error())
	}
	now := time.Now()
	if err := s.rows.Insert(ctx, kind, id.String(), key, fileName, mimeType, uploaderID.String(), now); err != nil {
		return nil, err
	}
	return &PrepareResult{ID: id.String(), UploadURL: u.String(), FilePath: key}, nil
}

func (s *Service) Confirm(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	row, err := s.rows.GetOwned(ctx, kind, id, uploaderID.String())
	if err != nil {
		return err
	}
	info, err := s.store.Stat(ctx, row.FilePath)
	if err != nil {
		return errx.FailedPrecond("OBJECT_MISSING", "upload not found in object storage")
	}
	return s.rows.UpdateSize(ctx, kind, id, info.Size, info.ContentType)
}

func (s *Service) Delete(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	row, err := s.rows.GetOwned(ctx, kind, id, uploaderID.String())
	if err != nil {
		return err
	}
	_ = s.store.Remove(ctx, row.FilePath)
	return s.rows.SoftDelete(ctx, kind, id)
}

func (s *Service) Open(ctx context.Context, kind, id string) (io.ReadCloser, string, string, error) {
	kind = NormalizeKind(kind)
	row, err := s.rows.GetAny(ctx, kind, id)
	if err != nil {
		return nil, "", "", err
	}
	obj, err := s.store.Get(ctx, row.FilePath)
	if err != nil {
		return nil, "", "", err
	}
	return obj, row.MimeType, row.FileName, nil
}

type ListItem = kinds.ListItem

func (s *Service) Get(ctx context.Context, kind, id string) (*ListItem, error) {
	return s.rows.GetFromView(ctx, NormalizeKind(kind), id)
}

func (s *Service) GetMany(ctx context.Context, kind string, ids []string) ([]ListItem, error) {
	return s.rows.GetManyFromView(ctx, NormalizeKind(kind), ids)
}

func (s *Service) List(ctx context.Context, uploaderID uuid.UUID, kind string) ([]ListItem, error) {
	return s.rows.ListFromView(ctx, NormalizeKind(kind), uploaderID.String())
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
