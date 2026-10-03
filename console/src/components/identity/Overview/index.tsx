import type { ReactNode } from "react";

import { CameraIcon, LinkIcon, MailFilledIcon, TelephoneIcon, UserIcon } from "nfx-ui/icons";
import { Badge, Box, Button, Flex, Grid, Section, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { EmptyState, PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { buildImageUrl, formatDateTime, safeArray, profileRoles, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "../Ledger";
import Masthead from "../Masthead";
import styles from "./s.module.css";

function blank(value: Nilable<string>, fallback: string) {
  const next = safeStringable(value);
  return next || fallback;
}

function Fact({ label, value }: { label: string; value: ReactNode }) {
  return (
    <Flex direction="column" gap="1" minWidth="0">
      <Text size="1" color="gray" weight="medium" className={styles.factLabel}>
        {label}
      </Text>
      {typeof value === "string" ? (
        <Text size="2" className={styles.factValue}>
          {value}
        </Text>
      ) : (
        value
      )}
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
        <Section key={row.key} size="1" py="3" className={styles.recordRule}>
          <Flex align="center" justify="between" gap="3" wrap="wrap">
            <Text size="2" weight="medium" truncate>
              {row.title}
            </Text>
            <Text size="1" color="gray">
              {row.meta}
            </Text>
          </Flex>
        </Section>
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
  const editAction = (
    <Button size="2" variant="outline" onClick={() => routerEventEmitter.navigate({ to: paths.edit })}>
      {t("actions.edit")}
    </Button>
  );

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
      <Masthead
        kind={kind}
        data={data}
        profile={profile}
        stats={[
          { label: t("sections.emails"), value: emails.length },
          { label: t("sections.phones"), value: phones.length },
          { label: t("sections.identities"), value: identities.length },
        ]}
      />

      <Grid columns={{ initial: "1", lg: "minmax(0, 1.5fr) minmax(0, 1fr)" }} gap="5" align="start">
        <Flex direction="column" gap="5" minWidth="0">
          <LedgerSection title={t("sections.profile")} description={t("sections.profileHint")}>
            <Flex direction="column" gap="5">
              <Grid columns={{ initial: "1", sm: "2", md: "3" }} gap="5">
                <Fact label={t("labels.displayName")} value={blank(profile?.displayName, empty)} />
                <Fact label={t("labels.firstName")} value={blank(profile?.firstName, empty)} />
                <Fact label={t("labels.lastName")} value={blank(profile?.lastName, empty)} />
                <Fact label={t("labels.gender")} value={blank(profile?.gender, empty)} />
                <Fact label={t("labels.birthday")} value={blank(profile?.birthday, empty)} />
                <Fact label={t("labels.language")} value={blank(profile?.profileLanguage, empty)} />
                <Fact label={t("labels.city")} value={blank(profile?.city, empty)} />
                <Fact label={t("labels.country")} value={blank(profile?.country, empty)} />
                <Fact label={t("labels.timezone")} value={blank(profile?.timezone, empty)} />
              </Grid>
              <Fact label={t("labels.website")} value={blank(profile?.website, empty)} />
              <Fact label={t("labels.bio")} value={blank(profile?.bio, empty)} />
              <Grid columns={{ initial: "1", sm: "2" }} gap="5">
                <Fact
                  label={t("labels.roles")}
                  value={
                    roles.length ? (
                      <Flex gap="2" wrap="wrap">
                        {roles.map((role) => (
                          <Badge key={role} variant="outline" color="gray" size="1">
                            {role}
                          </Badge>
                        ))}
                      </Flex>
                    ) : (
                      empty
                    )
                  }
                />
                <Fact
                  label={t("labels.loginNotification")}
                  value={
                    <Box>
                      <Badge size="1" variant="surface" color={profile?.settings?.loginNotification ? "green" : "gray"}>
                        {profile?.settings?.loginNotification ? t("labels.on") : t("labels.off")}
                      </Badge>
                    </Box>
                  }
                />
                <Fact label={t("labels.created")} value={profile?.createdAt ? formatDateTime(profile.createdAt) : empty} />
                <Fact label={t("labels.updated")} value={profile?.updatedAt ? formatDateTime(profile.updatedAt) : empty} />
              </Grid>
            </Flex>
          </LedgerSection>

          <LedgerSection title={t("sections.account")} description={t("sections.accountHint")}>
            <FieldList>
              <FieldRow label={t("labels.accountId")} value={blank(data?.account.id, empty)} />
              <FieldRow label={t("labels.profileId")} value={blank(profile?.profileId, empty)} />
              <FieldRow label={t("labels.accountStatus")} value={blank(data?.account.accountStatus, empty)} />
              <FieldRow label={t("labels.signupPlatform")} value={blank(data?.account.signupPlatform, empty)} />
              <FieldRow label={t("labels.accountCreated")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : empty} />
              <FieldRow label={t("labels.accountUpdated")} value={data?.account.updatedAt ? formatDateTime(data.account.updatedAt) : empty} />
            </FieldList>
          </LedgerSection>
        </Flex>

        <Flex direction="column" gap="5" minWidth="0">
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
              <EmptyState icon={MailFilledIcon} title={t("empty.emails")} action={editAction} />
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
              <EmptyState icon={TelephoneIcon} title={t("empty.phones")} action={editAction} />
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
              <EmptyState icon={LinkIcon} title={t("empty.identities")} action={editAction} />
            )}
          </LedgerSection>

          <LedgerSection title={t("sections.media")}>
            <Flex direction="column" gap="5">
              {avatars.length ? (
                <Fact
                  label={t("labels.avatars")}
                  value={
                    <Grid columns="4" gap="2">
                      {avatars.map((item) => (
                        <Box key={item.id} className={styles.thumb}>
                          <img src={buildImageUrl(item.imageId)} alt="" className={styles.thumbImage} />
                        </Box>
                      ))}
                    </Grid>
                  }
                />
              ) : (
                <EmptyState icon={CameraIcon} title={t("labels.avatars")} description={empty} action={editAction} />
              )}
              {backgrounds.length ? (
                <Fact
                  label={t("labels.backgrounds")}
                  value={
                    <Grid columns="2" gap="2">
                      {backgrounds.map((item) => (
                        <Box key={item.id} className={styles.coverThumb}>
                          <img src={buildImageUrl(item.imageId)} alt="" className={styles.thumbImage} />
                        </Box>
                      ))}
                    </Grid>
                  }
                />
              ) : (
                <EmptyState icon={CameraIcon} title={t("labels.backgrounds")} description={empty} action={editAction} />
              )}
            </Flex>
          </LedgerSection>
        </Flex>
      </Grid>
    </PageFrame>
  );
}
