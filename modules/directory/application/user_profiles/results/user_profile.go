package results

import (
	"time"

	"nfxid/modules/directory/domain/user_profiles"

	"github.com/google/uuid"
)

type UserProfileRO struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Role         *string
	FirstName    *string
	LastName     *string
	Nickname     *string
	DisplayName  *string
	AvatarID     *uuid.UUID
	BackgroundID *uuid.UUID
	BackgroundIDs []uuid.UUID
	Bio          *string
	Birthday     *time.Time
	Age          *int
	Gender       *string
	Location     *string
	Website      *string
	Github       *string
	SocialLinks  map[string]interface{}
	Skills       map[string]interface{}
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// UserProfileMapper 将 Domain UserProfile 转换为 Application UserProfileRO
func UserProfileMapper(up *user_profiles.UserProfile) UserProfileRO {
	if up == nil {
		return UserProfileRO{}
	}

	return UserProfileRO{
		ID:           up.ID(),
		UserID:       up.UserID(),
		Role:         up.Role(),
		FirstName:    up.FirstName(),
		LastName:     up.LastName(),
		Nickname:     up.Nickname(),
		DisplayName:  up.DisplayName(),
		AvatarID:     up.AvatarID(),
		BackgroundID: up.BackgroundID(),
		BackgroundIDs: up.BackgroundIDs(),
		Bio:          up.Bio(),
		Birthday:     up.Birthday(),
		Age:          up.Age(),
		Gender:       up.Gender(),
		Location:     up.Location(),
		Website:      up.Website(),
		Github:       up.Github(),
		SocialLinks:  up.SocialLinks(),
		Skills:       up.Skills(),
		CreatedAt:    up.CreatedAt(),
		UpdatedAt:    up.UpdatedAt(),
		DeletedAt:    up.DeletedAt(),
	}
}
