package mapper

import (
	userAvatarAppResult "nfxid/modules/directory/application/user_avatars/results"
	useravatarpb "nfxid/protos/gen/directory/user_avatar"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserAvatarROToProto 将 UserAvatarRO 转换为 proto UserAvatar 消息
func UserAvatarROToProto(v *userAvatarAppResult.UserAvatarRO) *useravatarpb.UserAvatar {
	if v == nil {
		return nil
	}

	return &useravatarpb.UserAvatar{
		UserId:    v.UserID.String(),
		ImageId:   v.ImageID.String(),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}
}
