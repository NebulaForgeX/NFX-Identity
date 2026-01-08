package trusted_devices

import (
	"time"

	"github.com/google/uuid"
)

type TrustedDevice struct {
	state TrustedDeviceState
}

type TrustedDeviceState struct {
	ID                 uuid.UUID
	DeviceID           string
	UserID             uuid.UUID
	TenantID           uuid.UUID
	DeviceFingerprintHash string
	DeviceName         *string
	TrustedUntil       time.Time
	LastUsedAt         time.Time
	IP                 *string
	UAHash             *string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (td *TrustedDevice) ID() uuid.UUID              { return td.state.ID }
func (td *TrustedDevice) DeviceID() string           { return td.state.DeviceID }
func (td *TrustedDevice) UserID() uuid.UUID          { return td.state.UserID }
func (td *TrustedDevice) TenantID() uuid.UUID        { return td.state.TenantID }
func (td *TrustedDevice) DeviceFingerprintHash() string { return td.state.DeviceFingerprintHash }
func (td *TrustedDevice) DeviceName() *string        { return td.state.DeviceName }
func (td *TrustedDevice) TrustedUntil() time.Time    { return td.state.TrustedUntil }
func (td *TrustedDevice) LastUsedAt() time.Time      { return td.state.LastUsedAt }
func (td *TrustedDevice) IP() *string                { return td.state.IP }
func (td *TrustedDevice) UAHash() *string            { return td.state.UAHash }
func (td *TrustedDevice) CreatedAt() time.Time       { return td.state.CreatedAt }
func (td *TrustedDevice) UpdatedAt() time.Time       { return td.state.UpdatedAt }
