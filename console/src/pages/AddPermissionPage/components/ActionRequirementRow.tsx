import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Button } from "@/components";
import { useAction } from "@/hooks/useAccess";

import styles from "../styles.module.css";

interface ActionRequirementRowProps {
  actionId: string;
  onRemove: () => void;
  isRemoving: boolean;
}

const ActionRequirementRow = memo(({ actionId, onRemove, isRemoving }: ActionRequirementRowProps) => {
  const { t } = useTranslation("AddPermissionPage");
  const { data: action } = useAction({ id: actionId });
  if (!action) return null;
  return (
    <div className={styles.permissionItem}>
      <span className={styles.permissionName}>{action.name}</span>
      {action.isSystem && <span className={styles.badgeSmall}>{t("systemAction")}</span>}
      <span className={styles.permissionKey}>({action.key})</span>
      <Button variant="ghost" size="small" onClick={onRemove} loading={isRemoving}>
        {t("remove")}
      </Button>
    </div>
  );
});
ActionRequirementRow.displayName = "ActionRequirementRow";

export default ActionRequirementRow;
