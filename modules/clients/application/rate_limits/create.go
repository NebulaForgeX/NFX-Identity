package rate_limits

import (
	"context"
	rateLimitCommands "nfxid/modules/clients/application/rate_limits/commands"
	rateLimitDomain "nfxid/modules/clients/domain/rate_limits"

	"github.com/google/uuid"
)

// CreateRateLimit 创建速率限制
func (s *Service) CreateRateLimit(ctx context.Context, cmd rateLimitCommands.CreateRateLimitCmd) (uuid.UUID, error) {
	// Check if rate limit already exists for this app and limit type
	if exists, _ := s.rateLimitRepo.Check.ByAppIDAndLimitType(ctx, cmd.AppID, cmd.LimitType); exists {
		return uuid.Nil, rateLimitDomain.ErrRateLimitAlreadyExists
	}

	// Create domain entity
	rateLimit, err := rateLimitDomain.NewRateLimit(rateLimitDomain.NewRateLimitParams{
		AppID:         cmd.AppID,
		LimitType:     cmd.LimitType,
		LimitValue:    cmd.LimitValue,
		WindowSeconds: cmd.WindowSeconds,
		Description:   cmd.Description,
		Status:        cmd.Status,
		CreatedBy:     cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.rateLimitRepo.Create.New(ctx, rateLimit); err != nil {
		return uuid.Nil, err
	}

	return rateLimit.ID(), nil
}
