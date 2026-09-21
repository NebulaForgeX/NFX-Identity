import { Bell, Save } from "lucide-react";
import { useEffect, useState } from "react";
import { Button, Checkbox, Flex, Text } from "@radix-ui/themes";
import { useCurrentProfile, useUpdateProfileSettings } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { LucideIcon } from "@/components";

import { LedgerSection } from "../Ledger";

export default function SystemSettings() {
  const { t } = useTranslation("pages.User.Setting", { keyPrefix: "systemSettings" });
  const { profile } = useCurrentProfile();
  const update = useUpdateProfileSettings();
  const [loginNotification, setLoginNotification] = useState(true);

  useEffect(() => {
    if (profile?.settings) {
      setLoginNotification(profile.settings.loginNotification);
      return;
    }
    setLoginNotification(true);
  }, [profile]);

  const baseline = profile?.settings?.loginNotification ?? true;
  const dirty = loginNotification !== baseline;

  return (
    <LedgerSection
      title={t("cardTitle")}
      description={t("cardDesc")}
      actions={
        <Button size="2" disabled={!dirty || update.isPending} loading={update.isPending} onClick={() => update.mutate({ loginNotification })}>
          <LucideIcon icon={Save} size={14} />
          {t("save")}
        </Button>
      }
    >
      <Text as="label" size="2">
        <Flex align="center" gap="2">
          <Checkbox checked={loginNotification} onCheckedChange={(c) => setLoginNotification(c === true)} />
          <LucideIcon icon={Bell} size={14} />
          {t("loginEmailNotification")}
        </Flex>
      </Text>
      <Text as="p" size="1" color="gray" mt="2">
        {t("loginEmailNotificationDesc")}
      </Text>
    </LedgerSection>
  );
}
