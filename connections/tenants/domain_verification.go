package tenants

import (
	"context"
	"fmt"

	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"
)

// DomainVerificationClient DomainVerification 客户端
type DomainVerificationClient struct {
	client domainverificationpb.DomainVerificationServiceClient
}

// NewDomainVerificationClient 创建 DomainVerification 客户端
func NewDomainVerificationClient(client domainverificationpb.DomainVerificationServiceClient) *DomainVerificationClient {
	return &DomainVerificationClient{client: client}
}

// GetDomainVerificationByID 根据ID获取域名验证
func (c *DomainVerificationClient) GetDomainVerificationByID(ctx context.Context, id string) (*domainverificationpb.DomainVerification, error) {
	req := &domainverificationpb.GetDomainVerificationByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetDomainVerificationByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.DomainVerification, nil
}

// GetDomainVerificationByDomain 根据域名获取域名验证
func (c *DomainVerificationClient) GetDomainVerificationByDomain(ctx context.Context, domain string) (*domainverificationpb.DomainVerification, error) {
	req := &domainverificationpb.GetDomainVerificationByDomainRequest{
		Domain: domain,
	}

	resp, err := c.client.GetDomainVerificationByDomain(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.DomainVerification, nil
}

// GetDomainVerificationsByTenantID 根据租户ID获取域名验证列表
func (c *DomainVerificationClient) GetDomainVerificationsByTenantID(ctx context.Context, tenantID string, status *domainverificationpb.TenantsVerificationStatus) ([]*domainverificationpb.DomainVerification, error) {
	req := &domainverificationpb.GetDomainVerificationsByTenantIDRequest{
		TenantId: tenantID,
		Status:   status,
	}

	resp, err := c.client.GetDomainVerificationsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.DomainVerifications, nil
}