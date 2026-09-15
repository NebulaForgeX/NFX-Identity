import { memo, useState } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useAuthRepository } from "nfx-ui/apis";
import { AuthStore, ensureDeviceIdStorage } from "nfx-ui/stores";
import { useTranslation } from "react-i18next";

import LoginEmailController from "../../controllers/LoginEmailController";
import LoginPasswordController from "../../controllers/LoginPasswordController";
import LoginTypeSwitch from "../LoginTypeSwitch";
import LoginPhoneForm from "../LoginPhoneForm";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { createEmailLoginSchema, type EmailLoginFormValues } from "../../schemas/loginSchema";

const LoginForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const auth = useAuthRepository();
  const [loginType, setLoginType] = useState<"email" | "phone">("email");

  const methods = useForm<EmailLoginFormValues>({
    resolver: zodResolver(createEmailLoginSchema((key: string) => t(key))),
    mode: "onChange",
    defaultValues: { loginType: "email", email: "", password: "" },
  });

  const login = useMutation({
    mutationFn: async (data: EmailLoginFormValues) => {
      const deviceId = await ensureDeviceIdStorage();
      const response = await auth.LoginWithEmail({ email: data.email, password: data.password, deviceId });
      AuthStore.getState().setTokens({ accessToken: response.accessToken, refreshToken: response.refreshToken });
      AuthStore.getState().setCurrentAccountId(response.accountId);
      sessionStorage.setItem("nfx-login-profiles", JSON.stringify(response.profiles ?? []));
      return response;
    },
  });

  const github = useMutation({
    mutationFn: async () => auth.GetGitHubAuthorizeUrl(),
    onSuccess: (data) => {
      if (data.authorizeUrl) window.location.href = data.authorizeUrl;
    },
  });

  if (loginType === "phone") {
    return (
      <Flex direction="column" gap="4">
        <LoginTypeSwitch loginType={loginType} onLoginTypeChange={setLoginType} />
        <LoginPhoneForm />
        <Button type="button" variant="outline" onClick={() => github.mutate()}>
          GitHub
        </Button>
      </Flex>
    );
  }

  return (
    <FormProvider {...methods}>
      <Flex direction="column" gap="4">
        <LoginTypeSwitch loginType={loginType} onLoginTypeChange={setLoginType} />
        <LoginEmailController />
        <LoginPasswordController />
        <Button type="button" size="3" loading={login.isPending} onClick={methods.handleSubmit((data) => login.mutateAsync(data))}>
          {t("login")}
        </Button>
        <Button type="button" variant="outline" onClick={() => github.mutate()}>
          GitHub
        </Button>
      </Flex>
    </FormProvider>
  );
});

LoginForm.displayName = "LoginForm";
export default LoginForm;
