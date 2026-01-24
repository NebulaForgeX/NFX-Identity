import type { ReactNode } from "react";

import { memo, useCallback } from "react";

import Sidebar from "../Sidebar";
import { useLayout } from "@/hooks/useLayout";

import styles from "./styles.module.css";

interface SideHideLayoutProps {
  children: ReactNode;
  headerHeight: number;
  footerHeight: number;
}


const SideHideLayout = memo(({ children, headerHeight, footerHeight }: SideHideLayoutProps) => {
  const { sidebarOpen, closeSidebar } = useLayout();

  const handleBackdropClick = useCallback(() => {
    closeSidebar();
  }, [closeSidebar]);

  return (

      <main
        className={styles.mainWrapper}
        style={{
          marginTop: `${headerHeight}px`,
          marginBottom: `${footerHeight}px`,
        }}
      >
        {/* Sidebar */}
        <Sidebar
          toggled={sidebarOpen}
          onBackdropClick={handleBackdropClick}
          breakPoint="all"
          className={styles.sidebar}
        />

        {/* Content */}
        <div className={styles.content} data-lenis-prevent>
          {children}
        </div>
      </main>
  );
});

SideHideLayout.displayName = "SideHideLayout";

export default SideHideLayout;
