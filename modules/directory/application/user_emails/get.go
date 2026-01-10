package user_emails

import (
	"context"
	userEmailResult "nfxid/modules/directory/application/user_emails/results"

	"github.com/google/uuid"
)

// GetUserEmail 根据ID获取用户邮箱
func (s *Service) GetUserEmail(ctx context.Context, userEmailID uuid.UUID) (userEmailResult.UserEmailRO, error) {
	domainEntity, err := s.userEmailRepo.Get.ByID(ctx, userEmailID)
	if err != nil {
		return userEmailResult.UserEmailRO{}, err
	}
	return userEmailResult.UserEmailMapper(domainEntity), nil
}

// GetUserEmailByEmail 根据邮箱地址获取用户邮箱
func (s *Service) GetUserEmailByEmail(ctx context.Context, email string) (userEmailResult.UserEmailRO, error) {
	domainEntity, err := s.userEmailRepo.Get.ByEmail(ctx, email)
	if err != nil {
		return userEmailResult.UserEmailRO{}, err
	}
	return userEmailResult.UserEmailMapper(domainEntity), nil
}

// GetUserEmailsByUserID 根据用户ID获取用户邮箱列表
func (s *Service) GetUserEmailsByUserID(ctx context.Context, userID uuid.UUID) ([]userEmailResult.UserEmailRO, error) {
	domainEntities, err := s.userEmailRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	results := make([]userEmailResult.UserEmailRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userEmailResult.UserEmailMapper(entity)
	}
	return results, nil
}
