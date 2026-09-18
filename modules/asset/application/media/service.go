package media

import (
	"context"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"nfxidentity/errors/src/asset"
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	filesDomain "nfxidentity/modules/asset/domain/files"
	imagesDomain "nfxidentity/modules/asset/domain/images"
	videosDomain "nfxidentity/modules/asset/domain/videos"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	audiosQuery "nfxidentity/modules/asset/query/audios"
	filesQuery "nfxidentity/modules/asset/query/files"
	imagesQuery "nfxidentity/modules/asset/query/images"
	videosQuery "nfxidentity/modules/asset/query/videos"
	"nfxidentity/pkgs/query"
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
	images *imagesDomain.Repo
	files  *filesDomain.Repo
	videos *videosDomain.Repo
	audios *audiosDomain.Repo
	imageQ *imagesQuery.Query
	fileQ  *filesQuery.Query
	videoQ *videosQuery.Query
	audioQ *audiosQuery.Query
}

func NewService(
	tx transaction.TxManager,
	store *objectstore.Store,
	images *imagesDomain.Repo,
	files *filesDomain.Repo,
	videos *videosDomain.Repo,
	audios *audiosDomain.Repo,
	imageQ *imagesQuery.Query,
	fileQ *filesQuery.Query,
	videoQ *videosQuery.Query,
	audioQ *audiosQuery.Query,
) *Service {
	return &Service{
		tx: tx, store: store,
		images: images, files: files, videos: videos, audios: audios,
		imageQ: imageQ, fileQ: fileQ, videoQ: videoQ, audioQ: audioQ,
	}
}

func (s *Service) EnsureBucket(ctx context.Context) error {
	return s.store.EnsureBucket(ctx)
}

type PrepareResult struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
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
		return nil, asset.ErrPresignFailed.WithCause(err)
	}
	now := time.Now().UTC()
	switch kind {
	case KindFiles:
		err = s.files.Create.New(ctx, filesDomain.NewFileFromState(filesDomain.FileState{
			ID: id, FilePath: key, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}))
	case KindVideos:
		err = s.videos.Create.New(ctx, videosDomain.NewVideoFromState(videosDomain.VideoState{
			ID: id, FilePath: key, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}))
	case KindAudios:
		err = s.audios.Create.New(ctx, audiosDomain.NewAudioFromState(audiosDomain.AudioState{
			ID: id, FilePath: key, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}))
	default:
		err = s.images.Create.New(ctx, imagesDomain.NewImageFromState(imagesDomain.ImageState{
			ID: id, FilePath: key, FileName: fileName, MimeType: mimeType, UploaderID: uploaderID, CreatedAt: now, UpdatedAt: now,
		}))
	}
	if err != nil {
		return nil, err
	}
	return &PrepareResult{ID: id.String(), UploadURL: u.String(), FilePath: key}, nil
}

func (s *Service) Confirm(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	uid, err := uuid.Parse(id)
	if err != nil {
		return asset.ErrInvalidImageID
	}
	info, path, err := s.ownedPath(ctx, kind, uid, uploaderID)
	if err != nil {
		return err
	}
	st, err := s.store.Stat(ctx, path)
	if err != nil {
		return asset.ErrObjectMissing
	}
	_ = info
	return s.updateSize(ctx, kind, uid, st.Size, st.ContentType)
}

func (s *Service) Delete(ctx context.Context, uploaderID uuid.UUID, kind, id string) error {
	kind = NormalizeKind(kind)
	uid, err := uuid.Parse(id)
	if err != nil {
		return asset.ErrInvalidImageID
	}
	_, filePath, err := s.ownedPath(ctx, kind, uid, uploaderID)
	if err != nil {
		return err
	}
	_ = s.store.Remove(ctx, filePath)
	now := time.Now().UTC()
	switch kind {
	case KindFiles:
		row, err := s.files.Get.ByID(ctx, uid)
		if err != nil {
			return err
		}
		row.Delete()
		return s.files.Update.Generic(ctx, row)
	case KindVideos:
		row, err := s.videos.Get.ByID(ctx, uid)
		if err != nil {
			return err
		}
		row.Delete()
		return s.videos.Update.Generic(ctx, row)
	case KindAudios:
		row, err := s.audios.Get.ByID(ctx, uid)
		if err != nil {
			return err
		}
		row.Delete()
		return s.audios.Update.Generic(ctx, row)
	default:
		row, err := s.images.Get.ByID(ctx, uid)
		if err != nil {
			return err
		}
		_ = now
		row.Delete()
		return s.images.Update.Generic(ctx, row)
	}
}

func (s *Service) Open(ctx context.Context, kind, id string) (io.ReadCloser, string, string, error) {
	kind = NormalizeKind(kind)
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, "", "", asset.ErrInvalidImageID
	}
	item, err := s.Get(ctx, kind, uid.String())
	if err != nil {
		return nil, "", "", err
	}
	obj, err := s.store.Get(ctx, item.FilePath)
	if err != nil {
		return nil, "", "", err
	}
	return obj, item.MimeType, item.FileName, nil
}

func (s *Service) Get(ctx context.Context, kind, id string) (*ListItem, error) {
	kind = NormalizeKind(kind)
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, asset.ErrInvalidImageID
	}
	switch kind {
	case KindFiles:
		vo, err := s.fileQ.Single.ByID(ctx, uid)
		if err != nil {
			return nil, err
		}
		return fileItem(vo), nil
	case KindVideos:
		vo, err := s.videoQ.Single.ByID(ctx, uid)
		if err != nil {
			return nil, err
		}
		return videoItem(vo), nil
	case KindAudios:
		vo, err := s.audioQ.Single.ByID(ctx, uid)
		if err != nil {
			return nil, err
		}
		return audioItem(vo), nil
	default:
		vo, err := s.imageQ.Single.ByID(ctx, uid)
		if err != nil {
			return nil, err
		}
		return imageItem(vo), nil
	}
}

func (s *Service) GetMany(ctx context.Context, kind string, ids []string) ([]ListItem, error) {
	out := make([]ListItem, 0, len(ids))
	for _, id := range ids {
		item, err := s.Get(ctx, kind, id)
		if err != nil {
			continue
		}
		out = append(out, *item)
	}
	return out, nil
}

func (s *Service) List(ctx context.Context, uploaderID uuid.UUID, kind string) ([]ListItem, error) {
	kind = NormalizeKind(kind)
	switch kind {
	case KindFiles:
		page, err := s.fileQ.List.Generic(ctx, filesQuery.ListQuery{
			DomainPagination: query.DomainPagination{All: true},
			UploaderIDs:      []uuid.UUID{uploaderID},
		})
		if err != nil {
			return nil, err
		}
		out := make([]ListItem, 0, len(page.Items))
		for i := range page.Items {
			out = append(out, *fileItem(&page.Items[i]))
		}
		return out, nil
	case KindVideos:
		page, err := s.videoQ.List.Generic(ctx, videosQuery.ListQuery{
			DomainPagination: query.DomainPagination{All: true},
			UploaderIDs:      []uuid.UUID{uploaderID},
		})
		if err != nil {
			return nil, err
		}
		out := make([]ListItem, 0, len(page.Items))
		for i := range page.Items {
			out = append(out, *videoItem(&page.Items[i]))
		}
		return out, nil
	case KindAudios:
		page, err := s.audioQ.List.Generic(ctx, audiosQuery.ListQuery{
			DomainPagination: query.DomainPagination{All: true},
			UploaderIDs:      []uuid.UUID{uploaderID},
		})
		if err != nil {
			return nil, err
		}
		out := make([]ListItem, 0, len(page.Items))
		for i := range page.Items {
			out = append(out, *audioItem(&page.Items[i]))
		}
		return out, nil
	default:
		page, err := s.imageQ.List.Generic(ctx, imagesQuery.ListQuery{
			DomainPagination: query.DomainPagination{All: true},
			UploaderIDs:      []uuid.UUID{uploaderID},
		})
		if err != nil {
			return nil, err
		}
		out := make([]ListItem, 0, len(page.Items))
		for i := range page.Items {
			out = append(out, *imageItem(&page.Items[i]))
		}
		return out, nil
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

func (s *Service) ownedPath(ctx context.Context, kind string, id, uploaderID uuid.UUID) (any, string, error) {
	switch kind {
	case KindFiles:
		row, err := s.files.Get.ByID(ctx, id)
		if err != nil || row.UploaderID() != uploaderID {
			return nil, "", asset.ErrAssetNotFound
		}
		return row, row.FilePath(), nil
	case KindVideos:
		row, err := s.videos.Get.ByID(ctx, id)
		if err != nil || row.UploaderID() != uploaderID {
			return nil, "", asset.ErrAssetNotFound
		}
		return row, row.FilePath(), nil
	case KindAudios:
		row, err := s.audios.Get.ByID(ctx, id)
		if err != nil || row.UploaderID() != uploaderID {
			return nil, "", asset.ErrAssetNotFound
		}
		return row, row.FilePath(), nil
	default:
		row, err := s.images.Get.ByID(ctx, id)
		if err != nil || row.UploaderID() != uploaderID {
			return nil, "", asset.ErrAssetNotFound
		}
		return row, row.FilePath(), nil
	}
}

func (s *Service) updateSize(ctx context.Context, kind string, id uuid.UUID, size int64, mime string) error {
	switch kind {
	case KindFiles:
		row, err := s.files.Get.ByID(ctx, id)
		if err != nil {
			return err
		}
		if err := row.UpdateFileMeta("", size, mime); err != nil {
			return err
		}
		return s.files.Update.Generic(ctx, row)
	case KindVideos:
		row, err := s.videos.Get.ByID(ctx, id)
		if err != nil {
			return err
		}
		if err := row.UpdateFileMeta("", size, mime); err != nil {
			return err
		}
		return s.videos.Update.Generic(ctx, row)
	case KindAudios:
		row, err := s.audios.Get.ByID(ctx, id)
		if err != nil {
			return err
		}
		if err := row.UpdateFileMeta("", size, mime); err != nil {
			return err
		}
		return s.audios.Update.Generic(ctx, row)
	default:
		row, err := s.images.Get.ByID(ctx, id)
		if err != nil {
			return err
		}
		if err := row.UpdateFileMeta("", size, mime); err != nil {
			return err
		}
		return s.images.Update.Generic(ctx, row)
	}
}

func imageItem(vo *imagesQuery.ImageVO) *ListItem {
	return &ListItem{ID: vo.ID.String(), FilePath: vo.FilePath, FileName: vo.FileName, FileSize: vo.FileSize, MimeType: vo.MimeType, UploaderID: vo.UploaderID.String(), CreatedAt: vo.CreatedAt, UpdatedAt: vo.UpdatedAt}
}
func fileItem(vo *filesQuery.FileVO) *ListItem {
	return &ListItem{ID: vo.ID.String(), FilePath: vo.FilePath, FileName: vo.FileName, FileSize: vo.FileSize, MimeType: vo.MimeType, UploaderID: vo.UploaderID.String(), CreatedAt: vo.CreatedAt, UpdatedAt: vo.UpdatedAt}
}
func videoItem(vo *videosQuery.VideoVO) *ListItem {
	return &ListItem{ID: vo.ID.String(), FilePath: vo.FilePath, FileName: vo.FileName, FileSize: vo.FileSize, MimeType: vo.MimeType, UploaderID: vo.UploaderID.String(), CreatedAt: vo.CreatedAt, UpdatedAt: vo.UpdatedAt}
}
func audioItem(vo *audiosQuery.AudioVO) *ListItem {
	return &ListItem{ID: vo.ID.String(), FilePath: vo.FilePath, FileName: vo.FileName, FileSize: vo.FileSize, MimeType: vo.MimeType, UploaderID: vo.UploaderID.String(), CreatedAt: vo.CreatedAt, UpdatedAt: vo.UpdatedAt}
}
