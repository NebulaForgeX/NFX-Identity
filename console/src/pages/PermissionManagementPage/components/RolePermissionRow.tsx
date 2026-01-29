import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Button } from "@/components";
import { usePermission } from "@/hooks/useAccess";

import styles from "../styles.module.css";

interface RolePermissionRowProps {
  permissionId: string;
  onRemove: () => void;
  isRemoving: boolean;
}

const RolePermissionRow = memo(({ permissionId, onRemove, isRemoving }: RolePermissionRowProps) => {
  const { t } = useTranslation("PermissionManagementPage");
  const { data: permission } = usePermission({ id: permissionId });
  if (!permission) return null;
  return (
    <div className={styles.permissionItem}>
      <span className={styles.permissionName}>{permission.name}</span>
      {permission.isSystem && <span className={styles.badgeSmall}>{t("systemPermission")}</span>}
      <span className={styles.permissionKey}>({permission.key})</span>
      <Button variant="ghost" size="small" onClick={onRemove} loading={isRemoving}>
        {t("remove")}
      </Button>
    </div>
  );
});
RolePermissionRow.displayName = "RolePermissionRow";

export default RolePermissionRow;
