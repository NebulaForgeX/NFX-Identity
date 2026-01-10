package rate_limits

import (
	"context"
	rateLimitCommands "nfxid/modules/clients/application/rate_limits/commands"
)

// DeleteRateLimit 删除速率限制
func (s *Service) DeleteRateLimit(ctx context.Context, cmd rateLimitCommands.DeleteRateLimitCmd) error {
	// Delete from repository (hard delete by id)
	return s.rateLimitRepo.Delete.ByID(ctx, cmd.RateLimitID)
}
