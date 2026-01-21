package system

import (
	systemstatepb "nfxid/protos/gen/system/system_state"
)

// Client System 服务客户端
type Client struct {
	SystemState *SystemStateClient
}

// NewClient 创建 System 客户端
func NewClient(systemStateClient systemstatepb.SystemStateServiceClient) *Client {
	return &Client{
		SystemState: NewSystemStateClient(systemStateClient),
	}
}