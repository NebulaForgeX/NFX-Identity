package commands

import (
	"github.com/google/uuid"
)

// CreateUserOccupationCmd 创建用户职业经历命令
type CreateUserOccupationCmd struct {
	UserID          uuid.UUID
	Company         string
	Position        string
	Department      *string
	Industry        *string
	Location        *string
	EmploymentType  *string
	StartDate       *string
	EndDate         *string
	IsCurrent       bool
	Description     *string
	Responsibilities *string
	Achievements    *string
	SkillsUsed      []string
}

// UpdateUserOccupationCmd 更新用户职业经历命令
type UpdateUserOccupationCmd struct {
	UserOccupationID uuid.UUID
	Company          string
	Position         string
	Department       *string
	Industry         *string
	Location         *string
	EmploymentType   *string
	StartDate        *string
	EndDate          *string
	IsCurrent        bool
	Description      *string
	Responsibilities *string
	Achievements     *string
	SkillsUsed       []string
}

// DeleteUserOccupationCmd 删除用户职业经历命令
type DeleteUserOccupationCmd struct {
	UserOccupationID uuid.UUID
}
