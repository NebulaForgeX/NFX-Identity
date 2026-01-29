import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Button, Input, IconButton, Suspense } from "@/components";
import { ArrowLeft } from "@/assets/icons/lucide";
import {
  useCreatePermission,
  useCreateRolePermission,
  useCreateActionRequirement,
  useDeleteRolePermission,
  useDeleteActionRequirement,
} from "@/hooks/useAccess";
import { GetActionByKey, GetPermissionByKey, GetRoleByKey } from "@/apis";
import { showError } from "@/stores/modalStore";
import { ROUTES } from "@/types/navigation";

import RolePermissionsLookupContent from "./components/RolePermissionsLookupContent";
import ActionRequirementsConfigContent from "./components/ActionRequirementsConfigContent";
import styles from "./styles.module.css";

const AddPermissionPage = memo(() => {
  const { t } = useTranslation("AddPermissionPage");
  const navigate = useNavigate();
  const [roleKey, setRoleKey] = useState("");
  const [permissionKey, setPermissionKey] = useState("");
  const [lookupInput, setLookupInput] = useState("");
  const [lookupRoleKey, setLookupRoleKey] = useState("");
  const [newPermKey, setNewPermKey] = useState("");
  const [newPermName, setNewPermName] = useState("");
  const [newPermDesc, setNewPermDesc] = useState("");
  const [actionConfigPermissionKey, setActionConfigPermissionKey] = useState("");
  const [lookupActionConfigKey, setLookupActionConfigKey] = useState("");
  const [actionKeyToAdd, setActionKeyToAdd] = useState("");

  const createPermission = useCreatePermission();
  const createRolePermission = useCreateRolePermission();
  const createActionRequirement = useCreateActionRequirement();
  const deleteRolePermission = useDeleteRolePermission();
  const deleteActionRequirement = useDeleteActionRequirement();

  const handleAssignPermission = async () => {
    if (!roleKey.trim() || !permissionKey.trim()) return;
    try {
      const roleRes = await GetRoleByKey(roleKey.trim());
      if (!roleRes) {
        showError(t("notFoundRole"));
        return;
      }
      const permRes = await GetPermissionByKey(permissionKey.trim());
      if (!permRes) {
        showError(t("notFoundPermission"));
        return;
      }
      await createRolePermission.mutateAsync({
        roleId: roleRes.id,
        permissionId: permRes.id,
      });
      setRoleKey("");
      setPermissionKey("");
      if (lookupRoleKey === roleKey) setLookupRoleKey("");
    } catch {
      // onError handled by mutation
    }
  };

  const handleCreatePermission = () => {
    if (!newPermKey.trim() || !newPermName.trim()) return;
    createPermission.mutate(
      {
        key: newPermKey.trim(),
        name: newPermName.trim(),
        description: newPermDesc.trim() || undefined,
        isSystem: false,
      },
      {
        onSuccess: () => {
          setNewPermKey("");
          setNewPermName("");
          setNewPermDesc("");
        },
      },
    );
  };

  const handleAddActionRequirement = async () => {
    if (!lookupActionConfigKey.trim() || !actionKeyToAdd.trim()) return;
    try {
      const permission = await GetPermissionByKey(lookupActionConfigKey.trim());
      if (!permission) {
        showError(t("notFoundPermission"));
        return;
      }
      const action = await GetActionByKey(actionKeyToAdd.trim());
      if (!action) {
        showError(t("notFoundAction"));
        return;
      }
      await createActionRequirement.mutateAsync({
        permissionId: permission.id,
        actionId: action.id,
      });
      setActionKeyToAdd("");
    } catch {
      // onError handled by mutation
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <IconButton
          variant="ghost"
          leftIcon={<ArrowLeft size={20} />}
          onClick={() => navigate(ROUTES.PERMISSION_MANAGEMENT)}
          className={styles.backButton}
        >
          {t("back")}
        </IconButton>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("createPermission")}</h3>
        <div className={styles.formRow}>
          <Input label={t("key")} value={newPermKey} onChange={(e) => setNewPermKey(e.target.value)} placeholder="e.g. directory.users.read" fullWidth />
          <Input label={t("name")} value={newPermName} onChange={(e) => setNewPermName(e.target.value)} placeholder="Users Read" fullWidth />
        </div>
        <Input label={t("description")} value={newPermDesc} onChange={(e) => setNewPermDesc(e.target.value)} fullWidth />
        <div className={styles.submitWrap}>
          <Button onClick={handleCreatePermission} loading={createPermission.isPending}>
            {t("createPermission")}
          </Button>
        </div>
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("assignPermissionToRole")}</h3>
        <div className={styles.formRow}>
          <Input label={t("roleKey")} value={roleKey} onChange={(e) => setRoleKey(e.target.value)} placeholder="e.g. system.admin" fullWidth />
          <Input label={t("permissionKey")} value={permissionKey} onChange={(e) => setPermissionKey(e.target.value)} placeholder="e.g. directory.*" fullWidth />
        </div>
        <Button onClick={handleAssignPermission} loading={createRolePermission.isPending}>
          {t("addPermission")}
        </Button>
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("viewRolePermissions")}</h3>
        <div className={styles.formRow}>
          <Input label={t("roleKey")} value={lookupInput} onChange={(e) => setLookupInput(e.target.value)} placeholder="e.g. system.admin" fullWidth />
          <Button variant="secondary" onClick={() => setLookupRoleKey(lookupInput.trim())} disabled={!lookupInput.trim()}>
            {t("lookup")}
          </Button>
        </div>
        {lookupRoleKey && (
          <Suspense loadingType="ecg" loadingText={t("loading")} loadingSize="small" loadingContainerClassName={styles.loading}>
            <RolePermissionsLookupContent roleKey={lookupRoleKey} deleteRolePermission={deleteRolePermission} />
          </Suspense>
        )}
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("configureActionRequirements")}</h3>
        <div className={styles.formRow}>
          <Input
            label={t("permissionKeyForActions")}
            value={actionConfigPermissionKey}
            onChange={(e) => setActionConfigPermissionKey(e.target.value)}
            placeholder="e.g. directory.users.read"
            fullWidth
          />
          <Button
            variant="secondary"
            onClick={() => setLookupActionConfigKey(actionConfigPermissionKey.trim())}
            disabled={!actionConfigPermissionKey.trim()}
          >
            {t("lookup")}
          </Button>
        </div>
        {lookupActionConfigKey && (
          <Suspense
            loadingType="ecg"
            loadingText={t("loading")}
            loadingSize="small"
            loadingContainerClassName={styles.loading}
          >
            <ActionRequirementsConfigContent
              permissionKey={lookupActionConfigKey}
              actionKeyToAdd={actionKeyToAdd}
              setActionKeyToAdd={setActionKeyToAdd}
              onAddAction={handleAddActionRequirement}
              createActionRequirement={createActionRequirement}
              deleteActionRequirement={deleteActionRequirement}
            />
          </Suspense>
        )}
      </div>
    </div>
  );
});

AddPermissionPage.displayName = "AddPermissionPage";

export default AddPermissionPage;
