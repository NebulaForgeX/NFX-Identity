package mapper

import (
	userOccupationAppResult "nfxid/modules/directory/application/user_occupations/results"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserOccupationROToProto 将 UserOccupationRO 转换为 proto UserOccupation 消息
func UserOccupationROToProto(v *userOccupationAppResult.UserOccupationRO) *useroccupationpb.UserOccupation {
	if v == nil {
		return nil
	}

	userOccupation := &useroccupationpb.UserOccupation{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		Company:   v.Company,
		Position:  v.Position,
		IsCurrent: v.IsCurrent,
		SkillsUsed: v.SkillsUsed,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Department != nil {
		userOccupation.Department = v.Department
	}

	if v.Industry != nil {
		userOccupation.Industry = v.Industry
	}

	if v.Location != nil {
		userOccupation.Location = v.Location
	}

	if v.EmploymentType != nil {
		userOccupation.EmploymentType = v.EmploymentType
	}

	if v.StartDate != nil {
		userOccupation.StartDate = timestamppb.New(*v.StartDate)
	}

	if v.EndDate != nil {
		userOccupation.EndDate = timestamppb.New(*v.EndDate)
	}

	if v.Description != nil {
		userOccupation.Description = v.Description
	}

	if v.Responsibilities != nil {
		userOccupation.Responsibilities = v.Responsibilities
	}

	if v.Achievements != nil {
		userOccupation.Achievements = v.Achievements
	}

	if v.DeletedAt != nil {
		userOccupation.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userOccupation
}

// UserOccupationListROToProto 批量转换 UserOccupationRO 到 proto UserOccupation
func UserOccupationListROToProto(results []userOccupationAppResult.UserOccupationRO) []*useroccupationpb.UserOccupation {
	userOccupations := make([]*useroccupationpb.UserOccupation, len(results))
	for i, v := range results {
		userOccupations[i] = UserOccupationROToProto(&v)
	}
	return userOccupations
}
