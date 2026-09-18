package mapper

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func ImageDomainToModel(e *imagesDomain.Image) *rdbmodels.Image {
	if e == nil {
		return nil
	}
	st := e.State()
	return &rdbmodels.Image{
		ID:         st.ID,
		FilePath:   st.FilePath,
		FileName:   st.FileName,
		FileSize:   st.FileSize,
		MimeType:   st.MimeType,
		Width:      st.Width,
		Height:     st.Height,
		AltText:    st.AltText,
		UploaderID: st.UploaderID,
		CreatedAt:  st.CreatedAt,
		UpdatedAt:  st.UpdatedAt,
		DeletedAt:  ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func ImageModelToDomain(m *rdbmodels.Image) *imagesDomain.Image {
	if m == nil {
		return nil
	}
	return imagesDomain.NewImageFromState(imagesDomain.ImageState{
		ID:         m.ID,
		FilePath:   m.FilePath,
		FileName:   m.FileName,
		FileSize:   m.FileSize,
		MimeType:   m.MimeType,
		Width:      m.Width,
		Height:     m.Height,
		AltText:    m.AltText,
		UploaderID: m.UploaderID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}
