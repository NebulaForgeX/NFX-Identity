import { memo } from "react";
import { useTranslation } from "react-i18next";

import { useActionByKey } from "@/hooks/useAccess";

import styles from "../styles.module.css";

interface ActionByKeyContentProps {
  actionKey: string;
}

const ActionByKeyContent = memo(({ actionKey }: ActionByKeyContentProps) => {
  const { t } = useTranslation("AddActionPage");
  const { data: action } = useActionByKey({ key: actionKey });
  if (!action) return <p className={styles.emptyText}>{t("notFound")}</p>;
  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>{action.name}</span>
        {action.isSystem && <span className={styles.badge}>{t("systemAction")}</span>}
      </div>
      <div className={styles.itemDetails}>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("key")}:</span>
          <span className={styles.itemValue}>{action.key}</span>
        </div>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("service")}:</span>
          <span className={styles.itemValue}>{action.service}</span>
        </div>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("status")}:</span>
          <span className={styles.itemValue}>{action.status}</span>
        </div>
        {action.description && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("description")}:</span>
            <span className={styles.itemValue}>{action.description}</span>
          </div>
        )}
      </div>
    </div>
  );
});
ActionByKeyContent.displayName = "ActionByKeyContent";

export default ActionByKeyContent;
