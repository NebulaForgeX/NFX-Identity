package mapper

import (
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func AudioDomainToModel(e *audiosDomain.Audio) *rdbmodels.Audio {
	if e == nil {
		return nil
	}
	st := e.State()
	return &rdbmodels.Audio{
		ID:              st.ID,
		FilePath:        st.FilePath,
		FileName:        st.FileName,
		FileSize:        st.FileSize,
		MimeType:        st.MimeType,
		DurationSeconds: st.DurationSeconds,
		UploaderID:      st.UploaderID,
		CreatedAt:       st.CreatedAt,
		UpdatedAt:       st.UpdatedAt,
		DeletedAt:       ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func AudioModelToDomain(m *rdbmodels.Audio) *audiosDomain.Audio {
	if m == nil {
		return nil
	}
	return audiosDomain.NewAudioFromState(audiosDomain.AudioState{
		ID:              m.ID,
		FilePath:        m.FilePath,
		FileName:        m.FileName,
		FileSize:        m.FileSize,
		MimeType:        m.MimeType,
		DurationSeconds: m.DurationSeconds,
		UploaderID:      m.UploaderID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}
