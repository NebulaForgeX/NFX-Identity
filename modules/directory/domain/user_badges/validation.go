package user_badges

import "github.com/google/uuid"

func (ub *UserBadge) Validate() error {
	if ub.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ub.BadgeID() == uuid.Nil {
		return ErrBadgeIDRequired
	}
	return nil
}
