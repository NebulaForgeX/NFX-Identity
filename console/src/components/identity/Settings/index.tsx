import { GearIcon } from "nfx-ui/icons";
import { useTranslation } from "react-i18next";

import { PageHeader, Suspense } from "@/components";
import { PageFrame } from "@/layouts";

import SystemSettings from "./SystemSettings";
import ThemeSettings from "./ThemeSettings";

export default function SettingsView() {
  const { t } = useTranslation("pages.User.Setting");
  return (
    <PageFrame>
      <PageHeader icon={GearIcon} title={t("title")} description={t("description")} />
      <ThemeSettings />
      <Suspense>
        <SystemSettings />
      </Suspense>
    </PageFrame>
  );
}
