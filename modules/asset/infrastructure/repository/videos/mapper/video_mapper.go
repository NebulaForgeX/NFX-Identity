package mapper

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func VideoDomainToModel(e *videosDomain.Video) *rdbmodels.Video {
	if e == nil {
		return nil
	}
	st := e.State()
	return &rdbmodels.Video{
		ID:              st.ID,
		FilePath:        st.FilePath,
		FileName:        st.FileName,
		FileSize:        st.FileSize,
		MimeType:        st.MimeType,
		DurationSeconds: st.DurationSeconds,
		Width:           st.Width,
		Height:          st.Height,
		UploaderID:      st.UploaderID,
		CreatedAt:       st.CreatedAt,
		UpdatedAt:       st.UpdatedAt,
		DeletedAt:       ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func VideoModelToDomain(m *rdbmodels.Video) *videosDomain.Video {
	if m == nil {
		return nil
	}
	return videosDomain.NewVideoFromState(videosDomain.VideoState{
		ID:              m.ID,
		FilePath:        m.FilePath,
		FileName:        m.FileName,
		FileSize:        m.FileSize,
		MimeType:        m.MimeType,
		DurationSeconds: m.DurationSeconds,
		Width:           m.Width,
		Height:          m.Height,
		UploaderID:      m.UploaderID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}
