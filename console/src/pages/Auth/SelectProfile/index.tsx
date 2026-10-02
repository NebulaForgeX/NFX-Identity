import { AnimatedIcon, RightChevron, ShieldCheck, UsersIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Avatar, Badge, Box, Button, Container, Flex, Grid, Heading, Section, Select, Spinner, Text, TextField } from "@radix-ui/themes";
import gsap from "gsap";
import { LanguageEnum, ProfileKindEnum } from "nfx-ui/enums";
import { useCreateAuthorityProfile, useCreateForgerProfile, useListProfiles, useSelectProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";
import type { Profile } from "nfx-ui/types";

import { EmptyState } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { profileHome } from "@/navigations";
import AuthChrome from "@/pages/Auth/shared/AuthChrome";
import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeArray } from "@/utils";

import styles from "./s.module.css";

gsap.registerPlugin(useGSAP);

export default function SelectProfilePage() {
  const { t } = useTranslation("pages.Account.SelectProfile");
  const { t: tIdentity } = useTranslation("pages.Profile.Identity");
  const { t: tHeader } = useTranslation("language");
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
      <Flex direction="column" className={isAuthority ? "js-col-authority" : "js-col-community"}>
        <Section size="1" py="4" className={styles.colHead}>
          <Container size="4" px="4" width="100%" maxWidth="100%">
            <Flex align="center" justify="between" gap="3">
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
          </Container>
        </Section>
        {items.length === 0 ? (
          <EmptyState icon={isAuthority ? ShieldCheck : UsersIcon} title={t(`empty.${kind}`)} />
        ) : (
          items.map((profile) => {
            const name = resolveAccountDisplayName(profile.displayName, profile.profileId);
            const initial = resolveAccountInitial(profile.displayName, profile.profileId);
            const roles = "authorityRoles" in profile ? safeArray(profile.authorityRoles) : safeArray((profile as Profile.Response.ForgerProfileItem).forgerRoles);
            return (
              <Section key={`${kind}:${profile.profileId}`} size="1" py="3" className={styles.row}>
                <Container size="1" px="4" width="100%" maxWidth="100%">
                  <Button
                    type="button"
                    variant="ghost"
                    color="gray"
                    className={styles.rowControl}
                    disabled={selectProfile.isPending}
                    onClick={() => void enter(profile.profileId, kind)}
                  >
                    <Flex align="center" width="100%" gap="3">
                      <Avatar size="3" radius="none" fallback={initial} src={profile.avatarImageId ? buildImageUrl(profile.avatarImageId) : undefined} />
                      <Flex direction="column" minWidth="0" flexGrow="1" gap="1">
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
                    </Flex>
                  </Button>
                </Container>
              </Section>
            );
          })
        )}
      </Flex>
    );
  };

  return (
    <AuthChrome>
      <Flex ref={rootRef} direction="column" height="100%" minHeight="0" width="100%">
        <Section size="1" pt="6" pb="5" className="js-head">
          <Container size="4" px="6" width="100%" maxWidth="100%">
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
          </Container>
        </Section>

        <Flex direction="column" flexGrow="1" minHeight="0" overflow="auto" className={styles.columnsRule}>
          <Grid columns="2" width="100%" className={styles.columns}>
            <Box className={styles.colCommunity}>{renderColumn(ProfileKindEnum.COMMUNITY, community)}</Box>
            <Box className={styles.colAuthority}>{renderColumn(ProfileKindEnum.AUTHORITY, authority)}</Box>
          </Grid>
        </Flex>

        <Section size="1" py="4" position="sticky" bottom="0" className={`${styles.dock} js-dock`}>
          <Container size="4" px="6" width="100%" maxWidth="100%">
            <Flex asChild direction="column" gap="3">
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
                <Grid gap="3" className={styles.dockGrid}>
                  <Flex direction="column" gap="1">
                    <Text as="label" size="1" color="gray">
                      {t("create.displayName")}
                    </Text>
                    <TextField.Root size="2" value={displayName} onChange={(e) => setDisplayName(e.target.value)} placeholder={t("create.displayName")} />
                  </Flex>
                  <Flex direction="column" gap="1">
                    <Text as="label" size="1" color="gray">
                      {tIdentity("labels.kind")}
                    </Text>
                    <Select.Root value={createKind} onValueChange={(v) => setCreateKind(v as ProfileKindEnum)}>
                      <Select.Trigger />
                      <Select.Content>
                        <Select.Item value={ProfileKindEnum.COMMUNITY}>{t("kind.community")}</Select.Item>
                        <Select.Item value={ProfileKindEnum.AUTHORITY}>{t("kind.authority")}</Select.Item>
                      </Select.Content>
                    </Select.Root>
                  </Flex>
                  <Flex direction="column" gap="1">
                    <Text as="label" size="1" color="gray">
                      {tHeader("header.language")}
                    </Text>
                    <Select.Root value={profileLanguage} onValueChange={(v) => setProfileLanguage(v as LanguageEnum)}>
                      <Select.Trigger />
                      <Select.Content>
                        <Select.Item value={LanguageEnum.EN}>English</Select.Item>
                        <Select.Item value={LanguageEnum.ZH}>中文</Select.Item>
                        <Select.Item value={LanguageEnum.FR}>Français</Select.Item>
                      </Select.Content>
                    </Select.Root>
                  </Flex>
                  <Flex align="end">
                    <Button type="submit" size="2" disabled={!displayName.trim()} loading={createForger.isPending || createAuthority.isPending}>
                      {t("create.submit")}
                    </Button>
                  </Flex>
                </Grid>
              </form>
            </Flex>
          </Container>
        </Section>
      </Flex>
    </AuthChrome>
  );
}
