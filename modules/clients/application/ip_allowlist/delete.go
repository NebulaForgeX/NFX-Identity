package ip_allowlist

import (
	"context"
	ipAllowlistCommands "nfxid/modules/clients/application/ip_allowlist/commands"
)

// DeleteIPAllowlist 删除IP白名单
func (s *Service) DeleteIPAllowlist(ctx context.Context, cmd ipAllowlistCommands.DeleteIPAllowlistCmd) error {
	// Delete from repository (hard delete by rule_id)
	return s.ipAllowlistRepo.Delete.ByRuleID(ctx, cmd.RuleID)
}
