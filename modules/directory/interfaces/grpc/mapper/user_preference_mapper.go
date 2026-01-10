package mapper

import (
	userPreferenceAppResult "nfxid/modules/directory/application/user_preferences/results"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserPreferenceROToProto 将 UserPreferenceRO 转换为 proto UserPreference 消息
func UserPreferenceROToProto(v *userPreferenceAppResult.UserPreferenceRO) *userpreferencepb.UserPreference {
	if v == nil {
		return nil
	}

	userPreference := &userpreferencepb.UserPreference{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Theme != "" {
		userPreference.Theme = &v.Theme
	}

	if v.Language != "" {
		userPreference.Language = &v.Language
	}

	if v.Timezone != "" {
		userPreference.Timezone = &v.Timezone
	}

	if v.Notifications != nil && len(v.Notifications) > 0 {
		if notificationsStruct, err := structpb.NewStruct(v.Notifications); err == nil {
			userPreference.Notifications = notificationsStruct
		}
	}

	if v.Privacy != nil && len(v.Privacy) > 0 {
		if privacyStruct, err := structpb.NewStruct(v.Privacy); err == nil {
			userPreference.Privacy = privacyStruct
		}
	}

	if v.Display != nil && len(v.Display) > 0 {
		if displayStruct, err := structpb.NewStruct(v.Display); err == nil {
			userPreference.Display = displayStruct
		}
	}

	if v.Other != nil && len(v.Other) > 0 {
		if otherStruct, err := structpb.NewStruct(v.Other); err == nil {
			userPreference.Other = otherStruct
		}
	}

	if v.DeletedAt != nil {
		userPreference.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userPreference
}
