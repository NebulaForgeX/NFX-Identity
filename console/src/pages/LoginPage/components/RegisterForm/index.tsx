import { memo } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useMutation } from "@tanstack/react-query";
import { AuthSignupPlatformEnum, LanguageEnum } from "nfx-ui/enums";
import { useAuthRepository } from "nfx-ui/apis";
import { AuthStore, ensureDeviceIdStorage } from "nfx-ui/stores";
import { showError, showSuccess } from "@/stores/modalStore";
import { useResendTimer } from "@/hooks/resendTimer";

import { createRegisterSchema, type RegisterFormValues } from "../../schemas/registerSchema";
import RegisterEmailController from "../../controllers/RegisterEmailController";
import RegisterPasswordController from "../../controllers/RegisterPasswordController";
import RegisterVerificationCodeController from "../../controllers/RegisterVerificationCodeController";

const RegisterForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const auth = useAuthRepository();
  const { timeLeft, canResend, startTimer } = useResendTimer();

  const methods = useForm<RegisterFormValues>({
    resolver: zodResolver(createRegisterSchema((key: string) => t(key))),
    mode: "onChange",
    defaultValues: { email: "", verificationCode: "", password: "" },
  });

  const email = methods.watch("email");
  const sendCode = useMutation({
    mutationFn: async () => auth.SendVerificationCode({ email, lang: LanguageEnum.ZH }),
  });
  const signup = useMutation({
    mutationFn: async (data: RegisterFormValues) => {
      const deviceId = await ensureDeviceIdStorage();
      const response = await auth.SignupWithEmail({
        email: data.email,
        password: data.password,
        verificationCode: data.verificationCode,
        lang: LanguageEnum.ZH,
        deviceId,
        signupPlatform: AuthSignupPlatformEnum.NFXIDENTITY,
      });
      AuthStore.getState().setTokens({ accessToken: response.accessToken, refreshToken: response.refreshToken });
      AuthStore.getState().setCurrentAccountId(response.accountId);
      sessionStorage.setItem("nfx-login-profiles", JSON.stringify(response.profiles ?? []));
      return response;
    },
  });

  return (
    <FormProvider {...methods}>
      <Flex direction="column" gap="4">
        <Flex gap="2" align="end">
          <RegisterEmailController />
          <Button
            type="button"
            variant="soft"
            size="3"
            disabled={!email || !canResend || sendCode.isPending}
            loading={sendCode.isPending}
            onClick={async () => {
              try {
                await sendCode.mutateAsync();
                startTimer(60);
                showSuccess(t("codeSentToEmail"));
              } catch {
                showError(t("registerFailed"));
              }
            }}
          >
            {canResend ? t("sendCode") : `${timeLeft}s`}
          </Button>
        </Flex>
        <RegisterVerificationCodeController />
        <RegisterPasswordController />
        <Button type="button" size="3" loading={signup.isPending} onClick={methods.handleSubmit((data) => signup.mutateAsync(data))}>
          {t("register")}
        </Button>
      </Flex>
    </FormProvider>
  );
});

RegisterForm.displayName = "RegisterForm";
export default RegisterForm;
