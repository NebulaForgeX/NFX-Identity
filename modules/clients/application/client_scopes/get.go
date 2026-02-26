package client_scopes

import (
	"context"
	clientScopeResult "nfxidentity/modules/clients/application/client_scopes/results"

	"github.com/google/uuid"
)

// GetClientScope 根据ID获取客户端作用域
func (s *Service) GetClientScope(ctx context.Context, clientScopeID uuid.UUID) (clientScopeResult.ClientScopeRO, error) {
	domainEntity, err := s.clientScopeRepo.Get.ByID(ctx, clientScopeID)
	if err != nil {
		return clientScopeResult.ClientScopeRO{}, err
	}
	return clientScopeResult.ClientScopeMapper(domainEntity), nil
}
