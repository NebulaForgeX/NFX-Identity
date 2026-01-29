import { memo } from "react";
import { useTranslation } from "react-i18next";

import { useRoleByKey } from "@/hooks/useAccess";
import type { useDeleteRolePermission } from "@/hooks/useAccess";

import RolePermissionsByRolePanel from "./RolePermissionsByRolePanel";
import styles from "../styles.module.css";

interface RolePermissionsLookupContentProps {
  roleKey: string;
  deleteRolePermission: ReturnType<typeof useDeleteRolePermission>;
}

const RolePermissionsLookupContent = memo(({ roleKey, deleteRolePermission }: RolePermissionsLookupContentProps) => {
  const { t } = useTranslation("AddPermissionPage");
  const { data: role } = useRoleByKey({ key: roleKey });
  if (!role) return <p className={styles.emptyText}>{t("notFoundRole")}</p>;
  return <RolePermissionsByRolePanel role={role} deleteRolePermission={deleteRolePermission} />;
});
RolePermissionsLookupContent.displayName = "RolePermissionsLookupContent";

export default RolePermissionsLookupContent;
