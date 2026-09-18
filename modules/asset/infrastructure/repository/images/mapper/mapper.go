package mapper

import (
	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"
)

func ToModel(img *images.Image) *models.Image {
	st := img.State()
	return &models.Image{
		ID:         st.ID,
		FilePath:   st.FilePath,
		FileName:   st.FileName,
		FileSize:   st.FileSize,
		MimeType:   st.MimeType,
		UploaderID: st.UploaderID,
		CreatedAt:  st.CreatedAt,
		UpdatedAt:  st.UpdatedAt,
		DeletedAt:  timex.TimeToGormDeletedAt(st.DeletedAt),
	}
}

func ToDomain(m *models.Image) *images.Image {
	return images.NewFromState(images.ImageState{
		ID:         m.ID,
		FilePath:   m.FilePath,
		FileName:   m.FileName,
		FileSize:   m.FileSize,
		MimeType:   m.MimeType,
		UploaderID: m.UploaderID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  timex.GormDeletedAtToTime(m.DeletedAt),
	})
}
