import { AnimatedIcon, RightChevron, ShieldCheck, UsersIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Avatar, Badge, Box, Button, Flex, Heading, Select, Spinner, Text, TextField } from "@radix-ui/themes";
import gsap from "gsap";
import { LanguageEnum, ProfileKindEnum } from "nfx-ui/enums";
import { useCreateAuthorityProfile, useCreateForgerProfile, useListProfiles, useSelectProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import type { Profile } from "nfx-ui/types";

import { routerEventEmitter } from "@/events/router";
import { profileHome } from "@/navigations";
import AuthChrome from "@/pages/Auth/shared/AuthChrome";
import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeArray } from "@/utils";

import styles from "./s.module.css";

gsap.registerPlugin(useGSAP);

export default function SelectProfilePage() {
  const { t } = useTranslation("pages.Account.SelectProfile");
  const selectProfile = useSelectProfile();
  const createForger = useCreateForgerProfile();
  const createAuthority = useCreateAuthorityProfile();
  const communityProfiles = useListProfiles(ProfileKindEnum.COMMUNITY);
  const authorityProfiles = useListProfiles(ProfileKindEnum.AUTHORITY);
  const [displayName, setDisplayName] = useState("");
  const [profileLanguage, setProfileLanguage] = useState<LanguageEnum>(LanguageEnum.EN);
  const [createKind, setCreateKind] = useState<ProfileKindEnum>(ProfileKindEnum.COMMUNITY);
  const rootRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) return;
      gsap.set(".js-col-community", { autoAlpha: 0, x: -36 });
      gsap.set(".js-col-authority", { autoAlpha: 0, x: 36 });
      gsap.set(".js-dock", { autoAlpha: 0, y: 20 });
      gsap.set(".js-head", { autoAlpha: 0, y: 12 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-head", { autoAlpha: 1, y: 0, duration: 0.45 })
        .to(".js-col-community", { autoAlpha: 1, x: 0, duration: 0.55 }, "-=0.15")
        .to(".js-col-authority", { autoAlpha: 1, x: 0, duration: 0.55 }, "<")
        .to(".js-dock", { autoAlpha: 1, y: 0, duration: 0.45 }, "-=0.2");
    },
    { scope: rootRef },
  );

  const community = safeArray(communityProfiles.data?.items);
  const authority = safeArray(authorityProfiles.data?.items);

  const enter = async (profileId: string, kind: ProfileKindEnum) => {
    await selectProfile.mutateAsync({ profileId, kind });
    routerEventEmitter.navigate({ to: profileHome(kind), replace: true });
  };

  const renderColumn = (kind: ProfileKindEnum, items: Array<Profile.Response.ForgerProfileItem | Profile.Response.AuthorityProfileItem>) => {
    const isAuthority = kind === ProfileKindEnum.AUTHORITY;
    return (
      <Box asChild className={`${styles.col} ${isAuthority ? styles.colAuthority : styles.colCommunity} ${isAuthority ? "js-col-authority" : "js-col-community"}`}>
        <section>
          <Box className={styles.colPx}>
            <Box className={styles.colPy}>
          <Box className={styles.colHeadHairline}>
            <Box className={styles.colHeadPy}>
          <Flex className={styles.colHead} align="baseline" justify="between" gap="4">
            <Flex align="center" gap="2">
              <AnimatedIcon icon={isAuthority ? ShieldCheck : UsersIcon} size={16} />
              <Text size="4" weight="bold" className={styles.colName}>
                {t(`kind.${kind}`)}
              </Text>
            </Flex>
            <Text size="1" className={styles.colHint}>
              {t(`kindHint.${kind}`)}
            </Text>
          </Flex>
            </Box>
          </Box>
          {items.length === 0 ? (
            <Text as="p" size="2" className={styles.empty}>
              {t(`empty.${kind}`)}
            </Text>
          ) : (
            items.map((profile) => {
              const name = resolveAccountDisplayName(profile.displayName, profile.profileId);
              const initial = resolveAccountInitial(profile.displayName, profile.profileId);
              const roles = "authorityRoles" in profile ? safeArray(profile.authorityRoles) : safeArray((profile as Profile.Response.ForgerProfileItem).forgerRoles);
              return (
                <Button
                  key={`${kind}:${profile.profileId}`}
                  type="button"
                  variant="ghost"
                  className={styles.row}
                  disabled={selectProfile.isPending}
                  onClick={() => void enter(profile.profileId, kind)}
                >
                  <Avatar size="3" radius="none" fallback={initial} src={profile.avatarImageId ? buildImageUrl(profile.avatarImageId) : undefined} />
                  <Flex direction="column" gap="1" className={styles.rowMeta}>
                    <Text size="3" weight="bold">
                      {name}
                    </Text>
                    <Flex gap="1" wrap="wrap">
                      {roles.map((role) => (
                        <Badge key={role} variant="outline" size="1">
                          {t(`roles.${role}`, { defaultValue: role })}
                        </Badge>
                      ))}
                    </Flex>
                  </Flex>
                  {selectProfile.isPending ? <Spinner size="2" /> : <AnimatedIcon icon={RightChevron} size={18} />}
                </Button>
              );
            })
          )}
            </Box>
          </Box>
        </section>
      </Box>
    );
  };

  return (
    <AuthChrome>
      <Box ref={rootRef} className={styles.ledger}>
        <Box className={`${styles.head} js-head`}>
          <Box className={styles.headPx}>
            <Box className={styles.headPy}>
          <Flex direction="column" gap="2">
            <Text as="p" size="1" weight="bold" className={styles.kicker}>
              {t("eyebrow")}
            </Text>
            <Heading as="h1" size="7" className={styles.title}>
              {t("title")}
            </Heading>
            <Text as="p" size="2" className={styles.lede}>
              {t("subtitle")}
            </Text>
          </Flex>
            </Box>
          </Box>
        </Box>

        <Box className={styles.columns}>
          {renderColumn(ProfileKindEnum.COMMUNITY, community)}
          {renderColumn(ProfileKindEnum.AUTHORITY, authority)}
        </Box>

        <Box className={`${styles.dock} js-dock`}>
          <Box className={styles.dockPx}>
            <Box className={styles.dockPy}>
        <Flex
          asChild
          align="center"
          gap="3"
        >
          <form
            onSubmit={async (e) => {
              e.preventDefault();
              const name = displayName.trim();
              if (!name) return;
              const created =
                createKind === ProfileKindEnum.AUTHORITY
                  ? await createAuthority.mutateAsync({ displayName: name, profileLanguage })
                  : await createForger.mutateAsync({ displayName: name, profileLanguage });
              setDisplayName("");
              await enter(created.profileId, createKind);
            }}
          >
            <Text as="p" size="1" weight="bold" className={styles.dockLabel}>
              {t("create.title")}
            </Text>
            <TextField.Root size="2" value={displayName} onChange={(e) => setDisplayName(e.target.value)} placeholder={t("create.displayName")} />
            <Select.Root value={createKind} onValueChange={(v) => setCreateKind(v as ProfileKindEnum)}>
              <Select.Trigger />
              <Select.Content>
                <Select.Item value={ProfileKindEnum.COMMUNITY}>{t("kind.community")}</Select.Item>
                <Select.Item value={ProfileKindEnum.AUTHORITY}>{t("kind.authority")}</Select.Item>
              </Select.Content>
            </Select.Root>
            <Select.Root value={profileLanguage} onValueChange={(v) => setProfileLanguage(v as LanguageEnum)}>
              <Select.Trigger />
              <Select.Content>
                <Select.Item value={LanguageEnum.EN}>English</Select.Item>
                <Select.Item value={LanguageEnum.ZH}>中文</Select.Item>
                <Select.Item value={LanguageEnum.FR}>Français</Select.Item>
              </Select.Content>
            </Select.Root>
            <Button type="submit" size="2" disabled={!displayName.trim()} loading={createForger.isPending || createAuthority.isPending}>
              {t("create.submit")}
            </Button>
          </form>
        </Flex>
            </Box>
          </Box>
        </Box>
      </Box>
    </AuthChrome>
  );
}
