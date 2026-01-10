package commands

import (
	"github.com/google/uuid"
)

// CreateTrustedDeviceCmd 创建受信任设备命令
type CreateTrustedDeviceCmd struct {
	DeviceID             string
	UserID               uuid.UUID
	TenantID             uuid.UUID
	DeviceFingerprintHash string
	DeviceName           *string
	TrustedUntil         string
	IP                   *string
	UAHash               *string
}

// UpdateLastUsedCmd 更新最后使用时间命令
type UpdateLastUsedCmd struct {
	DeviceID string
}

// UpdateTrustedUntilCmd 更新信任到期时间命令
type UpdateTrustedUntilCmd struct {
	DeviceID     string
	TrustedUntil string
}

// DeleteTrustedDeviceCmd 删除受信任设备命令
type DeleteTrustedDeviceCmd struct {
	DeviceID string
}
