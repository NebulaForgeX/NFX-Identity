import { Box, Container, Flex, Section, Text } from "@radix-ui/themes";
import { AnimatedIcon, RightChevron } from "nfx-ui/icons";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import { useLocation } from "react-router";

import { PreferencesPopover } from "@/components";
import { scopePaths } from "@/navigations";

import styles from "./s.module.css";

function useCrumbs(): string[] {
  const { t } = useTranslation("language");
  const { kind } = useCurrentProfile();
  const { pathname } = useLocation();
  const paths = scopePaths(kind);
  const under = (to: string) => pathname === to || pathname.startsWith(`${to}/`);

  if (under(paths.desk)) return [t("sidebar.desk")];
  if (under(paths.overview)) return [t("sidebar.profile"), t("sidebar.profileOverview")];
  if (under(paths.edit)) return [t("sidebar.profile"), t("sidebar.profileEdit")];
  if (under(paths.identity)) return [t("sidebar.profile"), t("sidebar.profileIdentities")];
  if (under(paths.security)) return [t("sidebar.profile"), t("sidebar.profileSecurity")];
  if (under(paths.assets)) return [t("sidebar.assets")];
  if (paths.directory && under(paths.directory)) return [t("sidebar.directory")];
  if (under(paths.settings)) return [t("sidebar.settingsItem")];
  return [];
}

export default function UserTopBar() {
  const { t } = useTranslation("language");
  const crumbs = useCrumbs();

  return (
    <Box position="sticky" top="0" className={styles.bar}>
      <Container size="4" width="100%" maxWidth="1440px" px={{ initial: "4", md: "6" }}>
        <Section size="1" py="3">
          <Flex align="center" justify="between" gap="3">
            <Flex align="center" gap="2" minWidth="0" asChild>
              <nav aria-label={t("header.breadcrumb")}>
                {crumbs.map((crumb, index) => (
                  <Flex key={crumb} align="center" gap="2" minWidth="0">
                    {index > 0 ? <AnimatedIcon icon={RightChevron} size={12} className={styles.divider} /> : null}
                    <Text size="2" weight={index === crumbs.length - 1 ? "medium" : "regular"} color={index === crumbs.length - 1 ? undefined : "gray"} truncate>
                      {crumb}
                    </Text>
                  </Flex>
                ))}
              </nav>
            </Flex>
            <PreferencesPopover />
          </Flex>
        </Section>
      </Container>
    </Box>
  );
}
