package handler

import (
	"context"

	grantApp "nfxid/modules/access/application/grants"
	grantDomain "nfxid/modules/access/domain/grants"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	grantpb "nfxid/protos/gen/access/grant"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrantHandler struct {
	grantpb.UnimplementedGrantServiceServer
	grantAppSvc *grantApp.Service
}

func NewGrantHandler(grantAppSvc *grantApp.Service) *GrantHandler {
	return &GrantHandler{
		grantAppSvc: grantAppSvc,
	}
}

// GetGrantByID 根据ID获取授权
func (h *GrantHandler) GetGrantByID(ctx context.Context, req *grantpb.GetGrantByIDRequest) (*grantpb.GetGrantByIDResponse, error) {
	grantID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid grant_id: %v", err)
	}

	grantView, err := h.grantAppSvc.GetGrant(ctx, grantID)
	if err != nil {
		logx.S().Errorf("failed to get grant by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "grant not found: %v", err)
	}

	grant := mapper.GrantROToProto(&grantView)
	return &grantpb.GetGrantByIDResponse{Grant: grant}, nil
}

// GetGrantsBySubject 根据主体获取授权列表
func (h *GrantHandler) GetGrantsBySubject(ctx context.Context, req *grantpb.GetGrantsBySubjectRequest) (*grantpb.GetGrantsBySubjectResponse, error) {
	subjectID, err := uuid.Parse(req.SubjectId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid subject_id: %v", err)
	}

	subjectType := protoToSubjectType(req.SubjectType)
	grantViews, err := h.grantAppSvc.GetGrantsBySubject(ctx, subjectType, subjectID)
	if err != nil {
		logx.S().Errorf("failed to get grants by subject: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get grants: %v", err)
	}

	protoGrants := mapper.GrantListROToProto(grantViews)
	return &grantpb.GetGrantsBySubjectResponse{Grants: protoGrants}, nil
}

// BatchGetGrants 批量获取授权
func (h *GrantHandler) BatchGetGrants(ctx context.Context, req *grantpb.BatchGetGrantsRequest) (*grantpb.BatchGetGrantsResponse, error) {
	grantIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		grantIDs = append(grantIDs, id)
	}

	grants := make([]*grantpb.Grant, 0, len(grantIDs))
	for _, grantID := range grantIDs {
		grantView, err := h.grantAppSvc.GetGrant(ctx, grantID)
		if err != nil {
			logx.S().Warnf("failed to get grant %s: %v", grantID, err)
			continue
		}
		grant := mapper.GrantROToProto(&grantView)
		grants = append(grants, grant)
	}

	return &grantpb.BatchGetGrantsResponse{Grants: grants}, nil
}

func protoToSubjectType(pt grantpb.AccessSubjectType) grantDomain.SubjectType {
	switch pt {
	case grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER:
		return grantDomain.SubjectTypeUser
	case grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT:
		return grantDomain.SubjectTypeClient
	default:
		return grantDomain.SubjectTypeUser
	}
}
