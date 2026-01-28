package http

import (
	"nfxid/modules/directory/interfaces/http/handler"
)

type Registry struct {
	User            *handler.UserHandler
	Badge           *handler.BadgeHandler
	UserBadge       *handler.UserBadgeHandler
	UserEducation   *handler.UserEducationHandler
	UserEmail       *handler.UserEmailHandler
	UserOccupation  *handler.UserOccupationHandler
	UserPhone       *handler.UserPhoneHandler
	UserPreference  *handler.UserPreferenceHandler
	UserProfile     *handler.UserProfileHandler
	UserAvatar      *handler.UserAvatarHandler
	UserImage       *handler.UserImageHandler
}
