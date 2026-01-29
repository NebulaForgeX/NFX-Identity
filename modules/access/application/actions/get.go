package actions

import (
	"context"

	actionResult "nfxid/modules/access/application/actions/results"

	"github.com/google/uuid"
)

func (s *Service) GetAction(ctx context.Context, actionID uuid.UUID) (actionResult.ActionRO, error) {
	a, err := s.actionRepo.Get.ByID(ctx, actionID)
	if err != nil {
		return actionResult.ActionRO{}, err
	}
	return actionResult.ActionMapper(a), nil
}

func (s *Service) GetActionByKey(ctx context.Context, key string) (actionResult.ActionRO, error) {
	a, err := s.actionRepo.Get.ByKey(ctx, key)
	if err != nil {
		return actionResult.ActionRO{}, err
	}
	return actionResult.ActionMapper(a), nil
}
