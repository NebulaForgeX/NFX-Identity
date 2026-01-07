package mapper

import (
	imageAppViews "nfxid/modules/image/application/image/views"
	imageDomainViews "nfxid/modules/image/domain/image/views"
	imagepb "nfxid/protos/gen/image/image"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ImageDomainViewToProto 将 domain ImageView 转换为 proto Image 消息（推荐使用）
func ImageDomainViewToProto(v *imageDomainViews.ImageView) *imagepb.Image {
	if v == nil {
		return nil
	}

	image := &imagepb.Image{
		Id:               v.ID.String(),
		Filename:         v.Filename,
		OriginalFilename: v.OriginalFilename,
		MimeType:         v.MimeType,
		Size:             v.Size,
		StoragePath:      v.StoragePath,
		IsPublic:         v.IsPublic,
		CreatedAt:        timestamppb.New(v.CreatedAt),
		UpdatedAt:        timestamppb.New(v.UpdatedAt),
	}

	if v.TypeID != nil {
		typeIDStr := v.TypeID.String()
		image.TypeId = &typeIDStr
	}
	if v.UserID != nil {
		userIDStr := v.UserID.String()
		image.UserId = &userIDStr
	}
	if v.SourceDomain != nil {
		image.SourceDomain = v.SourceDomain
	}
	if v.Width != nil {
		w := int32(*v.Width)
		image.Width = &w
	}
	if v.Height != nil {
		h := int32(*v.Height)
		image.Height = &h
	}
	if v.URL != nil {
		image.Url = v.URL
	}
	if v.Metadata != nil {
		// 将 map[string]interface{} 转换为 map[string]string
		metadataStr := make(map[string]string)
		for k, val := range v.Metadata {
			if str, ok := val.(string); ok {
				metadataStr[k] = str
			} else {
				// 尝试序列化为字符串
				metadataStr[k] = ""
			}
		}
		image.Metadata = metadataStr
	}

	return image
}

// ImageViewToProto 将 application ImageView 转换为 proto Image 消息
func ImageViewToProto(v *imageAppViews.ImageView) *imagepb.Image {
	if v == nil {
		return nil
	}

	image := &imagepb.Image{
		Id:               v.ID,
		Filename:         v.Filename,
		OriginalFilename: v.OriginalFilename,
		MimeType:         v.MimeType,
		Size:             v.Size,
		StoragePath:      v.StoragePath,
		IsPublic:         v.IsPublic,
	}

	if v.TypeID != "" {
		image.TypeId = &v.TypeID
	}
	if v.UserID != "" {
		image.UserId = &v.UserID
	}
	if v.SourceDomain != "" {
		image.SourceDomain = &v.SourceDomain
	}
	if v.Width != nil {
		w := int32(*v.Width)
		image.Width = &w
	}
	if v.Height != nil {
		h := int32(*v.Height)
		image.Height = &h
	}
	if v.URL != "" {
		image.Url = &v.URL
	}
	if v.Metadata != nil {
		// 将 map[string]interface{} 转换为 map[string]string
		metadataStr := make(map[string]string)
		for k, val := range v.Metadata {
			if str, ok := val.(string); ok {
				metadataStr[k] = str
			}
		}
		image.Metadata = metadataStr
	}

	// 解析时间字符串（如果 application view 中有时间字符串）
	// 注意：这里需要从 domain view 获取 time.Time，或者在这里解析字符串

	return image
}
