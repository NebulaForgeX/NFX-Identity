package mapper

import (
	grantAppResult "nfxid/modules/access/application/grants/results"
	grantDomain "nfxid/modules/access/domain/grants"
	grantpb "nfxid/protos/gen/access/grant"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// GrantROToProto 将 GrantRO 转换为 proto Grant 消息
func GrantROToProto(v *grantAppResult.GrantRO) *grantpb.Grant {
	if v == nil {
		return nil
	}

	grant := &grantpb.Grant{
		Id:         v.ID.String(),
		SubjectId:  v.SubjectID.String(),
		GrantRefId: v.GrantRefID.String(),
		Effect:     grantEffectToProto(v.Effect),
		CreatedAt:  timestamppb.New(v.CreatedAt),
	}

	grant.SubjectType = subjectTypeToProto(v.SubjectType)
	grant.GrantType = grantTypeToProto(v.GrantType)

	if v.TenantID != nil {
		tenantIDStr := v.TenantID.String()
		grant.TenantId = &tenantIDStr
	}
	if v.AppID != nil {
		appIDStr := v.AppID.String()
		grant.AppId = &appIDStr
	}
	if v.ResourceType != nil {
		grant.ResourceType = v.ResourceType
	}
	if v.ResourceID != nil {
		resourceIDStr := v.ResourceID.String()
		grant.ResourceId = &resourceIDStr
	}
	if v.ExpiresAt != nil {
		grant.ExpiresAt = timestamppb.New(*v.ExpiresAt)
	}
	if v.CreatedBy != nil {
		createdByStr := v.CreatedBy.String()
		grant.CreatedBy = &createdByStr
	}
	if v.RevokedAt != nil {
		grant.RevokedAt = timestamppb.New(*v.RevokedAt)
	}
	if v.RevokedBy != nil {
		revokedByStr := v.RevokedBy.String()
		grant.RevokedBy = &revokedByStr
	}
	if v.RevokeReason != nil {
		grant.RevokeReason = v.RevokeReason
	}

	return grant
}

// GrantListROToProto 批量转换 GrantRO 到 proto Grant
func GrantListROToProto(results []grantAppResult.GrantRO) []*grantpb.Grant {
	grants := make([]*grantpb.Grant, len(results))
	for i, v := range results {
		grants[i] = GrantROToProto(&v)
	}
	return grants
}

func subjectTypeToProto(subjectType grantDomain.SubjectType) grantpb.AccessSubjectType {
	switch subjectType {
	case grantDomain.SubjectTypeUser:
		return grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER
	case grantDomain.SubjectTypeClient:
		return grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT
	default:
		return grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_UNSPECIFIED
	}
}

func grantTypeToProto(grantType grantDomain.GrantType) grantpb.AccessGrantType {
	switch grantType {
	case grantDomain.GrantTypeRole:
		return grantpb.AccessGrantType_ACCESS_GRANT_TYPE_ROLE
	case grantDomain.GrantTypePermission:
		return grantpb.AccessGrantType_ACCESS_GRANT_TYPE_PERMISSION
	default:
		return grantpb.AccessGrantType_ACCESS_GRANT_TYPE_UNSPECIFIED
	}
}

func grantEffectToProto(effect grantDomain.GrantEffect) grantpb.AccessGrantEffect {
	switch effect {
	case grantDomain.GrantEffectAllow:
		return grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_ALLOW
	case grantDomain.GrantEffectDeny:
		return grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_DENY
	default:
		return grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_UNSPECIFIED
	}
}
