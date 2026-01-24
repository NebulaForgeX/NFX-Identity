import { Navigate, Route, Routes } from "react-router-dom";

import { LayoutSwitcher } from "@/layouts";
import {
  AddEducationPage,
  AddOccupationPage,
  DashboardPage,
  EditEducationPage,
  EditOccupationPage,
  EditPreferencePage,
  EditProfilePage,
  LoginPage,
  NotFoundPage,
  ProfilePage,
} from "@/pages";
import { ROUTES } from "@/types/navigation";

import "./App.module.css";

import { useAuthStore } from "./stores/authStore";
import { useUserPreferenceSync } from "./hooks/useUserPreferenceSync";

function App() {
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  
  // Sync and apply user preferences on login
  useUserPreferenceSync();

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
        <Route path={ROUTES.PROFILE} element={<ProfilePage />} />
        <Route path={ROUTES.EDIT_PROFILE} element={<EditProfilePage />} />
        <Route path={ROUTES.ADD_EDUCATION} element={<AddEducationPage />} />
        <Route path={ROUTES.ADD_OCCUPATION} element={<AddOccupationPage />} />
        <Route path={ROUTES.EDIT_EDUCATION} element={<EditEducationPage />} />
        <Route path={ROUTES.EDIT_OCCUPATION} element={<EditOccupationPage />} />
        <Route path={ROUTES.EDIT_PREFERENCE} element={<EditPreferencePage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </LayoutSwitcher>
  );
}

export default App;
