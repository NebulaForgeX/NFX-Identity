package profile

import (
	"context"

	"github.com/google/uuid"
)

type DeleteProfileCmd struct {
	ProfileID uuid.UUID
}

func (s *Service) DeleteProfile(ctx context.Context, cmd DeleteProfileCmd) error {
	p, err := s.profileRepo.Get.ByID(ctx, cmd.ProfileID)
	if err != nil {
		return err
	}

	if err := p.Delete(); err != nil {
		return err
	}

	if err := s.profileRepo.Update.Generic(ctx, p); err != nil {
		return err
	}

	return nil
}

// DeleteByUserID 根据 UserID 删除 Profile（用于事件处理）
func (s *Service) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	// 检查 profile 是否存在
	exists, err := s.profileRepo.Check.ByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if !exists {
		// profile 不存在，直接返回成功
		return nil
	}

	// 获取 profile
	p, err := s.profileRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// 使用 domain 的软删除方法
	if err := p.Delete(); err != nil {
		return err
	}

	// 更新 profile（软删除）
	if err := s.profileRepo.Update.Generic(ctx, p); err != nil {
		return err
	}

	return nil
}
