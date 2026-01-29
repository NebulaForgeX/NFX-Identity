import { memo, useState } from "react";
import { useTranslation } from "react-i18next";

import { Button, Input } from "@/components";
import { useCreateRole } from "@/hooks/useAccess";

import styles from "../styles.module.css";

const CreateRoleForm = memo(() => {
  const { t } = useTranslation("AddRolePage");
  const [newRoleKey, setNewRoleKey] = useState("");
  const [newRoleName, setNewRoleName] = useState("");
  const [newRoleDesc, setNewRoleDesc] = useState("");
  const [newRoleScope, setNewRoleScope] = useState("global");

  const createRole = useCreateRole();

  const handleCreateRole = () => {
    if (!newRoleKey.trim() || !newRoleName.trim()) return;
    createRole.mutate(
      {
        key: newRoleKey.trim(),
        name: newRoleName.trim(),
        description: newRoleDesc.trim() || undefined,
        scopeType: newRoleScope,
        isSystem: false,
      },
      {
        onSuccess: () => {
          setNewRoleKey("");
          setNewRoleName("");
          setNewRoleDesc("");
        },
      },
    );
  };

  return (
    <div className={styles.card}>
      <h3 className={styles.cardTitle}>{t("createRole")}</h3>
      <div className={styles.formRow}>
        <Input
          label={t("key")}
          value={newRoleKey}
          onChange={(e) => setNewRoleKey(e.target.value)}
          placeholder="e.g. app.editor"
          fullWidth
        />
        <Input
          label={t("name")}
          value={newRoleName}
          onChange={(e) => setNewRoleName(e.target.value)}
          placeholder="Editor"
          fullWidth
        />
      </div>
      <div className={styles.formRow}>
        <Input
          label={t("description")}
          value={newRoleDesc}
          onChange={(e) => setNewRoleDesc(e.target.value)}
          fullWidth
        />
        <Input
          label={t("scopeType")}
          value={newRoleScope}
          onChange={(e) => setNewRoleScope(e.target.value)}
          fullWidth
        />
      </div>
      <Button onClick={handleCreateRole} loading={createRole.isPending}>
        {t("createRole")}
      </Button>
    </div>
  );
});
CreateRoleForm.displayName = "CreateRoleForm";

export default CreateRoleForm;
