import type { ReactNode } from "react";

import { Box, Flex, Text } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { useTranslation } from "react-i18next";

import { Logo, PreferencesPopover } from "@/components";

import styles from "./s.module.css";

export default function AuthChrome({ children }: { children: ReactNode }) {
  const { t } = useTranslation("pages.Account.AuthShell");

  return (
    <Flex direction="column" className={styles.page} asChild>
      <Box>
        <Box className={styles.header}>
          <Box className={styles.headerPx}>
            <Box className={styles.headerPy}>
        <Flex asChild align="center" justify="between" gap="4">
          <header>
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
        </Flex>
            </Box>
          </Box>
        </Box>
        <Box className={styles.body}>{children}</Box>
      </Box>
    </Flex>
  );
}
