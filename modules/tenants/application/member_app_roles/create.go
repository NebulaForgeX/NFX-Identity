package member_app_roles

import (
	"context"
	"time"
	memberAppRoleCommands "nfxid/modules/tenants/application/member_app_roles/commands"
	memberAppRoleDomain "nfxid/modules/tenants/domain/member_app_roles"

	"github.com/google/uuid"
)

// CreateMemberAppRole 创建成员应用角色
func (s *Service) CreateMemberAppRole(ctx context.Context, cmd memberAppRoleCommands.CreateMemberAppRoleCmd) (uuid.UUID, error) {
	// Check if member app role already exists
	if exists, _ := s.memberAppRoleRepo.Check.ByMemberIDAndAppIDAndRoleID(ctx, cmd.MemberID, cmd.AppID, cmd.RoleID); exists {
		return uuid.Nil, memberAppRoleDomain.ErrMemberAppRoleAlreadyExists
	}

	var expiresAt *time.Time
	if cmd.ExpiresAt != nil && *cmd.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.ExpiresAt)
		if err != nil {
			return uuid.Nil, err
		}
		expiresAt = &parsed
	}

	// Create domain entity
	memberAppRole, err := memberAppRoleDomain.NewMemberAppRole(memberAppRoleDomain.NewMemberAppRoleParams{
		MemberID:   cmd.MemberID,
		AppID:      cmd.AppID,
		RoleID:     cmd.RoleID,
		AssignedBy: cmd.AssignedBy,
		ExpiresAt:  expiresAt,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.memberAppRoleRepo.Create.New(ctx, memberAppRole); err != nil {
		return uuid.Nil, err
	}

	return memberAppRole.ID(), nil
}
