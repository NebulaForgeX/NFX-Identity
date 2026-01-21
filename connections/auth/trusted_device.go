package auth

import (
	"context"
	"fmt"

	trusteddevicepb "nfxid/protos/gen/auth/trusted_device"
)

// TrustedDeviceClient TrustedDevice 客户端
type TrustedDeviceClient struct {
	client trusteddevicepb.TrustedDeviceServiceClient
}

// NewTrustedDeviceClient 创建 TrustedDevice 客户端
func NewTrustedDeviceClient(client trusteddevicepb.TrustedDeviceServiceClient) *TrustedDeviceClient {
	return &TrustedDeviceClient{client: client}
}

// GetTrustedDeviceByID 根据ID获取信任设备
func (c *TrustedDeviceClient) GetTrustedDeviceByID(ctx context.Context, id string) (*trusteddevicepb.TrustedDevice, error) {
	req := &trusteddevicepb.GetTrustedDeviceByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetTrustedDeviceByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TrustedDevice, nil
}

// GetTrustedDeviceByDeviceID 根据设备ID获取信任设备
func (c *TrustedDeviceClient) GetTrustedDeviceByDeviceID(ctx context.Context, deviceID, userID, tenantID string) (*trusteddevicepb.TrustedDevice, error) {
	req := &trusteddevicepb.GetTrustedDeviceByDeviceIDRequest{
		DeviceId: deviceID,
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetTrustedDeviceByDeviceID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TrustedDevice, nil
}

// GetTrustedDevicesByUserID 根据用户ID获取信任设备列表
func (c *TrustedDeviceClient) GetTrustedDevicesByUserID(ctx context.Context, userID, tenantID string) ([]*trusteddevicepb.TrustedDevice, error) {
	req := &trusteddevicepb.GetTrustedDevicesByUserIDRequest{
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetTrustedDevicesByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TrustedDevices, nil
}