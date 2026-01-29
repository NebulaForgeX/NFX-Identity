package mapper

import (
	arAppResult "nfxid/modules/access/application/action_requirements/results"
	actionrequirementpb "nfxid/protos/gen/access/action_requirement"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ActionRequirementROToProto(v *arAppResult.ActionRequirementRO) *actionrequirementpb.ActionRequirement {
	if v == nil {
		return nil
	}
	return &actionrequirementpb.ActionRequirement{
		Id:           v.ID.String(),
		ActionId:     v.ActionID.String(),
		PermissionId: v.PermissionID.String(),
		GroupId:      int32(v.GroupID),
		CreatedAt:    timestamppb.New(v.CreatedAt),
	}
}

func ActionRequirementListROToProto(results []arAppResult.ActionRequirementRO) []*actionrequirementpb.ActionRequirement {
	out := make([]*actionrequirementpb.ActionRequirement, len(results))
	for i := range results {
		out[i] = ActionRequirementROToProto(&results[i])
	}
	return out
}
