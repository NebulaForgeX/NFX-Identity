package tenant_roles

import (
	"context"

	"github.com/google/uuid"
)

// Update 更新租户角色
func (s *Service) Update(ctx context.Context, id uuid.UUID, roleKey string, name *string) error {
	r, err := s.repo.Get.ByID(ctx, id)
	if err != nil {
		return err
	}
	if err := r.Update(roleKey, name); err != nil {
		return err
	}
	return s.repo.Update.Generic(ctx, r)
}

