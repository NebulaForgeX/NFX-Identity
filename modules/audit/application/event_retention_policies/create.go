package event_retention_policies

import (
	"context"
	eventRetentionPolicyCommands "nfxid/modules/audit/application/event_retention_policies/commands"
	eventRetentionPolicyDomain "nfxid/modules/audit/domain/event_retention_policies"

	"github.com/google/uuid"
)

// CreateEventRetentionPolicy 创建事件保留策略
func (s *Service) CreateEventRetentionPolicy(ctx context.Context, cmd eventRetentionPolicyCommands.CreateEventRetentionPolicyCmd) (uuid.UUID, error) {
	// Check if policy name already exists
	if exists, _ := s.eventRetentionPolicyRepo.Check.ByPolicyName(ctx, cmd.PolicyName); exists {
		return uuid.Nil, eventRetentionPolicyDomain.ErrPolicyNameAlreadyExists
	}

	// Create domain entity
	eventRetentionPolicy, err := eventRetentionPolicyDomain.NewEventRetentionPolicy(eventRetentionPolicyDomain.NewEventRetentionPolicyParams{
		PolicyName:        cmd.PolicyName,
		TenantID:          cmd.TenantID,
		ActionPattern:     cmd.ActionPattern,
		DataClassification: cmd.DataClassification,
		RiskLevel:         cmd.RiskLevel,
		RetentionDays:     cmd.RetentionDays,
		RetentionAction:   cmd.RetentionAction,
		ArchiveLocation:   cmd.ArchiveLocation,
		Status:            cmd.Status,
		CreatedBy:         cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.eventRetentionPolicyRepo.Create.New(ctx, eventRetentionPolicy); err != nil {
		return uuid.Nil, err
	}

	return eventRetentionPolicy.ID(), nil
}
