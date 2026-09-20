package reqdto

import (
	asseterrs "nfxidentity/errors/src/asset"
	filesApp "nfxidentity/modules/asset/application/files"

	"github.com/google/uuid"
)

type PrepareFileUpload struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

func (r PrepareFileUpload) ToInput(accountID uuid.UUID) filesApp.PrepareUploadInput {
	return filesApp.PrepareUploadInput{
		AccountID: accountID,
		FileName:  r.FileName,
		MimeType:  r.MimeType,
	}
}

type PrepareFileUploadData struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
}

func PrepareFileUploadDataFrom(out *filesApp.PrepareUploadOutput) PrepareFileUploadData {
	return PrepareFileUploadData{
		ID:        out.FileID.String(),
		UploadURL: out.UploadURL,
		FilePath:  out.ObjectKey,
	}
}

type PrepareFileUploadItem struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

type PrepareFileUploads struct {
	Items []PrepareFileUploadItem `json:"items"`
}

func (r PrepareFileUploads) Validate() error {
	if len(r.Items) == 0 {
		return asseterrs.ErrFileRequestBodyInvalid
	}
	return nil
}

func (r PrepareFileUploads) ToInput(accountID uuid.UUID) filesApp.PrepareUploadsInput {
	items := make([]filesApp.PrepareUploadItemInput, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, filesApp.PrepareUploadItemInput{FileName: item.FileName, MimeType: item.MimeType})
	}
	return filesApp.PrepareUploadsInput{AccountID: accountID, Items: items}
}

func PrepareFileUploadsDataFrom(out *filesApp.PrepareUploadsOutput) []PrepareFileUploadData {
	results := make([]PrepareFileUploadData, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, PrepareFileUploadDataFrom(&result))
	}
	return results
}

type ConfirmFileUpload struct {
	ID string `json:"id"`
}

func (r ConfirmFileUpload) ToInput(accountID uuid.UUID) (filesApp.ConfirmUploadInput, error) {
	fileID, err := uuid.Parse(r.ID)
	if err != nil || fileID == uuid.Nil {
		return filesApp.ConfirmUploadInput{}, asseterrs.ErrInvalidFileID.WithCause(err)
	}
	return filesApp.ConfirmUploadInput{AccountID: accountID, FileID: fileID}, nil
}

type ConfirmFileUploads struct {
	IDs []string `json:"ids"`
}

func (r ConfirmFileUploads) Validate() error {
	if len(r.IDs) == 0 {
		return asseterrs.ErrFileRequestBodyInvalid
	}
	return nil
}

func (r ConfirmFileUploads) ToInput(accountID uuid.UUID) (filesApp.ConfirmUploadsInput, error) {
	ids := make([]uuid.UUID, 0, len(r.IDs))
	for _, raw := range r.IDs {
		id, err := uuid.Parse(raw)
		if err != nil || id == uuid.Nil {
			return filesApp.ConfirmUploadsInput{}, asseterrs.ErrInvalidFileID.WithCause(err)
		}
		ids = append(ids, id)
	}
	return filesApp.ConfirmUploadsInput{AccountID: accountID, FileIDs: ids}, nil
}
