import { Navigate, Route, Routes, useLocation } from "react-router";
import { useAuthStore, hasSelectedProfile } from "nfx-ui/stores";

import { ConsoleLayout } from "@/layouts";
import { DashboardPage, GitHubCallbackPage, ImagesPage, LoginPage, NotFoundPage, OwnerDirectoryPage, ProfilePage, SelectProfilePage, SettingsPage } from "@/pages";
import { ROUTES } from "@/navigations";

import "./App.module.css";

function App() {
  const location = useLocation();
  const accessToken = useAuthStore((state) => state.accessToken);
  const profileId = useAuthStore((state) => state.currentProfileId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);

  if (location.pathname === ROUTES.LOGIN_GITHUB_CALLBACK) {
    return (
      <Routes>
        <Route path={ROUTES.LOGIN_GITHUB_CALLBACK} element={<GitHubCallbackPage />} />
      </Routes>
    );
  }

  if (!accessToken) {
    return (
      <Routes>
        <Route path={ROUTES.LOGIN} element={<LoginPage />} />
        <Route path="*" element={<Navigate to={ROUTES.LOGIN} replace />} />
      </Routes>
    );
  }

  if (!isAuthValid || !hasSelectedProfile(profileId)) {
    return (
      <Routes>
        <Route path={ROUTES.SELECT_PROFILE} element={<SelectProfilePage />} />
        <Route path="*" element={<Navigate to={ROUTES.SELECT_PROFILE} replace />} />
      </Routes>
    );
  }

  return (
    <ConsoleLayout>
      <Routes>
        <Route path={ROUTES.HOME} element={<DashboardPage />} />
        <Route path={ROUTES.DASHBOARD} element={<DashboardPage />} />
        <Route path={ROUTES.PROFILE} element={<ProfilePage />} />
        <Route path={ROUTES.IMAGES} element={<ImagesPage />} />
        <Route path={ROUTES.SETTINGS} element={<SettingsPage />} />
        <Route path={ROUTES.OWNER} element={<OwnerDirectoryPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </ConsoleLayout>
  );
}

export default App;
