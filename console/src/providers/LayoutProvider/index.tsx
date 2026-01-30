import type { ReactNode } from "react";

import { createContext } from "react";
import useAction from "./hooks/useAction";
import useSet from "./hooks/useSet";

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
  const { sidebarOpen, setSidebarOpen, toggleSidebar, closeSidebar } = useAction();

  const { layoutMode, setLayoutMode } = useSet(defaultLayoutMode, sidebarOpen);


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
