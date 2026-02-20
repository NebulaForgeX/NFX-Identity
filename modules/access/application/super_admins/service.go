package super_admins

import (
	"context"

	domain "nfxid/modules/access/domain/super_admins"
	"github.com/google/uuid"
)

type Service struct {
	repo *domain.Repo
}

func NewService(repo *domain.Repo) *Service {
	return &Service{repo: repo}
}

// Create 将指定用户设为超级管理员（bootstrap 或运维调用）
func (s *Service) Create(ctx context.Context, userID uuid.UUID) error {
	exists, err := s.repo.Check.ByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if exists {
		return nil // 已是 super_admin，幂等
	}
	sa := domain.NewSuperAdmin(domain.NewSuperAdminParams{UserID: userID})
	return s.repo.Create.New(ctx, sa)
}

// GetByUserID 按用户 ID 查询超级管理员
func (s *Service) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.SuperAdmin, error) {
	return s.repo.Get.ByUserID(ctx, userID)
}

// List 分页列表
func (s *Service) List(ctx context.Context, limit, offset int) ([]*domain.SuperAdmin, error) {
	return s.repo.Get.All(ctx, limit, offset)
}
