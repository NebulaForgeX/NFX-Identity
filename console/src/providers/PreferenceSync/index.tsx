import { useEffect, useMemo } from "react";
import { useQuery } from "@tanstack/react-query";
import { useAuthRepository } from "nfx-ui/apis";
import { configurePreferenceSync, useApplyPreferenceOnLoad } from "nfx-ui/hooks";
import { AuthStore, parseServerPreference, toServerPreference, useAuthStore } from "nfx-ui/stores";
import { ProfileKindEnum } from "nfx-ui/enums";

export function PreferenceSync() {
  const auth = useAuthRepository();
  const isAuthValid = useAuthStore((s) => s.isAuthValid);
  const kind = useAuthStore((s) => s.currentProfileKind) || ProfileKindEnum.FORGER;
  const profileId = useAuthStore((s) => s.currentProfileId);

  const { data } = useQuery({
    queryKey: ["identity-preference", kind, profileId],
    enabled: isAuthValid && !!profileId,
    queryFn: () => auth.GetCurrentFullAccountInformationWithProfile(kind),
  });

  const serverPreference = useMemo(() => {
    const profile = data && "forgerProfile" in data ? data.forgerProfile : data && "authorityProfile" in data ? data.authorityProfile : null;
    if (!profile?.preference) return null;
    return parseServerPreference(profile.preference);
  }, [data]);

  useEffect(() => {
    if (!isAuthValid) {
      configurePreferenceSync(undefined);
      return;
    }
    configurePreferenceSync((state) => {
      const currentKind = AuthStore.getState().currentProfileKind || ProfileKindEnum.FORGER;
      return auth.UpdatePreference(currentKind, JSON.stringify(toServerPreference(state)));
    });
    return () => configurePreferenceSync(undefined);
  }, [auth, isAuthValid]);

  useApplyPreferenceOnLoad({
    isAuthValid,
    applyKey: profileId,
    serverPreference,
  });

  return null;
}

export default PreferenceSync;
