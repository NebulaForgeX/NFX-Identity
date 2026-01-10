package event_retention_policies

import (
	"context"
	eventRetentionPolicyCommands "nfxid/modules/audit/application/event_retention_policies/commands"
)

// DeleteEventRetentionPolicy 删除事件保留策略
func (s *Service) DeleteEventRetentionPolicy(ctx context.Context, cmd eventRetentionPolicyCommands.DeleteEventRetentionPolicyCmd) error {
	// Delete from repository (hard delete)
	return s.eventRetentionPolicyRepo.Delete.ByID(ctx, cmd.EventRetentionPolicyID)
}
