import { memo } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useMutation } from "@tanstack/react-query";
import { useAuthRepository } from "nfx-ui/apis";
import { AuthStore, ensureDeviceIdStorage } from "nfx-ui/stores";

import { createPhoneLoginSchema, type PhoneLoginFormValues } from "../../schemas/loginSchema";
import LoginPhoneController from "../../controllers/LoginPhoneController";
import LoginPasswordController from "../../controllers/LoginPasswordController";

const LoginPhoneForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const auth = useAuthRepository();
  const login = useMutation({
    mutationFn: async (data: PhoneLoginFormValues) => {
      const deviceId = await ensureDeviceIdStorage();
      const phone = data.code ? `+${data.code}${data.phone}` : data.phone;
      const response = await auth.LoginWithPhone({ phone, password: data.password, deviceId });
      AuthStore.getState().setTokens({ accessToken: response.accessToken, refreshToken: response.refreshToken });
      AuthStore.getState().setCurrentAccountId(response.accountId);
      sessionStorage.setItem("nfx-login-profiles", JSON.stringify(response.profiles ?? []));
      return response;
    },
  });

  const methods = useForm<PhoneLoginFormValues>({
    resolver: zodResolver(createPhoneLoginSchema((key: string) => t(key))),
    mode: "onChange",
    defaultValues: { loginType: "phone", code: "86", phone: "", password: "" },
  });

  return (
    <FormProvider {...methods}>
      <Flex direction="column" gap="4">
        <LoginPhoneController />
        <LoginPasswordController />
        <Button type="button" size="3" loading={login.isPending} onClick={methods.handleSubmit((data) => login.mutateAsync(data))}>
          {t("login")}
        </Button>
      </Flex>
    </FormProvider>
  );
});

LoginPhoneForm.displayName = "LoginPhoneForm";
export default LoginPhoneForm;
