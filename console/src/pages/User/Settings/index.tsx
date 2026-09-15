import type { ReactNode } from "react";
import { Box, Flex, Heading, Section, Text } from "@radix-ui/themes";
import { Settings2 } from "lucide-react";
import { useTranslation } from "react-i18next";
import { PageFrame } from "nfx-ui/layouts";
import { PageHeader, ThemeSettings } from "nfx-ui/components";

function SettingsSection({ id, title, description, children }: { id: string; title: string; description: string; children: ReactNode }) {
  return (
    <Section size="1" py="0" aria-labelledby={id}>
      <Box mb="3">
        <Heading as="h2" id={id} size="4" mb="1">
          {title}
        </Heading>
        <Text as="p" size="2" color="gray">
          {description}
        </Text>
      </Box>
      {children}
    </Section>
  );
}

export default function SettingsPage() {
  const { t } = useTranslation("EditPreferencePage");

  return (
    <PageFrame>
      <PageHeader icon={Settings2} title={t("title")} description={t("subtitle")} />
      <Flex direction="column" gap="6" width="100%">
        <SettingsSection id="settings-theme" title={t("title")} description={t("subtitle")}>
          <ThemeSettings />
        </SettingsSection>
      </Flex>
    </PageFrame>
  );
}
