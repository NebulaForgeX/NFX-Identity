import { Navigate, Route, Routes } from "react-router-dom";

import { LayoutSwitcher } from "@/layouts";
import { DashboardPage, LoginPage, NotFoundPage } from "@/pages";
import { ROUTES } from "@/types/navigation";

import "./App.module.css";

import { useAuthStore } from "./stores/authStore";

function App() {
  const isAuthValid = useAuthStore((state) => state.isAuthValid);

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
