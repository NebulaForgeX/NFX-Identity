import { UsersIcon } from "nfx-ui/icons";
import type { Profile } from "nfx-ui/types";

import { useState } from "react";
import { Avatar, Badge, Box, Button, Flex, Select, Table, Text, TextField } from "@radix-ui/themes";
import { LanguageEnum, ProfileKindEnum } from "nfx-ui/enums";
import {
  useCreateAuthorityProfile,
  useCreateEmail,
  useCreateForgerProfile,
  useCreatePhone,
  useDeleteEmail,
  useDeletePhone,
  useDeleteProfile,
  useListEmails,
  useListPhones,
  useListProfiles,
  useSearchAuthorityProfiles,
  useSearchForgerProfiles,
  useSelectProfile,
  useSendEmailVerificationCode,
  useSendPhoneVerificationCode,
  useSetPrimaryEmail,
  useSetPrimaryPhone,
  useUpdateEmail,
  useUpdatePhone,
  useVerifyEmail,
  useVerifyPhone,
} from "nfx-ui/hooks";
import { useAuthStore } from "nfx-ui/stores";
import { isVerificationCodeComplete, normalizeVerificationCode } from "nfx-ui/utils";
import { useTranslation } from "react-i18next";

import { PageHeader, Suspense } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { profileHome } from "@/navigations";
import { buildAvatarImageSrc, formatDateTime, safeArray, safeStringable } from "@/utils";

import { LedgerSection } from "./Ledger";

function EmailsPanel() {
  const { t } = useTranslation("pages.Profile.Identity");
  const emails = useListEmails();
  const createEmail = useCreateEmail();
  const deleteEmail = useDeleteEmail();
  const setPrimary = useSetPrimaryEmail();
  const sendCode = useSendEmailVerificationCode();
  const verify = useVerifyEmail();
  const updateEmail = useUpdateEmail();
  const [newEmail, setNewEmail] = useState("");
  const [codes, setCodes] = useState<Record<string, string>>({});
  const [edits, setEdits] = useState<Record<string, string>>({});
  const items = safeArray(emails.data?.items);

  return (
    <LedgerSection title={t("sections.emails.title")} description={t("sections.emails.description")}>
      <Box pb="3">
      <Flex gap="2" wrap="wrap">
        <Box minWidth="220px" flexGrow="1">
          <TextField.Root size="2" value={newEmail} onChange={(e) => setNewEmail(e.target.value)} placeholder={t("labels.emailPlaceholder")} />
        </Box>
        <Button size="2" onClick={() => createEmail.mutate({ email: newEmail }, { onSuccess: () => setNewEmail("") })}>
          {t("actions.addEmail")}
        </Button>
      </Flex>
      </Box>
      {items.length ? (
        <Table.Root variant="surface">
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell>{t("labels.email")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.status")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.actions")}</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {items.map((item) => {
              const verified = Boolean(item.verifiedAt);
              return (
                <Table.Row key={item.id}>
                  <Table.Cell>
                    <Flex direction="column" gap="1">
                      <Text size="2">{item.email}</Text>
                      <TextField.Root
                        size="1"
                        value={edits[item.id] ?? item.email}
                        onChange={(e) => setEdits((prev) => ({ ...prev, [item.id]: e.target.value }))}
                      />
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="1" wrap="wrap">
                      {item.isPrimary ? <Badge variant="outline">{t("labels.primary")}</Badge> : null}
                      <Badge variant="outline" color={verified ? "green" : "gray"}>
                        {verified ? formatDateTime(item.verifiedAt as string) : t("labels.unverified")}
                      </Badge>
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="2" wrap="wrap">
                      {!verified ? (
                        <>
                          <Button size="1" variant="outline" onClick={() => sendCode.mutate({ emailId: item.id })}>
                            {t("actions.sendCode")}
                          </Button>
                          <TextField.Root
                            size="1"
                            style={{ width: 108 }}
                            value={codes[item.id] ?? ""}
                            onChange={(e) => setCodes((prev) => ({ ...prev, [item.id]: e.target.value }))}
                            placeholder={t("labels.code")}
                          />
                          <Button
                            size="1"
                            disabled={!isVerificationCodeComplete(codes[item.id] ?? "")}
                            onClick={() => verify.mutate({ emailId: item.id, verificationCode: normalizeVerificationCode(codes[item.id] ?? "") })}
                          >
                            {t("actions.verify")}
                          </Button>
                        </>
                      ) : null}
                      <Button
                        size="1"
                        variant="outline"
                        disabled={!edits[item.id] || edits[item.id] === item.email}
                        onClick={() => updateEmail.mutate({ emailId: item.id, email: (edits[item.id] ?? "").trim() })}
                      >
                        {t("actions.save")}
                      </Button>
                      {!item.isPrimary ? (
                        <Button size="1" variant="outline" onClick={() => setPrimary.mutate(item.id)}>
                          {t("actions.setPrimary")}
                        </Button>
                      ) : null}
                      <Button size="1" variant="outline" color="red" onClick={() => deleteEmail.mutate(item.id)}>
                        {t("actions.remove")}
                      </Button>
                    </Flex>
                  </Table.Cell>
                </Table.Row>
              );
            })}
          </Table.Body>
        </Table.Root>
      ) : (
        <Text size="2" color="gray">
          {t("empty.emails")}
        </Text>
      )}
    </LedgerSection>
  );
}

function PhonesPanel() {
  const { t } = useTranslation("pages.Profile.Identity");
  const phones = useListPhones();
  const createPhone = useCreatePhone();
  const deletePhone = useDeletePhone();
  const setPrimary = useSetPrimaryPhone();
  const sendCode = useSendPhoneVerificationCode();
  const verify = useVerifyPhone();
  const updatePhone = useUpdatePhone();
  const [newPhone, setNewPhone] = useState("");
  const [codes, setCodes] = useState<Record<string, string>>({});
  const [edits, setEdits] = useState<Record<string, string>>({});
  const items = safeArray(phones.data?.items);

  return (
    <LedgerSection title={t("sections.phones.title")} description={t("sections.phones.description")}>
      <Box pb="3">
      <Flex gap="2" wrap="wrap">
        <Box minWidth="220px" flexGrow="1">
          <TextField.Root size="2" value={newPhone} onChange={(e) => setNewPhone(e.target.value)} placeholder={t("labels.phonePlaceholder")} />
        </Box>
        <Button size="2" onClick={() => createPhone.mutate({ phone: newPhone }, { onSuccess: () => setNewPhone("") })}>
          {t("actions.addPhone")}
        </Button>
      </Flex>
      </Box>
      {items.length ? (
        <Table.Root variant="surface">
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell>{t("labels.phone")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.status")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.actions")}</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {items.map((item) => {
              const verified = Boolean(item.verifiedAt);
              return (
                <Table.Row key={item.id}>
                  <Table.Cell>
                    <Flex direction="column" gap="1">
                      <Text size="2">{item.phone}</Text>
                      <TextField.Root
                        size="1"
                        value={edits[item.id] ?? item.phone}
                        onChange={(e) => setEdits((prev) => ({ ...prev, [item.id]: e.target.value }))}
                      />
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="1" wrap="wrap">
                      {item.isPrimary ? <Badge variant="outline">{t("labels.primary")}</Badge> : null}
                      <Badge variant="outline" color={verified ? "green" : "gray"}>
                        {verified ? formatDateTime(item.verifiedAt as string) : t("labels.unverified")}
                      </Badge>
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="2" wrap="wrap">
                      {!verified ? (
                        <>
                          <Button size="1" variant="outline" onClick={() => sendCode.mutate(item.id)}>
                            {t("actions.sendCode")}
                          </Button>
                          <TextField.Root
                            size="1"
                            style={{ width: 108 }}
                            value={codes[item.id] ?? ""}
                            onChange={(e) => setCodes((prev) => ({ ...prev, [item.id]: e.target.value }))}
                            placeholder={t("labels.code")}
                          />
                          <Button
                            size="1"
                            disabled={!isVerificationCodeComplete(codes[item.id] ?? "")}
                            onClick={() => verify.mutate({ phoneId: item.id, verificationCode: normalizeVerificationCode(codes[item.id] ?? "") })}
                          >
                            {t("actions.verify")}
                          </Button>
                        </>
                      ) : null}
                      <Button
                        size="1"
                        variant="outline"
                        disabled={!edits[item.id] || edits[item.id] === item.phone}
                        onClick={() => updatePhone.mutate({ phoneId: item.id, phone: (edits[item.id] ?? "").trim() })}
                      >
                        {t("actions.save")}
                      </Button>
                      {!item.isPrimary ? (
                        <Button size="1" variant="outline" onClick={() => setPrimary.mutate(item.id)}>
                          {t("actions.setPrimary")}
                        </Button>
                      ) : null}
                      <Button size="1" variant="outline" color="red" onClick={() => deletePhone.mutate(item.id)}>
                        {t("actions.remove")}
                      </Button>
                    </Flex>
                  </Table.Cell>
                </Table.Row>
              );
            })}
          </Table.Body>
        </Table.Root>
      ) : (
        <Text size="2" color="gray">
          {t("empty.phones")}
        </Text>
      )}
    </LedgerSection>
  );
}

function ProfilesPanel() {
  const { t } = useTranslation("pages.Profile.Identity");
  const currentProfileId = useAuthStore((s) => s.currentProfileId);
  const currentProfileKind = useAuthStore((s) => s.currentProfileKind);
  const communityProfiles = useListProfiles(ProfileKindEnum.COMMUNITY);
  const authorityProfiles = useListProfiles(ProfileKindEnum.AUTHORITY);
  const [query, setQuery] = useState("");
  const searchedForger = useSearchForgerProfiles(query);
  const searchedAuthority = useSearchAuthorityProfiles(query);
  const createForger = useCreateForgerProfile();
  const createAuthority = useCreateAuthorityProfile();
  const deleteProfile = useDeleteProfile();
  const selectProfile = useSelectProfile();
  const [forgerName, setForgerName] = useState("");
  const [authorityName, setAuthorityName] = useState("");
  const [profileLanguage, setProfileLanguage] = useState<LanguageEnum>(LanguageEnum.EN);
  const [switchingId, setSwitchingId] = useState<Nullable<string>>(null);

  const searching = query.trim().length > 0;
  const community = searching ? safeArray(searchedForger.data?.items) : safeArray(communityProfiles.data?.items);
  const authority = searching ? safeArray(searchedAuthority.data?.items) : safeArray(authorityProfiles.data?.items);
  const rows: Array<{ profileId: string; displayName: Nullable<string>; kind: ProfileKindEnum; avatarImageId: Nullable<string> }> = [
    ...community.map((item: Profile.Response.ForgerProfileItem) => ({
      profileId: item.profileId,
      displayName: item.displayName,
      kind: ProfileKindEnum.COMMUNITY,
      avatarImageId: item.avatarImageId,
    })),
    ...authority.map((item: Profile.Response.AuthorityProfileItem) => ({
      profileId: item.profileId,
      displayName: item.displayName,
      kind: ProfileKindEnum.AUTHORITY,
      avatarImageId: item.avatarImageId,
    })),
  ];

  const handleSwitch = async (profileId: string, kind: ProfileKindEnum) => {
    if ((profileId === currentProfileId && kind === currentProfileKind) || selectProfile.isPending) return;
    setSwitchingId(profileId);
    try {
      await selectProfile.mutateAsync({ profileId, kind });
      routerEventEmitter.navigate({ to: profileHome(kind), replace: true });
    } finally {
      setSwitchingId(null);
    }
  };

  return (
    <LedgerSection title={t("sections.profiles.title")} description={t("sections.profiles.description")}>
      <TextField.Root size="2" mb="3" value={query} onChange={(e) => setQuery(e.target.value)} placeholder={t("labels.searchProfiles")} />
      {rows.length ? (
        <Table.Root variant="surface">
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell>{t("labels.profile")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.kind")}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{t("labels.actions")}</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {rows.map((row) => {
              const isCurrent = row.profileId === currentProfileId && row.kind === currentProfileKind;
              const name = safeStringable(row.displayName) || t("labels.emptyName");
              return (
                <Table.Row key={`${row.kind}:${row.profileId}`}>
                  <Table.Cell>
                    <Flex align="center" gap="2">
                      <Avatar size="2" src={row.avatarImageId ? buildAvatarImageSrc(row.avatarImageId) : undefined} fallback={name.slice(0, 1)} />
                      <Flex direction="column">
                        <Text size="2">{name}</Text>
                        <Text size="1" color="gray">
                          {row.profileId}
                        </Text>
                      </Flex>
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="2" wrap="wrap" align="center">
                      <Badge variant="outline">{row.kind === ProfileKindEnum.COMMUNITY ? t("labels.scopeCommunity") : t("labels.scopeAuthority")}</Badge>
                      {isCurrent ? (
                        <Badge color="green" variant="outline">
                          {t("labels.current")}
                        </Badge>
                      ) : null}
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>
                    <Flex gap="2">
                      {isCurrent ? null : (
                        <Button size="1" variant="outline" loading={switchingId === row.profileId} onClick={() => void handleSwitch(row.profileId, row.kind)}>
                          {t("actions.switch")}
                        </Button>
                      )}
                      <Button
                        size="1"
                        variant="outline"
                        color="red"
                        disabled={isCurrent}
                        onClick={() => {
                          if (!window.confirm(t("labels.deleteConfirmBody", { name }))) return;
                          deleteProfile.mutate({ profileId: row.profileId, kind: row.kind });
                        }}
                      >
                        {t("actions.deleteProfile")}
                      </Button>
                    </Flex>
                  </Table.Cell>
                </Table.Row>
              );
            })}
          </Table.Body>
        </Table.Root>
      ) : (
        <Text size="2" color="gray">
          {t("empty.profiles")}
        </Text>
      )}

      <Box pt="4">
      <Flex direction={{ initial: "column", md: "row" }} gap="4">
        <Flex direction="column" gap="2" flexGrow="1">
          <Text size="2" weight="bold">
            {t("labels.newForgerProfile")}
          </Text>
          <TextField.Root size="2" value={forgerName} onChange={(e) => setForgerName(e.target.value)} placeholder={t("labels.displayName")} />
          <Button
            size="2"
            disabled={!forgerName.trim()}
            loading={createForger.isPending}
            onClick={() => createForger.mutate({ displayName: forgerName.trim(), profileLanguage }, { onSuccess: () => setForgerName("") })}
          >
            {t("actions.createForger")}
          </Button>
        </Flex>
        <Flex direction="column" gap="2" flexGrow="1">
          <Text size="2" weight="bold">
            {t("labels.newAuthorityProfile")}
          </Text>
          <TextField.Root size="2" value={authorityName} onChange={(e) => setAuthorityName(e.target.value)} placeholder={t("labels.displayName")} />
          <Select.Root value={profileLanguage} onValueChange={(v) => setProfileLanguage(v as LanguageEnum)}>
            <Select.Trigger />
            <Select.Content>
              <Select.Item value={LanguageEnum.EN}>{t("labels.langEn")}</Select.Item>
              <Select.Item value={LanguageEnum.ZH}>{t("labels.langZh")}</Select.Item>
              <Select.Item value={LanguageEnum.FR}>{t("labels.langFr")}</Select.Item>
            </Select.Content>
          </Select.Root>
          <Button
            size="2"
            disabled={!authorityName.trim()}
            loading={createAuthority.isPending}
            onClick={() => createAuthority.mutate({ displayName: authorityName.trim(), profileLanguage }, { onSuccess: () => setAuthorityName("") })}
          >
            {t("actions.createAuthority")}
          </Button>
        </Flex>
      </Flex>
      </Box>
    </LedgerSection>
  );
}

export default function IdentityView() {
  const { t } = useTranslation("pages.Profile.Identity");
  return (
    <PageFrame>
      <PageHeader icon={UsersIcon} title={t("title")} description={t("description")} />
      <Suspense loadingText={t("labels.loading")}>
        <ProfilesPanel />
        <EmailsPanel />
        <PhonesPanel />
      </Suspense>
    </PageFrame>
  );
}
