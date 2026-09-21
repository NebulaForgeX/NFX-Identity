import { AnimatedIcon, ArrowNarrowRightIcon } from "nfx-ui/icons";
import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { Button, Flex, Heading, Link, Text } from "@radix-ui/themes";
import gsap from "gsap";
import { APP_NAME } from "nfx-ui/config";
import { AuthSignupPlatformEnum, LanguageEnum } from "nfx-ui/enums";
import { useSendVerificationCode, useSignupWithEmail } from "nfx-ui/hooks";
import { SignupFormData, useInitSignupForm } from "nfx-ui/schemas";
import { usePreferenceStore } from "nfx-ui/stores";
import { FormProvider, SubmitHandler } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { routerEventEmitter } from "@/events/router";
import {
  SignupConfirmPasswordController,
  SignupEmailController,
  SignupPasswordController,
  SignupRememberController,
  SignupVerificationCodeController,
} from "@/features/account";
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
            <div className={styles.stepPx}>
              <div className={styles.stepPy}>
                <div className={styles.stepStack}>
            <span className={styles.stepIndex}>01</span>
            <span className={styles.stepLabel}>{t("stepEmail")}</span>
                </div>
              </div>
            </div>
          </div>
          <div className={`${styles.step} js-step`}>
            <div className={styles.stepPx}>
              <div className={styles.stepPy}>
                <div className={styles.stepStack}>
            <span className={styles.stepIndex}>02</span>
            <span className={styles.stepLabel}>{t("stepCode")}</span>
                </div>
              </div>
            </div>
          </div>
          <div className={`${styles.step} js-step`}>
            <div className={styles.stepPx}>
              <div className={styles.stepPy}>
                <div className={styles.stepStack}>
            <span className={styles.stepIndex}>03</span>
            <span className={styles.stepLabel}>{t("stepPassword")}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div className={styles.split}>
          <aside className={`${styles.copy} js-copy`}>
            <div className={styles.copyPx}>
              <div className={styles.copyPy}>
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
              </div>
            </div>
          </aside>

          <div className={styles.form}>
            <div className={styles.formPx}>
              <div className={styles.formPy}>
            <FormProvider {...form}>
              <Flex asChild direction="column" gap="4">
                <form noValidate onSubmit={form.handleSubmit(onSubmit)}>
                  <div className="js-field">
                    <SignupEmailController helperText={t("emailHint")} />
                  </div>

                  <div className="js-field">
                    <Flex direction="column" gap="2">
                      <SignupVerificationCodeController />
                      <Button
                        type="button"
                        size="3"
                        variant="outline"
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
                  </div>

                  <div className={`${styles.passRow} js-field`}>
                    <SignupPasswordController />
                    <SignupConfirmPasswordController />
                  </div>

                  <div className="js-field">
                    <SignupRememberController />
                  </div>

                  <div className="js-field">
                    <Button type="submit" size="3" loading={signup.isPending} className={styles.fullWidth}>
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
        </div>
      </div>
    </AuthChrome>
  );
}
