import { useCallback, useState } from "react";

const useAction = () => {

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
const setSidebarOpen = useCallback((open: boolean) => {
    setSidebarOpenState(open);
  }, []);

  const toggleSidebar = useCallback(() => {
    setSidebarOpenState((prev) => !prev);
  }, []);

  const closeSidebar = useCallback(() => {
    setSidebarOpenState(false);
  }, []);
  return {
    sidebarOpen,
    setSidebarOpen,
    toggleSidebar,
    closeSidebar,
  };
};

export default useAction;