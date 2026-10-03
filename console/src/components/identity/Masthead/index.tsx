import type { ReactNode } from "react";

import { Avatar, Badge, Box, Container, Flex, Grid, Heading, Section, Text } from "@radix-ui/themes";
import { ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";

import { Surface } from "@/components";
import { buildImageUrl, profileRoles, resolveAccountDisplayName, resolveAccountInitial, safeArray, safeNullable, safeStringable } from "@/utils";

import styles from "./s.module.css";

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
    <Surface tone="hero" py="0" px="0">
      <Box className={styles.cover} data-empty={coverId ? undefined : "true"}>
        {coverId ? <img src={buildImageUrl(coverId)} alt="" className={styles.coverImage} /> : null}
      </Box>
      <Container size="4" width="100%" maxWidth="100%" px={{ initial: "4", sm: "6" }}>
        <Section size="1" pt="0" pb="6">
          <Flex direction="column" gap="5">
            <Flex align={{ initial: "start", sm: "end" }} justify="between" gap="4" wrap="wrap">
              <Flex align="end" gap="4" minWidth="0">
                <Section size="1" py="0" mt="-8" flexShrink="0">
                  <Avatar size="7" className={styles.avatar} src={avatarImageId ? buildImageUrl(avatarImageId) : undefined} fallback={initial} />
                </Section>
                <Flex direction="column" gap="2" minWidth="0">
                  <Flex gap="2" wrap="wrap" align="center">
                    <Badge size="1" variant="surface">
                      {kind}
                    </Badge>
                    {roles.map((role) => (
                      <Badge key={role} size="1" variant="outline" color="gray">
                        {role}
                      </Badge>
                    ))}
                  </Flex>
                  <Heading as="h2" size="7" weight="bold" className={styles.name}>
                    {name}
                  </Heading>
                  {place ? (
                    <Text size="2" color="gray">
                      {place}
                    </Text>
                  ) : null}
                </Flex>
              </Flex>
              {action}
            </Flex>
            <Grid columns={{ initial: "1", xs: "3" }} gap="3" width="100%">
              {stats.map((stat) => (
                <Surface key={stat.label} tone="inset" py="3" px="4">
                  <Flex direction="column" gap="1">
                    <Text size="1" color="gray" weight="medium" className={styles.statLabel}>
                      {stat.label}
                    </Text>
                    <Text as="p" size="6" weight="bold" className={styles.statValue}>
                      {stat.value}
                    </Text>
                  </Flex>
                </Surface>
              ))}
            </Grid>
          </Flex>
        </Section>
      </Container>
    </Surface>
  );
}
