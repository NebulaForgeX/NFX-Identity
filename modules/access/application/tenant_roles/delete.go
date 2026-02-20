package tenant_roles

import (
	"context"

	"github.com/google/uuid"
)

// DeleteByID 按 ID 删除
func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete.ByID(ctx, id)
}
