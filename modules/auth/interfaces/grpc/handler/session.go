package handler

import (
	"context"

	sessionApp "nfxid/modules/auth/application/sessions"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	sessionpb "nfxid/protos/gen/auth/session"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type SessionHandler struct {
	sessionpb.UnimplementedSessionServiceServer
	sessionAppSvc *sessionApp.Service
}

func NewSessionHandler(sessionAppSvc *sessionApp.Service) *SessionHandler {
	return &SessionHandler{
		sessionAppSvc: sessionAppSvc,
	}
}

// GetSessionByID 根据ID获取会话
func (h *SessionHandler) GetSessionByID(ctx context.Context, req *sessionpb.GetSessionByIDRequest) (*sessionpb.GetSessionByIDResponse, error) {
	sessionID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	sessionView, err := h.sessionAppSvc.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	session := mapper.SessionROToProto(&sessionView)
	return &sessionpb.GetSessionByIDResponse{Session: session}, nil
}

// GetSessionBySessionID 根据会话标识符获取会话
func (h *SessionHandler) GetSessionBySessionID(ctx context.Context, req *sessionpb.GetSessionBySessionIDRequest) (*sessionpb.GetSessionBySessionIDResponse, error) {
	sessionView, err := h.sessionAppSvc.GetSessionBySessionID(ctx, req.SessionId)
	if err != nil {
		return nil, err
	}

	session := mapper.SessionROToProto(&sessionView)
	return &sessionpb.GetSessionBySessionIDResponse{Session: session}, nil
}

// GetSessionsByUserID 根据用户ID获取会话列表
func (h *SessionHandler) GetSessionsByUserID(ctx context.Context, req *sessionpb.GetSessionsByUserIDRequest) (*sessionpb.GetSessionsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	sessionViews, err := h.sessionAppSvc.GetSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	sessions := mapper.SessionListROToProto(sessionViews)
	return &sessionpb.GetSessionsByUserIDResponse{Sessions: sessions}, nil
}

// BatchGetSessions 批量获取会话
func (h *SessionHandler) BatchGetSessions(ctx context.Context, req *sessionpb.BatchGetSessionsRequest) (*sessionpb.BatchGetSessionsResponse, error) {
	sessionIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		sessionIDs = append(sessionIDs, id)
	}

	sessions := make([]*sessionpb.Session, 0, len(sessionIDs))
	for _, sessionID := range sessionIDs {
		sessionView, err := h.sessionAppSvc.GetSession(ctx, sessionID)
		if err != nil {
			continue
		}
		session := mapper.SessionROToProto(&sessionView)
		sessions = append(sessions, session)
	}

	return &sessionpb.BatchGetSessionsResponse{Sessions: sessions}, nil
}
