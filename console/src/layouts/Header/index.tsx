import { useEffect, useRef, useState } from "react";
import { Avatar, Box, Button, Container, DropdownMenu, Flex, Section, Text } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { useCurrentProfile } from "nfx-ui/hooks";
import { setHeaderHeight, useAuthStore } from "nfx-ui/stores";
import { useTranslation } from "react-i18next";

import { Logo, PreferencesPopover } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { profileHome, ROUTES, scopePaths } from "@/navigations";
import { buildImageUrl, logoutSession, resolveAccountDisplayName, resolveAccountInitial, safeNullable } from "@/utils";

import styles from "./s.module.css";

function Header() {
  const headerRef = useRef<Nullable<HTMLElement>>(null);
  const { t } = useTranslation("language");
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { data: accountInfo, profile, kind } = useCurrentProfile();
  const [elevated, setElevated] = useState(false);
  const paths = scopePaths(kind);

  const accountId = safeNullable(accountInfo?.account.id);
  const displayName = resolveAccountDisplayName(profile?.displayName, accountId);
  const initial = resolveAccountInitial(profile?.displayName, accountId);
  const avatarImageId = safeNullable(profile?.avatars?.[0]?.imageId);

  useEffect(() => {
    if (!headerRef.current) return;
    const element = headerRef.current;
    const emitHeight = () => {
      const rect = element.getBoundingClientRect();
      const computed = getComputedStyle(element);
      const marginBottom = parseFloat(computed.marginBottom) || 0;
      setHeaderHeight(rect.bottom + marginBottom);
    };
    emitHeight();
    const observer = new ResizeObserver(emitHeight);
    observer.observe(element);
    window.addEventListener("resize", emitHeight);
    return () => {
      observer.disconnect();
      window.removeEventListener("resize", emitHeight);
    };
  }, []);

  useEffect(() => {
    const onScroll = () => setElevated(window.scrollY > 12);
    onScroll();
    window.addEventListener("scroll", onScroll, { passive: true });
    return () => window.removeEventListener("scroll", onScroll);
  }, []);

  return (
    <Box asChild className={styles.header}>
      <header ref={headerRef}>
        <Container size="4" width="100%" maxWidth="100%" px="4">
          <Section size="1" py="3">
            <Flex justify="center" width="100%">
              <Box width="100%" maxWidth="1120px" className={styles.bar} data-elevated={elevated ? "true" : "false"}>
                <Container size="4" width="100%" maxWidth="100%" px="3">
                  <Section size="1" py="2">
                    <Flex align="center" justify="between" gap="3">
                      <Logo variant="glassSquare" size="small" title={<Text className={styles.brandWord}>{APP_NAME}</Text>} subtitle="Identity" />

                      <Flex align="center" gap="2" flexShrink="0">
                        <PreferencesPopover />

                        {isAuthValid ? (
                          <DropdownMenu.Root modal={false}>
                            <DropdownMenu.Trigger>
                              <Button variant="outline" color="gray" highContrast>
                                <Avatar size="1" radius="full" src={avatarImageId ? buildImageUrl(avatarImageId) : undefined} fallback={initial} />
                                <Box className={styles.accountName}>
                                  <Text size="2" truncate>
                                    {displayName}
                                  </Text>
                                </Box>
                              </Button>
                            </DropdownMenu.Trigger>
                            <DropdownMenu.Content align="end" sideOffset={8} size="2" className={styles.accountMenu}>
                              <DropdownMenu.Label>
                                <Box className={styles.accountMenuName}>
                                  <Text size="1" color="gray" truncate>
                                    {displayName}
                                  </Text>
                                </Box>
                              </DropdownMenu.Label>
                              <DropdownMenu.Item onSelect={() => routerEventEmitter.navigate({ to: profileHome(kind) })}>
                                {t("header.panel")}
                              </DropdownMenu.Item>
                              <DropdownMenu.Item onSelect={() => routerEventEmitter.navigate({ to: paths.overview })}>
                                {t("header.profile")}
                              </DropdownMenu.Item>
                              <DropdownMenu.Item onSelect={() => routerEventEmitter.navigate({ to: paths.settings })}>
                                {t("sidebar.settingsItem")}
                              </DropdownMenu.Item>
                              <DropdownMenu.Separator />
                              <DropdownMenu.Item
                                onSelect={() => {
                                  void logoutSession().then(() => routerEventEmitter.navigate({ to: ROUTES.LOGIN }));
                                }}
                              >
                                {t("header.logout")}
                              </DropdownMenu.Item>
                            </DropdownMenu.Content>
                          </DropdownMenu.Root>
                        ) : (
                          <>
                            <Button variant="outline" color="gray" onClick={() => routerEventEmitter.navigate({ to: ROUTES.LOGIN })}>
                              {t("header.login")}
                            </Button>
                            <Button onClick={() => routerEventEmitter.navigate({ to: ROUTES.SIGNUP })}>{t("header.signup")}</Button>
                          </>
                        )}
                      </Flex>
                    </Flex>
                  </Section>
                </Container>
              </Box>
            </Flex>
          </Section>
        </Container>
      </header>
    </Box>
  );
}

export default Header;
