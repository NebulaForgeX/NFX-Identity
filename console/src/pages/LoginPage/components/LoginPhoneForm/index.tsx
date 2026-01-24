import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { useLoginByPhone } from "@/hooks/useAuth";

import { createPhoneLoginSchema, type PhoneLoginFormValues } from "../../schemas/loginSchema";
import LoginPhoneController from "../LoginPhoneController";
import LoginPasswordController from "../LoginPasswordController";
import styles from "./styles.module.css";

const LoginPhoneForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const { mutateAsync: loginByPhone, isPending: isPhoneLoginPending } = useLoginByPhone();

  // 创建手机号登录 schema
  const phoneLoginSchema = createPhoneLoginSchema((key: string) => t(key));

  const methods = useForm<PhoneLoginFormValues>({
    resolver: zodResolver(phoneLoginSchema),
    mode: "onChange",
    defaultValues: {
      loginType: "phone",
      code: "86",
      phone: "",
      password: "",
    },
  });

  const { handleSubmit } = methods;

  const onSubmit = async (data: PhoneLoginFormValues) => {
    try {
      await loginByPhone({
        phone: data.phone,
        password: data.password,
        countryCode: data.code, // 后端仍使用 countryCode 字段名
      });
    } catch (error) {
      // 错误已在 useAuth hook 中处理
    }
  };

  return (
    <FormProvider {...methods}>
      <div className={styles.formContent}>
        <LoginPhoneController />
        <LoginPasswordController />

        <button type="button" className={styles.submitBtn} onClick={handleSubmit(onSubmit)} disabled={isPhoneLoginPending}>
          {isPhoneLoginPending ? t("loggingIn") : t("login")}
        </button>

        <span className={styles.bottomText}>
          {t("noAccount")}{" "}
          <label htmlFor="register_toggle" className={styles.switch}>
            {t("registerNow")}
          </label>
        </span>
      </div>
    </FormProvider>
  );
});

LoginPhoneForm.displayName = "LoginPhoneForm";

export default LoginPhoneForm;
