package mapper

import (
	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func ToModel(img *images.Image) *models.Image {
	st := img.State()
	return &models.Image{
		ID: st.ID.String(), FilePath: st.FilePath, FileName: st.FileName, FileSize: st.FileSize,
		MimeType: st.MimeType, UploaderID: st.UploaderID.String(), CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
}

func ToDomain(m *models.Image) *images.Image {
	id, _ := uuid.Parse(m.ID)
	uid, _ := uuid.Parse(m.UploaderID)
	return images.NewFromState(images.ImageState{
		ID: id, FilePath: m.FilePath, FileName: m.FileName, FileSize: m.FileSize,
		MimeType: m.MimeType, UploaderID: uid, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt,
	})
}
