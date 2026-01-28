import { memo, useCallback, useMemo, useState } from "react";
import { useTranslation } from "react-i18next";

import { Bell, Mail, Search } from "@/assets/icons/lucide";
import { LanguageSwitcher } from "@/components";
import { authEventEmitter, authEvents } from "@/events/auth";
import { routerEventEmitter, routerEvents } from "@/events/router";
import { useUser, useUserEmailsByUserID, useUserAvatar } from "@/hooks/useDirectory";
import { useChatStore } from "@/stores/chatStore";
import { showSearch } from "@/stores/modalStore";
import { useAuthStore } from "@/stores/authStore";


import { buildImageUrl } from "@/utils/image";

import styles from "./styles.module.css";

const RightContainer = memo(() => {
  const unreadCount = useChatStore((state) => state.unreadCount);

  return (
    <div className={styles.headerContainer}>
      <div className={styles.actions}>
        {/* 语言切换器 */}
        <LanguageSwitcher status="default" />
        <div className={styles.separator}></div>
        {/* 邮箱信息 */}
        <UserEmail />
        <div className={styles.separator}></div>
        {/* 搜索按钮 */}
        <button className={`${styles.action} ${styles.controlItem}`} onClick={() => showSearch()}>
          <Search size={20} />
        </button>
        <div className={styles.separator}></div>
        {/* 通知按钮 */}
        <div className={`${styles.action} ${styles.controlItem}`} style={{ position: "relative" }}>
          <Bell size={20} />
          {unreadCount > 0 && <span className={styles.badge}>{unreadCount}</span>}
        </div>
        <div className={styles.separator}></div>
        {/* 用户菜单 */}
        <UserMenu />
      </div>
    </div>
  );
});

RightContainer.displayName = "RightContainer";
export default RightContainer;

const UserEmail = memo(() => {
  const { t } = useTranslation("components");
  const currentUserId = useAuthStore((state) => state.currentUserId);
  
  if (!currentUserId) {
    return (
      <div className={styles.contactInfo}>
        <Mail size={16} />
        <span className={styles.email}>{t("header.notSet")}</span>
      </div>
    );
  }
  
  const { data: userEmailsList } = useUserEmailsByUserID({
    userId: currentUserId,
  });
  
  const userEmails = Array.isArray(userEmailsList) && userEmailsList.length > 0
    ? (userEmailsList.find((email) => email?.isPrimary)?.email || userEmailsList[0]?.email || null)
    : null;
  
  return (
    <div className={styles.contactInfo}>
      <Mail size={16} />
      <span className={styles.email}>{userEmails || t("header.notSet")}</span>
    </div>
  );
});
UserEmail.displayName = "UserEmail";

const UserMenu = memo(() => {
  const { t } = useTranslation("components");
  const [userMenuOpen, setUserMenuOpen] = useState(false);
  const currentUserId = useAuthStore((state) => state.currentUserId);
  
  if (!currentUserId) return null;
  
  const { data: user } = useUser({
    id: currentUserId,
  });
  const { data: userEmailsList } = useUserEmailsByUserID({
    userId: currentUserId,
  });
  const { data: userAvatar } = useUserAvatar({
    userId: currentUserId,
  });
  
  const userEmails = Array.isArray(userEmailsList) && userEmailsList.length > 0
    ? (userEmailsList.find((email) => email?.isPrimary)?.email || userEmailsList[0]?.email || null)
    : null;
  
  const handleLogout = useCallback(() => {
    // 同时发送 auth 事件（用于清理）和 router 事件（用于导航）
    authEventEmitter.emit(authEvents.LOGOUT);
    routerEventEmitter.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }, []);

  const handleNavigateProfile = useCallback(() => {
    routerEventEmitter.navigateToProfile();
    setUserMenuOpen(false);
  }, []);

  const userMenu: Array<{ title: string; action: () => void; disabled?: boolean }> = useMemo(
    () => [
      { title: t("header.profile"), action: handleNavigateProfile },
      { title: t("header.logout"), action: handleLogout },
    ],
    [handleNavigateProfile, handleLogout, t],
  );

  return (
    <div className={`${styles.userAction} ${styles.controlItem}`}>
      <button
        className={styles.user}
        onClick={() => setUserMenuOpen(!userMenuOpen)}
      >
        <img
          src={userAvatar?.imageId ? buildImageUrl(userAvatar.imageId, "avatar") : "/default-avatar.png"}
          alt={user?.username || t("header.user")}
          className={styles.userPicture}
        />
        <span className={styles.userName}>{user?.username || userEmails || t("header.user")}</span>
      </button>

      {userMenuOpen && (
        <div className={styles.contextMenu}>
          {userMenu.map((item, index) => (
            <button
              key={index}
              className={styles.menuItem}
              onClick={() => {
                setUserMenuOpen(false);
                item.action();
              }}
              disabled={item.disabled}
            >
              {item.title}
              {item.disabled && t("header.loggingOut")}
            </button>
          ))}
        </div>
      )}
    </div>
  );
});

UserMenu.displayName = "UserMenu";
