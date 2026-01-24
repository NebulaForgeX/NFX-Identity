import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Menu } from "@/assets/icons/lucide";
import { LayoutSwitcher, ThemeSwitcher } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { useLayout } from "@/hooks/useLayout";
import styles from "./styles.module.css";

const LeftContainer = memo(() => {
  const { t } = useTranslation("components");
  const { toggleSidebar } = useLayout();
  
  return (
    <div className={styles.headerContainer}>
      {/* 侧边栏切换按钮 */}
      <button
        type="button"
        className={styles.sidebarToggle}
        onClick={toggleSidebar}
      >
        <Menu size={28} />
      </button>

      {/* Logo */}
      <button
        type="button"
        className={styles.logo}
        onClick={() => {
          routerEventEmitter.navigateToDashboard();
        }}
      >
        {t("header.platformName")}
      </button>

      {/* 主题/布局 切换器 */}
      <ThemeSwitcher status="primary" />
      <LayoutSwitcher status="default" />
    </div>
  );
});

LeftContainer.displayName = "LeftContainer";

export default LeftContainer;
