import { Bell, Save } from "lucide-react";
import { useState } from "react";
import { Box, Button, Flex, Switch, Text } from "@radix-ui/themes";
import { useCurrentProfile, useUpdateProfileSettings } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { LucideIcon, Surface } from "@/components";

import { LedgerSection } from "../../Ledger";
import styles from "./s.module.css";

export default function SystemSettings() {
  const { t } = useTranslation("pages.User.Setting", { keyPrefix: "systemSettings" });
  const { profile } = useCurrentProfile();
  const update = useUpdateProfileSettings({ successMsg: t("saveSuccess") });
  const baseline = profile?.settings?.loginNotification ?? true;
  const [edited, setEdited] = useState<Nullable<boolean>>(null);
  const loginNotification = edited ?? baseline;
  const dirty = loginNotification !== baseline;

  return (
    <LedgerSection
      title={t("cardTitle")}
      description={t("cardDesc")}
      actions={
        <Button size="2" disabled={!dirty || update.isPending} loading={update.isPending} onClick={() => update.mutate({ loginNotification }, { onSuccess: () => setEdited(null) })}>
          <LucideIcon icon={Save} size={14} />
          {t("save")}
        </Button>
      }
    >
      <Surface tone="inset" py="4" px="4">
        <Flex align="center" justify="between" gap="4">
          <Flex align="center" gap="3" minWidth="0">
            <Box className={styles.mark}>
              <Flex align="center" justify="center" width="100%" height="100%">
                <LucideIcon icon={Bell} size={16} />
              </Flex>
            </Box>
            <Flex direction="column" gap="1" minWidth="0">
              <Text as="label" htmlFor="login-notification" size="2" weight="medium">
                {t("loginEmailNotification")}
              </Text>
              <Text as="p" size="1" color="gray">
                {t("loginEmailNotificationDesc")}
              </Text>
            </Flex>
          </Flex>
          <Switch id="login-notification" size="2" checked={loginNotification} onCheckedChange={setEdited} />
        </Flex>
      </Surface>
    </LedgerSection>
  );
}
