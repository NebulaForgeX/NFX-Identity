import { Navigate, Route, Routes } from "react-router-dom";

import { TruckLoading } from "@/components";
import { LayoutSwitcher } from "@/layouts";
import { DashboardPage, LoginPage, NotFoundPage } from "@/pages";
import { ROUTES } from "@/types/navigation";

import "./App.css";

import styles from "./App.module.css";
import { useAuthInit } from "./hooks/useAuthInit";
import { useCacheInvalidation } from "./hooks/useCacheInvalidation";
import { useRouter } from "./hooks/useRouter";
import { useAuthStore } from "./stores/authStore";

function App() {
  useRouter();
  useCacheInvalidation(); // 监听缓存失效事件
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
