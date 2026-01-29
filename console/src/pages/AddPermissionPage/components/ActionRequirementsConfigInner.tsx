import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Button, Input } from "@/components";
import { useActionRequirementsByPermission } from "@/hooks/useAccess";
import type { useCreateActionRequirement, useDeleteActionRequirement } from "@/hooks/useAccess";

import ActionRequirementRow from "./ActionRequirementRow";
import styles from "../styles.module.css";

interface ActionRequirementsConfigInnerProps {
  permission: { id: string; name: string; key: string; isSystem?: boolean };
  actionKeyToAdd: string;
  setActionKeyToAdd: (v: string) => void;
  onAddAction: () => void;
  createActionRequirement: ReturnType<typeof useCreateActionRequirement>;
  deleteActionRequirement: ReturnType<typeof useDeleteActionRequirement>;
}

const ActionRequirementsConfigInner = memo(
  ({
    permission,
    actionKeyToAdd,
    setActionKeyToAdd,
    onAddAction,
    createActionRequirement,
    deleteActionRequirement,
  }: ActionRequirementsConfigInnerProps) => {
    const { t } = useTranslation("AddPermissionPage");
    const { data: actionRequirements } = useActionRequirementsByPermission({
      permissionId: permission.id,
    });
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
          <div className={styles.formRow} style={{ marginTop: "1rem" }}>
            <Input
              label={t("actionKeyToAdd")}
              value={actionKeyToAdd}
              onChange={(e) => setActionKeyToAdd(e.target.value)}
              placeholder="e.g. directory.users.read"
              fullWidth
            />
            <Button
              onClick={onAddAction}
              loading={createActionRequirement.isPending}
              disabled={!actionKeyToAdd.trim()}
            >
              {t("addAction")}
            </Button>
          </div>
          {actionRequirements && actionRequirements.length > 0 ? (
            <div className={styles.itemDetail}>
              <span className={styles.itemLabel}>{t("permissions")} (Actions):</span>
              <div className={styles.permissionList}>
                {actionRequirements.map((ar) => (
                  <ActionRequirementRow
                    key={ar.id}
                    actionId={ar.actionId}
                    onRemove={() => deleteActionRequirement.mutate(ar.id)}
                    isRemoving={deleteActionRequirement.isPending}
                  />
                ))}
              </div>
            </div>
          ) : (
            <p className={styles.emptyText}>{t("noActions")}</p>
          )}
        </div>
      </div>
    );
  },
);
ActionRequirementsConfigInner.displayName = "ActionRequirementsConfigInner";

export default ActionRequirementsConfigInner;
