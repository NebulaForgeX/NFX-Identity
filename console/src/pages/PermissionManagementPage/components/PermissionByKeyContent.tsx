import { memo } from "react";
import { useTranslation } from "react-i18next";

import { usePermissionByKey } from "@/hooks/useAccess";

import styles from "../styles.module.css";

const PermissionByKeyContent = memo(({ permissionKey }: { permissionKey: string }) => {
  const { t } = useTranslation("PermissionManagementPage");
  const { data: permission } = usePermissionByKey({ key: permissionKey });
  if (!permission) return <p className={styles.emptyText}>{t("notFound")}</p>;
  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>{permission.name}</span>
        {permission.isSystem && <span className={styles.badge}>{t("systemPermission")}</span>}
      </div>
      <div className={styles.itemDetails}>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("key")}:</span>
          <span className={styles.itemValue}>{permission.key}</span>
        </div>
        {permission.description && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("description")}:</span>
            <span className={styles.itemValue}>{permission.description}</span>
          </div>
        )}
      </div>
    </div>
  );
});
PermissionByKeyContent.displayName = "PermissionByKeyContent";

export default PermissionByKeyContent;
