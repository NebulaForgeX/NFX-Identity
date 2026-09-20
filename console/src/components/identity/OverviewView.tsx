import { UserIcon } from "nfx-ui/icons";
import { Badge, Button, Flex, Table, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { buildImageUrl, formatDateTime, safeArray, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "./Ledger";
import Masthead, { profileRoles } from "./Masthead";

function blank(value: Nilable<string>, fallback: string) {
  const next = safeStringable(value);
  return next || fallback;
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

      <LedgerSection title={t("sections.account")} description={t("sections.accountHint")}>
        <FieldList>
          <FieldRow label={t("labels.accountId")} value={blank(data?.account.id, empty)} />
          <FieldRow label={t("labels.accountStatus")} value={blank(data?.account.accountStatus, empty)} />
          <FieldRow label={t("labels.signupPlatform")} value={blank(data?.account.signupPlatform, empty)} />
          <FieldRow label={t("labels.accountCreated")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : empty} />
          <FieldRow label={t("labels.accountUpdated")} value={data?.account.updatedAt ? formatDateTime(data.account.updatedAt) : empty} />
        </FieldList>
      </LedgerSection>

      <LedgerSection title={t("sections.profile")} description={t("sections.profileHint")}>
        <FieldList>
          <FieldRow label={t("labels.profileId")} value={blank(profile?.profileId, empty)} />
          <FieldRow label={t("labels.displayName")} value={blank(profile?.displayName, empty)} />
          <FieldRow label={t("labels.firstName")} value={blank(profile?.firstName, empty)} />
          <FieldRow label={t("labels.lastName")} value={blank(profile?.lastName, empty)} />
          <FieldRow label={t("labels.gender")} value={blank(profile?.gender, empty)} />
          <FieldRow label={t("labels.birthday")} value={blank(profile?.birthday, empty)} />
          <FieldRow label={t("labels.city")} value={blank(profile?.city, empty)} />
          <FieldRow label={t("labels.country")} value={blank(profile?.country, empty)} />
          <FieldRow label={t("labels.website")} value={blank(profile?.website, empty)} />
          <FieldRow label={t("labels.timezone")} value={blank(profile?.timezone, empty)} />
          <FieldRow label={t("labels.language")} value={blank(profile?.profileLanguage, empty)} />
          <FieldRow label={t("labels.bio")} value={blank(profile?.bio, empty)} />
          <FieldRow label={t("labels.created")} value={profile?.createdAt ? formatDateTime(profile.createdAt) : empty} />
          <FieldRow label={t("labels.updated")} value={profile?.updatedAt ? formatDateTime(profile.updatedAt) : empty} />
          <FieldRow label={t("labels.roles")}>
            <Flex gap="2" wrap="wrap">
              {roles.length ? roles.map((role) => <Badge key={role} variant="outline" size="1">{role}</Badge>) : <Text size="2">{empty}</Text>}
            </Flex>
          </FieldRow>
          <FieldRow label={t("labels.loginNotification")} value={profile?.settings?.loginNotification ? t("labels.on") : t("labels.off")} />
        </FieldList>
      </LedgerSection>

      <LedgerSection title={t("sections.emails")}>
        {emails.length ? (
          <Table.Root variant="surface">
            <Table.Header>
              <Table.Row>
                <Table.ColumnHeaderCell>{t("labels.email")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.primary")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.verified")}</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {emails.map((item) => (
                <Table.Row key={item.id}>
                  <Table.Cell>{item.email}</Table.Cell>
                  <Table.Cell>{item.isPrimary ? t("labels.yes") : t("labels.no")}</Table.Cell>
                  <Table.Cell>{item.verifiedAt ? formatDateTime(item.verifiedAt) : t("labels.unverified")}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table.Root>
        ) : (
          <Text size="2" color="gray">
            {t("empty.emails")}
          </Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.phones")}>
        {phones.length ? (
          <Table.Root variant="surface">
            <Table.Header>
              <Table.Row>
                <Table.ColumnHeaderCell>{t("labels.phone")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.primary")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.verified")}</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {phones.map((item) => (
                <Table.Row key={item.id}>
                  <Table.Cell>{item.phone}</Table.Cell>
                  <Table.Cell>{item.isPrimary ? t("labels.yes") : t("labels.no")}</Table.Cell>
                  <Table.Cell>{item.verifiedAt ? formatDateTime(item.verifiedAt) : t("labels.unverified")}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table.Root>
        ) : (
          <Text size="2" color="gray">
            {t("empty.phones")}
          </Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.identities")}>
        {identities.length ? (
          <Table.Root variant="surface">
            <Table.Header>
              <Table.Row>
                <Table.ColumnHeaderCell>{t("labels.provider")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.subject")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell>{t("labels.lastLogin")}</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {identities.map((item) => (
                <Table.Row key={`${item.identityProvider}:${item.providerSubject}`}>
                  <Table.Cell>{item.identityProvider}</Table.Cell>
                  <Table.Cell>{item.providerSubject}</Table.Cell>
                  <Table.Cell>{item.lastLoginAt ? formatDateTime(item.lastLoginAt) : empty}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table.Root>
        ) : (
          <Text size="2" color="gray">
            {t("empty.identities")}
          </Text>
        )}
      </LedgerSection>

      <LedgerSection title={t("sections.media")}>
        <FieldList>
          <FieldRow label={t("labels.avatars")}>
            <Flex gap="2" wrap="wrap">
              {avatars.length
                ? avatars.map((item) => (
                    <img key={item.id} src={buildImageUrl(item.imageId)} alt="" width={48} height={48} style={{ objectFit: "cover" }} />
                  ))
                : <Text size="2">{empty}</Text>}
            </Flex>
          </FieldRow>
          <FieldRow label={t("labels.backgrounds")}>
            <Flex gap="2" wrap="wrap">
              {backgrounds.length
                ? backgrounds.map((item) => (
                    <img key={item.id} src={buildImageUrl(item.imageId)} alt="" width={96} height={64} style={{ objectFit: "cover" }} />
                  ))
                : <Text size="2">{empty}</Text>}
            </Flex>
          </FieldRow>
        </FieldList>
      </LedgerSection>
    </PageFrame>
  );
}
