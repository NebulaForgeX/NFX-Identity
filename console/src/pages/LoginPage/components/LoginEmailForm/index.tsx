import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { useLoginByEmail } from "@/hooks/useAuth";

import { createEmailLoginSchema, type EmailLoginFormValues } from "../../schemas/loginSchema";
import LoginEmailController from "../LoginEmailController";
import LoginPasswordController from "../LoginPasswordController";
import styles from "./styles.module.css";

const LoginEmailForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const { mutateAsync: loginByEmail, isPending: isEmailLoginPending } = useLoginByEmail();

  // 创建邮箱登录 schema
  const emailLoginSchema = createEmailLoginSchema((key: string) => t(key));

  const methods = useForm<EmailLoginFormValues>({
    resolver: zodResolver(emailLoginSchema),
    mode: "onChange",
    defaultValues: {
      loginType: "email",
      email: "",
      password: "",
    },
  });

  const { handleSubmit } = methods;

  const onSubmit = async (data: EmailLoginFormValues) => {
    try {
      await loginByEmail({ email: data.email, password: data.password });
    } catch (error) {
      // 错误已在 useAuth hook 中处理
    }
  };

  return (
    <FormProvider {...methods}>
      <div className={styles.formContent}>
        <LoginEmailController />
        <LoginPasswordController />

        <button type="button" className={styles.submitBtn} onClick={handleSubmit(onSubmit)} disabled={isEmailLoginPending}>
          {isEmailLoginPending ? t("loggingIn") : t("login")}
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

LoginEmailForm.displayName = "LoginEmailForm";

export default LoginEmailForm;
