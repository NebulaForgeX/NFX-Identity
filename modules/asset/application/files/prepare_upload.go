package files

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	asseterrs "nfxidentity/errors/src/asset"
	filesDomain "nfxidentity/modules/asset/domain/files"

	"github.com/google/uuid"
)

type PrepareUploadInput struct {
	AccountID uuid.UUID
	FileName  string
	MimeType  string
}

type PrepareUploadOutput struct {
	FileID    uuid.UUID
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
		fileName = "file" + ext
	}
	mimeType := in.MimeType
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectKey := fmt.Sprintf("%s/files/%s%s", in.AccountID.String(), id.String(), ext)
	uploadURL, err := s.storage.PresignPut(ctx, objectKey, 15*time.Minute)
	if err != nil {
		return nil, asseterrs.ErrPresignFailed.WithCause(err)
	}
	img, err := filesDomain.NewFile(filesDomain.NewFileParams{
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
	if err := s.fileRepo.Create.New(ctx, img); err != nil {
		return nil, err
	}
	return &PrepareUploadOutput{
		FileID:    img.ID(),
		UploadURL: uploadURL.String(),
		ObjectKey: objectKey,
	}, nil
}
