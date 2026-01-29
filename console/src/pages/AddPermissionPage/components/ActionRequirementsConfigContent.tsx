import { memo } from "react";
import { useTranslation } from "react-i18next";

import { usePermissionByKey } from "@/hooks/useAccess";
import type { useCreateActionRequirement, useDeleteActionRequirement } from "@/hooks/useAccess";

import ActionRequirementsConfigInner from "./ActionRequirementsConfigInner";
import styles from "../styles.module.css";

interface ActionRequirementsConfigContentProps {
  permissionKey: string;
  actionKeyToAdd: string;
  setActionKeyToAdd: (v: string) => void;
  onAddAction: () => void;
  createActionRequirement: ReturnType<typeof useCreateActionRequirement>;
  deleteActionRequirement: ReturnType<typeof useDeleteActionRequirement>;
}

const ActionRequirementsConfigContent = memo(
  ({
    permissionKey,
    actionKeyToAdd,
    setActionKeyToAdd,
    onAddAction,
    createActionRequirement,
    deleteActionRequirement,
  }: ActionRequirementsConfigContentProps) => {
    const { t } = useTranslation("AddPermissionPage");
    const { data: permission } = usePermissionByKey({ key: permissionKey });
    if (!permission) return <p className={styles.emptyText}>{t("notFoundPermission")}</p>;
    return (
      <ActionRequirementsConfigInner
        permission={permission}
        actionKeyToAdd={actionKeyToAdd}
        setActionKeyToAdd={setActionKeyToAdd}
        onAddAction={onAddAction}
        createActionRequirement={createActionRequirement}
        deleteActionRequirement={deleteActionRequirement}
      />
    );
  },
);
ActionRequirementsConfigContent.displayName = "ActionRequirementsConfigContent";

export default ActionRequirementsConfigContent;
