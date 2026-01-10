package client_credentials

import (
	"context"
	clientCredentialResult "nfxid/modules/clients/application/client_credentials/results"

	"github.com/google/uuid"
)

// GetClientCredential 根据ID获取客户端凭证
func (s *Service) GetClientCredential(ctx context.Context, clientCredentialID uuid.UUID) (clientCredentialResult.ClientCredentialRO, error) {
	domainEntity, err := s.clientCredentialRepo.Get.ByID(ctx, clientCredentialID)
	if err != nil {
		return clientCredentialResult.ClientCredentialRO{}, err
	}
	return clientCredentialResult.ClientCredentialMapper(domainEntity), nil
}

// GetClientCredentialByClientID 根据ClientID获取客户端凭证
func (s *Service) GetClientCredentialByClientID(ctx context.Context, clientID string) (clientCredentialResult.ClientCredentialRO, error) {
	domainEntity, err := s.clientCredentialRepo.Get.ByClientID(ctx, clientID)
	if err != nil {
		return clientCredentialResult.ClientCredentialRO{}, err
	}
	return clientCredentialResult.ClientCredentialMapper(domainEntity), nil
}
