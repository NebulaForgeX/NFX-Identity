package mapper

import (
	scopeAppResult "nfxid/modules/access/application/scopes/results"
	scopepb "nfxid/protos/gen/access/scope"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ScopeROToProto 将 ScopeRO 转换为 proto Scope 消息
func ScopeROToProto(v *scopeAppResult.ScopeRO) *scopepb.Scope {
	if v == nil {
		return nil
	}

	scope := &scopepb.Scope{
		Scope:     v.Scope,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		scope.Description = v.Description
	}

	if v.DeletedAt != nil {
		scope.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return scope
}

// ScopeListROToProto 批量转换 ScopeRO 到 proto Scope
func ScopeListROToProto(results []scopeAppResult.ScopeRO) []*scopepb.Scope {
	scopes := make([]*scopepb.Scope, len(results))
	for i, v := range results {
		scopes[i] = ScopeROToProto(&v)
	}
	return scopes
}
