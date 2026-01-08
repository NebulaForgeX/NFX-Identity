package user_badges

import (
	"time"
)

func (ub *UserBadge) Update(description string, level int) error {
	if description != "" {
		ub.state.Description = description
	}
	if level > 0 {
		ub.state.Level = level
	}
	ub.state.UpdatedAt = time.Now().UTC()
	return nil
}
