package reqdto

import (
	asseterrs "nfxidentity/errors/src/asset"
	imagesApp "nfxidentity/modules/asset/application/images"

	"github.com/google/uuid"
)

type PrepareImageUpload struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

func (r PrepareImageUpload) ToInput(accountID uuid.UUID) imagesApp.PrepareUploadInput {
	return imagesApp.PrepareUploadInput{
		AccountID: accountID,
		FileName:  r.FileName,
		MimeType:  r.MimeType,
	}
}

type PrepareImageUploadData struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	FilePath  string `json:"file_path"`
}

func PrepareImageUploadDataFrom(out *imagesApp.PrepareUploadOutput) PrepareImageUploadData {
	return PrepareImageUploadData{
		ID:        out.ImageID.String(),
		UploadURL: out.UploadURL,
		FilePath:  out.ObjectKey,
	}
}

type PrepareImageUploadItem struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

type PrepareImageUploads struct {
	Items []PrepareImageUploadItem `json:"items"`
}

func (r PrepareImageUploads) Validate() error {
	if len(r.Items) == 0 {
		return asseterrs.ErrImageRequestBodyInvalid
	}
	return nil
}

func (r PrepareImageUploads) ToInput(accountID uuid.UUID) imagesApp.PrepareUploadsInput {
	items := make([]imagesApp.PrepareUploadItemInput, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, imagesApp.PrepareUploadItemInput{FileName: item.FileName, MimeType: item.MimeType})
	}
	return imagesApp.PrepareUploadsInput{AccountID: accountID, Items: items}
}

func PrepareImageUploadsDataFrom(out *imagesApp.PrepareUploadsOutput) []PrepareImageUploadData {
	results := make([]PrepareImageUploadData, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, PrepareImageUploadDataFrom(&result))
	}
	return results
}

type ConfirmImageUpload struct {
	ID string `json:"id"`
}

func (r ConfirmImageUpload) ToInput(accountID uuid.UUID) (imagesApp.ConfirmUploadInput, error) {
	imageID, err := uuid.Parse(r.ID)
	if err != nil || imageID == uuid.Nil {
		return imagesApp.ConfirmUploadInput{}, asseterrs.ErrInvalidImageID.WithCause(err)
	}
	return imagesApp.ConfirmUploadInput{AccountID: accountID, ImageID: imageID}, nil
}

type ConfirmImageUploads struct {
	IDs []string `json:"ids"`
}

func (r ConfirmImageUploads) Validate() error {
	if len(r.IDs) == 0 {
		return asseterrs.ErrImageRequestBodyInvalid
	}
	return nil
}

func (r ConfirmImageUploads) ToInput(accountID uuid.UUID) (imagesApp.ConfirmUploadsInput, error) {
	ids := make([]uuid.UUID, 0, len(r.IDs))
	for _, raw := range r.IDs {
		id, err := uuid.Parse(raw)
		if err != nil || id == uuid.Nil {
			return imagesApp.ConfirmUploadsInput{}, asseterrs.ErrInvalidImageID.WithCause(err)
		}
		ids = append(ids, id)
	}
	return imagesApp.ConfirmUploadsInput{AccountID: accountID, ImageIDs: ids}, nil
}
