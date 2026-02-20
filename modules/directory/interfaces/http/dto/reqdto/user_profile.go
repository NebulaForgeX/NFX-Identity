package reqdto

import (
	userProfileAppCommands "nfxid/modules/directory/application/user_profiles/commands"

	"github.com/google/uuid"
)

type UserProfileCreateRequestDTO struct {
	UserID      uuid.UUID              `json:"user_id" validate:"required,uuid"`
	Role        *string                `json:"role,omitempty"`
	FirstName   *string                `json:"first_name,omitempty"`
	LastName    *string                `json:"last_name,omitempty"`
	Nickname    *string                `json:"nickname,omitempty"`
	DisplayName *string                `json:"display_name,omitempty"`
	Bio         *string                `json:"bio,omitempty"`
	Birthday    *string                `json:"birthday,omitempty"`
	Age         *int                   `json:"age,omitempty"`
	Gender      *string                `json:"gender,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Github      *string                `json:"github,omitempty"`
	SocialLinks map[string]interface{} `json:"social_links,omitempty"`
	Skills      map[string]interface{} `json:"skills,omitempty"`
}

type UserProfileByIDRequestDTO struct {
	UserProfileID uuid.UUID `uri:"user_profile_id" validate:"required,uuid"`
}

type UserProfileUpdateRequestDTO struct {
	UserProfileID uuid.UUID              `uri:"user_profile_id" validate:"required,uuid"`
	Role        *string                `json:"role,omitempty"`
	FirstName   *string                `json:"first_name,omitempty"`
	LastName    *string                `json:"last_name,omitempty"`
	Nickname    *string                `json:"nickname,omitempty"`
	DisplayName *string                `json:"display_name,omitempty"`
	Bio         *string                `json:"bio,omitempty"`
	Birthday    *string                `json:"birthday,omitempty"`
	Age         *int                   `json:"age,omitempty"`
	Gender      *string                `json:"gender,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Github      *string                `json:"github,omitempty"`
	SocialLinks map[string]interface{} `json:"social_links,omitempty"`
	Skills      map[string]interface{} `json:"skills,omitempty"`
}

func (r *UserProfileCreateRequestDTO) ToCreateCmd() userProfileAppCommands.CreateUserProfileCmd {
	return userProfileAppCommands.CreateUserProfileCmd{
		UserID:      r.UserID,
		Role:        r.Role,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Nickname:    r.Nickname,
		DisplayName: r.DisplayName,
		Bio:         r.Bio,
		Birthday:    r.Birthday,
		Age:         r.Age,
		Gender:      r.Gender,
		Location:    r.Location,
		Website:     r.Website,
		Github:      r.Github,
		SocialLinks: r.SocialLinks,
		Skills:      r.Skills,
	}
}

func (r *UserProfileUpdateRequestDTO) ToUpdateCmd() userProfileAppCommands.UpdateUserProfileCmd {
	return userProfileAppCommands.UpdateUserProfileCmd{
		UserProfileID: r.UserProfileID,
		Role:          r.Role,
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		Nickname:      r.Nickname,
		DisplayName:   r.DisplayName,
		Bio:           r.Bio,
		Birthday:      r.Birthday,
		Age:           r.Age,
		Gender:        r.Gender,
		Location:      r.Location,
		Website:       r.Website,
		Github:        r.Github,
		SocialLinks:   r.SocialLinks,
		Skills:        r.Skills,
	}
}
