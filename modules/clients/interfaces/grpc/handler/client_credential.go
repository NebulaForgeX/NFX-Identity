package handler

import (
	"context"

	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientcredentialpb "nfxid/protos/gen/clients/client_credential"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (h *ClientCredentialHandler) GetClientCredentialByID(ctx context.Context, req *clientcredentialpb.GetClientCredentialByIDRequest) (*clientcredentialpb.GetClientCredentialByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientCredentialByID not implemented")
}
