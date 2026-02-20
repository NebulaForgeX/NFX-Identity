package mapper

import (
	domain "nfxid/modules/access/domain/super_admins"
	superadminpb "nfxid/protos/gen/access/super_admin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SuperAdminDomainToProto domain -> proto
func SuperAdminDomainToProto(s *domain.SuperAdmin) *superadminpb.SuperAdmin {
	if s == nil {
		return nil
	}
	return &superadminpb.SuperAdmin{
		UserId:    s.UserID().String(),
		CreatedAt: timestamppb.New(s.CreatedAt()),
	}
}
