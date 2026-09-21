package handler

import (
	"context"

	"nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/interface/grpc/mapper"
	accountpb "nfxidentity/protos/gen/auth/account"
	authorityprofilepb "nfxidentity/protos/gen/auth/authority_profile"

	"github.com/google/uuid"
)

type AuthorityHandler struct {
	authorityprofilepb.UnimplementedAuthorityProfileServiceServer
	svc *account.Service
}

func NewAuthorityHandler(svc *account.Service) *AuthorityHandler {
	return &AuthorityHandler{svc: svc}
}

func (h *AuthorityHandler) ConfirmAuthorityProfileAvatar(ctx context.Context, req *authorityprofilepb.ConfirmAuthorityProfileAvatarRequest) (*authorityprofilepb.ConfirmAuthorityProfileAvatarResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	imageID, err := uuid.Parse(req.GetImageId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.ConfirmAuthorityProfileAvatar(ctx, account.ConfirmAuthorityProfileAvatarInput{
		AccountID: aid,
		ProfileID: pid,
		ImageID:   imageID,
	}); err != nil {
		return nil, err
	}
	return &authorityprofilepb.ConfirmAuthorityProfileAvatarResponse{}, nil
}

func (h *AuthorityHandler) ConfirmAuthorityProfileBackgrounds(ctx context.Context, req *authorityprofilepb.ConfirmAuthorityProfileBackgroundsRequest) (*authorityprofilepb.ConfirmAuthorityProfileBackgroundsResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	images := make([]account.ConfirmAuthorityProfileBackgroundItemInput, 0, len(req.GetImages()))
	for _, img := range req.GetImages() {
		imageID, err := uuid.Parse(img.GetImageId())
		if err != nil {
			return nil, err
		}
		images = append(images, account.ConfirmAuthorityProfileBackgroundItemInput{
			ImageID:   imageID,
			SortOrder: int(img.GetSortOrder()),
		})
	}
	if _, err := h.svc.ConfirmAuthorityProfileBackgrounds(ctx, account.ConfirmAuthorityProfileBackgroundsInput{
		AccountID: aid,
		ProfileID: pid,
		Images:    images,
	}); err != nil {
		return nil, err
	}
	return &authorityprofilepb.ConfirmAuthorityProfileBackgroundsResponse{}, nil
}

func (h *AuthorityHandler) BatchGetPublicProfileCards(ctx context.Context, req *authorityprofilepb.BatchGetPublicProfileCardsRequest) (*authorityprofilepb.BatchGetPublicProfileCardsResponse, error) {
	ids := make([]uuid.UUID, 0, len(req.GetProfileIds()))
	for _, raw := range req.GetProfileIds() {
		pid, err := uuid.Parse(raw)
		if err != nil {
			continue
		}
		ids = append(ids, pid)
	}
	out, err := h.svc.BatchGetAuthorityPublicProfileCards(ctx, account.BatchGetAuthorityPublicProfileCardsInput{ProfileIDs: ids})
	if err != nil {
		return nil, err
	}
	return &authorityprofilepb.BatchGetPublicProfileCardsResponse{Profiles: mapper.AuthorityProfileItemsToPublicProfileCardProto(out.Profiles)}, nil
}

func (h *AuthorityHandler) GetAuthorityRoles(ctx context.Context, req *authorityprofilepb.GetAuthorityRolesRequest) (*authorityprofilepb.GetAuthorityRolesResponse, error) {
	pid, err := uuid.Parse(req.GetProfileId())
	if err != nil {
		return nil, err
	}
	roles, err := h.svc.AuthorityRoles(ctx, pid)
	if err != nil {
		return nil, err
	}
	out := make([]accountpb.AuthorityRole, 0, len(roles))
	for _, r := range roles {
		switch r {
		case "auditor":
			out = append(out, accountpb.AuthorityRole_AUTHORITY_ROLE_AUDITOR)
		case "administrator":
			out = append(out, accountpb.AuthorityRole_AUTHORITY_ROLE_ADMINISTRATOR)
		case "owner":
			out = append(out, accountpb.AuthorityRole_AUTHORITY_ROLE_OWNER)
		}
	}
	return &authorityprofilepb.GetAuthorityRolesResponse{AuthorityRoles: out}, nil
}

func (h *AuthorityHandler) HasAuthorityRole(ctx context.Context, req *authorityprofilepb.HasAuthorityRoleRequest) (*authorityprofilepb.HasAuthorityRoleResponse, error) {
	ok, err := h.svc.HasAuthorityRole(ctx, req.GetProfileId(), "auditor")
	if err != nil {
		return nil, err
	}
	return &authorityprofilepb.HasAuthorityRoleResponse{Allowed: ok}, nil
}

func (h *AuthorityHandler) IsAuditor(ctx context.Context, req *authorityprofilepb.HasAuthorityRoleRequest) (*authorityprofilepb.HasAuthorityRoleResponse, error) {
	ok, err := h.svc.HasAuthorityRole(ctx, req.GetProfileId(), "auditor")
	if err != nil {
		return nil, err
	}
	return &authorityprofilepb.HasAuthorityRoleResponse{Allowed: ok}, nil
}

func (h *AuthorityHandler) IsAdministrator(ctx context.Context, req *authorityprofilepb.HasAuthorityRoleRequest) (*authorityprofilepb.HasAuthorityRoleResponse, error) {
	ok, err := h.svc.HasAuthorityRole(ctx, req.GetProfileId(), "administrator")
	if err != nil {
		return nil, err
	}
	return &authorityprofilepb.HasAuthorityRoleResponse{Allowed: ok}, nil
}

func (h *AuthorityHandler) IsOwner(ctx context.Context, req *authorityprofilepb.HasAuthorityRoleRequest) (*authorityprofilepb.HasAuthorityRoleResponse, error) {
	ok, err := h.svc.HasAuthorityRole(ctx, req.GetProfileId(), "owner")
	if err != nil {
		return nil, err
	}
	return &authorityprofilepb.HasAuthorityRoleResponse{Allowed: ok}, nil
}
