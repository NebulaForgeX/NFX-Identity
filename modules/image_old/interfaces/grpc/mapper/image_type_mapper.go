package mapper

import (
	imageTypeAppViews "nfxid/modules/image/application/image_type/views"
	imageTypeDomainViews "nfxid/modules/image/domain/image_type/views"
	imagetypepb "nfxid/protos/gen/image/image_type"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ImageTypeDomainViewToProto 将 domain ImageTypeView 转换为 proto ImageType 消息（推荐使用）
func ImageTypeDomainViewToProto(v *imageTypeDomainViews.ImageTypeView) *imagetypepb.ImageType {
	if v == nil {
		return nil
	}

	isSystem := v.IsSystem
	imageType := &imagetypepb.ImageType{
		Id:        v.ID.String(),
		Key:       v.Key,
		IsSystem:  &isSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		imageType.Description = v.Description
	}
	if v.MaxWidth != nil {
		w := int32(*v.MaxWidth)
		imageType.MaxWidth = &w
	}
	if v.MaxHeight != nil {
		h := int32(*v.MaxHeight)
		imageType.MaxHeight = &h
	}
	if v.AspectRatio != nil {
		imageType.AspectRatio = v.AspectRatio
	}

	return imageType
}

// ImageTypeViewToProto 将 application ImageTypeView 转换为 proto ImageType 消息
func ImageTypeViewToProto(v *imageTypeAppViews.ImageTypeView) *imagetypepb.ImageType {
	if v == nil {
		return nil
	}

	isSystem := v.IsSystem
	imageType := &imagetypepb.ImageType{
		Id:       v.ID,
		Key:      v.Key,
		IsSystem: &isSystem,
	}

	if v.Description != "" {
		imageType.Description = &v.Description
	}
	if v.MaxWidth != nil {
		w := int32(*v.MaxWidth)
		imageType.MaxWidth = &w
	}
	if v.MaxHeight != nil {
		h := int32(*v.MaxHeight)
		imageType.MaxHeight = &h
	}
	if v.AspectRatio != "" {
		imageType.AspectRatio = &v.AspectRatio
	}

	return imageType
}
