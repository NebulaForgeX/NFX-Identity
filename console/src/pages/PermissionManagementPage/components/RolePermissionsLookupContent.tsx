import { memo } from "react";

import { useRoleByKey } from "@/hooks/useAccess";
import type { useDeleteRolePermission } from "@/hooks/useAccess";

import RolePermissionsByRolePanel from "./RolePermissionsByRolePanel";

interface RolePermissionsLookupContentProps {
  roleKey: string;
  deleteRolePermission: ReturnType<typeof useDeleteRolePermission>;
}

const RolePermissionsLookupContent = memo(({ roleKey, deleteRolePermission }: RolePermissionsLookupContentProps) => {
  const { data: role } = useRoleByKey({ key: roleKey });
  if (!role) return null;
  return (
    <RolePermissionsByRolePanel
      role={role}
      deleteRolePermission={deleteRolePermission}
    />
  );
});
RolePermissionsLookupContent.displayName = "RolePermissionsLookupContent";

export default RolePermissionsLookupContent;
