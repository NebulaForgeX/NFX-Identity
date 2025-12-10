package events

import "nebulaid/pkgs/eventbus"

type EventType = eventbus.EventType

const (
	// =============== Image -> Auth Events ===============
	ETImageToAuth_ImageDelete  EventType = "image_to_auth.image_delete"
	ETImageToAuth_ImageSuccess EventType = "image_to_auth.image_success"
	ETImageToAuth_ImageTest    EventType = "image_to_auth.image_test"

	// =============== Auth -> Image Events ===============
	ETAuthToImage_ImageDelete  EventType = "auth_to_image.image_delete"
	ETAuthToImage_ImageSuccess EventType = "auth_to_image.image_success"
	ETAuthToImage_ImageTest    EventType = "auth_to_image.image_test"

	// =============== Auth -> Auth Events (Internal) ===============
	ETAuthToAuth_Success                     EventType = "auth_to_auth.success"
	ETAuthToAuth_Test                        EventType = "auth_to_auth.test"
	ETAuthToAuth_UserInvalidateCache         EventType = "auth_to_auth.user.invalidate_cache"
	ETAuthToAuth_ProfileInvalidateCache      EventType = "auth_to_auth.profile.invalidate_cache"
	ETAuthToAuth_ProfileBadgeInvalidateCache EventType = "auth_to_auth.profile_badge.invalidate_cache"
	ETAuthToAuth_RoleInvalidateCache         EventType = "auth_to_auth.role.invalidate_cache"
	ETAuthToAuth_BadgeInvalidateCache        EventType = "auth_to_auth.badge.invalidate_cache"
	ETAuthToAuth_EducationInvalidateCache    EventType = "auth_to_auth.education.invalidate_cache"
	ETAuthToAuth_OccupationInvalidateCache   EventType = "auth_to_auth.occupation.invalidate_cache"
)
