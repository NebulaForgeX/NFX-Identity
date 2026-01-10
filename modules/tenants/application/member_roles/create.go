package member_roles

import (
	"context"
	"time"
	memberRoleCommands "nfxid/modules/tenants/application/member_roles/commands"
	memberRoleDomain "nfxid/modules/tenants/domain/member_roles"

	"github.com/google/uuid"
)

// CreateMemberRole 创建成员角色
func (s *Service) CreateMemberRole(ctx context.Context, cmd memberRoleCommands.CreateMemberRoleCmd) (uuid.UUID, error) {
	// Check if member role already exists
	if exists, _ := s.memberRoleRepo.Check.ByTenantIDAndMemberIDAndRoleID(ctx, cmd.TenantID, cmd.MemberID, cmd.RoleID); exists {
		return uuid.Nil, memberRoleDomain.ErrMemberRoleAlreadyExists
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
	memberRole, err := memberRoleDomain.NewMemberRole(memberRoleDomain.NewMemberRoleParams{
		TenantID:   cmd.TenantID,
		MemberID:   cmd.MemberID,
		RoleID:     cmd.RoleID,
		AssignedBy: cmd.AssignedBy,
		ExpiresAt:  expiresAt,
		Scope:      cmd.Scope,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.memberRoleRepo.Create.New(ctx, memberRole); err != nil {
		return uuid.Nil, err
	}

	return memberRole.ID(), nil
}
