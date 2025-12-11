package mapper

import (
	educationAppViews "nfxid/modules/auth/application/profile_education/views"
	profileeducationpb "nfxid/protos/gen/auth/profile_education"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// EducationViewToProto 将 EducationView 转换为 proto ProfileEducation 消息
func EducationViewToProto(v *educationAppViews.EducationView) *profileeducationpb.ProfileEducation {
	if v == nil {
		return nil
	}

	education := &profileeducationpb.ProfileEducation{
		Id:        v.ID.String(),
		ProfileId: v.ProfileID.String(),
		School:    v.School,
		IsCurrent: v.IsCurrent,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Degree != nil {
		education.Degree = v.Degree
	}
	if v.Major != nil {
		education.Major = v.Major
	}
	if v.FieldOfStudy != nil {
		education.FieldOfStudy = v.FieldOfStudy
	}
	if v.StartDate != nil {
		startDateStr := v.StartDate.Format("2006-01-02")
		education.StartDate = &startDateStr
	}
	if v.EndDate != nil {
		endDateStr := v.EndDate.Format("2006-01-02")
		education.EndDate = &endDateStr
	}
	if v.Description != nil {
		education.Description = v.Description
	}
	if v.Grade != nil {
		education.Grade = v.Grade
	}
	if v.Activities != nil {
		education.Activities = v.Activities
	}
	if v.Achievements != nil {
		education.Achievements = v.Achievements
	}
	if v.DeletedAt != nil {
		education.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return education
}
