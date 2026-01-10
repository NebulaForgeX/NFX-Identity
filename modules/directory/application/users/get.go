package users

import (
	"context"
	userResult "nfxid/modules/directory/application/users/results"

	"github.com/google/uuid"
)

// GetUser 根据ID获取用户
func (s *Service) GetUser(ctx context.Context, userID uuid.UUID) (userResult.UserRO, error) {
	domainEntity, err := s.userRepo.Get.ByID(ctx, userID)
	if err != nil {
		return userResult.UserRO{}, err
	}
	return userResult.UserMapper(domainEntity), nil
}

// GetUserByUsername 根据Username获取用户
func (s *Service) GetUserByUsername(ctx context.Context, username string) (userResult.UserRO, error) {
	domainEntity, err := s.userRepo.Get.ByUsername(ctx, username)
	if err != nil {
		return userResult.UserRO{}, err
	}
	return userResult.UserMapper(domainEntity), nil
}
