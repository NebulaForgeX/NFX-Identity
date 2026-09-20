package reqdto

import (
	asseterrs "nfxidentity/errors/src/asset"
	videosApp "nfxidentity/modules/asset/application/videos"

	"github.com/google/uuid"
)

type PrepareVideoUpload struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

func (r PrepareVideoUpload) ToInput(accountID uuid.UUID) videosApp.PrepareUploadInput {
	return videosApp.PrepareUploadInput{
		AccountID: accountID,
		FileName:  r.FileName,
		MimeType:  r.MimeType,
	}
}

type PrepareVideoUploadData struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
}

func PrepareVideoUploadDataFrom(out *videosApp.PrepareUploadOutput) PrepareVideoUploadData {
	return PrepareVideoUploadData{
		ID:        out.VideoID.String(),
		UploadURL: out.UploadURL,
		FilePath:  out.ObjectKey,
	}
}

type PrepareVideoUploadItem struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

type PrepareVideoUploads struct {
	Items []PrepareVideoUploadItem `json:"items"`
}

func (r PrepareVideoUploads) Validate() error {
	if len(r.Items) == 0 {
		return asseterrs.ErrVideoRequestBodyInvalid
	}
	return nil
}

func (r PrepareVideoUploads) ToInput(accountID uuid.UUID) videosApp.PrepareUploadsInput {
	items := make([]videosApp.PrepareUploadItemInput, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, videosApp.PrepareUploadItemInput{FileName: item.FileName, MimeType: item.MimeType})
	}
	return videosApp.PrepareUploadsInput{AccountID: accountID, Items: items}
}

func PrepareVideoUploadsDataFrom(out *videosApp.PrepareUploadsOutput) []PrepareVideoUploadData {
	results := make([]PrepareVideoUploadData, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, PrepareVideoUploadDataFrom(&result))
	}
	return results
}

type ConfirmVideoUpload struct {
	ID string `json:"id"`
}

func (r ConfirmVideoUpload) ToInput(accountID uuid.UUID) (videosApp.ConfirmUploadInput, error) {
	videoID, err := uuid.Parse(r.ID)
	if err != nil || videoID == uuid.Nil {
		return videosApp.ConfirmUploadInput{}, asseterrs.ErrInvalidVideoID.WithCause(err)
	}
	return videosApp.ConfirmUploadInput{AccountID: accountID, VideoID: videoID}, nil
}

type ConfirmVideoUploads struct {
	IDs []string `json:"ids"`
}

func (r ConfirmVideoUploads) Validate() error {
	if len(r.IDs) == 0 {
		return asseterrs.ErrVideoRequestBodyInvalid
	}
	return nil
}

func (r ConfirmVideoUploads) ToInput(accountID uuid.UUID) (videosApp.ConfirmUploadsInput, error) {
	ids := make([]uuid.UUID, 0, len(r.IDs))
	for _, raw := range r.IDs {
		id, err := uuid.Parse(raw)
		if err != nil || id == uuid.Nil {
			return videosApp.ConfirmUploadsInput{}, asseterrs.ErrInvalidVideoID.WithCause(err)
		}
		ids = append(ids, id)
	}
	return videosApp.ConfirmUploadsInput{AccountID: accountID, VideoIDs: ids}, nil
}
