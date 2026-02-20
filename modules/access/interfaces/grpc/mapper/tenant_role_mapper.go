package mapper

import (
	domain "nfxid/modules/access/domain/tenant_roles"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TenantRoleDomainToProto(r *domain.TenantRole) *tenantrolepb.TenantRole {
	if r == nil {
		return nil
	}
	return &tenantrolepb.TenantRole{
		Id:        r.ID().String(),
		TenantId:  r.TenantID().String(),
		RoleKey:   r.RoleKey(),
		Name:      r.Name(),
		CreatedAt: timestamppb.New(r.CreatedAt()),
	}
}
