import { memo, useCallback, useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";

import { GetUser, GetUserEmailsByUserID, GetUserProfile } from "@/apis/directory.api";
import { Bell, Mail, Search } from "@/assets/icons/lucide";
import { LanguageSwitcher } from "@/components";
import { useAuthStore } from "@/stores/authStore";
import { useChatStore } from "@/stores/chatStore";
import { showError, showSearch } from "@/stores/modalStore";
import { ROUTES } from "@/types/navigation";
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
  const { data: userEmails } = useQuery({
    queryKey: ["user-emails", currentUserId],
    queryFn: () => (currentUserId ? GetUserEmailsByUserID(currentUserId) : null),
    enabled: !!currentUserId,
    select: (data) => {
      // 优先返回主邮箱，否则返回第一个邮箱
      if (!data || data.length === 0) return null;
      const primaryEmail = data.find((email) => email.isPrimary);
      return primaryEmail?.email || data[0]?.email || null;
    },
  });
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
  const clearAuth = useAuthStore((state) => state.clearAuth);
  const { data: user } = useQuery({
    queryKey: ["user", currentUserId],
    queryFn: () => (currentUserId ? GetUser(currentUserId) : null),
    enabled: !!currentUserId,
  });
  const { data: userProfile } = useQuery({
    queryKey: ["user-profile", currentUserId],
    queryFn: () => (currentUserId ? GetUserProfile(currentUserId) : null),
    enabled: !!currentUserId,
  });
  const { data: userEmails } = useQuery({
    queryKey: ["user-emails", currentUserId],
    queryFn: () => (currentUserId ? GetUserEmailsByUserID(currentUserId) : null),
    enabled: !!currentUserId,
    select: (data) => {
      // 优先返回主邮箱，否则返回第一个邮箱
      if (!data || data.length === 0) return null;
      const primaryEmail = data.find((email) => email.isPrimary);
      return primaryEmail?.email || data[0]?.email || null;
    },
  });
  const navigate = useNavigate();

  const handleLogout = useCallback(() => {
    try {
      clearAuth();
      navigate(ROUTES.LOGIN);
    } catch (error) {
      showError(t("header.logoutFailed") + error);
    }
  }, [clearAuth, navigate, t]);

  const handleNavigateProfile = useCallback(() => {
    navigate(ROUTES.PROFILE);
    setUserMenuOpen(false);
  }, [navigate]);

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
          src={buildImageUrl(userProfile?.avatarId, "avatar") || "/default-avatar.png"}
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
