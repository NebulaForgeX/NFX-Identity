package tenant_settings

import "github.com/google/uuid"

func (ts *TenantSetting) Validate() error {
	if ts.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	return nil
}
