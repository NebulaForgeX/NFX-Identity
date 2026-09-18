package mapper

import (
	"nfxidentity/modules/asset/query/files"
	filepb "nfxidentity/protos/gen/asset/file"
)

func FileVOToProto(vo *files.FileVO) *filepb.File {
	if vo == nil {
		return nil
	}
	return &filepb.File{
		Id:         vo.ID.String(),
		FilePath:   vo.FilePath,
		FileName:   vo.FileName,
		FileSize:   vo.FileSize,
		MimeType:   vo.MimeType,
		UploaderId: vo.UploaderID.String(),
	}
}
