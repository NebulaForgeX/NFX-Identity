package scopes

import (
	"context"
	scopeResult "nfxid/modules/access/application/scopes/results"
)

// GetScope 根据Scope获取作用域
func (s *Service) GetScope(ctx context.Context, scope string) (scopeResult.ScopeRO, error) {
	domainEntity, err := s.scopeRepo.Get.ByScope(ctx, scope)
	if err != nil {
		return scopeResult.ScopeRO{}, err
	}
	return scopeResult.ScopeMapper(domainEntity), nil
}
