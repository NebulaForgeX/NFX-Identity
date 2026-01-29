import { memo } from "react";
import { useTranslation } from "react-i18next";

import { useRolePermissionsByRole } from "@/hooks/useAccess";
import type { useDeleteRolePermission } from "@/hooks/useAccess";

import RolePermissionRow from "./RolePermissionRow";
import styles from "../styles.module.css";

interface RolePermissionsByRolePanelProps {
  role: { id: string; name: string; key: string; isSystem?: boolean };
  deleteRolePermission: ReturnType<typeof useDeleteRolePermission>;
}

const RolePermissionsByRolePanel = memo(({ role, deleteRolePermission }: RolePermissionsByRolePanelProps) => {
  const { t } = useTranslation("AddPermissionPage");
  const { data: rolePermissions } = useRolePermissionsByRole({ roleId: role.id });
  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>{role.name}</span>
        {role.isSystem && <span className={styles.badge}>{t("systemRole")}</span>}
      </div>
      <div className={styles.itemDetails}>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("key")}:</span>
          <span className={styles.itemValue}>{role.key}</span>
        </div>
        {rolePermissions && rolePermissions.length > 0 ? (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("permissions")}:</span>
            <div className={styles.permissionList}>
              {rolePermissions.map((rp) => (
                <RolePermissionRow
                  key={rp.id}
                  permissionId={rp.permissionId}
                  onRemove={() => deleteRolePermission.mutate(rp.id)}
                  isRemoving={deleteRolePermission.isPending}
                />
              ))}
            </div>
          </div>
        ) : (
          <p className={styles.emptyText}>{t("noPermissions")}</p>
        )}
      </div>
    </div>
  );
});
RolePermissionsByRolePanel.displayName = "RolePermissionsByRolePanel";

export default RolePermissionsByRolePanel;
