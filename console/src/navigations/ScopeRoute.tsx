import { ProfileKindEnum } from "nfx-ui/enums";
import { hasSelectedProfile, useAuthStore } from "nfx-ui/stores";
import { Navigate, Outlet, useLocation } from "react-router";

import { profileHome, ROUTES } from "./routes";

export function IdentityGuestRoute() {
  const isAuthValid = useAuthStore((s) => s.isAuthValid);
  const currentProfileId = useAuthStore((s) => s.currentProfileId);
  const kind = useAuthStore((s) => s.currentProfileKind);

  if (isAuthValid && hasSelectedProfile(currentProfileId)) {
    return <Navigate to={profileHome(kind)} replace />;
  }
  if (isAuthValid) {
    return <Navigate to={ROUTES.SELECT_PROFILE} replace />;
  }
  return <Outlet />;
}

export function SelectProfileRoute() {
  const location = useLocation();
  const isAuthValid = useAuthStore((s) => s.isAuthValid);
  const currentProfileId = useAuthStore((s) => s.currentProfileId);
  const kind = useAuthStore((s) => s.currentProfileKind);

  if (!isAuthValid) {
    return <Navigate to={ROUTES.LOGIN} replace state={{ from: location }} />;
  }
  if (hasSelectedProfile(currentProfileId)) {
    return <Navigate to={profileHome(kind)} replace />;
  }
  return <Outlet />;
}

export function ScopeRoute({ scope }: { scope: ProfileKindEnum }) {
  const selectedKind = useAuthStore((s) => s.currentProfileKind);

  if (selectedKind !== scope) {
    return <Navigate to={profileHome(selectedKind)} replace />;
  }

  return <Outlet />;
}
