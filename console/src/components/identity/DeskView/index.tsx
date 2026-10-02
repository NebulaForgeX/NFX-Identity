import { LayoutDashboardIcon } from "nfx-ui/icons";
import type { ReactNode } from "react";
import { Badge, Button, Container, Flex, Grid, Section, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { formatDateTime, safeArray, safeNullable, safeStringable } from "@/utils";

import Masthead, { profileRoles } from "../Masthead";
import styles from "./s.module.css";

function Stat({ label, value }: { label: string; value: ReactNode }) {
  return (
    <Section size="1" py="4" className={styles.card}>
      <Container size="2" px="4">
        <Flex direction="column" gap="1">
          <Text size="1" color="gray">
            {label}
          </Text>
          {typeof value === "string" ? (
            <Text size="3" className={styles.value}>
              {value}
            </Text>
          ) : (
            value
          )}
        </Flex>
      </Container>
    </Section>
  );
}

export default function DeskView({ paths }: { paths: ScopePaths }) {
  const { t } = useTranslation("pages.Desk");
  const { data, profile, kind } = useCurrentProfile();
  const roles = profileRoles(kind, data);
  const emails = safeArray(data?.emails);
  const phones = safeArray(data?.phones);
  const identities = safeArray(data?.identities);

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
      <Grid columns={{ initial: "1", sm: "2" }} gap="4">
        <Stat label={t("session.accountId")} value={safeNullable(data?.account.id) || t("empty")} />
        <Stat label={t("session.profileId")} value={safeNullable(profile?.profileId) || t("empty")} />
        <Stat label={t("session.status")} value={safeStringable(data?.account.accountStatus) || t("empty")} />
        <Stat label={t("session.platform")} value={safeStringable(data?.account.signupPlatform) || t("empty")} />
        <Stat label={t("session.kind")} value={kind} />
        <Stat label={t("session.roles")} value={roles.length ? <Flex gap="2" wrap="wrap">{roles.map((role) => <Badge key={role} variant="outline" size="1">{role}</Badge>)}</Flex> : t("empty")} />
        <Stat label={t("session.loginNotification")} value={profile?.settings?.loginNotification ? t("on") : t("off")} />
        <Stat label={t("session.created")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : t("empty")} />
      </Grid>
    </PageFrame>
  );
}
