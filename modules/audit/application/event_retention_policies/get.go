package event_retention_policies

import (
	"context"
	eventRetentionPolicyResult "nfxid/modules/audit/application/event_retention_policies/results"

	"github.com/google/uuid"
)

// GetEventRetentionPolicy 根据ID获取事件保留策略
func (s *Service) GetEventRetentionPolicy(ctx context.Context, eventRetentionPolicyID uuid.UUID) (eventRetentionPolicyResult.EventRetentionPolicyRO, error) {
	domainEntity, err := s.eventRetentionPolicyRepo.Get.ByID(ctx, eventRetentionPolicyID)
	if err != nil {
		return eventRetentionPolicyResult.EventRetentionPolicyRO{}, err
	}
	return eventRetentionPolicyResult.EventRetentionPolicyMapper(domainEntity), nil
}
