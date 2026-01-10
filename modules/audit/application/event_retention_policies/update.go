package event_retention_policies

import (
	"context"
	eventRetentionPolicyCommands "nfxid/modules/audit/application/event_retention_policies/commands"
)

// UpdateEventRetentionPolicy 更新事件保留策略
func (s *Service) UpdateEventRetentionPolicy(ctx context.Context, cmd eventRetentionPolicyCommands.UpdateEventRetentionPolicyCmd) error {
	// Get domain entity
	eventRetentionPolicy, err := s.eventRetentionPolicyRepo.Get.ByID(ctx, cmd.EventRetentionPolicyID)
	if err != nil {
		return err
	}

	// Update domain entity
	retentionDays := cmd.RetentionDays
	if retentionDays <= 0 {
		retentionDays = eventRetentionPolicy.RetentionDays() // Keep existing value
	}
	if err := eventRetentionPolicy.Update(cmd.ActionPattern, cmd.DataClassification, cmd.RiskLevel, retentionDays, cmd.RetentionAction, cmd.ArchiveLocation); err != nil {
		return err
	}

	// Save to repository
	return s.eventRetentionPolicyRepo.Update.Generic(ctx, eventRetentionPolicy)
}

// UpdateEventRetentionPolicyStatus 更新事件保留策略状态
func (s *Service) UpdateEventRetentionPolicyStatus(ctx context.Context, cmd eventRetentionPolicyCommands.UpdateEventRetentionPolicyStatusCmd) error {
	// Get domain entity
	eventRetentionPolicy, err := s.eventRetentionPolicyRepo.Get.ByID(ctx, cmd.EventRetentionPolicyID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := eventRetentionPolicy.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.eventRetentionPolicyRepo.Update.Status(ctx, cmd.EventRetentionPolicyID, cmd.Status)
}
