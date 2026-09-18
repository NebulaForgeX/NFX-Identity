package handler

import (
	"context"

	"nfxidentity/modules/auth/application/platform"
	accountpb "nfxidentity/protos/gen/auth/account"
	authorityprofilepb "nfxidentity/protos/gen/auth/authority_profile"

	"github.com/google/uuid"
)

type AuthorityHandler struct {
	authorityprofilepb.UnimplementedAuthorityProfileServiceServer
	svc *platform.Service
}

func NewAuthorityHandler(svc *platform.Service) *AuthorityHandler {
	return &AuthorityHandler{svc: svc}
}

func (h *AuthorityHandler) ConfirmAuthorityProfileAvatar(ctx context.Context, req *authorityprofilepb.ConfirmAuthorityProfileAvatarRequest) (*authorityprofilepb.ConfirmAuthorityProfileAvatarResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	if err := h.svc.ConfirmAvatar(ctx, aid, pid, "authority", req.GetImageId()); err != nil {
		return nil, err
	}
	return &authorityprofilepb.ConfirmAuthorityProfileAvatarResponse{}, nil
}

func (h *AuthorityHandler) ConfirmAuthorityProfileBackgrounds(ctx context.Context, req *authorityprofilepb.ConfirmAuthorityProfileBackgroundsRequest) (*authorityprofilepb.ConfirmAuthorityProfileBackgroundsResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	items := make([]struct {
		ImageID   string
		SortOrder int
	}, 0, len(req.GetImages()))
	for _, img := range req.GetImages() {
		items = append(items, struct {
			ImageID   string
			SortOrder int
		}{ImageID: img.GetImageId(), SortOrder: int(img.GetSortOrder())})
	}
	if err := h.svc.ConfirmBackgrounds(ctx, aid, pid, "authority", items); err != nil {
		return nil, err
	}
	return &authorityprofilepb.ConfirmAuthorityProfileBackgroundsResponse{}, nil
}

func (h *AuthorityHandler) BatchGetPublicProfileCards(ctx context.Context, req *authorityprofilepb.BatchGetPublicProfileCardsRequest) (*authorityprofilepb.BatchGetPublicProfileCardsResponse, error) {
	cards := h.svc.BatchPublicCards(ctx, req.GetProfileIds())
	out := make([]*authorityprofilepb.PublicProfileCard, 0, len(cards))
	for _, c := range cards {
		out = append(out, &authorityprofilepb.PublicProfileCard{ProfileId: asString(c["profile_id"])})
	}
	return &authorityprofilepb.BatchGetPublicProfileCardsResponse{Profiles: out}, nil
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
