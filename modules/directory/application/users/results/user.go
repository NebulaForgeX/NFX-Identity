package results

import (
	"time"

	"nfxid/modules/directory/domain/users"

	"github.com/google/uuid"
)

type UserRO struct {
	ID          uuid.UUID
	Username    string
	Status      users.UserStatus
	IsVerified  bool
	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// UserMapper 将 Domain User 转换为 Application UserRO
func UserMapper(u *users.User) UserRO {
	if u == nil {
		return UserRO{}
	}

	return UserRO{
		ID:          u.ID(),
		Username:    u.Username(),
		Status:      u.Status(),
		IsVerified:  u.IsVerified(),
		LastLoginAt: u.LastLoginAt(),
		CreatedAt:   u.CreatedAt(),
		UpdatedAt:   u.UpdatedAt(),
		DeletedAt:   u.DeletedAt(),
	}
}
