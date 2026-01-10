package respdto

import (
	"time"

	userPhoneAppResult "nfxid/modules/directory/application/user_phones/results"

	"github.com/google/uuid"
)

type UserPhoneDTO struct {
	ID                    uuid.UUID  `json:"id"`
	UserID                uuid.UUID  `json:"user_id"`
	Phone                 string     `json:"phone"`
	CountryCode           *string    `json:"country_code,omitempty"`
	IsPrimary             bool       `json:"is_primary"`
	IsVerified            bool       `json:"is_verified"`
	VerifiedAt            *time.Time `json:"verified_at,omitempty"`
	VerificationCode      *string    `json:"verification_code,omitempty"`
	VerificationExpiresAt *time.Time `json:"verification_expires_at,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty"`
}

// UserPhoneROToDTO converts application UserPhoneRO to response DTO
func UserPhoneROToDTO(v *userPhoneAppResult.UserPhoneRO) *UserPhoneDTO {
	if v == nil {
		return nil
	}

	return &UserPhoneDTO{
		ID:                    v.ID,
		UserID:                v.UserID,
		Phone:                 v.Phone,
		CountryCode:           v.CountryCode,
		IsPrimary:             v.IsPrimary,
		IsVerified:            v.IsVerified,
		VerifiedAt:            v.VerifiedAt,
		VerificationCode:      v.VerificationCode,
		VerificationExpiresAt: v.VerificationExpiresAt,
		CreatedAt:             v.CreatedAt,
		UpdatedAt:             v.UpdatedAt,
		DeletedAt:             v.DeletedAt,
	}
}

// UserPhoneListROToDTO converts list of UserPhoneRO to DTOs
func UserPhoneListROToDTO(results []userPhoneAppResult.UserPhoneRO) []UserPhoneDTO {
	dtos := make([]UserPhoneDTO, len(results))
	for i, v := range results {
		if dto := UserPhoneROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
