import { Avatar, Badge, Box, Flex, Heading, Text } from "@radix-ui/themes";
import { ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";

import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeArray, safeNullable, safeStringable } from "@/utils";

import styles from "./masthead.module.css";

export function profileRoles(kind: ProfileKindEnum, data: Maybe<Profile.Response.FullAccountInformationWithForgerProfile | Profile.Response.FullAccountInformationWithAuthorityProfile>): string[] {
  if (!data) return [];
  if (kind === ProfileKindEnum.AUTHORITY) {
    return safeArray((data as Profile.Response.FullAccountInformationWithAuthorityProfile).authorityProfile?.authorityRoles);
  }
  return safeArray((data as Profile.Response.FullAccountInformationWithForgerProfile).forgerProfile?.forgerRoles);
}

export default function Masthead({
  kind,
  data,
  profile,
  action,
}: {
  kind: ProfileKindEnum;
  data: Maybe<Profile.Response.FullAccountInformationWithForgerProfile | Profile.Response.FullAccountInformationWithAuthorityProfile>;
  profile: Nullable<Profile.Response.ProfileBase>;
  action?: React.ReactNode;
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
    <Box className={styles.wrap}>
      <Box className={styles.cover}>
        {coverId ? <img src={buildImageUrl(coverId)} alt="" className={styles.coverImage} /> : null}
      </Box>
      <Flex align="end" justify="between" gap="4" wrap="wrap" className={styles.body}>
        <Flex align="end" gap="4" minWidth="0">
          <Avatar size="6" radius="none" className={styles.avatar} src={avatarImageId ? buildImageUrl(avatarImageId) : undefined} fallback={initial} />
          <Flex direction="column" gap="1" minWidth="0" pb="1">
            <Text size="1" color="gray" className={styles.kind}>
              {kind}
            </Text>
            <Heading as="h1" size="7" className={styles.name}>
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
        {action ? <Box pb="1">{action}</Box> : null}
      </Flex>
    </Box>
  );
}
