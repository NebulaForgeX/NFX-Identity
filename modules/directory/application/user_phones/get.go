package user_phones

import (
	"context"
	userPhoneResult "nfxid/modules/directory/application/user_phones/results"

	"github.com/google/uuid"
)

// GetUserPhone 根据ID获取用户手机号
func (s *Service) GetUserPhone(ctx context.Context, userPhoneID uuid.UUID) (userPhoneResult.UserPhoneRO, error) {
	domainEntity, err := s.userPhoneRepo.Get.ByID(ctx, userPhoneID)
	if err != nil {
		return userPhoneResult.UserPhoneRO{}, err
	}
	return userPhoneResult.UserPhoneMapper(domainEntity), nil
}

// GetUserPhoneByPhone 根据手机号获取用户手机
func (s *Service) GetUserPhoneByPhone(ctx context.Context, phone string) (userPhoneResult.UserPhoneRO, error) {
	domainEntity, err := s.userPhoneRepo.Get.ByPhone(ctx, phone)
	if err != nil {
		return userPhoneResult.UserPhoneRO{}, err
	}
	return userPhoneResult.UserPhoneMapper(domainEntity), nil
}

// GetUserPhonesByUserID 根据用户ID获取用户手机列表
func (s *Service) GetUserPhonesByUserID(ctx context.Context, userID uuid.UUID) ([]userPhoneResult.UserPhoneRO, error) {
	domainEntities, err := s.userPhoneRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	results := make([]userPhoneResult.UserPhoneRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userPhoneResult.UserPhoneMapper(entity)
	}
	return results, nil
}
