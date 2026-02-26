package ip_allowlist

import (
	"context"
	clientsErr "nfxidentity/errors/src/clients"
	ipAllowlistCommands "nfxidentity/modules/clients/application/ip_allowlist/commands"
	ipAllowlistDomain "nfxidentity/modules/clients/domain/ip_allowlist"

	"github.com/google/uuid"
)

// CreateIPAllowlist 创建IP白名单
func (s *Service) CreateIPAllowlist(ctx context.Context, cmd ipAllowlistCommands.CreateIPAllowlistCmd) (uuid.UUID, error) {
	// Check if rule_id already exists
	if exists, _ := s.ipAllowlistRepo.Check.ByRuleID(ctx, cmd.RuleID); exists {
		return uuid.Nil, clientsErr.ErrRuleIDAlreadyExists
	}

	// Create domain entity
	ipAllowlist, err := ipAllowlistDomain.NewIPAllowlist(ipAllowlistDomain.NewIPAllowlistParams{
		RuleID:      cmd.RuleID,
		AppID:       cmd.AppID,
		CIDR:        cmd.CIDR,
		Description: cmd.Description,
		Status:      cmd.Status,
		CreatedBy:   cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.ipAllowlistRepo.Create.New(ctx, ipAllowlist); err != nil {
		return uuid.Nil, err
	}

	return ipAllowlist.ID(), nil
}
