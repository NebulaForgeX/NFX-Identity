import type { ReactNode } from "react";

import { Box, Container, Flex, Section, Text } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { useTranslation } from "react-i18next";

import { Logo, PreferencesPopover } from "@/components";

import styles from "./s.module.css";

export default function AuthChrome({ children }: { children: ReactNode }) {
  const { t } = useTranslation("pages.Account.AuthShell");

  return (
    <Box className={styles.page}>
      <Flex direction="column" height="100%">
        <Section size="1" py="0.9rem" position="sticky" top="0" className={styles.header}>
          <Container size="4" px="1.25rem" width="100%" maxWidth="100%">
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
          </Container>
        </Section>
        <Box position="relative" overflow="auto" className={styles.body}>
          {children}
        </Box>
      </Flex>
    </Box>
  );
}
