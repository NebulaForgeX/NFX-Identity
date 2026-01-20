package respdto

import (
	"time"

	userAppResult "nfxid/modules/directory/application/users/results"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID          uuid.UUID  `json:"id"`
	Username    string     `json:"username"`
	Status      string     `json:"status"`
	IsVerified  bool       `json:"is_verified"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// UserROToDTO converts application UserRO to response DTO
func UserROToDTO(v *userAppResult.UserRO) *UserDTO {
	if v == nil {
		return nil
	}

	return &UserDTO{
		ID:          v.ID,
		Username:    v.Username,
		Status:      string(v.Status),
		IsVerified:  v.IsVerified,
		LastLoginAt: v.LastLoginAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// UserListROToDTO converts list of UserRO to DTOs
func UserListROToDTO(results []userAppResult.UserRO) []UserDTO {
	dtos := make([]UserDTO, len(results))
	for i, v := range results {
		if dto := UserROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
