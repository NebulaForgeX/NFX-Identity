import { memo } from "react";
import { useTranslation } from "react-i18next";

import { useRoleByKey } from "@/hooks/useAccess";

import styles from "../styles.module.css";

const RoleByKeyContent = memo(({ roleKey }: { roleKey: string }) => {
  const { t } = useTranslation("PermissionManagementPage");
  const { data: role } = useRoleByKey({ key: roleKey });
  if (!role) return null;
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
        {role.description && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("description")}:</span>
            <span className={styles.itemValue}>{role.description}</span>
          </div>
        )}
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("scopeType")}:</span>
          <span className={styles.itemValue}>{role.scopeType}</span>
        </div>
      </div>
    </div>
  );
});
RoleByKeyContent.displayName = "RoleByKeyContent";

export default RoleByKeyContent;
