import { useEffect, useRef, useState } from "react";
import { Flex, Heading, Text } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useNavigate, useSearchParams } from "react-router";
import { useTranslation } from "react-i18next";
import { AuthSignupPlatformEnum } from "nfx-ui/enums";
import { useAuthRepository } from "nfx-ui/apis";
import { AuthStore, ensureDeviceIdStorage } from "nfx-ui/stores";

import AuthShell from "./AuthShell";
import { ROUTES } from "@/navigations";

export default function GitHubCallbackPage() {
  const { t } = useTranslation("LoginPage");
  const auth = useAuthRepository();
  const navigate = useNavigate();
  const [params] = useSearchParams();
  const [error, setError] = useState<string | null>(null);

  const run = useMutation({
    mutationFn: async () => {
      const code = params.get("code")?.trim() ?? "";
      const state = params.get("state")?.trim() ?? "";
      if (!code || !state) {
        throw new Error(t("githubMissingCode"));
      }
      const accessToken = AuthStore.getState().accessToken;
      if (accessToken) {
        await auth.LinkGitHub({ code, state });
        navigate(ROUTES.PROFILE, { replace: true });
        return;
      }
      const deviceId = await ensureDeviceIdStorage();
      const response = await auth.LoginWithGitHub({
        code,
        state,
        deviceId,
        signupPlatform: AuthSignupPlatformEnum.NFXIDENTITY,
      });
      AuthStore.getState().setTokens({ accessToken: response.accessToken, refreshToken: response.refreshToken });
      AuthStore.getState().setCurrentAccountId(response.accountId);
      sessionStorage.setItem("nfx-login-profiles", JSON.stringify(response.profiles ?? []));
      navigate(ROUTES.SELECT_PROFILE, { replace: true });
    },
    onError: (err: Error) => {
      setError(err.message || t("githubCallbackFailed"));
    },
  });

  const started = useRef(false);

  useEffect(() => {
    if (started.current) return;
    started.current = true;
    void run.mutateAsync();
    // GitHub authorization code is single-use.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AuthShell brandEyebrow="NFX Identity" brandTitle={t("githubCallback")} heroFooter={t("githubCallbackHint")}>
      <Flex direction="column" gap="3">
        <Heading as="h2" size="6">
          {t("githubCallback")}
        </Heading>
        <Text as="p" size="2" color={error ? "red" : "gray"}>
          {error || t("githubCallbackHint")}
        </Text>
        {error ? (
          <Text as="p" size="2">
            <a href={AuthStore.getState().accessToken ? ROUTES.PROFILE : ROUTES.LOGIN}>{t("signInNow")}</a>
          </Text>
        ) : null}
      </Flex>
    </AuthShell>
  );
}
