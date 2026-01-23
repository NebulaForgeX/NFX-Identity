package mapper

import (
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// TrustedDeviceDomainToModel 将 Domain TrustedDevice 转换为 Model TrustedDevice
func TrustedDeviceDomainToModel(td *trusted_devices.TrustedDevice) *models.TrustedDevice {
	if td == nil {
		return nil
	}

	return &models.TrustedDevice{
		ID:                    td.ID(),
		DeviceID:              td.DeviceID(),
		UserID:                td.UserID(),
		DeviceFingerprintHash: td.DeviceFingerprintHash(),
		DeviceName:            td.DeviceName(),
		TrustedUntil:          td.TrustedUntil(),
		LastUsedAt:            td.LastUsedAt(),
		IP:                    td.IP(),
		UaHash:                td.UAHash(), // Model 使用 UaHash，Domain 使用 UAHash
		CreatedAt:             td.CreatedAt(),
		UpdatedAt:             td.UpdatedAt(),
	}
}

// TrustedDeviceModelToDomain 将 Model TrustedDevice 转换为 Domain TrustedDevice
func TrustedDeviceModelToDomain(m *models.TrustedDevice) *trusted_devices.TrustedDevice {
	if m == nil {
		return nil
	}

	state := trusted_devices.TrustedDeviceState{
		ID:                    m.ID,
		DeviceID:              m.DeviceID,
		UserID:                m.UserID,
		DeviceFingerprintHash: m.DeviceFingerprintHash,
		DeviceName:            m.DeviceName,
		TrustedUntil:          m.TrustedUntil,
		LastUsedAt:            m.LastUsedAt,
		IP:                    m.IP,
		UAHash:                m.UaHash, // Model 使用 UaHash，Domain 使用 UAHash
		CreatedAt:             m.CreatedAt,
		UpdatedAt:             m.UpdatedAt,
	}

	return trusted_devices.NewTrustedDeviceFromState(state)
}

// TrustedDeviceModelToUpdates 将 Model TrustedDevice 转换为更新字段映射
func TrustedDeviceModelToUpdates(m *models.TrustedDevice) map[string]any {
	return map[string]any{
		models.TrustedDeviceCols.DeviceFingerprintHash: m.DeviceFingerprintHash,
		models.TrustedDeviceCols.DeviceName:            m.DeviceName,
		models.TrustedDeviceCols.TrustedUntil:          m.TrustedUntil,
		models.TrustedDeviceCols.LastUsedAt:            m.LastUsedAt,
		models.TrustedDeviceCols.IP:                    m.IP,
		models.TrustedDeviceCols.UaHash:                m.UaHash,
		models.TrustedDeviceCols.UpdatedAt:             m.UpdatedAt,
	}
}
