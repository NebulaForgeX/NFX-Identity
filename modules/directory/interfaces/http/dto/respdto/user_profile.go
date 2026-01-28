package respdto

import (
	"time"

	userProfileAppResult "nfxid/modules/directory/application/user_profiles/results"

	"github.com/google/uuid"
)

type UserProfileDTO struct {
	ID          uuid.UUID              `json:"id"`
	UserID      uuid.UUID              `json:"user_id"`
	Role        *string                `json:"role,omitempty"`
	FirstName   *string                `json:"first_name,omitempty"`
	LastName    *string                `json:"last_name,omitempty"`
	Nickname    *string                `json:"nickname,omitempty"`
	DisplayName *string                `json:"display_name,omitempty"`
	Bio         *string                `json:"bio,omitempty"`
	Birthday    *time.Time             `json:"birthday,omitempty"`
	Age         *int                   `json:"age,omitempty"`
	Gender      *string                `json:"gender,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Github      *string                `json:"github,omitempty"`
	SocialLinks map[string]interface{} `json:"social_links,omitempty"`
	Skills      map[string]interface{} `json:"skills,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	DeletedAt   *time.Time             `json:"deleted_at,omitempty"`
}

// UserProfileROToDTO converts application UserProfileRO to response DTO
func UserProfileROToDTO(v *userProfileAppResult.UserProfileRO) *UserProfileDTO {
	if v == nil {
		return nil
	}

	return &UserProfileDTO{
		ID:          v.ID,
		UserID:      v.UserID,
		Role:        v.Role,
		FirstName:   v.FirstName,
		LastName:    v.LastName,
		Nickname:    v.Nickname,
		DisplayName: v.DisplayName,
		Bio:         v.Bio,
		Birthday:    v.Birthday,
		Age:         v.Age,
		Gender:      v.Gender,
		Location:    v.Location,
		Website:     v.Website,
		Github:      v.Github,
		SocialLinks: v.SocialLinks,
		Skills:      v.Skills,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// UserProfileListROToDTO converts list of UserProfileRO to DTOs
func UserProfileListROToDTO(results []userProfileAppResult.UserProfileRO) []UserProfileDTO {
	dtos := make([]UserProfileDTO, len(results))
	for i, v := range results {
		if dto := UserProfileROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
