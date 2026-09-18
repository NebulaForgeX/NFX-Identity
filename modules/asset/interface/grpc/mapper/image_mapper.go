package mapper

import (
	"nfxidentity/modules/asset/query/images"
	imagepb "nfxidentity/protos/gen/asset/image"
)

func ImageVOToProto(vo *images.ImageVO) *imagepb.Image {
	if vo == nil {
		return nil
	}
	return &imagepb.Image{
		Id:         vo.ID.String(),
		FilePath:   vo.FilePath,
		FileName:   vo.FileName,
		FileSize:   vo.FileSize,
		MimeType:   vo.MimeType,
		UploaderId: vo.UploaderID.String(),
	}
}
