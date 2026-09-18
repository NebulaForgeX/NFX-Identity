import { useEffect, useRef } from "react";
import { Flex, Spinner, Text } from "@radix-ui/themes";
import { AuthSignupPlatformEnum } from "nfx-ui/enums";
import { authEventEmitter, authEvents, systemEventEmitter } from "nfx-ui/events";
import { useLinkGitHub, useLoginWithGitHub, useSelectProfile } from "nfx-ui/hooks";
import { AuthStore, hasSelectedProfile, useAuthStore } from "nfx-ui/stores";
import { useTranslation } from "react-i18next";
import { useSearchParams } from "react-router";

import { routerEventEmitter } from "@/events/router";
import { ROUTES } from "@/navigations";
import { safeArray } from "@/utils";

export default function GitHubCallbackPage() {
  const { t } = useTranslation("pages.Account.Login");
  const [params] = useSearchParams();
  const login = useLoginWithGitHub();
  const link = useLinkGitHub();
  const selectProfile = useSelectProfile();
  const isAuthValid = useAuthStore((s) => s.isAuthValid);
  const currentProfileId = useAuthStore((s) => s.currentProfileId);
  const started = useRef(false);

  useEffect(() => {
    if (started.current) return;
    started.current = true;
    const code = params.get("code") ?? "";
    const state = params.get("state") ?? "";
    if (!code || !state) {
      systemEventEmitter.showError(t("github.missingParams"));
      routerEventEmitter.navigate({ to: ROUTES.LOGIN, replace: true });
      return;
    }
    void (async () => {
      try {
        if (isAuthValid && hasSelectedProfile(currentProfileId)) {
          await link.mutateAsync({ code, state });
          authEventEmitter.emit(authEvents.UPDATE_ACCOUNT_SUCCESS, AuthStore.getState().currentAccountId);
          routerEventEmitter.navigate({ to: ROUTES.USER_PROFILE_IDENTITIES, replace: true });
          return;
        }
        const result = await login.mutateAsync({
          code,
          state,
          signupPlatform: AuthSignupPlatformEnum.NFXIDENTITY,
        });
        const profiles = safeArray(result?.profiles);
        if (profiles.length === 1 && profiles[0]) {
          await selectProfile.mutateAsync({
            profileId: profiles[0].profileId,
            kind: profiles[0].kind,
          });
          routerEventEmitter.navigate({ to: ROUTES.USER_OVERVIEW, replace: true });
          return;
        }
        routerEventEmitter.navigate({ to: ROUTES.LOGIN, replace: true });
      } catch {
        routerEventEmitter.navigate({ to: isAuthValid && hasSelectedProfile(currentProfileId) ? ROUTES.USER_PROFILE_IDENTITIES : ROUTES.LOGIN, replace: true });
      }
    })();
  }, [currentProfileId, isAuthValid, link, login, params, selectProfile, t]);

  return (
    <Flex align="center" justify="center" minHeight="100dvh">
      <Flex direction="column" align="center" gap="3">
        <Spinner size="3" />
        <Text size="2" color="gray">
          {t("github.completing")}
        </Text>
      </Flex>
    </Flex>
  );
}
