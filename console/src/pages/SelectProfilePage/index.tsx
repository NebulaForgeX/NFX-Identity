import { Button, Flex, Heading, Text } from "@radix-ui/themes";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useAuthRepository } from "nfx-ui/apis";
import { AuthStore, ensureDeviceIdStorage } from "nfx-ui/stores";
import { AuthAuthorityRoleEnum, LanguageEnum, ProfileKind, ProfileKindEnum } from "nfx-ui/enums";
import type { Login, Profile } from "nfx-ui/types";

import AuthShell from "@/pages/LoginPage/AuthShell";

export default function SelectProfilePage() {
  const auth = useAuthRepository();
  const forgers = useQuery({
    queryKey: ["me-forger-profiles"],
    queryFn: () => auth.ListProfiles(ProfileKindEnum.FORGER, { limit: 50, offset: 0 }),
  });
  const authorities = useQuery({
    queryKey: ["me-authority-profiles"],
    queryFn: () => auth.ListProfiles(ProfileKindEnum.AUTHORITY, { limit: 50, offset: 0 }),
  });

  const profiles: Login.ProfileItem[] = [
    ...(forgers.data?.items ?? []).map((item: Profile.Response.ForgerProfileItem) => ({
      profileId: item.profileId,
      kind: ProfileKindEnum.FORGER,
      roles: item.forgerRoles ?? [],
      displayName: item.displayName,
      avatarImageId: item.avatarImageId,
      city: item.city,
      country: item.country,
    })),
    ...(authorities.data?.items ?? []).map((item: Profile.Response.AuthorityProfileItem) => ({
      profileId: item.profileId,
      kind: ProfileKindEnum.AUTHORITY,
      roles: item.authorityRoles ?? [],
      displayName: item.displayName,
      avatarImageId: item.avatarImageId,
      city: item.city,
      country: item.country,
    })),
  ];

  const select = useMutation({
    mutationFn: async (profile: Login.ProfileItem) => {
      const deviceId = await ensureDeviceIdStorage();
      const out = await auth.SelectProfile({ profileId: profile.profileId, kind: ProfileKind(profile.kind), deviceId });
      AuthStore.getState().setTokens({ accessToken: out.accessToken, refreshToken: out.refreshToken });
      AuthStore.getState().setCurrentAccountId(out.accountId);
      AuthStore.getState().setCurrentProfileId(out.profileId);
      AuthStore.getState().setCurrentProfileKind(ProfileKind(profile.kind));
      AuthStore.getState().setIsAuthValid(true);
    },
  });

  const create = useMutation({
    mutationFn: async () => {
      const created = await auth.CreateForgerProfile({ displayName: "Forger", profileLanguage: LanguageEnum.ZH });
      const deviceId = await ensureDeviceIdStorage();
      const out = await auth.SelectProfile({ profileId: created.profileId, kind: ProfileKind(ProfileKindEnum.FORGER), deviceId });
      AuthStore.getState().setTokens({ accessToken: out.accessToken, refreshToken: out.refreshToken });
      AuthStore.getState().setCurrentAccountId(out.accountId);
      AuthStore.getState().setCurrentProfileId(out.profileId);
      AuthStore.getState().setCurrentProfileKind(ProfileKind(ProfileKindEnum.FORGER));
      AuthStore.getState().setIsAuthValid(true);
    },
  });

  const createAuthority = useMutation({
    mutationFn: async () => {
      const created = await auth.CreateAuthorityProfile({ displayName: "Administrator", profileLanguage: LanguageEnum.ZH });
      const deviceId = await ensureDeviceIdStorage();
      const out = await auth.SelectProfile({ profileId: created.profileId, kind: ProfileKind(ProfileKindEnum.AUTHORITY), deviceId });
      AuthStore.getState().setTokens({ accessToken: out.accessToken, refreshToken: out.refreshToken });
      AuthStore.getState().setCurrentAccountId(out.accountId);
      AuthStore.getState().setCurrentProfileId(out.profileId);
      AuthStore.getState().setCurrentProfileKind(ProfileKind(ProfileKindEnum.AUTHORITY));
      AuthStore.getState().setIsAuthValid(true);
    },
  });

  const canCreateAuthority = (authorities.data?.items ?? []).some((item: Profile.Response.AuthorityProfileItem) =>
    (item.authorityRoles ?? []).includes(AuthAuthorityRoleEnum.OWNER),
  );
  const error = (select.error || create.error || createAuthority.error) as Error | null;

  return (
    <AuthShell brandEyebrow="NFX Identity" brandTitle="Select a profile" heroFooter="Identity issues one token per profile. Empty accounts create a Forger here.">
      <Flex direction="column" gap="4">
        <Heading as="h2" size="5">
          Profiles
        </Heading>
        <Text as="p" size="2">
          Choose a Forger or Authority profile, or create the first Forger on this account.
        </Text>
        {profiles.map((profile) => (
          <Flex key={profile.profileId} gap="2">
            <Button style={{ flex: 1 }} variant="soft" onClick={() => select.mutate(profile)} loading={select.isPending}>
              {profile.displayName || profile.profileId} ({profile.kind})
            </Button>
            {profiles.length > 1 && !(profile.roles ?? []).includes(AuthAuthorityRoleEnum.OWNER) ? (
              <Button
                color="red"
                variant="soft"
                onClick={() => {
                  void auth.DeleteProfile(ProfileKind(profile.kind), profile.profileId).then(() => {
                    void forgers.refetch();
                    void authorities.refetch();
                  });
                }}
              >
                ×
              </Button>
            ) : null}
          </Flex>
        ))}
        {profiles.length === 0 ? (
          <Button onClick={() => create.mutate()} loading={create.isPending}>
            Create Forger profile
          </Button>
        ) : null}
        {canCreateAuthority ? (
          <Button variant="outline" onClick={() => createAuthority.mutate()} loading={createAuthority.isPending}>
            Create Authority profile
          </Button>
        ) : null}
        {error ? (
          <Text size="2" color="red">
            {error.message}
          </Text>
        ) : null}
      </Flex>
    </AuthShell>
  );
}
