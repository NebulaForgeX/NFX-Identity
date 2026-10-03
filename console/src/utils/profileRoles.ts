import { ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";
import { safeArray } from "nfx-ui/utils";

export function profileRoles(kind: ProfileKindEnum, data: Maybe<Profile.Response.FullAccountInformationWithCommunityProfile | Profile.Response.FullAccountInformationWithAuthorityProfile>): string[] {
  if (!data) return [];
  if (kind === ProfileKindEnum.AUTHORITY) {
    return safeArray((data as Profile.Response.FullAccountInformationWithAuthorityProfile).authorityProfile?.authorityRoles);
  }
  return safeArray((data as Profile.Response.FullAccountInformationWithCommunityProfile).communityProfile?.forgerRoles);
}
