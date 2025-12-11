package mapper

import (
	occupationAppViews "nfxid/modules/auth/application/occupation/views"
	occupationpb "nfxid/protos/gen/auth/occupation"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// OccupationViewToProto 将 OccupationView 转换为 proto Occupation 消息
func OccupationViewToProto(v *occupationAppViews.OccupationView) *occupationpb.Occupation {
	if v == nil {
		return nil
	}

	occupation := &occupationpb.Occupation{
		Id:        v.ID.String(),
		ProfileId: v.ProfileID.String(),
		Company:   v.Company,
		Position:  v.Position,
		IsCurrent: v.IsCurrent,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Department != nil {
		occupation.Department = v.Department
	}
	if v.Industry != nil {
		occupation.Industry = v.Industry
	}
	if v.Location != nil {
		occupation.Location = v.Location
	}
	if v.EmploymentType != nil {
		occupation.EmploymentType = v.EmploymentType
	}
	if v.StartDate != nil {
		startDateStr := v.StartDate.Format("2006-01-02")
		occupation.StartDate = &startDateStr
	}
	if v.EndDate != nil {
		endDateStr := v.EndDate.Format("2006-01-02")
		occupation.EndDate = &endDateStr
	}
	if v.Description != nil {
		occupation.Description = v.Description
	}
	if v.Responsibilities != nil {
		occupation.Responsibilities = v.Responsibilities
	}
	if v.Achievements != nil {
		occupation.Achievements = v.Achievements
	}
	if v.SkillsUsed != nil {
		// Note: SkillsUsed is *string, but proto expects []string
		// For now, leave it empty or implement conversion if needed
		occupation.SkillsUsed = []string{}
	}
	if v.DeletedAt != nil {
		occupation.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return occupation
}
