import { AnimatedIcon, ArrowNarrowRightIcon, EyeIcon, EyeOffIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Button, Checkbox, Flex, Heading, Link, Text, TextField } from "@radix-ui/themes";
import gsap from "gsap";
import { APP_NAME } from "nfx-ui/config";
import { AuthSignupPlatformEnum, LanguageEnum } from "nfx-ui/enums";
import { useSendVerificationCode, useSignupWithEmail } from "nfx-ui/hooks";
import { SignupFormData, useInitSignupForm } from "nfx-ui/schemas";
import { usePreferenceStore } from "nfx-ui/stores";
import { Controller, FormProvider, SubmitHandler } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { routerEventEmitter } from "@/events/router";
import { ROUTES } from "@/navigations";
import AuthChrome from "@/pages/Auth/shared/AuthChrome";

import styles from "./s.module.css";

gsap.registerPlugin(useGSAP);

export default function SignupPage() {
  const { t } = useTranslation("pages.Account.Signup");
  const form = useInitSignupForm();
  const signup = useSignupWithEmail();
  const sendCode = useSendVerificationCode();
  const language = usePreferenceStore((s) => s.language);
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirm, setShowConfirm] = useState(false);
  const rootRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) return;
      gsap.set(".js-step", { autoAlpha: 0, x: -24 });
      gsap.set(".js-copy", { autoAlpha: 0, y: 16 });
      gsap.set(".js-field", { autoAlpha: 0, x: 28 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-step", { autoAlpha: 1, x: 0, duration: 0.5, stagger: 0.08 }).to(".js-copy", { autoAlpha: 1, y: 0, duration: 0.5 }, "-=0.25").to(".js-field", { autoAlpha: 1, x: 0, duration: 0.5, stagger: 0.06 }, "-=0.2");
    },
    { scope: rootRef },
  );

  const onSubmit: SubmitHandler<SignupFormData> = async (data) => {
    await signup.mutateAsync({
      email: data.email,
      password: data.password,
      verificationCode: data.verificationCode,
      lang: language ?? LanguageEnum.EN,
      rememberMe: data.rememberMe ?? false,
      signupPlatform: AuthSignupPlatformEnum.NFXIDENTITY,
    });
    routerEventEmitter.navigate({ to: ROUTES.SELECT_PROFILE, replace: true });
  };

  const email = form.watch("email");

  return (
    <AuthChrome>
      <div ref={rootRef} className={styles.issuance}>
        <div className={styles.ticket} aria-hidden={false}>
          <div className={`${styles.step} js-step`}>
            <span className={styles.stepIndex}>01</span>
            <span className={styles.stepLabel}>{t("stepEmail")}</span>
          </div>
          <div className={`${styles.step} js-step`}>
            <span className={styles.stepIndex}>02</span>
            <span className={styles.stepLabel}>{t("stepCode")}</span>
          </div>
          <div className={`${styles.step} js-step`}>
            <span className={styles.stepIndex}>03</span>
            <span className={styles.stepLabel}>{t("stepPassword")}</span>
          </div>
        </div>

        <div className={styles.split}>
          <aside className={`${styles.copy} js-copy`}>
            <Flex direction="column" gap="4">
              <Text as="p" size="1" weight="bold" className={styles.kicker}>
                {t("pageEyebrow")}
              </Text>
              <Heading as="h1" size="8" className={styles.title}>
                {t("pageTitle", { name: APP_NAME })}
              </Heading>
              <Text as="p" size="3" className={styles.lede}>
                {t("pageSubtitle")}
              </Text>
              <Text as="p" size="2" color="gray">
                {t("issuanceHint")}
              </Text>
            </Flex>
          </aside>

          <div className={styles.form}>
            <FormProvider {...form}>
              <Flex asChild direction="column" gap="4">
                <form noValidate onSubmit={form.handleSubmit(onSubmit)}>
                  <div className="js-field">
                    <Controller
                      name="email"
                      control={form.control}
                      render={({ field, fieldState }) => (
                        <Flex direction="column" gap="1">
                          <Text as="label" size="2" weight="medium" htmlFor="signup-email">
                            {t("emailLabel")}
                          </Text>
                          <TextField.Root id="signup-email" size="3" type="email" autoComplete="email" placeholder={t("emailPlaceholder")} {...field} />
                          <Text size="1" color="gray">
                            {t("emailHint")}
                          </Text>
                          {fieldState.error ? (
                            <Text size="1" color="red">
                              {fieldState.error.message}
                            </Text>
                          ) : null}
                        </Flex>
                      )}
                    />
                  </div>

                  <div className="js-field">
                    <Flex direction="column" gap="1">
                      <Text as="label" size="2" weight="medium" htmlFor="signup-code">
                        {t("codeLabel")}
                      </Text>
                      <Flex gap="2">
                        <Controller
                          name="verificationCode"
                          control={form.control}
                          render={({ field }) => <TextField.Root id="signup-code" size="3" placeholder={t("codePlaceholder")} style={{ flex: 1 }} {...field} />}
                        />
                        <Button
                          type="button"
                          size="3"
                          variant="soft"
                          loading={sendCode.isPending}
                          disabled={!email}
                          onClick={() =>
                            email &&
                            sendCode.mutate({
                              email,
                              lang: language ?? LanguageEnum.EN,
                            })
                          }
                        >
                          {t("sendCode")}
                        </Button>
                      </Flex>
                    </Flex>
                  </div>

                  <div className={`${styles.passRow} js-field`}>
                    <Controller
                      name="password"
                      control={form.control}
                      render={({ field, fieldState }) => (
                        <Flex direction="column" gap="1">
                          <Text as="label" size="2" weight="medium" htmlFor="signup-password">
                            {t("passwordLabel")}
                          </Text>
                          <TextField.Root
                            id="signup-password"
                            size="3"
                            type={showPassword ? "text" : "password"}
                            autoComplete="new-password"
                            placeholder={t("passwordPlaceholder")}
                            {...field}
                          >
                            <TextField.Slot side="right">
                              <Button type="button" size="1" variant="soft" color="gray" onClick={() => setShowPassword((v) => !v)} aria-label={showPassword ? t("hidePassword") : t("showPassword")}>
                                <AnimatedIcon icon={showPassword ? EyeOffIcon : EyeIcon} size={14} />
                              </Button>
                            </TextField.Slot>
                          </TextField.Root>
                          {fieldState.error ? (
                            <Text size="1" color="red">
                              {fieldState.error.message}
                            </Text>
                          ) : null}
                        </Flex>
                      )}
                    />

                    <Controller
                      name="confirmPassword"
                      control={form.control}
                      render={({ field, fieldState }) => (
                        <Flex direction="column" gap="1">
                          <Text as="label" size="2" weight="medium" htmlFor="signup-confirm">
                            {t("confirmLabel")}
                          </Text>
                          <TextField.Root
                            id="signup-confirm"
                            size="3"
                            type={showConfirm ? "text" : "password"}
                            autoComplete="new-password"
                            placeholder={t("confirmPlaceholder")}
                            {...field}
                          >
                            <TextField.Slot side="right">
                              <Button type="button" size="1" variant="soft" color="gray" onClick={() => setShowConfirm((v) => !v)} aria-label={showConfirm ? t("hidePassword") : t("showPassword")}>
                                <AnimatedIcon icon={showConfirm ? EyeOffIcon : EyeIcon} size={14} />
                              </Button>
                            </TextField.Slot>
                          </TextField.Root>
                          {fieldState.error ? (
                            <Text size="1" color="red">
                              {fieldState.error.message}
                            </Text>
                          ) : null}
                        </Flex>
                      )}
                    />
                  </div>

                  <div className="js-field">
                    <Controller
                      name="rememberMe"
                      control={form.control}
                      render={({ field }) => (
                        <Flex asChild align="center" gap="2">
                          <Text as="label" size="2">
                            <Checkbox checked={!!field.value} onCheckedChange={(v) => field.onChange(v === true)} />
                            {t("rememberMe")}
                          </Text>
                        </Flex>
                      )}
                    />
                  </div>

                  <div className="js-field">
                    <Button type="submit" size="3" loading={signup.isPending} style={{ width: "100%" }}>
                      {t("submit")}
                      <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                    </Button>
                  </div>
                </form>
              </Flex>
            </FormProvider>

            <div className={`${styles.foot} js-field`}>
              <Text as="p" size="2" color="gray">
                {t("hasAccount")}
              </Text>
              <Link
                href={ROUTES.LOGIN}
                size="2"
                onClick={(e) => {
                  e.preventDefault();
                  routerEventEmitter.navigate({ to: ROUTES.LOGIN });
                }}
              >
                {t("signIn")}
              </Link>
            </div>
          </div>
        </div>
      </div>
    </AuthChrome>
  );
}
