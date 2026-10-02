import { AnimatedIcon, ArrowNarrowRightIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Button, Container, Flex, Heading, Link, Section, Tabs, Text } from "@radix-ui/themes";
import gsap from "gsap";
import { APP_NAME } from "nfx-ui/config";
import { useLoginWithEmail, useLoginWithPhone } from "nfx-ui/hooks";
import { LoginFormData, LoginWithPhoneFormData, useInitLoginForm, useInitLoginWithPhoneForm } from "nfx-ui/schemas";
import { FormProvider, SubmitHandler } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { routerEventEmitter } from "@/events/router";
import { LoginEmailController, LoginPasswordController, LoginPhoneController, LoginRememberController } from "@/features/account";
import { ROUTES } from "@/navigations";
import AuthChrome from "@/pages/Auth/shared/AuthChrome";
import { safeOr } from "@/utils";

import styles from "./s.module.css";

gsap.registerPlugin(useGSAP);

export default function LoginPage() {
  const { t } = useTranslation("pages.Account.Login");
  const emailForm = useInitLoginForm();
  const phoneForm = useInitLoginWithPhoneForm();
  const loginEmail = useLoginWithEmail();
  const loginPhone = useLoginWithPhone();
  const [channel, setChannel] = useState<"email" | "phone">("email");
  const rootRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) return;
      gsap.set(".js-gate", { autoAlpha: 0, y: 16 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-gate", { autoAlpha: 1, y: 0, duration: 0.5, stagger: 0.08 });
    },
    { scope: rootRef },
  );

  const goSelect = () => routerEventEmitter.navigate({ to: ROUTES.SELECT_PROFILE, replace: true });

  const onEmail: SubmitHandler<LoginFormData> = async (data) => {
    await loginEmail.mutateAsync({
      email: data.email,
      password: data.password,
      rememberMe: safeOr(data.rememberMe, false),
    });
    goSelect();
  };

  const onPhone: SubmitHandler<LoginWithPhoneFormData> = async (data) => {
    await loginPhone.mutateAsync({
      phone: data.phone,
      password: data.password,
      rememberMe: safeOr(data.rememberMe, false),
    });
    goSelect();
  };

  return (
    <AuthChrome>
      <Flex ref={rootRef} direction="column" height="100%" minHeight="0" overflow="hidden" width="100%">
        <Flex direction="column" flexGrow="1" minHeight="0" overflow="auto">
          <Container size="2" mx="auto" px="6" width="100%" maxWidth="32rem">
            <Section size="1" py="8">
              <Flex direction="column" gap="6">
                <Flex direction="column" gap="3" className="js-gate">
                  <Text as="p" size="1" weight="bold" className={styles.index}>
                    01 — {t("form.welcomeBack")}
                  </Text>
                  <Heading as="h1" size="8" className={styles.title}>
                    {t("form.title", { name: APP_NAME })}
                  </Heading>
                  <Text as="p" size="2" className={styles.lede}>
                    {t("form.subtitle")}
                  </Text>
                </Flex>

                <Tabs.Root value={channel} onValueChange={(value) => setChannel(value as "email" | "phone")} className="js-gate">
                  <Tabs.List>
                    <Tabs.Trigger value="email">{t("form.channelEmail")}</Tabs.Trigger>
                    <Tabs.Trigger value="phone">{t("form.channelPhone")}</Tabs.Trigger>
                  </Tabs.List>

                  <Tabs.Content value="email">
                    <FormProvider {...emailForm}>
                      <Flex asChild direction="column" gap="4">
                        <form noValidate onSubmit={emailForm.handleSubmit(onEmail)}>
                          <LoginEmailController />
                          <LoginPasswordController />
                          <LoginRememberController />
                          <Button type="submit" size="3" loading={loginEmail.isPending} className={styles.submit}>
                            {t("form.submit")}
                            <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                          </Button>
                        </form>
                      </Flex>
                    </FormProvider>
                  </Tabs.Content>

                  <Tabs.Content value="phone">
                    <FormProvider {...phoneForm}>
                      <Flex asChild direction="column" gap="4">
                        <form noValidate onSubmit={phoneForm.handleSubmit(onPhone)}>
                          <LoginPhoneController />
                          <LoginPasswordController />
                          <LoginRememberController />
                          <Button type="submit" size="3" loading={loginPhone.isPending} className={styles.submit}>
                            {t("form.submit")}
                            <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                          </Button>
                        </form>
                      </Flex>
                    </FormProvider>
                  </Tabs.Content>
                </Tabs.Root>

                <Section size="1" py="4" className={`${styles.foot} js-gate`}>
                  <Flex justify="between" align="center" gap="4">
                    <Text as="p" size="2" color="gray">
                      {t("promo.newTo", { name: APP_NAME })}
                    </Text>
                    <Link
                      href={ROUTES.SIGNUP}
                      size="2"
                      onClick={(e) => {
                        e.preventDefault();
                        routerEventEmitter.navigate({ to: ROUTES.SIGNUP });
                      }}
                    >
                      {t("promo.createAccount")}
                    </Link>
                  </Flex>
                </Section>
              </Flex>
            </Section>
          </Container>
        </Flex>
      </Flex>
    </AuthChrome>
  );
}
