package mapper

import (
	userImageAppResult "nfxid/modules/directory/application/user_images/results"
	userimagepb "nfxid/protos/gen/directory/user_image"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserImageROToProto 将 UserImageRO 转换为 proto UserImage 消息
func UserImageROToProto(v *userImageAppResult.UserImageRO) *userimagepb.UserImage {
	if v == nil {
		return nil
	}

	userImage := &userimagepb.UserImage{
		Id:           v.ID.String(),
		UserId:       v.UserID.String(),
		ImageId:      v.ImageID.String(),
		DisplayOrder: int32(v.DisplayOrder),
		CreatedAt:    timestamppb.New(v.CreatedAt),
		UpdatedAt:    timestamppb.New(v.UpdatedAt),
	}

	if v.DeletedAt != nil {
		userImage.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userImage
}

// UserImageListROToProto 批量转换 UserImageRO 到 proto UserImage
func UserImageListROToProto(results []userImageAppResult.UserImageRO) []*userimagepb.UserImage {
	userImages := make([]*userimagepb.UserImage, len(results))
	for i, v := range results {
		userImages[i] = UserImageROToProto(&v)
	}
	return userImages
}
