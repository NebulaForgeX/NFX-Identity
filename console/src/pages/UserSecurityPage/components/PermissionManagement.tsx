import { memo, useState } from "react";
import { useTranslation } from "react-i18next";

import { Button, Input, Suspense } from "@/components";
import {
  useCreateRole,
  useCreatePermission,
  useCreateRolePermission,
  useDeleteRolePermission,
  useRoleByKey,
  useRolePermissionsByRole,
  usePermission,
} from "@/hooks/useAccess";
import { GetPermissionByKey, GetRoleByKey } from "@/apis";

import styles from "../styles.module.css";

const PermissionManagement = memo(() => {
  const { t } = useTranslation("UserSecurityPage");
  const [roleKey, setRoleKey] = useState("");
  const [permissionKey, setPermissionKey] = useState("");
  const [lookupInput, setLookupInput] = useState("");
  const [lookupRoleKey, setLookupRoleKey] = useState("");
  // Create role form
  const [newRoleKey, setNewRoleKey] = useState("");
  const [newRoleName, setNewRoleName] = useState("");
  const [newRoleDesc, setNewRoleDesc] = useState("");
  const [newRoleScope, setNewRoleScope] = useState("global");
  // Create permission form
  const [newPermKey, setNewPermKey] = useState("");
  const [newPermName, setNewPermName] = useState("");
  const [newPermDesc, setNewPermDesc] = useState("");

  const createRole = useCreateRole();
  const createPermission = useCreatePermission();
  const createRolePermission = useCreateRolePermission();
  const deleteRolePermission = useDeleteRolePermission();

  const handleAssignPermission = async () => {
    if (!roleKey.trim() || !permissionKey.trim()) return;
    try {
      const roleRes = await GetRoleByKey(roleKey.trim());
      const permRes = await GetPermissionByKey(permissionKey.trim());
      await createRolePermission.mutateAsync({
        roleId: roleRes.id,
        permissionId: permRes.id,
      });
      setRoleKey("");
      setPermissionKey("");
      if (lookupRoleKey === roleKey) setLookupRoleKey(""); // force refetch if viewing same role
    } catch (e) {
      // useCreateRolePermission onError already shows error
    }
  };

  const handleCreateRole = () => {
    if (!newRoleKey.trim() || !newRoleName.trim()) return;
    createRole.mutate({
      key: newRoleKey.trim(),
      name: newRoleName.trim(),
      description: newRoleDesc.trim() || undefined,
      scopeType: newRoleScope,
      isSystem: false,
    });
    setNewRoleKey("");
    setNewRoleName("");
    setNewRoleDesc("");
  };

  const handleCreatePermission = () => {
    if (!newPermKey.trim() || !newPermName.trim()) return;
    createPermission.mutate({
      key: newPermKey.trim(),
      name: newPermName.trim(),
      description: newPermDesc.trim() || undefined,
      isSystem: false,
    });
    setNewPermKey("");
    setNewPermName("");
    setNewPermDesc("");
  };

  return (
    <div className={styles.content}>
      <div className={styles.mgmtHeader}>
        <h2 className={styles.mgmtTitle}>{t("mgmtTitle")}</h2>
        <p className={styles.mgmtSubtitle}>{t("mgmtSubtitle")}</p>
      </div>

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

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("createPermission")}</h3>
        <div className={styles.formRow}>
          <Input
            label={t("key")}
            value={newPermKey}
            onChange={(e) => setNewPermKey(e.target.value)}
            placeholder="e.g. directory.users.read"
            fullWidth
          />
          <Input
            label={t("name")}
            value={newPermName}
            onChange={(e) => setNewPermName(e.target.value)}
            placeholder="Users Read"
            fullWidth
          />
        </div>
        <Input
          label={t("description")}
          value={newPermDesc}
          onChange={(e) => setNewPermDesc(e.target.value)}
          fullWidth
        />
        <Button onClick={handleCreatePermission} loading={createPermission.isPending}>
          {t("createPermission")}
        </Button>
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("assignPermissionToRole")}</h3>
        <div className={styles.formRow}>
          <Input
            label={t("roleKey")}
            value={roleKey}
            onChange={(e) => setRoleKey(e.target.value)}
            placeholder="e.g. system.admin"
            fullWidth
          />
          <Input
            label={t("permissionKey")}
            value={permissionKey}
            onChange={(e) => setPermissionKey(e.target.value)}
            placeholder="e.g. directory.*"
            fullWidth
          />
        </div>
        <Button onClick={handleAssignPermission} loading={createRolePermission.isPending}>
          {t("addPermission")}
        </Button>
      </div>

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("viewRolePermissions")}</h3>
        <div className={styles.formRow}>
          <Input
            label={t("roleKey")}
            value={lookupInput}
            onChange={(e) => setLookupInput(e.target.value)}
            placeholder="e.g. system.admin"
            fullWidth
          />
          <Button
            variant="secondary"
            onClick={() => setLookupRoleKey(lookupInput.trim())}
            disabled={!lookupInput.trim()}
          >
            {t("lookup")}
          </Button>
        </div>
        {lookupRoleKey && (
          <Suspense
            loadingType="ecg"
            loadingText={t("loading")}
            loadingSize="small"
            loadingContainerClassName={styles.loading}
          >
            <RolePermissionsLookupContent
              roleKey={lookupRoleKey}
              deleteRolePermission={deleteRolePermission}
            />
          </Suspense>
        )}
      </div>
    </div>
  );
});

PermissionManagement.displayName = "PermissionManagement";

/** 按 roleKey 查看角色及其权限，内部用 suspense hooks，需包在 Suspense 内 */
const RolePermissionsLookupContent = memo(
  ({
    roleKey,
    deleteRolePermission,
  }: {
    roleKey: string;
    deleteRolePermission: ReturnType<typeof useDeleteRolePermission>;
  }) => {
    const { t } = useTranslation("UserSecurityPage");
    const { data: role } = useRoleByKey({ key: roleKey });
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
  },
);
RolePermissionsLookupContent.displayName = "RolePermissionsLookupContent";

const RolePermissionRow = memo(
  ({
    permissionId,
    onRemove,
    isRemoving,
  }: {
    permissionId: string;
    onRemove: () => void;
    isRemoving: boolean;
  }) => {
    const { t } = useTranslation("UserSecurityPage");
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
  },
);

RolePermissionRow.displayName = "RolePermissionRow";

export default PermissionManagement;
