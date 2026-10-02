import type { ReactNode } from "react";

import { useEffect, useRef, useState } from "react";
import { Avatar, Box, Button, Container, Flex, IconButton, Section, Text } from "@radix-ui/themes";
import { AnimatedIcon, ArrowNarrowLeftIcon, ArrowNarrowUpIcon, CameraIcon, DownChevron, GearIcon, LayoutDashboardIcon, LogoutIcon, PassportIcon, PenIcon, RightChevron, ShieldCheck, UnorderedListIcon, UserIcon } from "nfx-ui/icons";
import { ProfileKindEnum } from "nfx-ui/enums";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import { Menu, Sidebar as ProSidebar } from "react-pro-sidebar";
import { Link, Outlet, useLocation } from "react-router";

import UserTopBar from "@/layouts/UserTopBar";
import { scopePaths } from "@/navigations";
import { buildImageUrl, logoutSession, resolveAccountDisplayName, safeNullable } from "@/utils";

import { MenuItem, SidebarMenuState, SubMenu } from "./Menu";
import styles from "./s.module.css";

const SIDEBAR_WIDTH = "234px";
const SIDEBAR_COLLAPSED_WIDTH = "84px";

function MenuLabel({ children, active = false }: { children: ReactNode; active?: boolean }) {
  return (
    <Text as="span" size="2" weight={active ? "bold" : "medium"}>
      {children}
    </Text>
  );
}

function MenuTrunk({ collapsed, children }: { collapsed: boolean; children: ReactNode }) {
  if (collapsed) return children;
  return <Box className={styles.trunk}>{children}</Box>;
}

function SectionTitle({ label, collapsed }: { label: string; collapsed: boolean }) {
  return (
    <Section size="1" py="3" className={styles.sectionRule}>
      <Flex align="center" justify={collapsed ? "center" : "start"} gap="2">
        <Text as="span" size="1" weight="bold" className={styles.sectionTitle}>
          {label}
        </Text>
      </Flex>
    </Section>
  );
}

function createMenuItemStyles(collapsed: boolean) {
  return {
    button: ({ active, level = 0 }: { active: boolean; level?: number }) => ({
      height: level > 0 ? "34px" : "40px",
      margin: level > 0 ? (collapsed ? "var(--space-1) var(--space-2)" : "var(--space-1) 0 var(--space-1) var(--space-6)") : "var(--space-2) 0",
      borderRadius: "var(--radius-chip)",
      paddingLeft: "var(--space-3)",
      paddingRight: "var(--space-3)",
      fontSize: level > 0 ? "14px" : "15px",
      fontWeight: 400,
      color: active ? "var(--accent-11)" : "var(--gray-11)",
      backgroundColor: active ? "var(--accent-a3)" : "transparent",
      transition: "background-color 150ms ease, color 150ms ease",
      "&:hover": { backgroundColor: active ? "var(--accent-a3)" : "var(--gray-a3)", color: active ? "var(--accent-11)" : "var(--gray-12)" },
      "&:focus-visible": {
        outline: "2px solid var(--accent-8)",
        outlineOffset: "2px",
      },
    }),
    icon: ({ level = 0 }: { level?: number }) => ({
      width: "var(--space-5)",
      minWidth: "var(--space-5)",
      height: "var(--space-5)",
      marginRight: "var(--space-2)",
      color: "inherit",
      ...(level > 0 ? { display: "none" } : {}),
    }),
    subMenuContent: {
      backgroundColor: collapsed ? "var(--color-panel-solid)" : "transparent",
      padding: collapsed ? "var(--space-1) 0" : "0",
      ...(collapsed
        ? {
            zIndex: 1000,
            minWidth: "156px",
            width: "max-content",
            maxWidth: "260px",
            maxHeight: "calc(100dvh - 32px)",
            overflowY: "auto" as const,
            borderRadius: "var(--radius-4)",
            border: "1px solid var(--gray-a5)",
            boxShadow: "var(--shadow-5)",
          }
        : {}),
    },
  };
}

interface SectionProps {
  collapsed: boolean;
  broken: boolean;
  onMobileClose: () => void;
}

function OverviewSection({ collapsed, broken, onMobileClose, paths }: SectionProps & { paths: ReturnType<typeof scopePaths> }) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const active = location.pathname === paths.desk || location.pathname.startsWith(`${paths.desk}/`);

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} menuItemStyles={createMenuItemStyles(collapsed)} closeOnClick>
      <MenuItem component={<Link to={paths.desk} />} icon={<AnimatedIcon icon={LayoutDashboardIcon} size={18} />} active={active} onClick={() => broken && onMobileClose()}>
        <MenuLabel active={active}>{t("sidebar.desk")}</MenuLabel>
      </MenuItem>
    </Menu>
  );
}

function MainMenuSection({ collapsed, broken, onMobileClose, paths }: SectionProps & { paths: ReturnType<typeof scopePaths> }) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const [profileOpen, setProfileOpen] = useState(() => location.pathname.startsWith(paths.profile));

  const isActive = (to: string) => location.pathname === to || location.pathname.startsWith(`${to}/`);

  const profileSubItems = [
    { key: "profileOverview", to: paths.overview, icon: <AnimatedIcon icon={UserIcon} size={16} />, label: t("sidebar.profileOverview") },
    { key: "profileEdit", to: paths.edit, icon: <AnimatedIcon icon={PenIcon} size={16} />, label: t("sidebar.profileEdit") },
    { key: "profileIdentities", to: paths.identity, icon: <AnimatedIcon icon={PassportIcon} size={16} />, label: t("sidebar.profileIdentities") },
    { key: "profileSecurity", to: paths.security, icon: <AnimatedIcon icon={ShieldCheck} size={16} />, label: t("sidebar.profileSecurity") },
  ];

  const isProfileChildActive = profileSubItems.some((item) => isActive(item.to));

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} menuItemStyles={createMenuItemStyles(collapsed)} closeOnClick>
      <SectionTitle label={t("sidebar.mainMenu")} collapsed={collapsed} />
      <SubMenu label={t("sidebar.profile")} icon={<AnimatedIcon icon={UserIcon} size={18} />} open={profileOpen} onOpenChange={setProfileOpen} active={isProfileChildActive}>
        {profileSubItems.map((item) => (
          <MenuItem key={item.key} component={<Link to={item.to} />} icon={item.icon} active={isActive(item.to)} onClick={() => broken && onMobileClose()}>
            <MenuLabel active={isActive(item.to)}>{item.label}</MenuLabel>
          </MenuItem>
        ))}
      </SubMenu>
      <MenuItem component={<Link to={paths.assets} />} icon={<AnimatedIcon icon={CameraIcon} size={18} />} active={isActive(paths.assets)} onClick={() => broken && onMobileClose()}>
        <MenuLabel active={isActive(paths.assets)}>{t("sidebar.assets")}</MenuLabel>
      </MenuItem>
      {paths.directory ? (
        <MenuItem component={<Link to={paths.directory} />} icon={<AnimatedIcon icon={ShieldCheck} size={18} />} active={isActive(paths.directory)} onClick={() => broken && onMobileClose()}>
          <MenuLabel active={isActive(paths.directory)}>{t("sidebar.directory")}</MenuLabel>
        </MenuItem>
      ) : null}
    </Menu>
  );
}

function SettingsSection({ collapsed, broken, onMobileClose, paths }: SectionProps & { paths: ReturnType<typeof scopePaths> }) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const isActive = (to: string) => location.pathname === to || location.pathname.startsWith(`${to}/`);

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} menuItemStyles={createMenuItemStyles(collapsed)} closeOnClick>
      <SectionTitle label={t("sidebar.settings")} collapsed={collapsed} />
      <MenuItem
        component={<Link to={paths.settings} />}
        icon={<AnimatedIcon icon={GearIcon} size={18} />}
        active={isActive(paths.settings)}
        onClick={() => broken && onMobileClose()}
      >
        <MenuLabel active={isActive(paths.settings)}>{t("sidebar.settingsItem")}</MenuLabel>
      </MenuItem>
    </Menu>
  );
}

function Sidebar() {
  const { t } = useTranslation("language");
  const [desktopCollapsed, setCollapsed] = useState(false);
  const [toggled, setToggled] = useState(false);
  const [broken, setBroken] = useState(false);
  const collapsed = !broken && desktopCollapsed;
  const drawerRef = useRef<HTMLHtmlElement>(null);
  useEffect(() => {
    if (!broken || !toggled) return;
    const previousFocus = document.activeElement as HTMLElement | null;
    const previousOverflow = document.body.style.overflow;
    document.body.style.overflow = "hidden";
    const focusable = () =>
      Array.from(drawerRef.current?.querySelectorAll<HTMLElement>("button:not([disabled]), a[href], [tabindex='0']") ?? []).filter(
        (node) => node.getClientRects().length && getComputedStyle(node).visibility !== "hidden",
      );
    focusable()[0]?.focus();
    const onKeyDown = (event: KeyboardEvent) => {
      if (event.defaultPrevented) return;
      if (event.key === "Escape") {
        event.preventDefault();
        setToggled(false);
      }
      if (event.key !== "Tab") return;
      const nodes = focusable();
      const first = nodes[0];
      const last = nodes[nodes.length - 1];
      if (event.shiftKey && document.activeElement === first) {
        event.preventDefault();
        last?.focus();
      } else if (!event.shiftKey && document.activeElement === last) {
        event.preventDefault();
        first?.focus();
      }
    };
    document.addEventListener("keydown", onKeyDown);
    return () => {
      document.body.style.overflow = previousOverflow;
      document.removeEventListener("keydown", onKeyDown);
      previousFocus?.focus();
    };
  }, [broken, toggled]);

  const { kind, data, profile } = useCurrentProfile();
  const paths = scopePaths(kind);

  const accountId = safeNullable(data?.account.id);
  const displayName = resolveAccountDisplayName(profile?.displayName, accountId);
  const avatarImageId = safeNullable(profile?.avatars?.[0]?.imageId);

  const closeMobile = () => broken && setToggled(false);

  const handleLogout = () => {
    void logoutSession();
  };

  return (
    <Box className={styles.shell}>
      <Flex minHeight="100dvh" width="100%">
        <SidebarMenuState collapsed={collapsed}>
          <ProSidebar
            ref={drawerRef}
            inert={broken && !toggled ? true : undefined}
            aria-hidden={broken && !toggled ? true : undefined}
            collapsed={collapsed}
            toggled={toggled}
            onBackdropClick={() => setToggled(false)}
            onBreakPoint={setBroken}
            breakPoint="md"
            width={SIDEBAR_WIDTH}
            collapsedWidth={SIDEBAR_COLLAPSED_WIDTH}
            rootStyles={{
              border: 0,
              height: "100dvh",
              flexShrink: 0,
              position: "sticky",
              top: 0,
              margin: 0,
              zIndex: 200,
              overflow: "visible",
              "&.ps-broken.ps-toggled": { left: 0 },
              "&.ps-broken": { height: "100dvh", position: "fixed", margin: 0, top: 0, bottom: 0 },
              "& .ps-sidebar-container": {
                background: "transparent",
                height: "100%",
                overflow: "visible",
              },
            }}
          >
            <Box className={styles.panel}>
              <Flex direction="column" height="100%" minHeight="0">
                <Section size="1" py="4" className={styles.accountBand}>
                  <Container size="4" px="4" width="100%" maxWidth="100%">
                    <Flex align="center" justify="between" direction={collapsed ? "column" : "row"} gap="3">
                      <Button type="button" variant="ghost" className={styles.accountReset} aria-label={displayName}>
                        <Flex align="center" width="100%" gap="3">
                          <Avatar
                            size="3"
                            className={styles.avatar}
                            src={avatarImageId ? buildImageUrl(avatarImageId) : undefined}
                            fallback={<UserIcon size={20} />}
                            alt=""
                            aria-hidden="true"
                          />
                          {!collapsed && (
                            <Flex direction="column" flexGrow="1" minWidth="0" gap="1">
                              <Text as="span" size="1" className={styles.accountRole}>
                                {t(kind === ProfileKindEnum.AUTHORITY ? "sidebar.profileAuthority" : "sidebar.profileCommunity")}
                              </Text>
                              <Text as="span" size="2" className={styles.accountName}>
                                {displayName}
                              </Text>
                            </Flex>
                          )}
                        </Flex>
                      </Button>
                      <IconButton
                        variant="ghost"
                        size="1"
                        className={styles.toggle}
                        aria-label={collapsed ? t("sidebar.expand") : t("sidebar.collapse")}
                        aria-expanded={!collapsed}
                        onClick={() => (broken ? setToggled(false) : setCollapsed((value) => !value))}
                      >
                        <AnimatedIcon icon={collapsed ? RightChevron : ArrowNarrowLeftIcon} size={14} />
                      </IconButton>
                    </Flex>
                  </Container>
                </Section>

                <Flex direction="column" flexGrow="1" minHeight="0" overflowX="hidden" overflowY="auto" className={`${styles.menu} ${collapsed ? styles.menuCollapsed : styles.menuOpen}`}>
                  <Section size="1" py="3">
                    <Container size="4" px="3" width="100%" maxWidth="100%">
                      <MenuTrunk collapsed={collapsed}>
                        <OverviewSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                        <MainMenuSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                        <SettingsSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                      </MenuTrunk>
                    </Container>
                  </Section>
                </Flex>

                <Section size="1" py="3" mt="auto">
                  <Container size="4" px="3" width="100%" maxWidth="100%">
                    <Button
                      type="button"
                      variant="ghost"
                      className={styles.logout}
                      onClick={handleLogout}
                      aria-label={t("sidebar.logout")}
                      title={collapsed ? t("sidebar.logout") : undefined}
                    >
                      <Flex align="center" justify={collapsed ? "center" : "start"} gap="3" width="100%">
                        <AnimatedIcon icon={LogoutIcon} size={18} />
                        {!collapsed ? t("sidebar.logout") : null}
                      </Flex>
                    </Button>
                  </Container>
                </Section>
              </Flex>
            </Box>
          </ProSidebar>
        </SidebarMenuState>

        <Box minWidth="0" minHeight="0" height="100%" overflowX="hidden" position="relative" className={styles.content} >
          <Flex direction="column" width="100%" height="100%" minHeight="0" inert={broken && toggled ? true : undefined}>
            {broken ? (
              <Section size="1" my="3" pt="0" pb="0" position="sticky" top="3" className={styles.mobileStick}>
                <Container size="4" mx="3" px="0" width="100%" maxWidth="100%">
                  <Box className={styles.mobile}>
                    <IconButton variant="ghost" color="gray" size="3" className={styles.mobileControl} onClick={() => setToggled(true)} aria-label={t("sidebar.openMenu")}>
                      <AnimatedIcon icon={UnorderedListIcon} size={18} />
                    </IconButton>
                  </Box>
                </Container>
              </Section>
            ) : null}
            <UserTopBar />
            <Box minWidth="0" minHeight="0" height="100%" position="relative" width="100%" >
              <Flex direction="column" width="100%" height="100%">
                <Outlet />
              </Flex>
            </Box>
          </Flex>
        </Box>
      </Flex>
    </Box>
  );
}

export default Sidebar;
