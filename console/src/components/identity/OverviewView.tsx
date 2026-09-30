import type { ReactNode } from "react";

import { UserIcon } from "nfx-ui/icons";
import { Badge, Box, Button, Flex, Grid, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { buildImageUrl, formatDateTime, safeArray, safeStringable } from "@/utils";

import { LedgerSection } from "./Ledger";
import Masthead, { profileRoles } from "./Masthead";
import styles from "./Overview/s.module.css";

function blank(value: Nilable<string>, fallback: string) {
  const next = safeStringable(value);
  return next || fallback;
}

function Fact({ label, value }: { label: string; value: ReactNode }) {
  return (
    <Flex direction="column" gap="1">
      <Text size="1" color="gray">
        {label}
      </Text>
      {typeof value === "string" ? <Text size="2">{value}</Text> : value}
    </Flex>
  );
}

function RecordList({
  rows,
}: {
  rows: { key: string; title: string; meta: string }[];
}) {
  return (
    <Flex direction="column">
      {rows.map((row) => (
        <Box key={row.key} className={styles.recordRule}>
          <Box py="3">
            <Flex align="center" justify="between" gap="3" wrap="wrap">
              <Text size="2">{row.title}</Text>
              <Text size="1" color="gray">
                {row.meta}
              </Text>
            </Flex>
          </Box>
        </Box>
      ))}
    </Flex>
  );
}

export default function OverviewView({ paths }: { paths: ScopePaths }) {
  const { t } = useTranslation("pages.Profile.Overview");
  const { data, profile, kind } = useCurrentProfile();
  const roles = profileRoles(kind, data);
  const emails = safeArray(data?.emails);
  const phones = safeArray(data?.phones);
  const identities = safeArray(data?.identities);
  const avatars = safeArray(profile?.avatars);
  const backgrounds = safeArray(profile?.backgrounds)
    .slice()
    .sort((a, b) => a.sortOrder - b.sortOrder || a.imageId.localeCompare(b.imageId));
  const empty = t("labels.notSpecified");

  return (
    <PageFrame>
      <PageHeader
        icon={UserIcon}
        title={t("title")}
        description={t("description")}
        actions={
          <Button size="2" onClick={() => routerEventEmitter.navigate({ to: paths.edit })}>
            {t("actions.edit")}
          </Button>
        }
      />
      <Masthead kind={kind} data={data} profile={profile} />

      <LedgerSection title={t("sections.profile")} description={t("sections.profileHint")}>
        <Flex direction="column" gap="5">
          <Grid columns={{ initial: "1", sm: "2" }} gap="4">
            <Fact label={t("labels.displayName")} value={blank(profile?.displayName, empty)} />
            <Fact label={t("labels.language")} value={blank(profile?.profileLanguage, empty)} />
            <Fact label={t("labels.firstName")} value={blank(profile?.firstName, empty)} />
            <Fact label={t("labels.lastName")} value={blank(profile?.lastName, empty)} />
            <Fact label={t("labels.gender")} value={blank(profile?.gender, empty)} />
            <Fact label={t("labels.birthday")} value={blank(profile?.birthday, empty)} />
          </Grid>
          <Grid columns={{ initial: "1", sm: "2" }} gap="4">
            <Fact label={t("labels.city")} value={blank(profile?.city, empty)} />
            <Fact label={t("labels.country")} value={blank(profile?.country, empty)} />
            <Fact label={t("labels.timezone")} value={blank(profile?.timezone, empty)} />
            <Fact label={t("labels.website")} value={blank(profile?.website, empty)} />
          </Grid>
          <Fact label={t("labels.bio")} value={blank(profile?.bio, empty)} />
          <Grid columns={{ initial: "1", sm: "2" }} gap="4">
            <Fact label={t("labels.roles")} value={roles.length ? (
              <Flex gap="2" wrap="wrap">
                {roles.map((role) => (
                  <Badge key={role} variant="outline" size="1">
                    {role}
                  </Badge>
                ))}
              </Flex>
            ) : empty} />
            <Fact label={t("labels.loginNotification")} value={profile?.settings?.loginNotification ? t("labels.on") : t("labels.off")} />
            <Fact label={t("labels.created")} value={profile?.createdAt ? formatDateTime(profile.createdAt) : empty} />
            <Fact label={t("labels.updated")} value={profile?.updatedAt ? formatDateTime(profile.updatedAt) : empty} />
            <Fact label={t("labels.profileId")} value={blank(profile?.profileId, empty)} />
          </Grid>
        </Flex>
      </LedgerSection>

      <LedgerSection title={t("sections.account")} description={t("sections.accountHint")}>
        <Grid columns={{ initial: "1", sm: "2" }} gap="4">
          <Fact label={t("labels.accountId")} value={blank(data?.account.id, empty)} />
          <Fact label={t("labels.accountStatus")} value={blank(data?.account.accountStatus, empty)} />
          <Fact label={t("labels.signupPlatform")} value={blank(data?.account.signupPlatform, empty)} />
          <Fact label={t("labels.accountCreated")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : empty} />
          <Fact label={t("labels.accountUpdated")} value={data?.account.updatedAt ? formatDateTime(data.account.updatedAt) : empty} />
        </Grid>
      </LedgerSection>

      <LedgerSection title={t("sections.emails")}>
        {emails.length ? (
          <RecordList
            rows={emails.map((item) => ({
              key: item.id,
              title: item.email,
              meta: [item.isPrimary ? t("labels.primary") : null, item.verifiedAt ? formatDateTime(item.verifiedAt) : t("labels.unverified")].filter(Boolean).join(" · "),
            }))}
          />
        ) : (
          <Text size="2" color="gray">{t("empty.emails")}</Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.phones")}>
        {phones.length ? (
          <RecordList
            rows={phones.map((item) => ({
              key: item.id,
              title: item.phone,
              meta: [item.isPrimary ? t("labels.primary") : null, item.verifiedAt ? formatDateTime(item.verifiedAt) : t("labels.unverified")].filter(Boolean).join(" · "),
            }))}
          />
        ) : (
          <Text size="2" color="gray">{t("empty.phones")}</Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.identities")}>
        {identities.length ? (
          <RecordList
            rows={identities.map((item) => ({
              key: `${item.identityProvider}:${item.providerSubject}`,
              title: item.identityProvider,
              meta: item.lastLoginAt ? formatDateTime(item.lastLoginAt) : empty,
            }))}
          />
        ) : (
          <Text size="2" color="gray">{t("empty.identities")}</Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.media")}>
        <Flex direction="column" gap="4">
          <Fact
            label={t("labels.avatars")}
            value={
              avatars.length ? (
                <Flex gap="2" wrap="wrap">
                  {avatars.map((item) => (
                    <Box key={item.id} className={styles.thumbSize}>
                      <Box className={styles.thumbClip}>
                        <img src={buildImageUrl(item.imageId)} alt="" className={styles.thumbImage} />
                      </Box>
                    </Box>
                  ))}
                </Flex>
              ) : empty
            }
          />
          <Fact
            label={t("labels.backgrounds")}
            value={
              backgrounds.length ? (
                <Flex gap="2" wrap="wrap">
                  {backgrounds.map((item) => (
                    <Box key={item.id} className={styles.coverThumbSize}>
                      <Box className={styles.thumbClip}>
                        <img src={buildImageUrl(item.imageId)} alt="" className={styles.thumbImage} />
                      </Box>
                    </Box>
                  ))}
                </Flex>
              ) : empty
            }
          />
        </Flex>
      </LedgerSection>
    </PageFrame>
  );
}
