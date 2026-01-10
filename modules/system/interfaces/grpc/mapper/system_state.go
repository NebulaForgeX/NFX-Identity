package mapper

import (
	systemStateApp "nfxid/modules/system/application/system_state/results"
	systemstatepb "nfxid/protos/gen/system/system_state"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SystemStateROToProto 将 Application SystemStateRO 转换为 Protobuf SystemState
func SystemStateROToProto(ro *systemStateApp.SystemStateRO) *systemstatepb.SystemState {
	if ro == nil {
		return nil
	}

	pb := &systemstatepb.SystemState{
		Id:          ro.ID.String(),
		Initialized: ro.Initialized,
		ResetCount:  int32(ro.ResetCount),
	}

	if ro.InitializedAt != nil {
		pb.InitializedAt = timestamppb.New(*ro.InitializedAt)
	}

	if ro.InitializationVersion != nil {
		pb.InitializationVersion = ro.InitializationVersion
	}

	if ro.LastResetAt != nil {
		pb.LastResetAt = timestamppb.New(*ro.LastResetAt)
	}

	if ro.LastResetBy != nil {
		resetByStr := ro.LastResetBy.String()
		pb.LastResetBy = &resetByStr
	}

	if ro.Metadata != nil {
		if metadataStruct, err := structpb.NewStruct(ro.Metadata); err == nil {
			pb.Metadata = metadataStruct
		}
	}

	pb.CreatedAt = timestamppb.New(ro.CreatedAt)
	pb.UpdatedAt = timestamppb.New(ro.UpdatedAt)

	return pb
}

// SystemStateListROToProto 将 Application SystemStateRO 列表转换为 Protobuf SystemState 列表
func SystemStateListROToProto(ros []systemStateApp.SystemStateRO) []*systemstatepb.SystemState {
	if len(ros) == 0 {
		return nil
	}

	pbs := make([]*systemstatepb.SystemState, len(ros))
	for i := range ros {
		pbs[i] = SystemStateROToProto(&ros[i])
	}
	return pbs
}
