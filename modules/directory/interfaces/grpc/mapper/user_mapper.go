package mapper

import (
	userAppResult "nfxid/modules/directory/application/users/results"
	userDomain "nfxid/modules/directory/domain/users"
	userpb "nfxid/protos/gen/directory/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserROToProto 将 UserRO 转换为 proto User 消息
func UserROToProto(v *userAppResult.UserRO) *userpb.User {
	if v == nil {
		return nil
	}

	user := &userpb.User{
		Id:         v.ID.String(),
		Username:   v.Username,
		Status:     userStatusToProto(v.Status),
		IsVerified: v.IsVerified,
		CreatedAt:  timestamppb.New(v.CreatedAt),
		UpdatedAt:  timestamppb.New(v.UpdatedAt),
	}

	if v.LastLoginAt != nil {
		user.LastLoginAt = timestamppb.New(*v.LastLoginAt)
	}

	if v.DeletedAt != nil {
		user.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return user
}

// UserListROToProto 批量转换 UserRO 到 proto User
func UserListROToProto(results []userAppResult.UserRO) []*userpb.User {
	users := make([]*userpb.User, len(results))
	for i, v := range results {
		users[i] = UserROToProto(&v)
	}
	return users
}

// userStatusToProto 将 domain UserStatus 转换为 proto DirectoryUserStatus
func userStatusToProto(status userDomain.UserStatus) userpb.DirectoryUserStatus {
	switch status {
	case userDomain.UserStatusPending:
		return userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_PENDING
	case userDomain.UserStatusActive:
		return userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_ACTIVE
	case userDomain.UserStatusDeactive:
		return userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_DEACTIVE
	default:
		return userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_UNSPECIFIED
	}
}
