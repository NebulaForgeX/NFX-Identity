import { GearIcon } from "nfx-ui/icons";
import { Button, Grid } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { PageHeader, Suspense } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { scopePaths } from "@/navigations";

import LanguageSettings from "./LanguageSettings";
import SystemSettings from "./SystemSettings";
import ThemeSettings from "./ThemeSettings";

export default function SettingsView() {
  const { t } = useTranslation("pages.User.Setting");
  const { kind } = useCurrentProfile();
  return (
    <PageFrame>
      <PageHeader
        icon={GearIcon}
        title={t("title")}
        description={t("description")}
        actions={
          <Button size="2" variant="outline" onClick={() => routerEventEmitter.navigate({ to: scopePaths(kind).overview })}>
            {t("actions.openProfile")}
          </Button>
        }
      />
      <ThemeSettings />
      <Grid columns={{ initial: "1", md: "2" }} gap="5" width="100%" align="start">
        <LanguageSettings />
        <Suspense>
          <SystemSettings />
        </Suspense>
      </Grid>
    </PageFrame>
  );
}
