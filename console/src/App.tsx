import { ProtectedRoute } from "nfx-ui/navigations";
import { ProfileKindEnum } from "nfx-ui/enums";
import { Navigate, Route, Routes } from "react-router";

import { Sidebar } from "@/layouts";
import { IdentityGuestRoute, ROUTES, ScopeRoute, SelectProfileRoute } from "@/navigations";
import {
  AuthorityAssetsPage,
  AuthorityDeskPage,
  AuthorityDirectoryPage,
  AuthorityProfileEditPage,
  AuthorityProfileIdentityPage,
  AuthorityProfileOverviewPage,
  AuthorityProfileSecurityPage,
  AuthoritySettingsPage,
  ForgerAssetsPage,
  ForgerDeskPage,
  ForgerProfileEditPage,
  ForgerProfileIdentityPage,
  ForgerProfileOverviewPage,
  ForgerProfileSecurityPage,
  ForgerSettingsPage,
  LoginPage,
  NotFoundPage,
  SelectProfilePage,
  SignupPage,
} from "@/pages";

export default function App() {
  return (
    <Routes>
      <Route element={<IdentityGuestRoute />}>
        <Route index element={<LoginPage />} />
        <Route path={ROUTES.LOGIN} element={<LoginPage />} />
        <Route path={ROUTES.SIGNUP} element={<SignupPage />} />
      </Route>

      <Route element={<SelectProfileRoute />}>
        <Route path={ROUTES.SELECT_PROFILE} element={<SelectProfilePage />} />
      </Route>

      <Route element={<ProtectedRoute redirectTo={ROUTES.LOGIN} />}>
        <Route element={<Sidebar />}>
          <Route element={<ScopeRoute scope={ProfileKindEnum.FORGER} />}>
            <Route path={ROUTES.FORGER} element={<Navigate to={ROUTES.FORGER_DESK} replace />} />
            <Route path={ROUTES.FORGER_DESK} element={<ForgerDeskPage />} />
            <Route path={ROUTES.FORGER_PROFILE} element={<Navigate to={ROUTES.FORGER_PROFILE_OVERVIEW} replace />} />
            <Route path={ROUTES.FORGER_PROFILE_OVERVIEW} element={<ForgerProfileOverviewPage />} />
            <Route path={ROUTES.FORGER_PROFILE_EDIT} element={<ForgerProfileEditPage />} />
            <Route path={ROUTES.FORGER_PROFILE_IDENTITY} element={<ForgerProfileIdentityPage />} />
            <Route path={ROUTES.FORGER_PROFILE_SECURITY} element={<ForgerProfileSecurityPage />} />
            <Route path={ROUTES.FORGER_ASSETS} element={<ForgerAssetsPage />} />
            <Route path={ROUTES.FORGER_SETTINGS} element={<ForgerSettingsPage />} />
          </Route>
          <Route element={<ScopeRoute scope={ProfileKindEnum.AUTHORITY} />}>
            <Route path={ROUTES.AUTHORITY} element={<Navigate to={ROUTES.AUTHORITY_DESK} replace />} />
            <Route path={ROUTES.AUTHORITY_DESK} element={<AuthorityDeskPage />} />
            <Route path={ROUTES.AUTHORITY_PROFILE} element={<Navigate to={ROUTES.AUTHORITY_PROFILE_OVERVIEW} replace />} />
            <Route path={ROUTES.AUTHORITY_PROFILE_OVERVIEW} element={<AuthorityProfileOverviewPage />} />
            <Route path={ROUTES.AUTHORITY_PROFILE_EDIT} element={<AuthorityProfileEditPage />} />
            <Route path={ROUTES.AUTHORITY_PROFILE_IDENTITY} element={<AuthorityProfileIdentityPage />} />
            <Route path={ROUTES.AUTHORITY_PROFILE_SECURITY} element={<AuthorityProfileSecurityPage />} />
            <Route path={ROUTES.AUTHORITY_ASSETS} element={<AuthorityAssetsPage />} />
            <Route path={ROUTES.AUTHORITY_SETTINGS} element={<AuthoritySettingsPage />} />
            <Route path={ROUTES.AUTHORITY_DIRECTORY} element={<AuthorityDirectoryPage />} />
          </Route>
        </Route>
      </Route>

      <Route path="*" element={<NotFoundPage />} />
    </Routes>
  );
}
