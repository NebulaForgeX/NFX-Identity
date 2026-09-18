package mapper

import (
	filesDomain "nfxidentity/modules/asset/domain/files"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func FileDomainToModel(e *filesDomain.File) *rdbmodels.File {
	if e == nil {
		return nil
	}
	st := e.State()
	return &rdbmodels.File{
		ID:         st.ID,
		FilePath:   st.FilePath,
		FileName:   st.FileName,
		FileSize:   st.FileSize,
		MimeType:   st.MimeType,
		UploaderID: st.UploaderID,
		CreatedAt:  st.CreatedAt,
		UpdatedAt:  st.UpdatedAt,
		DeletedAt:  ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func FileModelToDomain(m *rdbmodels.File) *filesDomain.File {
	if m == nil {
		return nil
	}
	return filesDomain.NewFileFromState(filesDomain.FileState{
		ID:         m.ID,
		FilePath:   m.FilePath,
		FileName:   m.FileName,
		FileSize:   m.FileSize,
		MimeType:   m.MimeType,
		UploaderID: m.UploaderID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}
