package mapper

import (
	domainVerificationAppResult "nfxid/modules/tenants/application/domain_verifications/results"
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DomainVerificationROToProto 将 DomainVerificationRO 转换为 proto DomainVerification 消息
func DomainVerificationROToProto(v *domainVerificationAppResult.DomainVerificationRO) *domainverificationpb.DomainVerification {
	if v == nil {
		return nil
	}

	domainVerification := &domainverificationpb.DomainVerification{
		Id:                 v.ID.String(),
		TenantId:           v.TenantID.String(),
		Domain:             v.Domain,
		VerificationMethod: verificationMethodToProto(v.VerificationMethod),
		Status:             verificationStatusToProto(v.Status),
		CreatedAt:          timestamppb.New(v.CreatedAt),
	}

	if v.VerificationToken != nil {
		domainVerification.VerificationToken = v.VerificationToken
	}

	if v.VerifiedAt != nil {
		domainVerification.VerifiedAt = timestamppb.New(*v.VerifiedAt)
	}

	if v.ExpiresAt != nil {
		domainVerification.ExpiresAt = timestamppb.New(*v.ExpiresAt)
	}

	if v.CreatedBy != nil {
		createdBy := v.CreatedBy.String()
		domainVerification.CreatedBy = &createdBy
	}

	if v.Metadata != nil && len(v.Metadata) > 0 {
		if metadataStruct, err := structpb.NewStruct(v.Metadata); err == nil {
			domainVerification.Metadata = metadataStruct
		}
	}

	return domainVerification
}

// DomainVerificationListROToProto 批量转换 DomainVerificationRO 到 proto DomainVerification
func DomainVerificationListROToProto(results []domainVerificationAppResult.DomainVerificationRO) []*domainverificationpb.DomainVerification {
	domainVerifications := make([]*domainverificationpb.DomainVerification, len(results))
	for i, v := range results {
		domainVerifications[i] = DomainVerificationROToProto(&v)
	}
	return domainVerifications
}

// verificationMethodToProto 将 domain VerificationMethod 转换为 proto TenantsVerificationMethod
func verificationMethodToProto(method domainVerificationDomain.VerificationMethod) domainverificationpb.TenantsVerificationMethod {
	switch method {
	case domainVerificationDomain.VerificationMethodDNS:
		return domainverificationpb.TenantsVerificationMethod_TENANTS_VERIFICATION_METHOD_DNS
	case domainVerificationDomain.VerificationMethodTXT:
		return domainverificationpb.TenantsVerificationMethod_TENANTS_VERIFICATION_METHOD_TXT
	case domainVerificationDomain.VerificationMethodHTML:
		return domainverificationpb.TenantsVerificationMethod_TENANTS_VERIFICATION_METHOD_HTML
	case domainVerificationDomain.VerificationMethodFILE:
		return domainverificationpb.TenantsVerificationMethod_TENANTS_VERIFICATION_METHOD_FILE
	default:
		return domainverificationpb.TenantsVerificationMethod_TENANTS_VERIFICATION_METHOD_UNSPECIFIED
	}
}

// verificationStatusToProto 将 domain VerificationStatus 转换为 proto TenantsVerificationStatus
func verificationStatusToProto(status domainVerificationDomain.VerificationStatus) domainverificationpb.TenantsVerificationStatus {
	switch status {
	case domainVerificationDomain.VerificationStatusPending:
		return domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_PENDING
	case domainVerificationDomain.VerificationStatusVerified:
		return domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_VERIFIED
	case domainVerificationDomain.VerificationStatusFailed:
		return domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_FAILED
	case domainVerificationDomain.VerificationStatusExpired:
		return domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_EXPIRED
	default:
		return domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_UNSPECIFIED
	}
}
