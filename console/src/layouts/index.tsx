import type { SidebarMenuItem } from "nfx-ui/layouts";
import type { ReactNode } from "react";

import { memo, useCallback, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { useLocation } from "react-router";

import { LayoutFrame } from "nfx-ui/layouts";
import { Logo } from "nfx-ui/components";

import { Home, Image, Settings, Shield, User } from "@/assets/icons/lucide";
import { routerEventEmitter } from "@/events/router";
import { ROUTES } from "@/navigations";
import { clearLocalData } from "@/utils/clearLocalData";
import RightContainer from "./Header/RightContainer";

interface ConsoleLayoutProps {
  children: ReactNode;
}

function useSidebarItems(): SidebarMenuItem[] {
  const { t } = useTranslation("components");
  return useMemo(
    () => [
      { label: t("sidebar.dashboard", { defaultValue: "Dashboard" }), path: ROUTES.DASHBOARD, icon: <Home size={20} /> },
      { label: t("sidebar.profile", { defaultValue: "Profile" }), path: ROUTES.PROFILE, icon: <User size={20} /> },
      { label: t("sidebar.assets", { defaultValue: "Assets" }), path: ROUTES.IMAGES, icon: <Image size={20} /> },
      { label: t("sidebar.owner", { defaultValue: "Owner directory" }), path: ROUTES.OWNER, icon: <Shield size={20} /> },
      { label: t("sidebar.settings", { defaultValue: "Settings" }), path: ROUTES.SETTINGS, icon: <Settings size={20} /> },
    ],
    [t],
  );
}

export const ConsoleLayout = memo(({ children }: ConsoleLayoutProps) => {
  const { t } = useTranslation("components");
  const location = useLocation();
  const sidebarItems = useSidebarItems();

  const onSidebarNavigate = useCallback((path: string) => {
    routerEventEmitter.navigate({ to: path });
  }, []);

  const onSidebarLogout = useCallback(() => {
    void clearLocalData();
  }, []);

  return (
    <LayoutFrame
      headerLeft={<Logo title="NFX" subtitle="Identity" alt="NFX" onClick={() => routerEventEmitter.navigateToDashboard()} />}
      headerRight={<RightContainer />}
      sidebarItems={sidebarItems}
      sidebarCurrentPathname={location.pathname}
      onSidebarNavigate={onSidebarNavigate}
      sidebarLogoutLabel={t("header.logout")}
      onSidebarLogout={onSidebarLogout}
    >
      {children}
    </LayoutFrame>
  );
});

ConsoleLayout.displayName = "ConsoleLayout";
export default ConsoleLayout;
