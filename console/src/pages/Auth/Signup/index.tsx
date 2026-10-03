import { AnimatedIcon, ArrowNarrowRightIcon } from "nfx-ui/icons";
import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { Button, Container, Flex, Grid, Heading, Link, Section, Text } from "@radix-ui/themes";
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

function Step({ index, label }: { index: string; label: string }) {
  return (
    <Section size="1" py="4" className={`${styles.step} js-step`}>
      <Container size="1" px="4" width="100%" maxWidth="100%">
        <Flex direction="column" gap="1">
          <Text as="span" size="1" className={styles.stepIndex}>
            {index}
          </Text>
          <Text as="span" className={styles.stepLabel}>
            {label}
          </Text>
        </Flex>
      </Container>
    </Section>
  );
}

export default function SignupPage() {
  const { t } = useTranslation("pages.Account.Signup");
  const form = useInitSignupForm(t);
  const signup = useSignupWithEmail();
  const sendCode = useSendVerificationCode({ successMsg: t("toasts.sendVerificationCodeSuccess") });
  const language = usePreferenceStore((s) => s.language);
  const rootRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) return;
      gsap.set(".js-step", { autoAlpha: 0, y: 12 });
      gsap.set(".js-copy", { autoAlpha: 0, y: 16 });
      gsap.set(".js-field", { autoAlpha: 0, y: 12 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-step", { autoAlpha: 1, y: 0, duration: 0.45, stagger: 0.08 }).to(".js-copy", { autoAlpha: 1, y: 0, duration: 0.5 }, "-=0.25").to(".js-field", { autoAlpha: 1, y: 0, duration: 0.45, stagger: 0.06 }, "-=0.2");
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
      <Flex ref={rootRef} direction="column" height="100%" minHeight="0" overflow="hidden" width="100%">
        <Grid columns="3" className={styles.legend}>
          <Step index="01" label={t("stepEmail")} />
          <Step index="02" label={t("stepCode")} />
          <Step index="03" label={t("stepPassword")} />
        </Grid>

        <Flex direction="column" flexGrow="1" minHeight="0" overflow="hidden">
          <Grid columns="2" width="100%" height="100%" className={styles.split}>
            <Section size="1" py="8" className={`${styles.copy} js-copy`}>
              <Container size="4" px="6" width="100%" maxWidth="100%">
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
              </Container>
            </Section>

            <Section size="1" py="8" className={styles.formPane}>
              <Container size="4" px="6" width="100%" maxWidth="100%">
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

                      <Grid columns="2" gap="4" className={`${styles.passRow} js-field`}>
                        <SignupPasswordController />
                        <SignupConfirmPasswordController />
                      </Grid>

                      <div className="js-field">
                        <SignupRememberController />
                      </div>

                      <div className="js-field">
                        <Button type="submit" size="3" loading={signup.isPending} className={styles.submit}>
                          {t("submit")}
                          <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                        </Button>
                      </div>
                    </form>
                  </Flex>
                </FormProvider>

                <Section size="1" py="4" className={`${styles.foot} js-field`}>
                  <Flex justify="between" align="center" gap="3">
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
                  </Flex>
                </Section>
              </Container>
            </Section>
          </Grid>
        </Flex>
      </Flex>
    </AuthChrome>
  );
}
