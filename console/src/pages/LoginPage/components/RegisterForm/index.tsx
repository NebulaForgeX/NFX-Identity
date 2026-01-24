import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { useSendVerificationCode, useSignup } from "@/hooks/useAuth";
import { useResendTimer } from "@/hooks/useResendTimer";
import { showError, showSuccess } from "@/stores/modalStore";

import { createRegisterSchema, type RegisterFormValues } from "../../schemas/registerSchema";
import RegisterEmailController from "../RegisterEmailController";
import RegisterPasswordController from "../RegisterPasswordController";
import RegisterVerificationCodeController from "../RegisterVerificationCodeController";
import styles from "./styles.module.css";

const RegisterForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const { mutateAsync: signup, isPending: isSigningUp } = useSignup();
  const { mutateAsync: sendCode, isPending: isSendingCode } = useSendVerificationCode();
  const { timeLeft, canResend, startTimer } = useResendTimer();

  // 创建带翻译的 schema
  const registerSchema = createRegisterSchema((key: string) => t(key));

  const methods = useForm<RegisterFormValues>({
    resolver: zodResolver(registerSchema),
    mode: "onChange",
    defaultValues: {
      email: "",
      verificationCode: "",
      password: "",
    },
  });

  const { handleSubmit, watch } = methods;
  const email = watch("email");

  const handleSendCode = async () => {
    if (!email || !canResend || isSendingCode) return;
    try {
      await sendCode({ email });
      startTimer(60);
      showSuccess(t("codeSentToEmail"));
    } catch (error) {
      showError(t("registerFailed"));
    }
  };

  const onSubmit = async (data: RegisterFormValues) => {
    try {
      await signup({
        email: data.email,
        password: data.password,
        verificationCode: data.verificationCode,
      });
    } catch (error) {
      showError(t("registerFailed"));
    }
  };

  const canSendCode = email && canResend && !isSendingCode;

  return (
    <div className={styles.formWrapper}>
      <div className={styles.form}>
        <span className={styles.title}>{t("register")}</span>

        <FormProvider {...methods}>
          <div className={styles.formContent}>
            {/* Email with Send Code Button */}
            <div className={styles.formControl}>
              <RegisterEmailController />
              <button
                type="button"
                className={styles.sendCodeBtn}
                onClick={handleSendCode}
                disabled={!canSendCode}
              >
                {canResend ? t("sendCode") : `${timeLeft}s`}
              </button>
            </div>

            {/* Verification Code */}
            <RegisterVerificationCodeController />

            {/* Password */}
            <RegisterPasswordController />

            <button type="button" className={styles.submitBtn} onClick={handleSubmit(onSubmit)} disabled={isSigningUp}>
              {isSigningUp ? t("registering") : t("register")}
            </button>

            <span className={styles.bottomText}>
              {t("hasAccount")}{" "}
              <label htmlFor="register_toggle" className={styles.switch}>
                {t("signInNow")}
              </label>
            </span>
          </div>
        </FormProvider>
      </div>
    </div>
  );
});

RegisterForm.displayName = "RegisterForm";

export default RegisterForm;
