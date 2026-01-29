package mapper

import (
	imageAppResult "nfxid/modules/image/application/images/results"
	imagepb "nfxid/protos/gen/image/image"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ImageROToProto 将 ImageRO 转换为 proto Image 消息
func ImageROToProto(ro imageAppResult.ImageRO) *imagepb.Image {
	out := &imagepb.Image{
		Id:               ro.ID.String(),
		Filename:         ro.Filename,
		OriginalFilename: ro.OriginalFilename,
		MimeType:         ro.MimeType,
		Size:             ro.Size,
		StoragePath:      ro.StoragePath,
		IsPublic:         ro.IsPublic,
		CreatedAt:        timestamppb.New(ro.CreatedAt),
		UpdatedAt:        timestamppb.New(ro.UpdatedAt),
	}
	if ro.TypeID != nil {
		s := ro.TypeID.String()
		out.TypeId = &s
	}
	if ro.UserID != nil {
		s := ro.UserID.String()
		out.UserId = &s
	}
	if ro.TenantID != nil {
		s := ro.TenantID.String()
		out.TenantId = &s
	}
	if ro.AppID != nil {
		s := ro.AppID.String()
		out.AppId = &s
	}
	if ro.SourceDomain != nil {
		out.SourceDomain = ro.SourceDomain
	}
	if ro.Width != nil {
		w := int32(*ro.Width)
		out.Width = &w
	}
	if ro.Height != nil {
		h := int32(*ro.Height)
		out.Height = &h
	}
	if ro.URL != nil {
		out.Url = ro.URL
	}
	if len(ro.Metadata) > 0 {
		if meta, err := structpb.NewStruct(ro.Metadata); err == nil {
			out.Metadata = meta
		}
	}
	if ro.DeletedAt != nil {
		out.DeletedAt = timestamppb.New(*ro.DeletedAt)
	}
	return out
}
