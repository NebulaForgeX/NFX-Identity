package audios

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	asseterrs "nfxidentity/errors/src/asset"
	audiosDomain "nfxidentity/modules/asset/domain/audios"

	"github.com/google/uuid"
)

type PrepareUploadInput struct {
	AccountID uuid.UUID
	FileName  string
	MimeType  string
}

type PrepareUploadOutput struct {
	AudioID   uuid.UUID
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
		fileName = "audio" + ext
	}
	mimeType := in.MimeType
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectKey := fmt.Sprintf("%s/audios/%s%s", in.AccountID.String(), id.String(), ext)
	uploadURL, err := s.storage.PresignPut(ctx, objectKey, 15*time.Minute)
	if err != nil {
		return nil, asseterrs.ErrPresignFailed.WithCause(err)
	}
	img, err := audiosDomain.NewAudio(audiosDomain.NewAudioParams{
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
	if err := s.audioRepo.Create.New(ctx, img); err != nil {
		return nil, err
	}
	return &PrepareUploadOutput{
		AudioID:   img.ID(),
		UploadURL: uploadURL.String(),
		ObjectKey: objectKey,
	}, nil
}
