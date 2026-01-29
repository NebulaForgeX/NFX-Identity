package action_requirements

import (
	"context"

	arResult "nfxid/modules/access/application/action_requirements/results"

	"github.com/google/uuid"
)

func (s *Service) GetActionRequirement(ctx context.Context, id uuid.UUID) (arResult.ActionRequirementRO, error) {
	ar, err := s.arRepo.Get.ByID(ctx, id)
	if err != nil {
		return arResult.ActionRequirementRO{}, err
	}
	return arResult.ActionRequirementMapper(ar), nil
}

func (s *Service) GetByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) (arResult.ActionRequirementRO, error) {
	ar, err := s.arRepo.Get.ByActionIDAndPermissionID(ctx, actionID, permissionID)
	if err != nil {
		return arResult.ActionRequirementRO{}, err
	}
	return arResult.ActionRequirementMapper(ar), nil
}

func (s *Service) GetByActionID(ctx context.Context, actionID uuid.UUID) ([]arResult.ActionRequirementRO, error) {
	list, err := s.arRepo.Get.ByActionID(ctx, actionID)
	if err != nil {
		return nil, err
	}
	out := make([]arResult.ActionRequirementRO, len(list))
	for i, ar := range list {
		out[i] = arResult.ActionRequirementMapper(ar)
	}
	return out, nil
}

func (s *Service) GetByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]arResult.ActionRequirementRO, error) {
	list, err := s.arRepo.Get.ByPermissionID(ctx, permissionID)
	if err != nil {
		return nil, err
	}
	out := make([]arResult.ActionRequirementRO, len(list))
	for i, ar := range list {
		out[i] = arResult.ActionRequirementMapper(ar)
	}
	return out, nil
}
