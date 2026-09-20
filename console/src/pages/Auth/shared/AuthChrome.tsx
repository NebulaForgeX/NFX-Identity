import type { ReactNode } from "react";

import { Flex, Text } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { useTranslation } from "react-i18next";

import { Logo, PreferencesPopover } from "@/components";

import styles from "./s.module.css";

export default function AuthChrome({ children }: { children: ReactNode }) {
  const { t } = useTranslation("pages.Account.AuthShell");

  return (
    <div className={styles.page}>
      <header className={styles.header}>
        <Flex align="center" gap="3" minWidth="0">
          <Logo variant="glassSquare" size="medium" alt={`${APP_NAME} logo`} />
          <Flex direction="column" gap="0" minWidth="0">
            <Text size="2" weight="bold">
              {APP_NAME}
            </Text>
            <Text size="1" className={styles.brandHome}>
              {t("brandHome")}
            </Text>
          </Flex>
        </Flex>
        <PreferencesPopover />
      </header>
      <div className={styles.body}>{children}</div>
    </div>
  );
}
