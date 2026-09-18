package mapper

import (
	"nfxidentity/modules/asset/query/audios"
	audiopb "nfxidentity/protos/gen/asset/audio"
)

func AudioVOToProto(vo *audios.AudioVO) *audiopb.Audio {
	if vo == nil {
		return nil
	}
	return &audiopb.Audio{
		Id:         vo.ID.String(),
		FilePath:   vo.FilePath,
		FileName:   vo.FileName,
		FileSize:   vo.FileSize,
		MimeType:   vo.MimeType,
		UploaderId: vo.UploaderID.String(),
	}
}
