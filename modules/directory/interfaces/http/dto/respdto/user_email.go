package respdto

import (
	"time"

	userEmailAppResult "nfxid/modules/directory/application/user_emails/results"

	"github.com/google/uuid"
)

type UserEmailDTO struct {
	ID                uuid.UUID  `json:"id"`
	UserID            uuid.UUID  `json:"user_id"`
	Email             string     `json:"email"`
	IsPrimary         bool       `json:"is_primary"`
	IsVerified        bool       `json:"is_verified"`
	VerifiedAt        *time.Time `json:"verified_at,omitempty"`
	VerificationToken *string    `json:"verification_token,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

// UserEmailROToDTO converts application UserEmailRO to response DTO
func UserEmailROToDTO(v *userEmailAppResult.UserEmailRO) *UserEmailDTO {
	if v == nil {
		return nil
	}

	return &UserEmailDTO{
		ID:                v.ID,
		UserID:            v.UserID,
		Email:             v.Email,
		IsPrimary:         v.IsPrimary,
		IsVerified:        v.IsVerified,
		VerifiedAt:        v.VerifiedAt,
		VerificationToken: v.VerificationToken,
		CreatedAt:         v.CreatedAt,
		UpdatedAt:         v.UpdatedAt,
		DeletedAt:         v.DeletedAt,
	}
}

// UserEmailListROToDTO converts list of UserEmailRO to DTOs
func UserEmailListROToDTO(results []userEmailAppResult.UserEmailRO) []UserEmailDTO {
	dtos := make([]UserEmailDTO, len(results))
	for i, v := range results {
		if dto := UserEmailROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
