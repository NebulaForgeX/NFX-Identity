package mapper

import (
	"nfxidentity/modules/asset/query/videos"
	videopb "nfxidentity/protos/gen/asset/video"
)

func VideoVOToProto(vo *videos.VideoVO) *videopb.Video {
	if vo == nil {
		return nil
	}
	return &videopb.Video{
		Id:         vo.ID.String(),
		FilePath:   vo.FilePath,
		FileName:   vo.FileName,
		FileSize:   vo.FileSize,
		MimeType:   vo.MimeType,
		UploaderId: vo.UploaderID.String(),
	}
}
