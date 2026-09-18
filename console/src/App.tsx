import { GuestRoute, ProtectedRoute } from "nfx-ui/navigations";
import { Navigate, Route, Routes } from "react-router";

import { Main, Sidebar } from "@/layouts";
import { ROUTES } from "@/navigations";
import {
  DashboardPage,
  ImagesPage,
  LoginPage,
  NotFoundPage,
  OwnerDirectoryPage,
  ProfileEditPage,
  ProfileIdentitiesPage,
  ProfileOverviewPage,
  SettingsPage,
  SignupPage,
} from "@/pages";

export default function App() {
  return (
    <Routes>
      <Route element={<Main />}>
        <Route path={ROUTES.HOME} element={<Navigate to={ROUTES.LOGIN} replace />} />
      </Route>

      <Route element={<ProtectedRoute redirectTo={ROUTES.LOGIN} />}>
        <Route element={<Sidebar />}>
          <Route path={ROUTES.USER} element={<Navigate to={ROUTES.USER_OVERVIEW} replace />} />
          <Route path={ROUTES.USER_OVERVIEW} element={<DashboardPage />} />
          <Route path={ROUTES.PROFILE} element={<Navigate to={ROUTES.USER_PROFILE_OVERVIEW} replace />} />
          <Route path={ROUTES.USER_PROFILE_OVERVIEW} element={<ProfileOverviewPage />} />
          <Route path={ROUTES.USER_PROFILE_EDIT} element={<ProfileEditPage />} />
          <Route path={ROUTES.USER_PROFILE_IDENTITIES} element={<ProfileIdentitiesPage />} />
          <Route path={ROUTES.USER_SETTINGS} element={<SettingsPage />} />
          <Route path={ROUTES.IMAGES} element={<ImagesPage />} />
          <Route path={ROUTES.OWNER} element={<OwnerDirectoryPage />} />
        </Route>
      </Route>

      <Route element={<GuestRoute redirectTo={ROUTES.USER_OVERVIEW} />}>
        <Route path={ROUTES.LOGIN} element={<LoginPage />} />
        <Route path={ROUTES.SIGNUP} element={<SignupPage />} />
      </Route>

      <Route path="*" element={<NotFoundPage />} />
    </Routes>
  );
}
