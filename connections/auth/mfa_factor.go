package auth

import (
	"context"
	"fmt"

	mfafactorpb "nfxid/protos/gen/auth/mfa_factor"
)

// MfaFactorClient MfaFactor 客户端
type MfaFactorClient struct {
	client mfafactorpb.MfaFactorServiceClient
}

// NewMfaFactorClient 创建 MfaFactor 客户端
func NewMfaFactorClient(client mfafactorpb.MfaFactorServiceClient) *MfaFactorClient {
	return &MfaFactorClient{client: client}
}

// GetMfaFactorByID 根据ID获取MFA因子
func (c *MfaFactorClient) GetMfaFactorByID(ctx context.Context, id string) (*mfafactorpb.MfaFactor, error) {
	req := &mfafactorpb.GetMfaFactorByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetMfaFactorByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MfaFactor, nil
}

// GetMfaFactorByFactorID 根据因子ID获取MFA因子
func (c *MfaFactorClient) GetMfaFactorByFactorID(ctx context.Context, factorID string) (*mfafactorpb.MfaFactor, error) {
	req := &mfafactorpb.GetMfaFactorByFactorIDRequest{
		FactorId: factorID,
	}

	resp, err := c.client.GetMfaFactorByFactorID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MfaFactor, nil
}

// GetMfaFactorsByUserID 根据用户ID获取MFA因子列表
func (c *MfaFactorClient) GetMfaFactorsByUserID(ctx context.Context, userID string, enabledOnly *bool) ([]*mfafactorpb.MfaFactor, error) {
	req := &mfafactorpb.GetMfaFactorsByUserIDRequest{
		UserId:      userID,
		EnabledOnly: enabledOnly,
	}

	resp, err := c.client.GetMfaFactorsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MfaFactors, nil
}

// BatchGetMfaFactors 批量获取MFA因子
func (c *MfaFactorClient) BatchGetMfaFactors(ctx context.Context, ids []string) ([]*mfafactorpb.MfaFactor, error) {
	req := &mfafactorpb.BatchGetMfaFactorsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetMfaFactors(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MfaFactors, nil
}