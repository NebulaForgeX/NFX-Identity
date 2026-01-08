package user_profiles

import (
	"time"

	"github.com/google/uuid"
)

func (up *UserProfile) Update(role, firstName, lastName, nickname, displayName *string, avatarID, backgroundID *uuid.UUID, backgroundIDs []uuid.UUID, bio, gender, location, website, github *string, birthday *time.Time, age *int, socialLinks, skills map[string]interface{}) error {
	if up.DeletedAt() != nil {
		return ErrUserProfileNotFound
	}

	if role != nil {
		up.state.Role = role
	}
	if firstName != nil {
		up.state.FirstName = firstName
	}
	if lastName != nil {
		up.state.LastName = lastName
	}
	if nickname != nil {
		up.state.Nickname = nickname
	}
	if displayName != nil {
		up.state.DisplayName = displayName
	}
	if avatarID != nil {
		up.state.AvatarID = avatarID
	}
	if backgroundID != nil {
		up.state.BackgroundID = backgroundID
	}
	if backgroundIDs != nil {
		up.state.BackgroundIDs = backgroundIDs
	}
	if bio != nil {
		up.state.Bio = bio
	}
	if birthday != nil {
		up.state.Birthday = birthday
	}
	if age != nil {
		up.state.Age = age
	}
	if gender != nil {
		up.state.Gender = gender
	}
	if location != nil {
		up.state.Location = location
	}
	if website != nil {
		up.state.Website = website
	}
	if github != nil {
		up.state.Github = github
	}
	if socialLinks != nil {
		up.state.SocialLinks = socialLinks
	}
	if skills != nil {
		up.state.Skills = skills
	}

	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserProfile) Delete() error {
	if up.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	up.state.DeletedAt = &now
	up.state.UpdatedAt = now
	return nil
}
