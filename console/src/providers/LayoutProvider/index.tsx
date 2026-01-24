import type { ReactNode } from "react";

import { createContext, useCallback, useEffect, useState } from "react";

export type LayoutMode = "show" | "hide";

export interface LayoutContextType {
  sidebarOpen: boolean;
  layoutMode: LayoutMode;
  setSidebarOpen: (open: boolean) => void;
  toggleSidebar: () => void;
  closeSidebar: () => void;
  setLayoutMode: (mode: LayoutMode) => void;
}

export const LayoutContext = createContext<LayoutContextType | undefined>(undefined);

interface LayoutProviderProps {
  children: ReactNode;
  defaultLayoutMode?: LayoutMode;
}

export function LayoutProvider({ children, defaultLayoutMode = "show" }: LayoutProviderProps) {
  // 从 localStorage 读取初始状态（兼容 zustand persist 格式）
  const [sidebarOpen, setSidebarOpenState] = useState<boolean>(() => {
    const saved = localStorage.getItem("layout-storage");
    if (saved) {
      try {
        const parsed = JSON.parse(saved);
        // 兼容 zustand persist 格式: { state: { ... } } 或直接 { ... }
        return parsed.state?.sidebarOpen ?? parsed.sidebarOpen ?? false;
      } catch {
        return false;
      }
    }
    return false;
  });

  const [layoutMode, setLayoutModeState] = useState<LayoutMode>(() => {
    const saved = localStorage.getItem("layout-storage");
    if (saved) {
      try {
        const parsed = JSON.parse(saved);
        // 兼容 zustand persist 格式: { state: { ... } } 或直接 { ... }
        const mode = parsed.state?.layoutMode ?? parsed.layoutMode;
        return mode === "show" || mode === "hide" ? mode : defaultLayoutMode;
      } catch {
        return defaultLayoutMode;
      }
    }
    return defaultLayoutMode;
  });

  // 同步到 localStorage（使用与 zustand persist 兼容的格式）
  useEffect(() => {
    const storage = {
      state: {
        sidebarOpen,
        layoutMode,
      },
    };
    localStorage.setItem("layout-storage", JSON.stringify(storage));
  }, [sidebarOpen, layoutMode]);

  const setSidebarOpen = useCallback((open: boolean) => {
    setSidebarOpenState(open);
  }, []);

  const toggleSidebar = useCallback(() => {
    setSidebarOpenState((prev) => !prev);
  }, []);

  const closeSidebar = useCallback(() => {
    setSidebarOpenState(false);
  }, []);

  const setLayoutMode = useCallback((mode: LayoutMode) => {
    setLayoutModeState(mode);
  }, []);

  return (
    <LayoutContext.Provider
      value={{
        sidebarOpen,
        layoutMode,
        setSidebarOpen,
        toggleSidebar,
        closeSidebar,
        setLayoutMode,
      }}
    >
      {children}
    </LayoutContext.Provider>
  );
}
