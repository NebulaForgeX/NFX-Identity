package user_preferences

import (
	"time"

	"github.com/google/uuid"
)

type UserPreference struct {
	state UserPreferenceState
}

type UserPreferenceState struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Theme        string
	Language     string
	Timezone     string
	Notifications map[string]interface{}
	Privacy      map[string]interface{}
	Display      map[string]interface{}
	Other        map[string]interface{}
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (up *UserPreference) ID() uuid.UUID                    { return up.state.ID }
func (up *UserPreference) UserID() uuid.UUID                { return up.state.UserID }
func (up *UserPreference) Theme() string                    { return up.state.Theme }
func (up *UserPreference) Language() string                 { return up.state.Language }
func (up *UserPreference) Timezone() string                 { return up.state.Timezone }
func (up *UserPreference) Notifications() map[string]interface{} { return up.state.Notifications }
func (up *UserPreference) Privacy() map[string]interface{}  { return up.state.Privacy }
func (up *UserPreference) Display() map[string]interface{}  { return up.state.Display }
func (up *UserPreference) Other() map[string]interface{}    { return up.state.Other }
func (up *UserPreference) CreatedAt() time.Time             { return up.state.CreatedAt }
func (up *UserPreference) UpdatedAt() time.Time             { return up.state.UpdatedAt }
func (up *UserPreference) DeletedAt() *time.Time            { return up.state.DeletedAt }
