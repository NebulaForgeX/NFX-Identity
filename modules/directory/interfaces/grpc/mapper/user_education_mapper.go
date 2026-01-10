package mapper

import (
	userEducationAppResult "nfxid/modules/directory/application/user_educations/results"
	usereducationpb "nfxid/protos/gen/directory/user_education"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserEducationROToProto 将 UserEducationRO 转换为 proto UserEducation 消息
func UserEducationROToProto(v *userEducationAppResult.UserEducationRO) *usereducationpb.UserEducation {
	if v == nil {
		return nil
	}

	userEducation := &usereducationpb.UserEducation{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		School:    v.School,
		IsCurrent: v.IsCurrent,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Degree != nil {
		userEducation.Degree = v.Degree
	}

	if v.Major != nil {
		userEducation.Major = v.Major
	}

	if v.FieldOfStudy != nil {
		userEducation.FieldOfStudy = v.FieldOfStudy
	}

	if v.StartDate != nil {
		userEducation.StartDate = timestamppb.New(*v.StartDate)
	}

	if v.EndDate != nil {
		userEducation.EndDate = timestamppb.New(*v.EndDate)
	}

	if v.Description != nil {
		userEducation.Description = v.Description
	}

	if v.Grade != nil {
		userEducation.Grade = v.Grade
	}

	if v.Activities != nil {
		userEducation.Activities = v.Activities
	}

	if v.Achievements != nil {
		userEducation.Achievements = v.Achievements
	}

	if v.DeletedAt != nil {
		userEducation.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userEducation
}

// UserEducationListROToProto 批量转换 UserEducationRO 到 proto UserEducation
func UserEducationListROToProto(results []userEducationAppResult.UserEducationRO) []*usereducationpb.UserEducation {
	userEducations := make([]*usereducationpb.UserEducation, len(results))
	for i, v := range results {
		userEducations[i] = UserEducationROToProto(&v)
	}
	return userEducations
}
