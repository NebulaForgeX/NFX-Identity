package user_images

import (
	"time"

	"github.com/google/uuid"
)

type UserImage struct {
	state UserImageState
}

type UserImageState struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ImageID      uuid.UUID
	DisplayOrder int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (ui *UserImage) ID() uuid.UUID         { return ui.state.ID }
func (ui *UserImage) UserID() uuid.UUID     { return ui.state.UserID }
func (ui *UserImage) ImageID() uuid.UUID    { return ui.state.ImageID }
func (ui *UserImage) DisplayOrder() int     { return ui.state.DisplayOrder }
func (ui *UserImage) CreatedAt() time.Time  { return ui.state.CreatedAt }
func (ui *UserImage) UpdatedAt() time.Time  { return ui.state.UpdatedAt }
func (ui *UserImage) DeletedAt() *time.Time { return ui.state.DeletedAt }
