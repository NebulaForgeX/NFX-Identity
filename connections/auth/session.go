package auth

import (
	"context"
	"fmt"

	sessionpb "nfxid/protos/gen/auth/session"
)

// SessionClient Session 客户端
type SessionClient struct {
	client sessionpb.SessionServiceClient
}

// NewSessionClient 创建 Session 客户端
func NewSessionClient(client sessionpb.SessionServiceClient) *SessionClient {
	return &SessionClient{client: client}
}

// GetSessionByID 根据ID获取会话
func (c *SessionClient) GetSessionByID(ctx context.Context, id string) (*sessionpb.Session, error) {
	req := &sessionpb.GetSessionByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetSessionByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Session, nil
}

// GetSessionBySessionID 根据会话标识符获取会话
func (c *SessionClient) GetSessionBySessionID(ctx context.Context, sessionID string) (*sessionpb.Session, error) {
	req := &sessionpb.GetSessionBySessionIDRequest{
		SessionId: sessionID,
	}

	resp, err := c.client.GetSessionBySessionID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Session, nil
}

// GetSessionsByUserID 根据用户ID获取会话列表
func (c *SessionClient) GetSessionsByUserID(ctx context.Context, userID string) ([]*sessionpb.Session, error) {
	req := &sessionpb.GetSessionsByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetSessionsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Sessions, nil
}

// BatchGetSessions 批量获取会话
func (c *SessionClient) BatchGetSessions(ctx context.Context, ids []string) ([]*sessionpb.Session, error) {
	req := &sessionpb.BatchGetSessionsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetSessions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Sessions, nil
}