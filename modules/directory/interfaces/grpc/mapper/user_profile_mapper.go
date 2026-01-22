package mapper

import (
	userProfileAppResult "nfxid/modules/directory/application/user_profiles/results"
	userprofilepb "nfxid/protos/gen/directory/user_profile"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserProfileROToProto 将 UserProfileRO 转换为 proto UserProfile 消息
func UserProfileROToProto(v *userProfileAppResult.UserProfileRO) *userprofilepb.UserProfile {
	if v == nil {
		return nil
	}

	userProfile := &userprofilepb.UserProfile{
		Id:        v.ID.String(), // id 直接引用 users.id
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Role != nil {
		userProfile.Role = v.Role
	}

	if v.FirstName != nil {
		userProfile.FirstName = v.FirstName
	}

	if v.LastName != nil {
		userProfile.LastName = v.LastName
	}

	if v.Nickname != nil {
		userProfile.Nickname = v.Nickname
	}

	if v.DisplayName != nil {
		userProfile.DisplayName = v.DisplayName
	}

	if v.AvatarID != nil {
		avatarID := v.AvatarID.String()
		userProfile.AvatarId = &avatarID
	}

	if v.BackgroundID != nil {
		backgroundID := v.BackgroundID.String()
		userProfile.BackgroundId = &backgroundID
	}

	if len(v.BackgroundIDs) > 0 {
		// proto定义background_ids是string，需要将UUID数组序列化为字符串
		// 这里简单处理，将第一个ID作为字符串，或者可以序列化为JSON字符串
		backgroundIDsStr := v.BackgroundIDs[0].String()
		userProfile.BackgroundIds = &backgroundIDsStr
	}

	if v.Bio != nil {
		userProfile.Bio = v.Bio
	}

	if v.Birthday != nil {
		userProfile.Birthday = timestamppb.New(*v.Birthday)
	}

	if v.Age != nil {
		age := int32(*v.Age)
		userProfile.Age = &age
	}

	if v.Gender != nil {
		userProfile.Gender = v.Gender
	}

	if v.Location != nil {
		userProfile.Location = v.Location
	}

	if v.Website != nil {
		userProfile.Website = v.Website
	}

	if v.Github != nil {
		userProfile.Github = v.Github
	}

	if len(v.SocialLinks) > 0 {
		if socialLinksStruct, err := structpb.NewStruct(v.SocialLinks); err == nil {
			userProfile.SocialLinks = socialLinksStruct
		}
	}

	if len(v.Skills) > 0 {
		if skillsStruct, err := structpb.NewStruct(v.Skills); err == nil {
			userProfile.Skills = skillsStruct
		}
	}

	if v.DeletedAt != nil {
		userProfile.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userProfile
}

// UserProfileListROToProto 批量转换 UserProfileRO 到 proto UserProfile
func UserProfileListROToProto(results []userProfileAppResult.UserProfileRO) []*userprofilepb.UserProfile {
	userProfiles := make([]*userprofilepb.UserProfile, len(results))
	for i, v := range results {
		userProfiles[i] = UserProfileROToProto(&v)
	}
	return userProfiles
}
