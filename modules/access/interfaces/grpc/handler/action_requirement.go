package handler

import (
	"context"

	arApp "nfxid/modules/access/application/action_requirements"
	arAppCommands "nfxid/modules/access/application/action_requirements/commands"
	arDomain "nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	actionrequirementpb "nfxid/protos/gen/access/action_requirement"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActionRequirementHandler struct {
	actionrequirementpb.UnimplementedActionRequirementServiceServer
	arAppSvc *arApp.Service
}

func NewActionRequirementHandler(arAppSvc *arApp.Service) *ActionRequirementHandler {
	return &ActionRequirementHandler{arAppSvc: arAppSvc}
}

func (h *ActionRequirementHandler) CreateActionRequirement(ctx context.Context, req *actionrequirementpb.CreateActionRequirementRequest) (*actionrequirementpb.CreateActionRequirementResponse, error) {
	actionID, err := uuid.Parse(req.ActionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action_id: %v", err)
	}
	permissionID, err := uuid.Parse(req.PermissionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}
	groupID := int(req.GroupId)
	if groupID == 0 {
		groupID = 1
	}
	cmd := arAppCommands.CreateActionRequirementCmd{
		ActionID:     actionID,
		PermissionID: permissionID,
		GroupID:      groupID,
	}
	id, err := h.arAppSvc.CreateActionRequirement(ctx, cmd)
	if err != nil {
		if err == arDomain.ErrActionRequirementAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "action requirement already exists")
		}
		logx.S().Errorf("failed to create action requirement: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create action requirement: %v", err)
	}
	view, err := h.arAppSvc.GetActionRequirement(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get created action requirement: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created action requirement: %v", err)
	}
	return &actionrequirementpb.CreateActionRequirementResponse{ActionRequirement: mapper.ActionRequirementROToProto(&view)}, nil
}

func (h *ActionRequirementHandler) GetActionRequirementByID(ctx context.Context, req *actionrequirementpb.GetActionRequirementByIDRequest) (*actionrequirementpb.GetActionRequirementByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}
	view, err := h.arAppSvc.GetActionRequirement(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get action requirement by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "action requirement not found: %v", err)
	}
	return &actionrequirementpb.GetActionRequirementByIDResponse{ActionRequirement: mapper.ActionRequirementROToProto(&view)}, nil
}

func (h *ActionRequirementHandler) GetByActionID(ctx context.Context, req *actionrequirementpb.GetByActionIDRequest) (*actionrequirementpb.GetByActionIDResponse, error) {
	actionID, err := uuid.Parse(req.ActionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action_id: %v", err)
	}
	list, err := h.arAppSvc.GetByActionID(ctx, actionID)
	if err != nil {
		logx.S().Errorf("failed to get action requirements by action_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get action requirements: %v", err)
	}
	return &actionrequirementpb.GetByActionIDResponse{ActionRequirements: mapper.ActionRequirementListROToProto(list)}, nil
}

func (h *ActionRequirementHandler) GetByPermissionID(ctx context.Context, req *actionrequirementpb.GetByPermissionIDRequest) (*actionrequirementpb.GetByPermissionIDResponse, error) {
	permissionID, err := uuid.Parse(req.PermissionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}
	list, err := h.arAppSvc.GetByPermissionID(ctx, permissionID)
	if err != nil {
		logx.S().Errorf("failed to get action requirements by permission_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get action requirements: %v", err)
	}
	return &actionrequirementpb.GetByPermissionIDResponse{ActionRequirements: mapper.ActionRequirementListROToProto(list)}, nil
}

func (h *ActionRequirementHandler) DeleteActionRequirement(ctx context.Context, req *actionrequirementpb.DeleteActionRequirementRequest) (*actionrequirementpb.DeleteActionRequirementResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}
	if err := h.arAppSvc.DeleteActionRequirement(ctx, arAppCommands.DeleteActionRequirementCmd{ActionRequirementID: id}); err != nil {
		logx.S().Errorf("failed to delete action requirement: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete action requirement: %v", err)
	}
	return &actionrequirementpb.DeleteActionRequirementResponse{}, nil
}

func (h *ActionRequirementHandler) DeleteByActionIDAndPermissionID(ctx context.Context, req *actionrequirementpb.DeleteByActionIDAndPermissionIDRequest) (*actionrequirementpb.DeleteByActionIDAndPermissionIDResponse, error) {
	actionID, err := uuid.Parse(req.ActionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action_id: %v", err)
	}
	permissionID, err := uuid.Parse(req.PermissionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}
	if err := h.arAppSvc.DeleteByActionIDAndPermissionID(ctx, arAppCommands.DeleteByActionIDAndPermissionIDCmd{ActionID: actionID, PermissionID: permissionID}); err != nil {
		logx.S().Errorf("failed to delete action requirement: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete action requirement: %v", err)
	}
	return &actionrequirementpb.DeleteByActionIDAndPermissionIDResponse{}, nil
}

func (h *ActionRequirementHandler) DeleteByActionID(ctx context.Context, req *actionrequirementpb.DeleteByActionIDRequest) (*actionrequirementpb.DeleteByActionIDResponse, error) {
	actionID, err := uuid.Parse(req.ActionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid action_id: %v", err)
	}
	if err := h.arAppSvc.DeleteByActionID(ctx, arAppCommands.DeleteByActionIDCmd{ActionID: actionID}); err != nil {
		logx.S().Errorf("failed to delete action requirements by action_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete action requirements: %v", err)
	}
	return &actionrequirementpb.DeleteByActionIDResponse{}, nil
}
