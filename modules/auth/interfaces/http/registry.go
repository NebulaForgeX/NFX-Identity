package http

import (
	"nebulaid/modules/auth/interfaces/http/handler"
)

type Registry struct {
	User         *handler.UserHandler
	Profile      *handler.ProfileHandler
	Role         *handler.RoleHandler
	Badge        *handler.BadgeHandler
	Education    *handler.EducationHandler
	Occupation   *handler.OccupationHandler
	ProfileBadge *handler.ProfileBadgeHandler
}
