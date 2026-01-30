import { useCallback, useEffect, useState } from "react";
import { LayoutMode } from "..";

const useSet = (defaultLayoutMode: LayoutMode, sidebarOpen: boolean) => {
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
    
      const setLayoutMode = useCallback((mode: LayoutMode) => {
        setLayoutModeState(mode);
      }, []);
      return {
        layoutMode,
        setLayoutMode,
      };
};

export default useSet;