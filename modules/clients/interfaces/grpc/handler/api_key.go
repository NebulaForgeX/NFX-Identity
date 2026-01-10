package handler

import (
	"context"

	apiKeyApp "nfxid/modules/clients/application/api_keys"
	apikeypb "nfxid/protos/gen/clients/api_key"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type APIKeyHandler struct {
	apikeypb.UnimplementedApiKeyServiceServer
	apiKeyAppSvc *apiKeyApp.Service
}

func NewAPIKeyHandler(apiKeyAppSvc *apiKeyApp.Service) *APIKeyHandler {
	return &APIKeyHandler{
		apiKeyAppSvc: apiKeyAppSvc,
	}
}

// GetApiKeyByID 根据ID获取API Key
func (h *APIKeyHandler) GetApiKeyByID(ctx context.Context, req *apikeypb.GetApiKeyByIDRequest) (*apikeypb.GetApiKeyByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiKeyByID not implemented")
}
