package handler

import (
	"context"

	actionApp "nfxid/modules/access/application/actions"
	actionAppCommands "nfxid/modules/access/application/actions/commands"
	actionDomain "nfxid/modules/access/domain/actions"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	actionpb "nfxid/protos/gen/access/action"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActionHandler struct {
	actionpb.UnimplementedActionServiceServer
	actionAppSvc *actionApp.Service
}

func NewActionHandler(actionAppSvc *actionApp.Service) *ActionHandler {
	return &ActionHandler{actionAppSvc: actionAppSvc}
}

func (h *ActionHandler) CreateAction(ctx context.Context, req *actionpb.CreateActionRequest) (*actionpb.CreateActionResponse, error) {
	statusVal := req.Status
	if statusVal == "" {
		statusVal = "active"
	}
	var desc *string
	if req.Description != nil {
		desc = req.Description
	}
	cmd := actionAppCommands.CreateActionCmd{
		Key:         req.Key,
		Service:     req.Service,
		Status:      statusVal,
		Name:        req.Name,
		Description: desc,
		IsSystem:    req.IsSystem,
	}
	actionID, err := h.actionAppSvc.CreateAction(ctx, cmd)
	if err != nil {
		if err == actionDomain.ErrActionKeyExists {
			return nil, status.Errorf(codes.AlreadyExists, "action key already exists: %s", req.Key)
		}
		logx.S().Errorf("failed to create action: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create action: %v", err)
	}
	view, err := h.actionAppSvc.GetAction(ctx, actionID)
	if err != nil {
		logx.S().Errorf("failed to get created action: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created action: %v", err)
	}
	return &actionpb.CreateActionResponse{Action: mapper.ActionROToProto(&view)}, nil
}

func (h *ActionHandler) GetActionByID(ctx context.Context, req *actionpb.GetActionByIDRequest) (*actionpb.GetActionByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action id: %v", err)
	}
	view, err := h.actionAppSvc.GetAction(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get action by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "action not found: %v", err)
	}
	return &actionpb.GetActionByIDResponse{Action: mapper.ActionROToProto(&view)}, nil
}

func (h *ActionHandler) GetActionByKey(ctx context.Context, req *actionpb.GetActionByKeyRequest) (*actionpb.GetActionByKeyResponse, error) {
	view, err := h.actionAppSvc.GetActionByKey(ctx, req.Key)
	if err != nil {
		logx.S().Errorf("failed to get action by key: %v", err)
		return nil, status.Errorf(codes.NotFound, "action not found: %v", err)
	}
	return &actionpb.GetActionByKeyResponse{Action: mapper.ActionROToProto(&view)}, nil
}

func (h *ActionHandler) UpdateAction(ctx context.Context, req *actionpb.UpdateActionRequest) (*actionpb.UpdateActionResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action id: %v", err)
	}
	var desc *string
	if req.Description != nil {
		desc = req.Description
	}
	cmd := actionAppCommands.UpdateActionCmd{
		ActionID:    id,
		Key:         req.Key,
		Service:     req.Service,
		Status:      req.Status,
		Name:        req.Name,
		Description: desc,
	}
	if err := h.actionAppSvc.UpdateAction(ctx, cmd); err != nil {
		logx.S().Errorf("failed to update action: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update action: %v", err)
	}
	view, err := h.actionAppSvc.GetAction(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get updated action: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get updated action: %v", err)
	}
	return &actionpb.UpdateActionResponse{Action: mapper.ActionROToProto(&view)}, nil
}

func (h *ActionHandler) DeleteAction(ctx context.Context, req *actionpb.DeleteActionRequest) (*actionpb.DeleteActionResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action id: %v", err)
	}
	if err := h.actionAppSvc.DeleteAction(ctx, actionAppCommands.DeleteActionCmd{ActionID: id}); err != nil {
		logx.S().Errorf("failed to delete action: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete action: %v", err)
	}
	return &actionpb.DeleteActionResponse{}, nil
}
