package user_badges

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (ub *UserBadge) Validate() error {
	if ub.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if ub.BadgeID() == uuid.Nil {
		return dirErr.ErrBadgeIDRequired
	}
	return nil
}
