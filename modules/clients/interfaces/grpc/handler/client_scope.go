package handler

import (
	"context"

	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	clientscopepb "nfxid/protos/gen/clients/client_scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientScopeHandler struct {
	clientscopepb.UnimplementedClientScopeServiceServer
	clientScopeAppSvc *clientScopeApp.Service
}

func NewClientScopeHandler(clientScopeAppSvc *clientScopeApp.Service) *ClientScopeHandler {
	return &ClientScopeHandler{
		clientScopeAppSvc: clientScopeAppSvc,
	}
}

// GetClientScopeByID 根据ID获取Client Scope
func (h *ClientScopeHandler) GetClientScopeByID(ctx context.Context, req *clientscopepb.GetClientScopeByIDRequest) (*clientscopepb.GetClientScopeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientScopeByID not implemented")
}
