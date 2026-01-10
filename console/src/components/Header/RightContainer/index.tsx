import { memo, useCallback, useMemo, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";

import { GetUser } from "@/apis/auth.api";
import { Bell, Mail, Search } from "@/assets/icons/lucide";
import { useLogOut } from "@/hooks/useAuth";
import { useAuthStore } from "@/stores/authStore";
// import { useSelfProfile } from "@/hooks/useProfile"; // TODO: 已删除 profile hooks，改用 GetUser
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
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const { data: user } = useQuery({
    queryKey: ["user", currentUserId],
    queryFn: () => (currentUserId ? GetUser(currentUserId) : null),
    enabled: !!currentUserId,
  });
  return (
    <div className={styles.contactInfo}>
      <Mail size={16} />
      <span className={styles.email}>{user?.email || "未设置"}</span>
    </div>
  );
});
UserEmail.displayName = "UserEmail";

const UserMenu = memo(() => {
  const [userMenuOpen, setUserMenuOpen] = useState(false);
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const { data: user } = useQuery({
    queryKey: ["user", currentUserId],
    queryFn: () => (currentUserId ? GetUser(currentUserId) : null),
    enabled: !!currentUserId,
  });
  const navigate = useNavigate();
  const { mutateAsync: logOut, isPending: isLoggingOut } = useLogOut();

  const handleLogout = useCallback(async () => {
    try {
      await logOut();
    } catch (error) {
      showError("退出登录失败: " + error);
    }
  }, [logOut]);

  const userMenu: Array<{ title: string; action: () => void; disabled?: boolean }> = useMemo(
    () => [
      { title: "个人资料", action: () => navigate(ROUTES.PROFILE) },
      { title: "退出登录", action: handleLogout, disabled: isLoggingOut },
    ],
    [navigate, handleLogout, isLoggingOut],
  );

  return (
    <div className={`${styles.userAction} ${styles.controlItem}`}>
      <button className={styles.user} onClick={() => setUserMenuOpen(!userMenuOpen)}>
        <img
          src={buildImageUrl(user?.profile?.avatarId, "avatar") || "/default-avatar.png"}
          alt={user?.username || "用户"}
          className={styles.userPicture}
        />
        <span className={styles.userName}>{user?.username || user?.email || "用户"}</span>
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
              {item.disabled && " (正在退出...)"}
            </button>
          ))}
        </div>
      )}
    </div>
  );
});

UserMenu.displayName = "UserMenu";
