import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useAuthStore } from "@/stores/authStore";
import { useGrantsBySubject, useRoleById, usePermission, useRolePermissionsByRole } from "@/hooks/useAccess";
import type { Grant } from "@/types";

import styles from "./styles.module.css";

const UserSecurityPage = memo(() => {
  const { t } = useTranslation("UserSecurityPage");
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    return (
      <div className={styles.container}>
        <div className={styles.errorContainer}>
          <p>{t("userNotFound")}</p>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <Suspense
        loadingType="ecg"
        loadingText={t("loading")}
        loadingSize="small"
        loadingContainerClassName={styles.loading}
      >
        <UserSecurityContent userId={currentUserId} />
      </Suspense>
    </div>
  );
});

UserSecurityPage.displayName = "UserSecurityPage";

const UserSecurityContent = memo(({ userId }: { userId: string }) => {
  const { data: grants } = useGrantsBySubject({
    subject_type: "USER",
    subject_id: userId,
  });

  // 分离角色授权和权限授权
  const roleGrants = useMemo(
    () => grants?.filter((g) => g.grantType === "ROLE") || [],
    [grants],
  );
  const permissionGrants = useMemo(
    () => grants?.filter((g) => g.grantType === "PERMISSION") || [],
    [grants],
  );

  return (
    <div className={styles.content}>
      {/* 角色信息卡片 */}
      <RolesCard roleGrants={roleGrants} />

      {/* 权限信息卡片 */}
      <PermissionsCard permissionGrants={permissionGrants} roleGrants={roleGrants} />
    </div>
  );
});

UserSecurityContent.displayName = "UserSecurityContent";

// 角色卡片组件
const RolesCard = memo(({ roleGrants }: { roleGrants: Grant[] }) => {
  const { t } = useTranslation("UserSecurityPage");

  return (
    <div className={styles.card}>
      <h3 className={styles.cardTitle}>{t("roles")}</h3>
      {roleGrants.length === 0 ? (
        <p className={styles.emptyText}>{t("noRoles")}</p>
      ) : (
        <div className={styles.list}>
          {roleGrants.map((grant) => (
            <RoleItem key={grant.id} grant={grant} />
          ))}
        </div>
      )}
    </div>
  );
});

RolesCard.displayName = "RolesCard";

// 角色项组件
const RoleItem = memo(({ grant }: { grant: Grant }) => {
  const { t } = useTranslation("UserSecurityPage");
  const { data: role } = useRoleById({ id: grant.grantRefId });

  if (!role) {
    return null;
  }

  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>{role.name}</span>
        {role.isSystem && (
          <span className={styles.badge}>{t("systemRole")}</span>
        )}
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
        {grant.expiresAt && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("expiresAt")}:</span>
            <span className={styles.itemValue}>
              {new Date(grant.expiresAt).toLocaleString()}
            </span>
          </div>
        )}
        {grant.revokedAt && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("revokedAt")}:</span>
            <span className={styles.itemValue}>
              {new Date(grant.revokedAt).toLocaleString()}
            </span>
          </div>
        )}
      </div>
    </div>
  );
});

RoleItem.displayName = "RoleItem";

// 权限卡片组件
const PermissionsCard = memo(
  ({
    permissionGrants,
    roleGrants,
  }: {
    permissionGrants: Grant[];
    roleGrants: Grant[];
  }) => {
    const { t } = useTranslation("UserSecurityPage");

    // 获取所有角色包含的权限
    const rolePermissionIds = useMemo(() => {
      const ids = new Set<string>();
      roleGrants.forEach((grant) => {
        ids.add(grant.grantRefId);
      });
      return ids;
    }, [roleGrants]);

    return (
      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("permissions")}</h3>
        {permissionGrants.length === 0 && rolePermissionIds.size === 0 ? (
          <p className={styles.emptyText}>{t("noPermissions")}</p>
        ) : (
          <div className={styles.list}>
            {/* 直接授予的权限 */}
            {permissionGrants.map((grant) => (
              <PermissionItem key={grant.id} grant={grant} />
            ))}
            {/* 通过角色获得的权限 */}
            {Array.from(rolePermissionIds).map((roleId) => (
              <RolePermissions key={roleId} roleId={roleId} />
            ))}
          </div>
        )}
      </div>
    );
  },
);

PermissionsCard.displayName = "PermissionsCard";

// 权限项组件
const PermissionItem = memo(({ grant }: { grant: Grant }) => {
  const { t } = useTranslation("UserSecurityPage");
  const { data: permission } = usePermission({ id: grant.grantRefId });

  if (!permission) {
    return null;
  }

  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>{permission.name}</span>
        {permission.isSystem && (
          <span className={styles.badge}>{t("systemPermission")}</span>
        )}
        <span className={styles.badgeSecondary}>{t("directGrant")}</span>
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
        {grant.expiresAt && (
          <div className={styles.itemDetail}>
            <span className={styles.itemLabel}>{t("expiresAt")}:</span>
            <span className={styles.itemValue}>
              {new Date(grant.expiresAt).toLocaleString()}
            </span>
          </div>
        )}
      </div>
    </div>
  );
});

PermissionItem.displayName = "PermissionItem";

// 角色权限组件（显示角色包含的所有权限）
const RolePermissions = memo(({ roleId }: { roleId: string }) => {
  const { t } = useTranslation("UserSecurityPage");
  const { data: role } = useRoleById({ id: roleId });
  const { data: rolePermissions } = useRolePermissionsByRole({ roleId });

  if (!role) {
    return null;
  }

  if (!rolePermissions || rolePermissions.length === 0) {
    return (
      <div className={styles.item}>
        <div className={styles.itemHeader}>
          <span className={styles.itemName}>
            {t("permissionsFromRole")}: {role.name}
          </span>
          <span className={styles.badgeSecondary}>{t("viaRole")}</span>
        </div>
        <div className={styles.itemDetails}>
          <div className={styles.itemDetail}>
            <span className={styles.itemValue}>{t("noPermissions")}</span>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.item}>
      <div className={styles.itemHeader}>
        <span className={styles.itemName}>
          {t("permissionsFromRole")}: {role.name}
        </span>
        <span className={styles.badgeSecondary}>{t("viaRole")}</span>
      </div>
      <div className={styles.itemDetails}>
        <div className={styles.itemDetail}>
          <span className={styles.itemLabel}>{t("permissions")}:</span>
          <div className={styles.permissionList}>
            {rolePermissions.map((rp) => (
              <RolePermissionItem key={rp.id} permissionId={rp.permissionId} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
});

// 角色权限项组件
const RolePermissionItem = memo(({ permissionId }: { permissionId: string }) => {
  const { t } = useTranslation("UserSecurityPage");
  const { data: permission } = usePermission({ id: permissionId });

  if (!permission) {
    return null;
  }

  return (
    <div className={styles.permissionItem}>
      <span className={styles.permissionName}>{permission.name}</span>
      {permission.isSystem && (
        <span className={styles.badgeSmall}>{t("systemPermission")}</span>
      )}
      <span className={styles.permissionKey}>({permission.key})</span>
    </div>
  );
});

RolePermissionItem.displayName = "RolePermissionItem";

RolePermissions.displayName = "RolePermissions";

export default UserSecurityPage;
