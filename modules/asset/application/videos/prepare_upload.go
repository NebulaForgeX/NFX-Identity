package videos

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	asseterrs "nfxidentity/errors/src/asset"
	videosDomain "nfxidentity/modules/asset/domain/videos"

	"github.com/google/uuid"
)

type PrepareUploadInput struct {
	AccountID uuid.UUID
	FileName  string
	MimeType  string
}

type PrepareUploadOutput struct {
	VideoID   uuid.UUID
	UploadURL string
	ObjectKey string
}

func (s *Service) PrepareUpload(ctx context.Context, in PrepareUploadInput) (*PrepareUploadOutput, error) {
	ext := filepath.Ext(in.FileName)
	if ext == "" {
		ext = ".bin"
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	fileName := in.FileName
	if fileName == "" {
		fileName = "video" + ext
	}
	mimeType := in.MimeType
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectKey := fmt.Sprintf("%s/videos/%s%s", in.AccountID.String(), id.String(), ext)
	uploadURL, err := s.storage.PresignPut(ctx, objectKey, 15*time.Minute)
	if err != nil {
		return nil, asseterrs.ErrPresignFailed.WithCause(err)
	}
	img, err := videosDomain.NewVideo(videosDomain.NewVideoParams{
		ID:         &id,
		FilePath:   objectKey,
		FileName:   fileName,
		FileSize:   1,
		MimeType:   mimeType,
		UploaderID: in.AccountID,
	})
	if err != nil {
		return nil, err
	}
	if err := s.videoRepo.Create.New(ctx, img); err != nil {
		return nil, err
	}
	return &PrepareUploadOutput{
		VideoID:   img.ID(),
		UploadURL: uploadURL.String(),
		ObjectKey: objectKey,
	}, nil
}
