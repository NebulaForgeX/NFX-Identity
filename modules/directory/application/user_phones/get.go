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
