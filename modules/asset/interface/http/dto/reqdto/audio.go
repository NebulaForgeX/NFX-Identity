package reqdto

import (
	asseterrs "nfxidentity/errors/src/asset"
	audiosApp "nfxidentity/modules/asset/application/audios"

	"github.com/google/uuid"
)

type PrepareAudioUpload struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

func (r PrepareAudioUpload) ToInput(accountID uuid.UUID) audiosApp.PrepareUploadInput {
	return audiosApp.PrepareUploadInput{
		AccountID: accountID,
		FileName:  r.FileName,
		MimeType:  r.MimeType,
	}
}

type PrepareAudioUploadData struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
}

func PrepareAudioUploadDataFrom(out *audiosApp.PrepareUploadOutput) PrepareAudioUploadData {
	return PrepareAudioUploadData{
		ID:        out.AudioID.String(),
		UploadURL: out.UploadURL,
		FilePath:  out.ObjectKey,
	}
}

type PrepareAudioUploadItem struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

type PrepareAudioUploads struct {
	Items []PrepareAudioUploadItem `json:"items"`
}

func (r PrepareAudioUploads) Validate() error {
	if len(r.Items) == 0 {
		return asseterrs.ErrAudioRequestBodyInvalid
	}
	return nil
}

func (r PrepareAudioUploads) ToInput(accountID uuid.UUID) audiosApp.PrepareUploadsInput {
	items := make([]audiosApp.PrepareUploadItemInput, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, audiosApp.PrepareUploadItemInput{FileName: item.FileName, MimeType: item.MimeType})
	}
	return audiosApp.PrepareUploadsInput{AccountID: accountID, Items: items}
}

func PrepareAudioUploadsDataFrom(out *audiosApp.PrepareUploadsOutput) []PrepareAudioUploadData {
	results := make([]PrepareAudioUploadData, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, PrepareAudioUploadDataFrom(&result))
	}
	return results
}

type ConfirmAudioUpload struct {
	ID string `json:"id"`
}

func (r ConfirmAudioUpload) ToInput(accountID uuid.UUID) (audiosApp.ConfirmUploadInput, error) {
	audioID, err := uuid.Parse(r.ID)
	if err != nil || audioID == uuid.Nil {
		return audiosApp.ConfirmUploadInput{}, asseterrs.ErrInvalidAudioID.WithCause(err)
	}
	return audiosApp.ConfirmUploadInput{AccountID: accountID, AudioID: audioID}, nil
}

type ConfirmAudioUploads struct {
	IDs []string `json:"ids"`
}

func (r ConfirmAudioUploads) Validate() error {
	if len(r.IDs) == 0 {
		return asseterrs.ErrAudioRequestBodyInvalid
	}
	return nil
}

func (r ConfirmAudioUploads) ToInput(accountID uuid.UUID) (audiosApp.ConfirmUploadsInput, error) {
	ids := make([]uuid.UUID, 0, len(r.IDs))
	for _, raw := range r.IDs {
		id, err := uuid.Parse(raw)
		if err != nil || id == uuid.Nil {
			return audiosApp.ConfirmUploadsInput{}, asseterrs.ErrInvalidAudioID.WithCause(err)
		}
		ids = append(ids, id)
	}
	return audiosApp.ConfirmUploadsInput{AccountID: accountID, AudioIDs: ids}, nil
}
