package images

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	asseterrs "nfxidentity/errors/src/asset"
	imagesDomain "nfxidentity/modules/asset/domain/images"

	"github.com/google/uuid"
)

type PrepareUploadInput struct {
	AccountID uuid.UUID
	FileName  string
	MimeType  string
}

type PrepareUploadOutput struct {
	ImageID   uuid.UUID
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
		fileName = "image" + ext
	}
	mimeType := in.MimeType
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectKey := fmt.Sprintf("%s/images/%s%s", in.AccountID.String(), id.String(), ext)
	uploadURL, err := s.storage.PresignPut(ctx, objectKey, 15*time.Minute)
	if err != nil {
		return nil, asseterrs.ErrPresignFailed.WithCause(err)
	}
	img, err := imagesDomain.NewImage(imagesDomain.NewImageParams{
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
	if err := s.imageRepo.Create.New(ctx, img); err != nil {
		return nil, err
	}
	return &PrepareUploadOutput{
		ImageID:   img.ID(),
		UploadURL: uploadURL.String(),
		ObjectKey: objectKey,
	}, nil
}
