import type { ReactNode } from "react";

import { useEffect, useRef, useState } from "react";
import { Avatar, Box, Button, Container, Flex, IconButton, Section, Text } from "@radix-ui/themes";
import { AnimatedIcon, type AnimatedIconComponent, ArrowNarrowLeftIcon, ArrowNarrowUpIcon, CameraIcon, DownChevron, GearIcon, LayersIcon, LayoutDashboardIcon, LogoutIcon, PassportIcon, PenIcon, RightChevron, ShieldCheck, UnorderedListIcon, UserIcon } from "nfx-ui/icons";
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
const SIDEBAR_COLLAPSED_WIDTH = "88px";

function MenuLabel({ children, active = false }: { children: ReactNode; active?: boolean }) {
  return (
    <Text as="span" size="2" weight={active ? "bold" : "medium"}>
      {children}
    </Text>
  );
}

function SectionTitle({ label, icon, collapsed }: { label: string; icon: AnimatedIconComponent; collapsed: boolean }) {
  if (collapsed) {
    return <Section size="1" mt="3" pt="3" pb="0" className={styles.sectionRule} />;
  }
  return (
    <Section size="1" mt="4" pt="5" pb="0" className={styles.sectionRule}>
      <Flex align="center" justify="between" gap="2">
        <Text as="span" size="2" weight="bold" className={styles.sectionLabel}>
          {label}
        </Text>
        <AnimatedIcon icon={icon} size={16} className={styles.sectionTitleIcon} />
      </Flex>
    </Section>
  );
}

interface SectionProps {
  collapsed: boolean;
  broken: boolean;
  onMobileClose: () => void;
  paths: ReturnType<typeof scopePaths>;
}

function OverviewSection({ broken, onMobileClose, paths }: SectionProps) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const active = location.pathname === paths.desk || location.pathname.startsWith(`${paths.desk}/`);

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} closeOnClick>
      <MenuItem component={<Link to={paths.desk} />} icon={<AnimatedIcon icon={LayoutDashboardIcon} size={18} />} active={active} onClick={() => broken && onMobileClose()}>
        <MenuLabel active={active}>{t("sidebar.desk")}</MenuLabel>
      </MenuItem>
    </Menu>
  );
}

function MainMenuSection({ collapsed, broken, onMobileClose, paths }: SectionProps) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const [profileOpen, setProfileOpen] = useState(() => location.pathname.startsWith(paths.profile));

  const isActive = (to: string) => location.pathname === to || location.pathname.startsWith(`${to}/`);

  const profileSubItems = [
    {
      key: "profileOverview",
      to: paths.overview,
      icon: <AnimatedIcon icon={UserIcon} size={16} />,
      label: t("sidebar.profileOverview"),
    },
    {
      key: "profileEdit",
      to: paths.edit,
      icon: <AnimatedIcon icon={PenIcon} size={16} />,
      label: t("sidebar.profileEdit"),
    },
    {
      key: "profileIdentities",
      to: paths.identity,
      icon: <AnimatedIcon icon={PassportIcon} size={16} />,
      label: t("sidebar.profileIdentities"),
    },
    {
      key: "profileSecurity",
      to: paths.security,
      icon: <AnimatedIcon icon={ShieldCheck} size={16} />,
      label: t("sidebar.profileSecurity"),
    },
  ];

  const isProfileChildActive = profileSubItems.some((item) => isActive(item.to));

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} closeOnClick>
      <SectionTitle label={t("sidebar.mainMenu")} icon={LayersIcon} collapsed={collapsed} />
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

function SettingsSection({ collapsed, broken, onMobileClose, paths }: SectionProps) {
  const { t } = useTranslation("language");
  const location = useLocation();
  const isActive = (to: string) => location.pathname === to || location.pathname.startsWith(`${to}/`);

  return (
    <Menu renderExpandIcon={({ open }) => <AnimatedIcon icon={open ? ArrowNarrowUpIcon : DownChevron} size={14} />} closeOnClick>
      <SectionTitle label={t("sidebar.settings")} icon={GearIcon} collapsed={collapsed} />
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
    <Flex className={styles.shell} minHeight="100dvh" width="100%">
      <SidebarMenuState collapsed={collapsed}>
        <Box className={styles.proSidebar}>
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
          >
            <Flex direction="column" className={styles.sidebar} height="100%" minHeight="0">
              <Box>
                <Container size="4" width="100%" maxWidth="100%" px="5">
                  <Box position="relative">
                    <Section size="1" pt="21px" pb="22px" className={styles.headerRule}>
                      <Button type="button" variant="ghost" className={styles.accountCard} aria-label={displayName}>
                        <Flex align="center" width="100%" gap="3">
                          <Box>
                            <Avatar
                              size="3"
                              src={avatarImageId ? buildImageUrl(avatarImageId) : undefined}
                              fallback={<UserIcon size={20} />}
                              alt=""
                              aria-hidden="true"
                            />
                          </Box>
                          {!collapsed && (
                            <Flex direction="column" flexGrow="1" minWidth="0" gap="1">
                              <span className={styles.accountRole}>{t(kind === ProfileKindEnum.AUTHORITY ? "sidebar.profileAuthority" : "sidebar.profileCommunity")}</span>
                              <span className={styles.accountName}>{displayName}</span>
                            </Flex>
                          )}
                        </Flex>
                      </Button>
                    </Section>
                    <IconButton
                      variant="outline"
                      size="1"
                      className={styles.toggle}
                      aria-label={collapsed ? t("sidebar.expand") : t("sidebar.collapse")}
                      aria-expanded={!collapsed}
                      onClick={() => (broken ? setToggled(false) : setCollapsed((value) => !value))}
                    >
                      <AnimatedIcon icon={collapsed ? RightChevron : ArrowNarrowLeftIcon} size={14} />
                    </IconButton>
                  </Box>
                </Container>
              </Box>

              <Box minHeight="0" className={styles.menuArea} data-collapsed={collapsed ? "true" : "false"}>
                <Section size="1" py="2">
                  <Container size="4" width="100%" maxWidth="100%" px="5">
                    <OverviewSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                    <MainMenuSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                    <SettingsSection collapsed={collapsed} broken={broken} onMobileClose={closeMobile} paths={paths} />
                  </Container>
                </Section>
              </Box>

              <Section size="1" pt="3" pb="5" className={styles.footer} data-collapsed={collapsed ? "true" : "false"}>
                <Container size="4" width="100%" maxWidth="100%" px="5">
                  <Button
                    variant="ghost"
                    className={styles.logout}
                    data-collapsed={collapsed ? "true" : "false"}
                    onClick={handleLogout}
                    aria-label={t("sidebar.logout")}
                    title={collapsed ? t("sidebar.logout") : undefined}
                  >
                    <Flex align="center" justify={collapsed ? "center" : "start"} gap="3" width="100%">
                      <AnimatedIcon icon={LogoutIcon} size={18} />
                      {!collapsed && t("sidebar.logout")}
                    </Flex>
                  </Button>
                </Container>
              </Section>
            </Flex>
          </ProSidebar>
        </Box>
      </SidebarMenuState>

      <Flex className={styles.content} direction="column" flexGrow="1" minWidth="0" minHeight="0" height="100%" position="relative" inert={broken && toggled ? true : undefined}>
        {broken ? (
          <IconButton variant="surface" color="gray" size="3" className={styles.mobileToggle} onClick={() => setToggled(true)} aria-label={t("sidebar.openMenu")}>
            <AnimatedIcon icon={UnorderedListIcon} size={18} />
          </IconButton>
        ) : null}
        <UserTopBar />
        <Box minWidth="0" minHeight="0" position="relative" width="100%" height="100%">
          <Outlet />
        </Box>
      </Flex>
    </Flex>
  );
}

export default Sidebar;
