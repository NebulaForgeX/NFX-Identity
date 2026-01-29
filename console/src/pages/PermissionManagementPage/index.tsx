import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { Tab, TabGroup, TabList, TabPanel, TabPanels } from "@headlessui/react";

import { Button, Input, Suspense } from "@/components";
import { useDeleteRolePermission } from "@/hooks/useAccess";

import RoleByKeyContent from "./components/RoleByKeyContent";
import PermissionByKeyContent from "./components/PermissionByKeyContent";
import RolePermissionsLookupContent from "./components/RolePermissionsLookupContent";
import styles from "./styles.module.css";

const PermissionManagementPage = memo(() => {
  const { t } = useTranslation("PermissionManagementPage");
  const deleteRolePermission = useDeleteRolePermission();
  const [rolesKey, setRolesKey] = useState("");
  const [permissionsKey, setPermissionsKey] = useState("");
  const [rolePermRoleKey, setRolePermRoleKey] = useState("");
  const [lookupRoleKey, setLookupRoleKey] = useState("");

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <TabGroup as="div" className={styles.tabGroup}>
        <TabList className={styles.tabs} aria-label={t("tabsLabel")}>
          <Tab className={styles.tab}>{t("tabRoles")}</Tab>
          <Tab className={styles.tab}>{t("tabPermissions")}</Tab>
          <Tab className={styles.tab}>{t("tabRolePermissions")}</Tab>
        </TabList>
        <TabPanels className={styles.tabPanels}>
          <TabPanel>
            <div className={styles.card}>
              <h3 className={styles.cardTitle}>{t("getRoleByKey")}</h3>
              <div className={styles.formRow}>
                <Input
                  label={t("roleKey")}
                  value={rolesKey}
                  onChange={(e) => setRolesKey(e.target.value)}
                  placeholder="e.g. system.admin"
                  fullWidth
                />
                <Button
                  variant="secondary"
                  onClick={() => setRolesKey((k) => k.trim())}
                  disabled={!rolesKey.trim()}
                >
                  {t("lookup")}
                </Button>
              </div>
              {rolesKey.trim() && (
                <Suspense
                  loadingType="ecg"
                  loadingText={t("loading")}
                  loadingSize="small"
                  loadingContainerClassName={styles.loading}
                >
                  <RoleByKeyContent roleKey={rolesKey.trim()} />
                </Suspense>
              )}
            </div>
          </TabPanel>
          <TabPanel>
            <div className={styles.card}>
              <h3 className={styles.cardTitle}>{t("getPermissionByKey")}</h3>
              <div className={styles.formRow}>
                <Input
                  label={t("permissionKey")}
                  value={permissionsKey}
                  onChange={(e) => setPermissionsKey(e.target.value)}
                  placeholder="e.g. directory.*"
                  fullWidth
                />
                <Button
                  variant="secondary"
                  onClick={() => setPermissionsKey((k) => k.trim())}
                  disabled={!permissionsKey.trim()}
                >
                  {t("lookup")}
                </Button>
              </div>
              {permissionsKey.trim() && (
                <Suspense
                  loadingType="ecg"
                  loadingText={t("loading")}
                  loadingSize="small"
                  loadingContainerClassName={styles.loading}
                >
                  <PermissionByKeyContent permissionKey={permissionsKey.trim()} />
                </Suspense>
              )}
            </div>
          </TabPanel>
          <TabPanel>
            <div className={styles.card}>
              <h3 className={styles.cardTitle}>{t("getRolePermissionsByRole")}</h3>
              <div className={styles.formRow}>
                <Input
                  label={t("roleKey")}
                  value={rolePermRoleKey}
                  onChange={(e) => setRolePermRoleKey(e.target.value)}
                  placeholder="e.g. system.admin"
                  fullWidth
                />
                <Button
                  variant="secondary"
                  onClick={() => setLookupRoleKey(rolePermRoleKey.trim())}
                  disabled={!rolePermRoleKey.trim()}
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
          </TabPanel>
        </TabPanels>
      </TabGroup>
    </div>
  );
});

PermissionManagementPage.displayName = "PermissionManagementPage";

export default PermissionManagementPage;
