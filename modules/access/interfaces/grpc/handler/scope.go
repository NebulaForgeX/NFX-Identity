package handler

import (
	"context"

	scopeApp "nfxid/modules/access/application/scopes"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	scopepb "nfxid/protos/gen/access/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScopeHandler struct {
	scopepb.UnimplementedScopeServiceServer
	scopeAppSvc *scopeApp.Service
}

func NewScopeHandler(scopeAppSvc *scopeApp.Service) *ScopeHandler {
	return &ScopeHandler{
		scopeAppSvc: scopeAppSvc,
	}
}

// GetScopeByScope 根据Scope标识符获取范围
func (h *ScopeHandler) GetScopeByScope(ctx context.Context, req *scopepb.GetScopeByScopeRequest) (*scopepb.GetScopeByScopeResponse, error) {
	scopeView, err := h.scopeAppSvc.GetScope(ctx, req.Scope)
	if err != nil {
		logx.S().Errorf("failed to get scope by scope: %v", err)
		return nil, status.Errorf(codes.NotFound, "scope not found: %v", err)
	}

	scope := mapper.ScopeROToProto(&scopeView)
	return &scopepb.GetScopeByScopeResponse{Scope: scope}, nil
}

// GetAllScopes 获取所有范围列表
func (h *ScopeHandler) GetAllScopes(ctx context.Context, req *scopepb.GetAllScopesRequest) (*scopepb.GetAllScopesResponse, error) {
	// 注意：domain repository 目前没有 GetAll 方法，需要先实现 repository 层的 GetAll 方法
	// 然后才能根据 is_system 进行过滤
	// 目前返回空列表，待 repository 层实现后补充
	scopes := []*scopepb.Scope{}
	return &scopepb.GetAllScopesResponse{Scopes: scopes}, nil
}

// BatchGetScopes 批量获取范围
func (h *ScopeHandler) BatchGetScopes(ctx context.Context, req *scopepb.BatchGetScopesRequest) (*scopepb.BatchGetScopesResponse, error) {
	scopes := make([]*scopepb.Scope, 0, len(req.Scopes))
	for _, scopeStr := range req.Scopes {
		scopeView, err := h.scopeAppSvc.GetScope(ctx, scopeStr)
		if err != nil {
			logx.S().Warnf("failed to get scope %s: %v", scopeStr, err)
			continue
		}
		scope := mapper.ScopeROToProto(&scopeView)
		scopes = append(scopes, scope)
	}

	return &scopepb.BatchGetScopesResponse{Scopes: scopes}, nil
}
