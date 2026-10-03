import type { AnimatedIconComponent } from "nfx-ui/icons";

import { useState } from "react";
import { Badge, Box, Button, Flex, Grid, Heading, IconButton, Text, Tooltip } from "@radix-ui/themes";
import { AnimatedIcon, ArrowNarrowRightIcon, CameraIcon, CheckedIcon, CopyIcon, FilledBellIcon, GearIcon, LayoutDashboardIcon, PassportIcon, PenIcon, ShieldCheck } from "nfx-ui/icons";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import { Link } from "react-router";

import { PageHeader, Surface } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { formatDateTime, safeArray, safeNullable, profileRoles, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "../Ledger";
import Masthead from "../Masthead";
import styles from "./s.module.css";

function CopyValue({ value, copyLabel, copiedLabel }: { value: string; copyLabel: string; copiedLabel: string }) {
  const [copied, setCopied] = useState(false);
  return (
    <Flex align="center" gap="2" minWidth="0">
      <Text size="2" className={styles.mono} truncate>
        {value}
      </Text>
      <Tooltip content={copied ? copiedLabel : copyLabel}>
        <IconButton
          size="1"
          variant="ghost"
          color="gray"
          aria-label={copyLabel}
          onClick={() =>
            void navigator.clipboard.writeText(value).then(() => {
              setCopied(true);
              window.setTimeout(() => setCopied(false), 1400);
            })
          }
        >
          <AnimatedIcon icon={copied ? CheckedIcon : CopyIcon} size={14} />
        </IconButton>
      </Tooltip>
    </Flex>
  );
}

function QuickTile({ to, icon, title, hint }: { to: string; icon: AnimatedIconComponent; title: string; hint: string }) {
  return (
    <Surface interactive py="4" px="4" className={styles.tile}>
      <Flex align="center" gap="3">
        <Box className={styles.tileMark}>
          <Flex align="center" justify="center" width="100%" height="100%">
            <AnimatedIcon icon={icon} size={18} />
          </Flex>
        </Box>
        <Flex direction="column" gap="1" minWidth="0" flexGrow="1">
          <Link to={to} className={styles.tileLink}>
            <Text size="2" weight="bold">
              {title}
            </Text>
          </Link>
          <Text size="1" color="gray" truncate>
            {hint}
          </Text>
        </Flex>
        <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} className={styles.tileArrow} />
      </Flex>
    </Surface>
  );
}

export default function DeskView({ paths }: { paths: ScopePaths }) {
  const { t } = useTranslation("pages.Desk");
  const { data, profile, kind } = useCurrentProfile();
  const roles = profileRoles(kind, data);
  const emails = safeArray(data?.emails);
  const phones = safeArray(data?.phones);
  const identities = safeArray(data?.identities);
  const accountId = safeNullable(data?.account.id);
  const profileId = safeNullable(profile?.profileId);
  const status = safeStringable(data?.account.accountStatus);
  const loginNotification = Boolean(profile?.settings?.loginNotification);

  return (
    <PageFrame>
      <PageHeader
        icon={LayoutDashboardIcon}
        title={t("title")}
        description={t("description")}
        actions={
          <Button size="2" onClick={() => routerEventEmitter.navigate({ to: paths.overview })}>
            {t("openRegistry")}
          </Button>
        }
      />
      <Masthead
        kind={kind}
        data={data}
        profile={profile}
        stats={[
          { label: t("session.emails"), value: emails.length },
          { label: t("session.phones"), value: phones.length },
          { label: t("session.links"), value: identities.length },
        ]}
      />
      <Grid columns={{ initial: "1", lg: "minmax(0, 1.55fr) minmax(0, 1fr)" }} gap="5" align="start">
        <Flex direction="column" gap="5" minWidth="0">
          <LedgerSection title={t("session.title")} description={t("session.description")}>
            <FieldList>
              <FieldRow label={t("session.accountId")}>
                {accountId ? <CopyValue value={accountId} copyLabel={t("actions.copy")} copiedLabel={t("actions.copied")} /> : <Text size="2">{t("empty")}</Text>}
              </FieldRow>
              <FieldRow label={t("session.profileId")}>
                {profileId ? <CopyValue value={profileId} copyLabel={t("actions.copy")} copiedLabel={t("actions.copied")} /> : <Text size="2">{t("empty")}</Text>}
              </FieldRow>
              <FieldRow label={t("session.status")}>
                {status ? (
                  <Badge size="2" variant="surface" color="green">
                    {status}
                  </Badge>
                ) : (
                  <Text size="2">{t("empty")}</Text>
                )}
              </FieldRow>
              <FieldRow label={t("session.platform")} value={safeStringable(data?.account.signupPlatform) || t("empty")} />
              <FieldRow label={t("session.created")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : t("empty")} />
            </FieldList>
          </LedgerSection>
          <Flex direction="column" gap="3">
            <Heading as="h2" size="3" weight="bold" data-reveal="">
              {t("quick.title")}
            </Heading>
            <Grid columns={{ initial: "1", sm: "2" }} gap="3">
              <QuickTile to={paths.edit} icon={PenIcon} title={t("quick.edit.title")} hint={t("quick.edit.hint")} />
              <QuickTile to={paths.identity} icon={PassportIcon} title={t("quick.identity.title")} hint={t("quick.identity.hint")} />
              <QuickTile to={paths.assets} icon={CameraIcon} title={t("quick.assets.title")} hint={t("quick.assets.hint")} />
              <QuickTile to={paths.settings} icon={GearIcon} title={t("quick.settings.title")} hint={t("quick.settings.hint")} />
            </Grid>
          </Flex>
        </Flex>
        <Flex direction="column" gap="5" minWidth="0">
          <LedgerSection title={t("access.title")} description={t("access.description")}>
            <Flex direction="column" gap="3">
              <Flex align="center" justify="between" gap="3">
                <Text size="2" color="gray">
                  {t("session.kind")}
                </Text>
                <Badge size="2" variant="surface">
                  {kind}
                </Badge>
              </Flex>
              <Flex direction="column" gap="2">
                <Text size="2" color="gray">
                  {t("session.roles")}
                </Text>
                {roles.length ? (
                  <Flex gap="2" wrap="wrap">
                    {roles.map((role) => (
                      <Badge key={role} size="2" variant="outline" color="gray">
                        {role}
                      </Badge>
                    ))}
                  </Flex>
                ) : (
                  <Text size="2">{t("empty")}</Text>
                )}
              </Flex>
            </Flex>
          </LedgerSection>
          <LedgerSection title={t("security.title")} description={t("security.description")}>
            <Flex direction="column" gap="4">
              <Surface tone="inset" py="3" px="4">
                <Flex align="center" justify="between" gap="3">
                  <Flex align="center" gap="3" minWidth="0">
                    <AnimatedIcon icon={loginNotification ? FilledBellIcon : ShieldCheck} size={18} className={styles.securityIcon} />
                    <Text size="2">{t("session.loginNotification")}</Text>
                  </Flex>
                  <Badge size="1" variant="surface" color={loginNotification ? "green" : "gray"}>
                    {loginNotification ? t("on") : t("off")}
                  </Badge>
                </Flex>
              </Surface>
              <Flex gap="2" wrap="wrap">
                <Button size="2" variant="outline" onClick={() => routerEventEmitter.navigate({ to: paths.security })}>
                  {t("security.password")}
                </Button>
                <Button size="2" variant="ghost" onClick={() => routerEventEmitter.navigate({ to: paths.settings })}>
                  {t("security.manage")}
                </Button>
              </Flex>
            </Flex>
          </LedgerSection>
        </Flex>
      </Grid>
    </PageFrame>
  );
}
