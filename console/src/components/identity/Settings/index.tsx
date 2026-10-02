import { GearIcon } from "nfx-ui/icons";
import { Grid } from "@radix-ui/themes";
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
      <Grid columns={{ initial: "1", lg: "minmax(0, 1.4fr) minmax(16rem, 0.6fr)" }} gap="6" width="100%" align="start">
        <ThemeSettings />
        <Suspense>
          <SystemSettings />
        </Suspense>
      </Grid>
    </PageFrame>
  );
}
