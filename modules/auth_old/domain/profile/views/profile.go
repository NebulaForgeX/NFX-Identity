package views

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ProfileView struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	Username        string
	Email           string
	UserPhone       *string
	UserStatus      string
	IsVerified      bool
	FirstName       *string
	LastName        *string
	Nickname        *string
	DisplayName     *string
	AvatarID        *uuid.UUID
	BackgroundID    *uuid.UUID
	BackgroundIds   []uuid.UUID
	Bio             *string
	Phone           *string
	Birthday        *time.Time
	Age             *int
	Gender          *string
	Location        *string
	Website         *string
	Github          *string
	SocialLinks     *datatypes.JSON
	Preferences     *datatypes.JSON
	Skills          *datatypes.JSON
	PrivacySettings *datatypes.JSON
	Occupations     *datatypes.JSON
	Educations      *datatypes.JSON
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

