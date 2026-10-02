import { GearIcon } from "nfx-ui/icons";
import { Button, Grid } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { ActionBar, PageHeader, Suspense } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { scopePaths } from "@/navigations";

import SystemSettings from "./SystemSettings";
import ThemeSettings from "./ThemeSettings";

export default function SettingsView() {
  const { t } = useTranslation("pages.User.Setting");
  const { kind } = useCurrentProfile();
  return (
    <PageFrame>
      <PageHeader icon={GearIcon} title={t("title")} description={t("description")} />
      <ActionBar>
        <Button size="2" variant="outline" onClick={() => routerEventEmitter.navigate({ to: scopePaths(kind).overview })}>
          {t("actions.openProfile")}
        </Button>
      </ActionBar>
      <Grid columns={{ initial: "1", lg: "minmax(0, 1.4fr) minmax(16rem, 0.6fr)" }} gap="6" width="100%" align="start">
        <ThemeSettings />
        <Suspense>
          <SystemSettings />
        </Suspense>
      </Grid>
    </PageFrame>
  );
}
