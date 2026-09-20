import { AnimatedIcon, RightChevron, ShieldCheck, UsersIcon } from "nfx-ui/icons";
import { useMemo, useState } from "react";
import { Avatar, Badge, Button, Flex, Heading, Spinner, Text, TextField, Select } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { LanguageEnum, ProfileKindEnum } from "nfx-ui/enums";
import { useCreateAuthorityProfile, useCreateForgerProfile, useListProfiles, useSelectProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import type { Profile } from "nfx-ui/types";

import { routerEventEmitter } from "@/events/router";
import { profileHome } from "@/navigations";
import AuthShell from "@/pages/Auth/shared/AuthShell";
import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeArray } from "@/utils";

import styles from "../Login/s.module.css";

export default function SelectProfilePage() {
  const { t } = useTranslation("pages.Account.SelectProfile");
  const selectProfile = useSelectProfile();
  const createForger = useCreateForgerProfile();
  const createAuthority = useCreateAuthorityProfile();
  const forgerProfiles = useListProfiles(ProfileKindEnum.FORGER);
  const authorityProfiles = useListProfiles(ProfileKindEnum.AUTHORITY);
  const [displayName, setDisplayName] = useState("");
  const [profileLanguage, setProfileLanguage] = useState<LanguageEnum>(LanguageEnum.EN);
  const [createKind, setCreateKind] = useState<ProfileKindEnum>(ProfileKindEnum.FORGER);

  const forger = safeArray(forgerProfiles.data?.items);
  const authority = safeArray(authorityProfiles.data?.items);
  const groups = useMemo(
    () =>
      [
        { kind: ProfileKindEnum.FORGER as const, items: forger },
        { kind: ProfileKindEnum.AUTHORITY as const, items: authority },
      ].filter((group) => group.items.length > 0),
    [forger, authority],
  );

  const enter = async (profileId: string, kind: ProfileKindEnum) => {
    await selectProfile.mutateAsync({ profileId, kind });
    routerEventEmitter.navigate({ to: profileHome(kind), replace: true });
  };

  return (
    <AuthShell brandEyebrow={t("brand.access", { name: APP_NAME })} brandTitle={t("brand.title")} heroFooter={t("heroFooter")}>
      <Flex direction="column" gap="5">
        <Flex direction="column" gap="1" className="js-auth-stagger">
          <Text as="p" size="1" weight="bold" className={styles.kicker}>
            {t("eyebrow")}
          </Text>
          <Heading as="h2" size="6">
            {t("title")}
          </Heading>
          <Text as="p" size="2" color="gray">
            {t("subtitle")}
          </Text>
        </Flex>

        <Flex direction="column" gap="4" className="js-auth-stagger">
          {groups.map((group) => {
            const isAuthority = group.kind === ProfileKindEnum.AUTHORITY;
            const kindLabel = t(`kind.${group.kind}`);
            return (
              <Flex key={group.kind} direction="column" gap="2">
                <Flex align="center" gap="2">
                  <AnimatedIcon icon={isAuthority ? ShieldCheck : UsersIcon} size={15} />
                  <Text size="2" weight="bold">
                    {kindLabel}
                  </Text>
                </Flex>
                {group.items.map((profile: Profile.Response.ForgerProfileItem | Profile.Response.AuthorityProfileItem) => {
                  const name = resolveAccountDisplayName(profile.displayName, profile.profileId);
                  const initial = resolveAccountInitial(profile.displayName, profile.profileId);
                  const roles = "authorityRoles" in profile ? safeArray(profile.authorityRoles) : safeArray((profile as Profile.Response.ForgerProfileItem).forgerRoles);
                  return (
                    <Button
                      key={`${group.kind}:${profile.profileId}`}
                      type="button"
                      variant="outline"
                      size="3"
                      disabled={selectProfile.isPending}
                      onClick={() => void enter(profile.profileId, group.kind)}
                    >
                      <Flex align="center" gap="3" width="100%">
                        <Avatar size="3" radius="none" fallback={initial} src={profile.avatarImageId ? buildImageUrl(profile.avatarImageId) : undefined} />
                        <Flex direction="column" align="start" minWidth="0" flexGrow="1">
                          <Text size="3" weight="bold">
                            {name}
                          </Text>
                          <Flex gap="1" wrap="wrap">
                            {roles.map((role) => (
                              <Badge key={role} variant="soft" size="1">
                                {t(`roles.${role}`, { defaultValue: role })}
                              </Badge>
                            ))}
                          </Flex>
                        </Flex>
                        {selectProfile.isPending ? <Spinner size="2" /> : <AnimatedIcon icon={RightChevron} size={18} />}
                      </Flex>
                    </Button>
                  );
                })}
              </Flex>
            );
          })}
        </Flex>

        <Flex direction="column" gap="3" className="js-auth-stagger">
          <Text size="2" weight="bold">
            {t("create.title")}
          </Text>
          <Select.Root value={createKind} onValueChange={(v) => setCreateKind(v as ProfileKindEnum)}>
            <Select.Trigger />
            <Select.Content>
              <Select.Item value={ProfileKindEnum.FORGER}>{t("kind.forger")}</Select.Item>
              <Select.Item value={ProfileKindEnum.AUTHORITY}>{t("kind.authority")}</Select.Item>
            </Select.Content>
          </Select.Root>
          <TextField.Root size="3" value={displayName} onChange={(e) => setDisplayName(e.target.value)} placeholder={t("create.displayName")} />
          <Select.Root value={profileLanguage} onValueChange={(v) => setProfileLanguage(v as LanguageEnum)}>
            <Select.Trigger />
            <Select.Content>
              <Select.Item value={LanguageEnum.EN}>English</Select.Item>
              <Select.Item value={LanguageEnum.ZH}>中文</Select.Item>
              <Select.Item value={LanguageEnum.FR}>Français</Select.Item>
            </Select.Content>
          </Select.Root>
          <Button
            size="3"
            disabled={!displayName.trim()}
            loading={createForger.isPending || createAuthority.isPending}
            onClick={async () => {
              const name = displayName.trim();
              const created =
                createKind === ProfileKindEnum.AUTHORITY
                  ? await createAuthority.mutateAsync({ displayName: name, profileLanguage })
                  : await createForger.mutateAsync({ displayName: name, profileLanguage });
              setDisplayName("");
              await enter(created.profileId, createKind);
            }}
          >
            {t("create.submit")}
          </Button>
        </Flex>
      </Flex>
    </AuthShell>
  );
}
