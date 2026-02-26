package directory

import (
	badgepb "nfxidentity/protos/gen/directory/badge"
	userpb "nfxidentity/protos/gen/directory/user"
	useravatarpb "nfxidentity/protos/gen/directory/user_avatar"
	userbadgepb "nfxidentity/protos/gen/directory/user_badge"
	usereducationpb "nfxidentity/protos/gen/directory/user_education"
	useremailpb "nfxidentity/protos/gen/directory/user_email"
	userimagepb "nfxidentity/protos/gen/directory/user_image"
	useroccupationpb "nfxidentity/protos/gen/directory/user_occupation"
	userphonepb "nfxidentity/protos/gen/directory/user_phone"
	userpreferencepb "nfxidentity/protos/gen/directory/user_preference"
	userprofilepb "nfxidentity/protos/gen/directory/user_profile"
)

// Client Directory 服务客户端
type Client struct {
	User           *UserClient
	UserProfile    *UserProfileClient
	UserEmail      *UserEmailClient
	UserPhone      *UserPhoneClient
	UserPreference *UserPreferenceClient
	UserEducation  *UserEducationClient
	UserOccupation *UserOccupationClient
	UserAvatar     *UserAvatarClient
	UserImage      *UserImageClient
	Badge          *BadgeClient
	UserBadge      *UserBadgeClient
}

// NewClient 创建 Directory 客户端
func NewClient(
	userClient userpb.UserServiceClient,
	userProfileClient userprofilepb.UserProfileServiceClient,
	userEmailClient useremailpb.UserEmailServiceClient,
	userPhoneClient userphonepb.UserPhoneServiceClient,
	userPreferenceClient userpreferencepb.UserPreferenceServiceClient,
	userEducationClient usereducationpb.UserEducationServiceClient,
	userOccupationClient useroccupationpb.UserOccupationServiceClient,
	userAvatarClient useravatarpb.UserAvatarServiceClient,
	userImageClient userimagepb.UserImageServiceClient,
	badgeClient badgepb.BadgeServiceClient,
	userBadgeClient userbadgepb.UserBadgeServiceClient,
) *Client {
	return &Client{
		User:           NewUserClient(userClient),
		UserProfile:    NewUserProfileClient(userProfileClient),
		UserEmail:      NewUserEmailClient(userEmailClient),
		UserPhone:      NewUserPhoneClient(userPhoneClient),
		UserPreference: NewUserPreferenceClient(userPreferenceClient),
		UserEducation:  NewUserEducationClient(userEducationClient),
		UserOccupation: NewUserOccupationClient(userOccupationClient),
		UserAvatar:     NewUserAvatarClient(userAvatarClient),
		UserImage:      NewUserImageClient(userImageClient),
		Badge:          NewBadgeClient(badgeClient),
		UserBadge:      NewUserBadgeClient(userBadgeClient),
	}
}
