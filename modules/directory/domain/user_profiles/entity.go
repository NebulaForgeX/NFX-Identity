package user_profiles

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	state UserProfileState
}

type UserProfileState struct {
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

func (up *UserProfile) ID() uuid.UUID                  { return up.state.ID }
func (up *UserProfile) UserID() uuid.UUID              { return up.state.UserID }
func (up *UserProfile) Role() *string                  { return up.state.Role }
func (up *UserProfile) FirstName() *string             { return up.state.FirstName }
func (up *UserProfile) LastName() *string              { return up.state.LastName }
func (up *UserProfile) Nickname() *string              { return up.state.Nickname }
func (up *UserProfile) DisplayName() *string           { return up.state.DisplayName }
func (up *UserProfile) AvatarID() *uuid.UUID           { return up.state.AvatarID }
func (up *UserProfile) BackgroundID() *uuid.UUID       { return up.state.BackgroundID }
func (up *UserProfile) BackgroundIDs() []uuid.UUID     { return up.state.BackgroundIDs }
func (up *UserProfile) Bio() *string                   { return up.state.Bio }
func (up *UserProfile) Birthday() *time.Time           { return up.state.Birthday }
func (up *UserProfile) Age() *int                      { return up.state.Age }
func (up *UserProfile) Gender() *string                { return up.state.Gender }
func (up *UserProfile) Location() *string              { return up.state.Location }
func (up *UserProfile) Website() *string               { return up.state.Website }
func (up *UserProfile) Github() *string                { return up.state.Github }
func (up *UserProfile) SocialLinks() map[string]interface{} { return up.state.SocialLinks }
func (up *UserProfile) Skills() map[string]interface{} { return up.state.Skills }
func (up *UserProfile) CreatedAt() time.Time           { return up.state.CreatedAt }
func (up *UserProfile) UpdatedAt() time.Time           { return up.state.UpdatedAt }
func (up *UserProfile) DeletedAt() *time.Time          { return up.state.DeletedAt }
