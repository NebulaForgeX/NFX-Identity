package event_retention_policies

import (
	eventRetentionPolicyDomain "nfxidentity/modules/audit/domain/event_retention_policies"
)

type Service struct {
	eventRetentionPolicyRepo *eventRetentionPolicyDomain.Repo
}

func NewService(
	eventRetentionPolicyRepo *eventRetentionPolicyDomain.Repo,
) *Service {
	return &Service{
		eventRetentionPolicyRepo: eventRetentionPolicyRepo,
	}
}
