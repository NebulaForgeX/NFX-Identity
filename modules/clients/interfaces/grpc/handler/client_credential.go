package handler

import (
	"context"

	clientCredentialApp "nfxidentity/modules/clients/application/client_credentials"
	"nfxidentity/pkgs/errx"
	clientcredentialpb "nfxidentity/protos/gen/clients/client_credential"
)

type ClientCredentialHandler struct {
	clientcredentialpb.UnimplementedClientCredentialServiceServer
	clientCredentialAppSvc *clientCredentialApp.Service
}

func NewClientCredentialHandler(clientCredentialAppSvc *clientCredentialApp.Service) *ClientCredentialHandler {
	return &ClientCredentialHandler{
		clientCredentialAppSvc: clientCredentialAppSvc,
	}
}

// GetClientCredentialByID 根据ID获取Client Credential
func (h *ClientCredentialHandler) GetClientCredentialByID(
	ctx context.Context,
	req *clientcredentialpb.GetClientCredentialByIDRequest,
) (*clientcredentialpb.GetClientCredentialByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetClientCredentialByID not implemented")
}
