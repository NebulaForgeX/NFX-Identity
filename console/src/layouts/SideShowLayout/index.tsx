import type { ReactNode } from "react";

import { memo, useCallback, useLayoutEffect, useRef, useState } from "react";

import { Sidebar } from "@/components";
import LayoutStore, { useLayoutStore } from "@/stores/layoutStore";
  
import styles from "./styles.module.css";

interface SideShowLayoutProps {
  children: ReactNode;
  headerHeight: number;
  footerHeight: number;
}

function useElementWidth<T extends HTMLElement>() {
  const ref = useRef<T | null>(null);
  const [width, setWidth] = useState(0);

  useLayoutEffect(() => {
    const node = ref.current;
    if (!node) return;

    const observer = new ResizeObserver(([entry]) => {
      const nextWidth = entry?.contentRect.width ?? 0;
      setWidth((prev) => (prev !== nextWidth ? nextWidth : prev));
    });

    observer.observe(node);
    return () => observer.disconnect();
  }, []);

  const callbackRef = useCallback((instance: T | null) => {
    if (instance) {
      ref.current = instance;
    }
  }, []);

  return [callbackRef, width] as const;
}

const SideShowLayout = memo(({ children, headerHeight, footerHeight }: SideShowLayoutProps) => {
  const sidebarOpen = useLayoutStore((state) => state.sidebarOpen);
  const closeSidebar = LayoutStore.getState().closeSidebar;
  const [sidebarRef, sidebarWidth] = useElementWidth<HTMLDivElement>();

  const handleBackdropClick = () => {
    closeSidebar();
  };

  return (
<>

      <div
        ref={sidebarRef}
        className={styles.sidebarContainer}
        style={{
          top: `${headerHeight}px`,
          height: `calc(100vh - ${headerHeight + footerHeight}px)`,
        }}
      >
        <Sidebar
          collapsed={sidebarOpen}
          toggled={sidebarOpen}
          onBackdropClick={handleBackdropClick}
          breakPoint="xs"
          className={styles.sidebar}
        />
      </div>
      {/* Main Content Area with Sidebar */}
      <main
        className={styles.mainWrapper}
        style={{
          marginTop: `${headerHeight}px`,
          marginBottom: `${footerHeight}px`,
          // transform: `translateX(${sidebarWidth}px)`,
          marginLeft: `${sidebarWidth}px`,
          width: `calc(100% - ${sidebarWidth}px)`,
        }}
      >
        {/* Content */}
        <div className={styles.content} data-lenis-prevent>
          {children}
        </div>
      </main>

      </>
  );
});

SideShowLayout.displayName = "SideShowLayout";

export default SideShowLayout;
