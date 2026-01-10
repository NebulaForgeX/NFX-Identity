import type { ReactNode } from "react";
import type { SidebarProps as ProSidebarProps } from "react-pro-sidebar";

import { memo } from "react";
import {
  Home,
} from "@/assets/icons/lucide";
import { Menu, MenuItem, Sidebar as ProSidebar, SubMenu } from "react-pro-sidebar";
import { Link, useLocation } from "react-router-dom";

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
    const location = useLocation();

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
                仪表盘
              </MenuItem>
            </Menu>
          )}
        </div>
      </ProSidebar>
    );
  },
);

Sidebar.displayName = "Sidebar";

export default Sidebar;
