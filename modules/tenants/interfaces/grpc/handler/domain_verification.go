package handler

import (
	"context"

	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DomainVerificationHandler struct {
	domainverificationpb.UnimplementedDomainVerificationServiceServer
	domainVerificationAppSvc *domainVerificationApp.Service
}

func NewDomainVerificationHandler(domainVerificationAppSvc *domainVerificationApp.Service) *DomainVerificationHandler {
	return &DomainVerificationHandler{
		domainVerificationAppSvc: domainVerificationAppSvc,
	}
}

// GetDomainVerificationByID 根据ID获取域名验证
func (h *DomainVerificationHandler) GetDomainVerificationByID(ctx context.Context, req *domainverificationpb.GetDomainVerificationByIDRequest) (*domainverificationpb.GetDomainVerificationByIDResponse, error) {
	domainVerificationID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid domain_verification_id: %v", err)
	}

	domainVerificationView, err := h.domainVerificationAppSvc.GetDomainVerification(ctx, domainVerificationID)
	if err != nil {
		logx.S().Errorf("failed to get domain verification by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "domain verification not found: %v", err)
	}

	domainVerification := mapper.DomainVerificationROToProto(&domainVerificationView)
	return &domainverificationpb.GetDomainVerificationByIDResponse{DomainVerification: domainVerification}, nil
}

// GetDomainVerificationByDomain 根据域名获取域名验证
func (h *DomainVerificationHandler) GetDomainVerificationByDomain(ctx context.Context, req *domainverificationpb.GetDomainVerificationByDomainRequest) (*domainverificationpb.GetDomainVerificationByDomainResponse, error) {
	domainVerificationView, err := h.domainVerificationAppSvc.GetDomainVerificationByDomain(ctx, req.Domain)
	if err != nil {
		logx.S().Errorf("failed to get domain verification by domain: %v", err)
		return nil, status.Errorf(codes.NotFound, "domain verification not found: %v", err)
	}

	domainVerification := mapper.DomainVerificationROToProto(&domainVerificationView)
	return &domainverificationpb.GetDomainVerificationByDomainResponse{DomainVerification: domainVerification}, nil
}

// GetDomainVerificationsByTenantID 根据租户ID获取域名验证列表
func (h *DomainVerificationHandler) GetDomainVerificationsByTenantID(ctx context.Context, req *domainverificationpb.GetDomainVerificationsByTenantIDRequest) (*domainverificationpb.GetDomainVerificationsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	var verificationStatus *domainVerificationDomain.VerificationStatus
	if req.Status != nil {
		statusVal := protoVerificationStatusToDomain(*req.Status)
		verificationStatus = &statusVal
	}

	domainVerificationViews, err := h.domainVerificationAppSvc.GetDomainVerificationsByTenantID(ctx, tenantID, verificationStatus)
	if err != nil {
		logx.S().Errorf("failed to get domain verifications by tenant_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get domain verifications: %v", err)
	}

	domainVerifications := mapper.DomainVerificationListROToProto(domainVerificationViews)
	return &domainverificationpb.GetDomainVerificationsByTenantIDResponse{DomainVerifications: domainVerifications}, nil
}

// protoVerificationStatusToDomain 将 proto TenantsVerificationStatus 转换为 domain VerificationStatus
func protoVerificationStatusToDomain(status domainverificationpb.TenantsVerificationStatus) domainVerificationDomain.VerificationStatus {
	switch status {
	case domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_PENDING:
		return domainVerificationDomain.VerificationStatusPending
	case domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_VERIFIED:
		return domainVerificationDomain.VerificationStatusVerified
	case domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_FAILED:
		return domainVerificationDomain.VerificationStatusFailed
	case domainverificationpb.TenantsVerificationStatus_TENANTS_VERIFICATION_STATUS_EXPIRED:
		return domainVerificationDomain.VerificationStatusExpired
	default:
		return domainVerificationDomain.VerificationStatusPending
	}
}
