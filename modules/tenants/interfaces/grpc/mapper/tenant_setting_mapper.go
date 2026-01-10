package mapper

import (
	tenantSettingAppResult "nfxid/modules/tenants/application/tenant_settings/results"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TenantSettingROToProto 将 TenantSettingRO 转换为 proto TenantSetting 消息
func TenantSettingROToProto(v *tenantSettingAppResult.TenantSettingRO) *tenantsettingpb.TenantSetting {
	if v == nil {
		return nil
	}

	tenantSetting := &tenantsettingpb.TenantSetting{
		Id:                  v.ID.String(),
		TenantId:           v.TenantID.String(),
		EnforceMfa:         v.EnforceMFA,
		AllowedEmailDomains: v.AllowedEmailDomains,
		CreatedAt:          timestamppb.New(v.CreatedAt),
		UpdatedAt:          timestamppb.New(v.UpdatedAt),
	}

	if v.SessionTTLMinutes != nil {
		ttl := int32(*v.SessionTTLMinutes)
		tenantSetting.SessionTtlMinutes = &ttl
	}

	if v.PasswordPolicy != nil && len(v.PasswordPolicy) > 0 {
		if policyStruct, err := structpb.NewStruct(v.PasswordPolicy); err == nil {
			tenantSetting.PasswordPolicy = policyStruct
		}
	}

	if v.LoginPolicy != nil && len(v.LoginPolicy) > 0 {
		if policyStruct, err := structpb.NewStruct(v.LoginPolicy); err == nil {
			tenantSetting.LoginPolicy = policyStruct
		}
	}

	if v.MFAPolicy != nil && len(v.MFAPolicy) > 0 {
		if policyStruct, err := structpb.NewStruct(v.MFAPolicy); err == nil {
			tenantSetting.MfaPolicy = policyStruct
		}
	}

	if v.UpdatedBy != nil {
		updatedBy := v.UpdatedBy.String()
		tenantSetting.UpdatedBy = &updatedBy
	}

	return tenantSetting
}
