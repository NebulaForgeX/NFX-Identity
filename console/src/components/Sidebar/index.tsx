import type { ReactNode } from "react";
import type { SidebarProps as ProSidebarProps } from "react-pro-sidebar";

import { memo, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { Menu, MenuItem, Sidebar as ProSidebar, SubMenu } from "react-pro-sidebar";
import { Link, useLocation } from "react-router-dom";

import { Home, User, Plus, Settings, Eye, LogOut } from "@/assets/icons/lucide";
import { authEventEmitter, authEvents } from "@/events/auth";
import { isActiveRoute, ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

interface SidebarProps extends ProSidebarProps {
  children?: ReactNode;
  collapsed?: boolean;
  toggled?: boolean;
  onBackdropClick?: () => void;
  className?: string;
}

const Sidebar = memo(
  ({ children, collapsed = false, toggled = false, onBackdropClick, breakPoint = "all", className }: SidebarProps) => {
    const { t } = useTranslation("components");
    const location = useLocation();

    const handleLogout = useCallback(() => {
      authEventEmitter.emit(authEvents.LOGOUT);
    }, []);

    return (
      <ProSidebar
        collapsed={collapsed}
        toggled={toggled}
        onBackdropClick={onBackdropClick}
        breakPoint={breakPoint}
        backgroundColor="var(--color-bg-2)"
        rootStyles={{
          border: "none",
          borderRight: "1px solid var(--color-separator)",
        }}
        className={`${styles.sidebar} ${className || ""}`}
      >
        <div className={styles.sidebarContent}>
          <div className={styles.menuWrapper}>
            {children || (
              <Menu
                key={`${collapsed}-${toggled}`} //! Prevent re-rendering when collapsed/toggled changes; Do not remove this!
                transitionDuration={300}
                closeOnClick
              menuItemStyles={{
                button: {
                  color: "var(--color-fg-text)",
                  backgroundColor: "transparent",
                  "&:hover": {
                    backgroundColor: "var(--color-bg-3)",
                    color: "var(--color-fg-text)",
                  },
                  "&.active": {
                    backgroundColor: "var(--color-primary)",
                    color: "#ffffff",
                  },
                },
                icon: {
                  color: "var(--color-fg-text)",
                  "&.active": {
                    color: "#ffffff",
                  },
                },
                label: {
                  color: "var(--color-fg-text)",
                  "&.active": {
                    color: "#ffffff",
                  },
                },
              }}
            >
              <MenuItem
                icon={<Home size={20} />}
                component={<Link to={ROUTES.DASHBOARD} />}
                active={
                  isActiveRoute(location.pathname, ROUTES.DASHBOARD) || isActiveRoute(location.pathname, ROUTES.HOME)
                }
              >
                {t("sidebar.dashboard")}
              </MenuItem>
              <SubMenu
                label={t("sidebar.profile")}
                icon={<User size={20} />}
                active={
                  isActiveRoute(location.pathname, ROUTES.PROFILE) ||
                  isActiveRoute(location.pathname, ROUTES.EDIT_PROFILE) ||
                  isActiveRoute(location.pathname, ROUTES.ADD_EDUCATION) ||
                  isActiveRoute(location.pathname, ROUTES.ADD_OCCUPATION) ||
                  isActiveRoute(location.pathname, ROUTES.EDIT_EDUCATION) ||
                  isActiveRoute(location.pathname, ROUTES.EDIT_OCCUPATION) ||
                  isActiveRoute(location.pathname, ROUTES.EDIT_PREFERENCE)
                }
              >
                <MenuItem
                  icon={<Eye size={18} />}
                  component={<Link to={ROUTES.PROFILE} />}
                  active={isActiveRoute(location.pathname, ROUTES.PROFILE)}
                >
                  {t("sidebar.profileView")}
                </MenuItem>
                <MenuItem
                  icon={<Settings size={18} />}
                  component={<Link to={ROUTES.EDIT_PROFILE} />}
                  active={isActiveRoute(location.pathname, ROUTES.EDIT_PROFILE)}
                >
                  {t("sidebar.editProfile")}
                </MenuItem>
                <MenuItem
                  icon={<Plus size={18} />}
                  component={<Link to={ROUTES.ADD_EDUCATION} />}
                  active={isActiveRoute(location.pathname, ROUTES.ADD_EDUCATION)}
                >
                  {t("sidebar.addEducation")}
                </MenuItem>
                <MenuItem
                  icon={<Plus size={18} />}
                  component={<Link to={ROUTES.ADD_OCCUPATION} />}
                  active={isActiveRoute(location.pathname, ROUTES.ADD_OCCUPATION)}
                >
                  {t("sidebar.addOccupation")}
                </MenuItem>
                <MenuItem
                  icon={<Settings size={18} />}
                  component={<Link to={ROUTES.EDIT_PREFERENCE} />}
                  active={isActiveRoute(location.pathname, ROUTES.EDIT_PREFERENCE)}
                >
                  {t("sidebar.editPreference")}
                </MenuItem>
              </SubMenu>
            </Menu>
            )}
          </div>
          {/* 退出登录按钮 - 固定在底部 */}
          <div className={styles.logoutContainer}>
            <button
              className={styles.logoutButton}
              onClick={handleLogout}
              title={t("header.logout")}
            >
              <LogOut size={20} />
              <span className={collapsed ? styles.hiddenText : styles.visibleText}>
                {t("header.logout")}
              </span>
            </button>
          </div>
        </div>
      </ProSidebar>
    );
  },
);

Sidebar.displayName = "Sidebar";

export default Sidebar;
