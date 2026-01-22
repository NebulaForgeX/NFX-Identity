package user_profiles

import (
	"time"

	"github.com/google/uuid"
)

type NewUserProfileParams struct {
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
}

func NewUserProfile(p NewUserProfileParams) (*UserProfile, error) {
	if err := validateUserProfileParams(p); err != nil {
		return nil, err
	}

	// id 必须等于 UserID（一对一关系，id 直接引用 users.id）
	now := time.Now().UTC()
	return NewUserProfileFromState(UserProfileState{
		ID:           p.UserID, // id 直接引用 users.id
		UserID:       p.UserID,
		Role:         p.Role,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Nickname:     p.Nickname,
		DisplayName:  p.DisplayName,
		AvatarID:     p.AvatarID,
		BackgroundID: p.BackgroundID,
		BackgroundIDs: p.BackgroundIDs,
		Bio:          p.Bio,
		Birthday:     p.Birthday,
		Age:          p.Age,
		Gender:       p.Gender,
		Location:     p.Location,
		Website:      p.Website,
		Github:       p.Github,
		SocialLinks:  p.SocialLinks,
		Skills:       p.Skills,
		CreatedAt:    now,
		UpdatedAt:    now,
	}), nil
}

func NewUserProfileFromState(st UserProfileState) *UserProfile {
	return &UserProfile{state: st}
}

func validateUserProfileParams(p NewUserProfileParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	return nil
}
