package directory

import (
	badgepb "nfxid/protos/gen/directory/badge"
	userpb "nfxid/protos/gen/directory/user"
	userbadgepb "nfxid/protos/gen/directory/user_badge"
	usereducationpb "nfxid/protos/gen/directory/user_education"
	useremailpb "nfxid/protos/gen/directory/user_email"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	userprofilepb "nfxid/protos/gen/directory/user_profile"
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
		Badge:          NewBadgeClient(badgeClient),
		UserBadge:      NewUserBadgeClient(userBadgeClient),
	}
}
