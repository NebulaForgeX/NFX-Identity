package mapper

import (
	actionAppResult "nfxid/modules/access/application/actions/results"
	actionpb "nfxid/protos/gen/access/action"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ActionROToProto(v *actionAppResult.ActionRO) *actionpb.Action {
	if v == nil {
		return nil
	}
	out := &actionpb.Action{
		Id:        v.ID.String(),
		Key:       v.Key,
		Service:   v.Service,
		Status:    v.Status,
		Name:      v.Name,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}
	if v.Description != nil {
		out.Description = v.Description
	}
	if v.DeletedAt != nil {
		out.DeletedAt = timestamppb.New(*v.DeletedAt)
	}
	return out
}
