import { Navigate, Route, Routes } from "react-router-dom";

import { TruckLoading } from "@/components";
import { LayoutSwitcher } from "@/layouts";
import { BootstrapPage, DashboardPage, LoginPage, NotFoundPage } from "@/pages";
import { ROUTES } from "@/types/navigation";

import "./App.css";

import styles from "./App.module.css";
import { useAuthInit } from "./hooks/useAuthInit";
import { useCacheInvalidation } from "./hooks/useCacheInvalidation";
import { useRouter } from "./hooks/useRouter";
import { useSystemInit } from "./hooks/useSystem";
import { useAuthStore } from "./stores/authStore";

function App() {
  useRouter();
  useCacheInvalidation(); // 监听缓存失效事件
  
  // 检查系统是否已初始化（公开接口，不需要认证）
  const { isInitialized: isSystemInitialized, isLoading: isSystemInitLoading } = useSystemInit();
  
  // 等待系统初始化检查完成
  if (isSystemInitLoading) {
    return (
      <div className={styles.loadingContainer}>
        <TruckLoading size="medium" />
        <p className={styles.loadingText}>检查系统状态...</p>
      </div>
    );
  }

  // 如果系统未初始化，显示 Bootstrap 页面
  if (isSystemInitialized === false) {
    return <BootstrapPage />;
  }

  const { isInitialized } = useAuthInit();
  const isAuthValid = useAuthStore((state) => state.isAuthValid);

  // 等待认证验证完成再渲染
  if (!isInitialized) {
    return (
      <div className={styles.loadingContainer}>
        <TruckLoading size="medium" />
        <p className={styles.loadingText}>验证中...</p>
      </div>
    );
  }

  if (!isAuthValid) {
    return (
      <Routes>
        <Route path={ROUTES.LOGIN} element={<LoginPage />} />
        <Route path="*" element={<Navigate to={ROUTES.LOGIN} replace />} />
      </Routes>
    );
  }

  return (
    <LayoutSwitcher>
      <Routes>
        <Route path={ROUTES.HOME} element={<DashboardPage />} />
        <Route path={ROUTES.DASHBOARD} element={<DashboardPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </LayoutSwitcher>
  );
}

export default App;
