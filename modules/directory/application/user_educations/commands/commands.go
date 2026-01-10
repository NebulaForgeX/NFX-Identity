package commands

import (
	"github.com/google/uuid"
)

// CreateUserEducationCmd 创建用户教育经历命令
type CreateUserEducationCmd struct {
	UserID      uuid.UUID
	School      string
	Degree      *string
	Major       *string
	FieldOfStudy *string
	StartDate   *string
	EndDate     *string
	IsCurrent   bool
	Description *string
	Grade       *string
	Activities  *string
	Achievements *string
}

// UpdateUserEducationCmd 更新用户教育经历命令
type UpdateUserEducationCmd struct {
	UserEducationID uuid.UUID
	School          string
	Degree          *string
	Major           *string
	FieldOfStudy    *string
	StartDate       *string
	EndDate         *string
	IsCurrent       bool
	Description     *string
	Grade           *string
	Activities      *string
	Achievements    *string
}

// DeleteUserEducationCmd 删除用户教育经历命令
type DeleteUserEducationCmd struct {
	UserEducationID uuid.UUID
}
