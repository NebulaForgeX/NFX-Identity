package auth

import "nfxid/pkgs/errx"

const (
	CodeTrustedDeviceNotFound         = "TRUSTED_DEVICE_NOT_FOUND"
	CodeDeviceIDRequired              = "DEVICE_ID_REQUIRED"
	CodeDeviceFingerprintHashRequired = "DEVICE_FINGERPRINT_HASH_REQUIRED"
	CodeTrustedUntilRequired          = "TRUSTED_UNTIL_REQUIRED"
	CodeTrustedDeviceAlreadyExists    = "TRUSTED_DEVICE_ALREADY_EXISTS"
)

var (
	ErrTrustedDeviceNotFound         = errx.NotFound(CodeTrustedDeviceNotFound, "trusted device not found")
	ErrDeviceIDRequired              = errx.InvalidArg(CodeDeviceIDRequired, "device id is required")
	ErrDeviceFingerprintHashRequired = errx.InvalidArg(CodeDeviceFingerprintHashRequired, "device fingerprint hash is required")
	ErrTrustedUntilRequired          = errx.InvalidArg(CodeTrustedUntilRequired, "trusted until is required")
	ErrTrustedDeviceAlreadyExists    = errx.Conflict(CodeTrustedDeviceAlreadyExists, "trusted device already exists")
)

/*
!TRUSTED_DEVICE_NOT_FOUND
*en<trusted device not found>
*zh<受信任设备不存在>
*fr<appareil de confiance introuvable>

!DEVICE_ID_REQUIRED
*en<device id required>
*zh<设备 ID 为必填>
*fr<id d'appareil requis>

!DEVICE_FINGERPRINT_HASH_REQUIRED
*en<device fingerprint hash required>
*zh<设备指纹哈希为必填>
*fr<hachage d'empreinte d'appareil requis>

!TRUSTED_UNTIL_REQUIRED
*en<trusted until required>
*zh<信任截止时间为必填>
*fr<date de confiance requise>

!TRUSTED_DEVICE_ALREADY_EXISTS
*en<trusted device already exists>
*zh<受信任设备已存在>
*fr<appareil de confiance existe déjà>

*/
