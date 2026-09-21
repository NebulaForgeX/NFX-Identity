import { AnimatedIcon, ArrowNarrowRightIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Box, Button, Flex, Heading, Link, Tabs, Text } from "@radix-ui/themes";
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
      gsap.set(".js-gate-copy", { autoAlpha: 0, y: 16 });
      gsap.set(".js-gate-slab", { autoAlpha: 0, y: 28 });
      gsap.set(".js-gate-fields", { autoAlpha: 0, y: 18 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-gate-copy", { autoAlpha: 1, y: 0, duration: 0.55, stagger: 0.06 }).to(".js-gate-slab", { autoAlpha: 1, y: 0, duration: 0.5, stagger: 0.08 }, "-=0.25").to(".js-gate-fields", { autoAlpha: 1, y: 0, duration: 0.5 }, "-=0.2");
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
      <Box ref={rootRef} className={styles.gate}>
        <Box pb="6">
        <Flex direction="column" gap="3">
          <Text as="p" size="1" weight="bold" className={`${styles.index} js-gate-copy`}>
            01 — {t("form.welcomeBack")}
          </Text>
          <Heading as="h1" size="8" className={`${styles.title} js-gate-copy`}>
            {t("form.title", { name: APP_NAME })}
          </Heading>
          <Text as="p" size="2" className={`${styles.lede} js-gate-copy`}>
            {t("form.subtitle")}
          </Text>
        </Flex>
        </Box>

        <Tabs.Root value={channel} onValueChange={(value) => setChannel(value as "email" | "phone")}>
          <Tabs.List className={`${styles.slabs} js-gate-slab`}>
            <Tabs.Trigger value="email" className={styles.slab}>
              <Box className={styles.slabPx}>
                <Box className={styles.slabPy}>
                  <Box className={styles.slabStack}>
              <Text as="span" size="1" className={styles.slabIndex}>
                01
              </Text>
              <Text as="span" className={styles.slabLabel}>
                {t("form.channelEmail")}
              </Text>
                  </Box>
                </Box>
              </Box>
            </Tabs.Trigger>
            <Tabs.Trigger value="phone" className={styles.slab}>
              <Box className={styles.slabPx}>
                <Box className={styles.slabPy}>
                  <Box className={styles.slabStack}>
              <Text as="span" size="1" className={styles.slabIndex}>
                02
              </Text>
              <Text as="span" className={styles.slabLabel}>
                {t("form.channelPhone")}
              </Text>
                  </Box>
                </Box>
              </Box>
            </Tabs.Trigger>
          </Tabs.List>

          <Box className={styles.fieldsPt}>
          <Box className={`${styles.fields} js-gate-fields`}>
            <Tabs.Content value="email">
              <FormProvider {...emailForm}>
                <Flex asChild direction="column" gap="4">
                  <form noValidate onSubmit={emailForm.handleSubmit(onEmail)}>
                    <LoginEmailController />
                    <LoginPasswordController />
                    <LoginRememberController />
                    <Button type="submit" size="3" loading={loginEmail.isPending} className={styles.fullWidth}>
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
                    <Button type="submit" size="3" loading={loginPhone.isPending} className={styles.fullWidth}>
                      {t("form.submit")}
                      <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                    </Button>
                  </form>
                </Flex>
              </FormProvider>
            </Tabs.Content>

            <Box className={styles.foot}>
              <Box className={styles.footPt}>
            <Flex justify="between" align="baseline" gap="4">
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
              </Box>
            </Box>
          </Box>
          </Box>
        </Tabs.Root>
      </Box>
    </AuthChrome>
  );
}
