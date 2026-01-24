import type { ReactNode } from "react";

import { memo, useCallback } from "react";

import { Sidebar } from "@/components";
import LayoutStore, { useLayoutStore } from "@/stores/layoutStore";



import styles from "./styles.module.css";

interface SideHideLayoutProps {
  children: ReactNode;
  headerHeight: number;
  footerHeight: number;
}


const SideHideLayout = memo(({ children, headerHeight, footerHeight }: SideHideLayoutProps) => {
  const sidebarOpen = useLayoutStore((state) => state.sidebarOpen);
  const closeSidebar = LayoutStore.getState().closeSidebar;


  const handleBackdropClick = useCallback(() => {
    closeSidebar();
  }, [closeSidebar]);

  console.log("headerHeight", headerHeight);
  console.log("footerHeight", footerHeight);

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
