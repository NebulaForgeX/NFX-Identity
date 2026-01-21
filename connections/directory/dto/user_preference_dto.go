package dto

import (
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	"google.golang.org/protobuf/types/known/structpb"
)

// CreateUserPreferenceDTO 创建用户偏好的 DTO
type CreateUserPreferenceDTO struct {
	UserID        string
	Theme         *string
	Language      *string
	Timezone      *string
	Notifications map[string]interface{}
	Privacy       map[string]interface{}
	Display       map[string]interface{}
	Other         map[string]interface{}
}

// ToCreateUserPreferenceRequest 转换为 protobuf 请求
func (d *CreateUserPreferenceDTO) ToCreateUserPreferenceRequest() (*userpreferencepb.CreateUserPreferenceRequest, error) {
	req := &userpreferencepb.CreateUserPreferenceRequest{
		UserId:   d.UserID,
		Theme:    d.Theme,
		Language: d.Language,
		Timezone: d.Timezone,
	}

	if d.Notifications != nil {
		notificationsStruct, err := structpb.NewStruct(d.Notifications)
		if err != nil {
			return nil, err
		}
		req.Notifications = notificationsStruct
	}

	if d.Privacy != nil {
		privacyStruct, err := structpb.NewStruct(d.Privacy)
		if err != nil {
			return nil, err
		}
		req.Privacy = privacyStruct
	}

	if d.Display != nil {
		displayStruct, err := structpb.NewStruct(d.Display)
		if err != nil {
			return nil, err
		}
		req.Display = displayStruct
	}

	if d.Other != nil {
		otherStruct, err := structpb.NewStruct(d.Other)
		if err != nil {
			return nil, err
		}
		req.Other = otherStruct
	}

	return req, nil
}
