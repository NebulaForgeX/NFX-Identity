import { Flex } from "@radix-ui/themes";
import { Settings2 } from "lucide-react";
import { useTranslation } from "react-i18next";
import { PageFrame } from "nfx-ui/layouts";
import { PageHeader, ThemeSettings } from "nfx-ui/components";

export default function SettingsPage() {
  const { t } = useTranslation("EditPreferencePage");

  return (
    <PageFrame>
      <PageHeader icon={Settings2} title={t("title")} description={t("subtitle")} />
      <Flex direction="column" gap="6" width="100%">
        <ThemeSettings />
      </Flex>
    </PageFrame>
  );
}
