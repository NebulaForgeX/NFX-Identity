package commands

import (
	"nfxid/modules/auth/domain/login_attempts"

	"github.com/google/uuid"
)

// CreateLoginAttemptCmd 创建登录尝试命令
type CreateLoginAttemptCmd struct {
	TenantID          uuid.UUID
	Identifier        string
	UserID            *uuid.UUID
	IP                *string
	UAHash            *string
	DeviceFingerprint *string
	Success           bool
	FailureCode       *login_attempts.FailureCode
	MFARequired       bool
	MFAVerified       bool
}

// DeleteLoginAttemptCmd 删除登录尝试命令
type DeleteLoginAttemptCmd struct {
	LoginAttemptID uuid.UUID
}
