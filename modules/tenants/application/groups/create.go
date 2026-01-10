package groups

import (
	"context"
	groupCommands "nfxid/modules/tenants/application/groups/commands"
	groupDomain "nfxid/modules/tenants/domain/groups"

	"github.com/google/uuid"
)

// CreateGroup 创建组
func (s *Service) CreateGroup(ctx context.Context, cmd groupCommands.CreateGroupCmd) (uuid.UUID, error) {
	// Check if group ID already exists
	if exists, _ := s.groupRepo.Check.ByGroupID(ctx, cmd.GroupID); exists {
		return uuid.Nil, groupDomain.ErrGroupIDAlreadyExists
	}

	// Create domain entity
	group, err := groupDomain.NewGroup(groupDomain.NewGroupParams{
		GroupID:       cmd.GroupID,
		TenantID:      cmd.TenantID,
		Name:          cmd.Name,
		Type:          cmd.Type,
		ParentGroupID: cmd.ParentGroupID,
		Description:   cmd.Description,
		CreatedBy:     cmd.CreatedBy,
		Metadata:      cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.groupRepo.Create.New(ctx, group); err != nil {
		return uuid.Nil, err
	}

	return group.ID(), nil
}
