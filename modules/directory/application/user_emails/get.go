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
