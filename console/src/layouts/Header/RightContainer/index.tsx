import { memo, useCallback, useEffect, useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import { PreferencesPopover } from "nfx-ui/components";
import { useAuthRepository } from "nfx-ui/apis";
import { useAuthStore } from "nfx-ui/stores";
import { Mail } from "@/assets/icons/lucide";
import { routerEventEmitter } from "@/events/router";
import { clearLocalData } from "@/utils/clearLocalData";
import styles from "./styles.module.css";

const RightContainer = memo(() => {
  return (
    <div className={styles.headerContainer}>
      <div className={styles.actions}>
        <PreferencesPopover />
        <div className={styles.separator}></div>
        <UserEmail />
        <div className={styles.separator}></div>
        <UserMenu />
      </div>
    </div>
  );
});

RightContainer.displayName = "RightContainer";
export default RightContainer;

const UserEmail = memo(() => {
  const { t } = useTranslation("components");
  const auth = useAuthRepository();
  const accountId = useAuthStore((s) => s.currentAccountId);
  const [email, setEmail] = useState<string | null>(null);

  useEffect(() => {
    void auth.ListEmails({ limit: 20, offset: 0 }).then((list) => {
      const primary = list.items.find((row: { isPrimary: boolean }) => row.isPrimary) ?? list.items[0];
      setEmail(primary?.email ?? null);
    });
  }, [accountId, auth]);

  return (
    <div className={styles.contactInfo}>
      <Mail size={16} />
      <span className={styles.email}>{email || t("header.notSet")}</span>
    </div>
  );
});
UserEmail.displayName = "UserEmail";

const UserMenu = memo(() => {
  const { t } = useTranslation("components");
  const [open, setOpen] = useState(false);
  const accountId = useAuthStore((s) => s.currentAccountId);

  const handleLogout = useCallback(() => {
    void clearLocalData();
  }, []);

  const items = useMemo(
    () => [
      { title: t("header.profile"), action: () => routerEventEmitter.navigateToProfile() },
      { title: t("header.logout"), action: handleLogout },
    ],
    [handleLogout, t],
  );

  return (
    <div className={`${styles.userAction} ${styles.controlItem}`}>
      <button className={styles.user} onClick={() => setOpen((v) => !v)}>
        <span className={styles.userName}>{accountId.slice(0, 8)}</span>
      </button>
      {open ? (
        <div className={styles.contextMenu}>
          {items.map((item) => (
            <button
              key={item.title}
              className={styles.menuItem}
              onClick={() => {
                setOpen(false);
                item.action();
              }}
            >
              {item.title}
            </button>
          ))}
        </div>
      ) : null}
    </div>
  );
});

UserMenu.displayName = "UserMenu";
