import { Bell, Save } from "lucide-react";
import { useEffect, useState } from "react";
import { Button, Checkbox, Container, Flex, Section, Text } from "@radix-ui/themes";
import { useCurrentProfile, useUpdateProfileSettings } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { LucideIcon } from "@/components";

import styles from "./s.module.css";

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
    <Section size="1" py="4" className={styles.notice}>
      <Container size="2" px="4">
        <Flex direction="column" gap="3">
          <Flex align="start" justify="between" gap="3">
            <Flex direction="column" gap="1" minWidth="0">
              <Text size="2" weight="bold">
                {t("cardTitle")}
              </Text>
              <Text size="1" color="gray">
                {t("cardDesc")}
              </Text>
            </Flex>
            <Button size="2" disabled={!dirty || update.isPending} loading={update.isPending} onClick={() => update.mutate({ loginNotification })}>
              <LucideIcon icon={Save} size={14} />
              {t("save")}
            </Button>
          </Flex>
          <Text as="label" size="2">
            <Flex align="center" gap="2">
              <Checkbox checked={loginNotification} onCheckedChange={(c) => setLoginNotification(c === true)} />
              <LucideIcon icon={Bell} size={14} />
              {t("loginEmailNotification")}
            </Flex>
          </Text>
          <Text as="p" size="1" color="gray">
            {t("loginEmailNotificationDesc")}
          </Text>
        </Flex>
      </Container>
    </Section>
  );
}
