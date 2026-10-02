import { MailFilledIcon, TelephoneIcon, UsersIcon } from "nfx-ui/icons";
import type { Profile } from "nfx-ui/types";

import { useState } from "react";
import { Avatar, Badge, Button, Container, Flex, Grid, Section, Select, Text, TextField } from "@radix-ui/themes";
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
import { useAuthStore, usePreferenceStore } from "nfx-ui/stores";
import { isVerificationCodeComplete, normalizeVerificationCode } from "nfx-ui/utils";
import { useTranslation } from "react-i18next";

import { EmptyState, PageHeader, Suspense } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { profileHome } from "@/navigations";
import { buildAvatarImageSrc, safeArray, safeStringable } from "@/utils";

import { LedgerSection } from "../Ledger";
import styles from "./s.module.css";

function EmailRow({
  item,
}: {
  item: {
    id: string;
    email: string;
    isPrimary: boolean;
    verifiedAt: Nullable<string>;
  };
}) {
  const { t } = useTranslation("pages.Profile.Identity");
  const currentLanguage = usePreferenceStore((s) => s.language);
  const deleteEmail = useDeleteEmail();
  const setPrimary = useSetPrimaryEmail();
  const sendCode = useSendEmailVerificationCode();
  const verify = useVerifyEmail();
  const updateEmail = useUpdateEmail();
  const [code, setCode] = useState("");
  const [nextEmail, setNextEmail] = useState(item.email);
  const [editing, setEditing] = useState(false);
  const verified = Boolean(item.verifiedAt);
  const hint = [item.isPrimary ? t("labels.primary") : null, verified ? t("labels.verified") : t("labels.unverified")].filter(Boolean).join(" · ");

  return (
    <Section size="1" py="4" className={styles.contact}>
      <Flex direction="column" gap="3">
        <Flex align="start" justify="between" gap="4" wrap="wrap">
          <Flex direction="column" gap="1" minWidth="0">
            <Text size="2" weight="bold">
              {item.email}
            </Text>
            <Text size="1" color="gray">
              {hint}
            </Text>
          </Flex>
          <Flex gap="2" wrap="wrap" align="center">
            {!verified ? (
              <Button size="1" variant="outline" loading={sendCode.isPending} onClick={() => sendCode.mutate({ emailId: item.id, lang: currentLanguage ?? LanguageEnum.EN })}>
                {t("actions.sendCode")}
              </Button>
            ) : null}
            {!item.isPrimary ? (
              <Button size="1" variant="outline" onClick={() => setPrimary.mutate(item.id)}>
                {t("actions.setPrimary")}
              </Button>
            ) : null}
            <Button size="1" variant="outline" onClick={() => setEditing((value) => !value)}>
              {editing ? t("actions.cancel") : t("actions.edit")}
            </Button>
            <Button size="1" variant="outline" onClick={() => deleteEmail.mutate(item.id)}>
              {t("actions.remove")}
            </Button>
          </Flex>
        </Flex>
        {editing ? (
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">
              {t("labels.newEmail")}
            </Text>
            <Flex align="center" gap="3" wrap="wrap">
              <TextField.Root size="2" value={nextEmail} onChange={(e) => setNextEmail(e.target.value)} />
              <Button
                size="1"
                loading={updateEmail.isPending}
                disabled={!nextEmail.trim() || nextEmail.trim() === item.email}
                onClick={() => updateEmail.mutate({ emailId: item.id, email: nextEmail.trim() }, { onSuccess: () => setEditing(false) })}
              >
                {t("actions.save")}
              </Button>
            </Flex>
          </Flex>
        ) : null}
        {!verified ? (
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">
              {t("labels.code")}
            </Text>
            <Flex align="center" gap="3" wrap="wrap">
              <TextField.Root
                size="2"
                value={code}
                onChange={(e) => setCode(normalizeVerificationCode(e.target.value))}
                placeholder={t("labels.code")}
              />
              <Button
                size="1"
                loading={verify.isPending}
                disabled={!isVerificationCodeComplete(code)}
                onClick={() => verify.mutate({ emailId: item.id, verificationCode: normalizeVerificationCode(code) }, { onSuccess: () => setCode("") })}
              >
                {t("actions.verify")}
              </Button>
            </Flex>
          </Flex>
        ) : null}
      </Flex>
    </Section>
  );
}

function EmailsPanel() {
  const { t } = useTranslation("pages.Profile.Identity");
  const emails = useListEmails();
  const createEmail = useCreateEmail();
  const [newEmail, setNewEmail] = useState("");
  const items = safeArray(emails.data?.items);

  return (
    <LedgerSection title={t("sections.emails.title")} description={t("sections.emails.description")}>
      <Flex direction="column" gap="4">
        <Flex align="center" gap="3">
          <Flex flexGrow="1" minWidth="0">
            <TextField.Root size="2" value={newEmail} onChange={(e) => setNewEmail(e.target.value)} placeholder={t("labels.emailPlaceholder")} className={styles.search} />
          </Flex>
          <Button size="2" onClick={() => createEmail.mutate({ email: newEmail }, { onSuccess: () => setNewEmail("") })}>
            {t("actions.addEmail")}
          </Button>
        </Flex>
        {items.length ? (
          items.map((item) => <EmailRow key={item.id} item={item} />)
        ) : (
          <EmptyState icon={MailFilledIcon} title={t("empty.emails")} />
        )}
      </Flex>
    </LedgerSection>
  );
}

function PhoneRow({
  item,
}: {
  item: {
    id: string;
    phone: string;
    isPrimary: boolean;
    verifiedAt: Nullable<string>;
  };
}) {
  const { t } = useTranslation("pages.Profile.Identity");
  const deletePhone = useDeletePhone();
  const setPrimary = useSetPrimaryPhone();
  const sendCode = useSendPhoneVerificationCode();
  const verify = useVerifyPhone();
  const updatePhone = useUpdatePhone();
  const [code, setCode] = useState("");
  const [nextPhone, setNextPhone] = useState(item.phone);
  const [editing, setEditing] = useState(false);
  const verified = Boolean(item.verifiedAt);
  const hint = [item.isPrimary ? t("labels.primary") : null, verified ? t("labels.verified") : t("labels.unverified")].filter(Boolean).join(" · ");

  return (
    <Section size="1" py="4" className={styles.contact}>
      <Flex direction="column" gap="3">
        <Flex align="start" justify="between" gap="4" wrap="wrap">
          <Flex direction="column" gap="1" minWidth="0">
            <Text size="2" weight="bold">
              {item.phone}
            </Text>
            <Text size="1" color="gray">
              {hint}
            </Text>
          </Flex>
          <Flex gap="2" wrap="wrap" align="center">
            {!verified ? (
              <Button size="1" variant="outline" loading={sendCode.isPending} onClick={() => sendCode.mutate(item.id)}>
                {t("actions.sendCode")}
              </Button>
            ) : null}
            {!item.isPrimary ? (
              <Button size="1" variant="outline" onClick={() => setPrimary.mutate(item.id)}>
                {t("actions.setPrimary")}
              </Button>
            ) : null}
            <Button size="1" variant="outline" onClick={() => setEditing((value) => !value)}>
              {editing ? t("actions.cancel") : t("actions.edit")}
            </Button>
            <Button size="1" variant="outline" onClick={() => deletePhone.mutate(item.id)}>
              {t("actions.remove")}
            </Button>
          </Flex>
        </Flex>
        {editing ? (
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">
              {t("labels.newPhone")}
            </Text>
            <Flex align="center" gap="3" wrap="wrap">
              <TextField.Root size="2" value={nextPhone} onChange={(e) => setNextPhone(e.target.value)} />
              <Button
                size="1"
                loading={updatePhone.isPending}
                disabled={!nextPhone.trim() || nextPhone.trim() === item.phone}
                onClick={() => updatePhone.mutate({ phoneId: item.id, phone: nextPhone.trim() }, { onSuccess: () => setEditing(false) })}
              >
                {t("actions.save")}
              </Button>
            </Flex>
          </Flex>
        ) : null}
        {!verified ? (
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">
              {t("labels.code")}
            </Text>
            <Flex align="center" gap="3" wrap="wrap">
              <TextField.Root size="2" value={code} onChange={(e) => setCode(normalizeVerificationCode(e.target.value))} placeholder={t("labels.code")} />
              <Button
                size="1"
                loading={verify.isPending}
                disabled={!isVerificationCodeComplete(code)}
                onClick={() => verify.mutate({ phoneId: item.id, verificationCode: normalizeVerificationCode(code) }, { onSuccess: () => setCode("") })}
              >
                {t("actions.verify")}
              </Button>
            </Flex>
          </Flex>
        ) : null}
      </Flex>
    </Section>
  );
}

function PhonesPanel() {
  const { t } = useTranslation("pages.Profile.Identity");
  const phones = useListPhones();
  const createPhone = useCreatePhone();
  const [newPhone, setNewPhone] = useState("");
  const items = safeArray(phones.data?.items);

  return (
    <LedgerSection title={t("sections.phones.title")} description={t("sections.phones.description")}>
      <Flex direction="column" gap="4">
        <Flex align="center" gap="3">
          <Flex flexGrow="1" minWidth="0">
            <TextField.Root size="2" value={newPhone} onChange={(e) => setNewPhone(e.target.value)} placeholder={t("labels.phonePlaceholder")} className={styles.search} />
          </Flex>
          <Button size="2" onClick={() => createPhone.mutate({ phone: newPhone }, { onSuccess: () => setNewPhone("") })}>
            {t("actions.addPhone")}
          </Button>
        </Flex>
        {items.length ? (
          items.map((item) => <PhoneRow key={item.id} item={item} />)
        ) : (
          <EmptyState icon={TelephoneIcon} title={t("empty.phones")} />
        )}
      </Flex>
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
      <Flex direction="column" gap="4">
        <Section size="1" py="3">
          <TextField.Root size="2" value={query} onChange={(e) => setQuery(e.target.value)} placeholder={t("labels.searchProfiles")} className={styles.search} />
        </Section>
        {rows.length ? (
          rows.map((row) => {
            const isCurrent = row.profileId === currentProfileId && row.kind === currentProfileKind;
            const name = safeStringable(row.displayName) || t("labels.emptyName");
            return (
              <Section key={`${row.kind}:${row.profileId}`} size="1" py="3" className={styles.row}>
                <Container size="2" px="4">
                  <Flex align="center" justify="between" gap="3" wrap="wrap">
                    <Flex align="center" gap="3" minWidth="0">
                      <Avatar size="2" src={row.avatarImageId ? buildAvatarImageSrc(row.avatarImageId) : undefined} fallback={name.slice(0, 1)} />
                      <Flex direction="column" gap="1" minWidth="0">
                        <Text size="2" weight="bold">
                          {name}
                        </Text>
                        <Text size="1" color="gray">
                          {row.profileId}
                        </Text>
                        <Flex gap="2" wrap="wrap" align="center">
                          <Badge variant="outline">{row.kind === ProfileKindEnum.COMMUNITY ? t("labels.scopeCommunity") : t("labels.scopeAuthority")}</Badge>
                          {isCurrent ? (
                            <Badge variant="outline">
                              {t("labels.current")}
                            </Badge>
                          ) : null}
                        </Flex>
                      </Flex>
                    </Flex>
                    <Flex gap="2">
                      {isCurrent ? null : (
                        <Button size="1" variant="outline" loading={switchingId === row.profileId} onClick={() => void handleSwitch(row.profileId, row.kind)}>
                          {t("actions.switch")}
                        </Button>
                      )}
                      <Button
                        size="1"
                        variant="outline"
                        disabled={isCurrent}
                        onClick={() => {
                          if (!window.confirm(t("labels.deleteConfirmBody", { name }))) return;
                          deleteProfile.mutate({ profileId: row.profileId, kind: row.kind });
                        }}
                      >
                        {t("actions.deleteProfile")}
                      </Button>
                    </Flex>
                  </Flex>
                </Container>
              </Section>
            );
          })
        ) : (
          <EmptyState icon={UsersIcon} title={t("empty.profiles")} />
        )}

        <Section size="1" py="4" className={styles.dock}>
          <Container size="2" px="4">
            <Flex direction="column" gap="4">
              <Flex direction="column" gap="1">
                <Text as="label" size="1" color="gray">
                  {t("labels.language")}
                </Text>
                <Select.Root value={profileLanguage} onValueChange={(v) => setProfileLanguage(v as LanguageEnum)}>
                  <Select.Trigger />
                  <Select.Content>
                    <Select.Item value={LanguageEnum.EN}>{t("labels.langEn")}</Select.Item>
                    <Select.Item value={LanguageEnum.ZH}>{t("labels.langZh")}</Select.Item>
                    <Select.Item value={LanguageEnum.FR}>{t("labels.langFr")}</Select.Item>
                  </Select.Content>
                </Select.Root>
              </Flex>
              <Grid columns={{ initial: "1", sm: "2" }} gap="4">
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">
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
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">
                    {t("labels.newAuthorityProfile")}
                  </Text>
                  <TextField.Root size="2" value={authorityName} onChange={(e) => setAuthorityName(e.target.value)} placeholder={t("labels.displayName")} />
                  <Button
                    size="2"
                    disabled={!authorityName.trim()}
                    loading={createAuthority.isPending}
                    onClick={() => createAuthority.mutate({ displayName: authorityName.trim(), profileLanguage }, { onSuccess: () => setAuthorityName("") })}
                  >
                    {t("actions.createAuthority")}
                  </Button>
                </Flex>
              </Grid>
            </Flex>
          </Container>
        </Section>
      </Flex>
    </LedgerSection>
  );
}

export default function IdentityView() {
  const { t } = useTranslation("pages.Profile.Identity");
  return (
    <PageFrame>
      <PageHeader icon={UsersIcon} title={t("title")} description={t("description")} />
      <Suspense loadingText={t("labels.loading")}>
        <Grid columns={{ initial: "1", lg: "minmax(0, 1.15fr) minmax(18rem, 0.85fr)" }} gap="6" align="start">
          <ProfilesPanel />
          <Flex direction="column" gap="6" minWidth="0">
            <EmailsPanel />
            <PhonesPanel />
          </Flex>
        </Grid>
      </Suspense>
    </PageFrame>
  );
}
