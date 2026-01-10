package user_credentials

import (
	"context"
	userCredentialResult "nfxid/modules/auth/application/user_credentials/results"

	"github.com/google/uuid"
)

// GetUserCredential 根据ID获取用户凭证
func (s *Service) GetUserCredential(ctx context.Context, userCredentialID uuid.UUID) (userCredentialResult.UserCredentialRO, error) {
	domainEntity, err := s.userCredentialRepo.Get.ByID(ctx, userCredentialID)
	if err != nil {
		return userCredentialResult.UserCredentialRO{}, err
	}
	return userCredentialResult.UserCredentialMapper(domainEntity), nil
}

// GetUserCredentialByUserID 根据UserID获取用户凭证
func (s *Service) GetUserCredentialByUserID(ctx context.Context, userID uuid.UUID) (userCredentialResult.UserCredentialRO, error) {
	domainEntity, err := s.userCredentialRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return userCredentialResult.UserCredentialRO{}, err
	}
	return userCredentialResult.UserCredentialMapper(domainEntity), nil
}
