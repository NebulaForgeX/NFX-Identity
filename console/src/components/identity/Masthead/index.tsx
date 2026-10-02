import type { ReactNode } from "react";

import { Avatar, Badge, Box, Container, Flex, Grid, Heading, Section, Text } from "@radix-ui/themes";
import { ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";

import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeArray, safeNullable, safeStringable } from "@/utils";

import styles from "./s.module.css";

export function profileRoles(kind: ProfileKindEnum, data: Maybe<Profile.Response.FullAccountInformationWithCommunityProfile | Profile.Response.FullAccountInformationWithAuthorityProfile>): string[] {
  if (!data) return [];
  if (kind === ProfileKindEnum.AUTHORITY) {
    return safeArray((data as Profile.Response.FullAccountInformationWithAuthorityProfile).authorityProfile?.authorityRoles);
  }
  return safeArray((data as Profile.Response.FullAccountInformationWithCommunityProfile).communityProfile?.forgerRoles);
}

export type MastheadStat = {
  label: string;
  value: number;
};

export default function Masthead({
  kind,
  data,
  profile,
  action,
  stats,
}: {
  kind: ProfileKindEnum;
  data: Maybe<Profile.Response.FullAccountInformationWithCommunityProfile | Profile.Response.FullAccountInformationWithAuthorityProfile>;
  profile: Nullable<Profile.Response.ProfileBase>;
  action?: ReactNode;
  stats: MastheadStat[];
}) {
  const accountId = safeNullable(data?.account.id);
  const name = resolveAccountDisplayName(profile?.displayName, accountId);
  const initial = resolveAccountInitial(profile?.displayName, accountId);
  const avatarImageId = safeNullable(profile?.avatars?.[0]?.imageId);
  const backgrounds = safeArray(profile?.backgrounds)
    .slice()
    .sort((a, b) => a.sortOrder - b.sortOrder || a.imageId.localeCompare(b.imageId));
  const coverId = backgrounds[0]?.imageId;
  const roles = profileRoles(kind, data);
  const place = [safeStringable(profile?.city), safeStringable(profile?.country)].filter(Boolean).join(" · ");

  return (
    <Flex direction="column" width="100%" gap="5">
      {coverId ? (
        <Box className={styles.cover}>
          <img src={buildImageUrl(coverId)} alt="" className={styles.coverImage} />
        </Box>
      ) : null}
      <Flex align="end" justify="between" gap="4" wrap="wrap">
        <Flex align="end" gap="4" minWidth="0">
          <Flex flexShrink="0" className={styles.avatar}>
            <Avatar size="6" radius="none" src={avatarImageId ? buildImageUrl(avatarImageId) : undefined} fallback={initial} />
          </Flex>
          <Flex direction="column" gap="1" minWidth="0">
            <Text size="1" color="gray" className={styles.kind}>
              {kind}
            </Text>
            <Heading as="h2" size="7" className={styles.name}>
              {name}
            </Heading>
            <Flex gap="2" wrap="wrap" align="center">
              {roles.map((role) => (
                <Badge key={role} variant="outline" size="1">
                  {role}
                </Badge>
              ))}
              {place ? (
                <Text size="1" color="gray">
                  {place}
                </Text>
              ) : null}
            </Flex>
          </Flex>
        </Flex>
        {action}
      </Flex>
      <Grid columns={{ initial: "1", sm: "3" }} gap="3" width="100%">
        {stats.map((stat) => (
          <Section key={stat.label} size="1" py="4" className={styles.stat}>
            <Container size="2" px="4">
              <Flex direction="column" gap="1">
                <Text size="1" color="gray">
                  {stat.label}
                </Text>
                <Text size="6" weight="bold">
                  {stat.value}
                </Text>
              </Flex>
            </Container>
          </Section>
        ))}
      </Grid>
    </Flex>
  );
}
