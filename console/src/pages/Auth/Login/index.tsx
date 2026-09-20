import { AnimatedIcon, ArrowNarrowRightIcon, EyeIcon, EyeOffIcon } from "nfx-ui/icons";
import { useRef, useState } from "react";
import { useGSAP } from "@gsap/react";
import { Button, Checkbox, Flex, Heading, Link, Text, TextField } from "@radix-ui/themes";
import clsx from "clsx";
import gsap from "gsap";
import { APP_NAME } from "nfx-ui/config";
import { useLoginWithEmail, useLoginWithPhone } from "nfx-ui/hooks";
import { LoginFormData, LoginWithPhoneFormData, useInitLoginForm, useInitLoginWithPhoneForm } from "nfx-ui/schemas";
import { Controller, FormProvider, SubmitHandler } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { routerEventEmitter } from "@/events/router";
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
  const [showPassword, setShowPassword] = useState(false);
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
      <div ref={rootRef} className={styles.gate}>
        <Flex direction="column" gap="3" mb="6">
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

        <div className={styles.slabs}>
          <button type="button" className={clsx(styles.slab, "js-gate-slab", channel === "email" && styles.slabActive)} aria-pressed={channel === "email"} onClick={() => setChannel("email")}>
            <span className={styles.slabIndex}>01</span>
            <span className={styles.slabLabel}>{t("form.channelEmail")}</span>
          </button>
          <button type="button" className={clsx(styles.slab, "js-gate-slab", channel === "phone" && styles.slabActive)} aria-pressed={channel === "phone"} onClick={() => setChannel("phone")}>
            <span className={styles.slabIndex}>02</span>
            <span className={styles.slabLabel}>{t("form.channelPhone")}</span>
          </button>
        </div>

        <div className={`${styles.fields} js-gate-fields`}>
          {channel === "email" ? (
            <FormProvider {...emailForm}>
              <Flex asChild direction="column" gap="4">
                <form noValidate onSubmit={emailForm.handleSubmit(onEmail)}>
                  <Controller
                    name="email"
                    control={emailForm.control}
                    render={({ field, fieldState }) => (
                      <Flex direction="column" gap="1">
                        <Text as="label" size="2" weight="medium" htmlFor="login-email">
                          {t("form.emailLabel")}
                        </Text>
                        <TextField.Root id="login-email" size="3" type="email" autoComplete="email" placeholder={t("form.emailPlaceholder")} {...field} />
                        {fieldState.error ? (
                          <Text size="1" color="red">
                            {fieldState.error.message}
                          </Text>
                        ) : null}
                      </Flex>
                    )}
                  />
                  <Controller
                    name="password"
                    control={emailForm.control}
                    render={({ field, fieldState }) => (
                      <Flex direction="column" gap="1">
                        <Text as="label" size="2" weight="medium" htmlFor="login-password">
                          {t("form.passwordLabel")}
                        </Text>
                        <TextField.Root
                          id="login-password"
                          size="3"
                          type={showPassword ? "text" : "password"}
                          autoComplete="current-password"
                          placeholder={t("form.passwordPlaceholder")}
                          {...field}
                        >
                          <TextField.Slot side="right">
                            <Button type="button" size="1" variant="soft" color="gray" onClick={() => setShowPassword((v) => !v)} aria-label={showPassword ? t("form.hidePassword") : t("form.showPassword")}>
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
                    name="rememberMe"
                    control={emailForm.control}
                    render={({ field }) => (
                      <Flex asChild align="center" gap="2">
                        <Text as="label" size="2">
                          <Checkbox checked={!!field.value} onCheckedChange={(v) => field.onChange(v === true)} />
                          {t("form.rememberMe")}
                        </Text>
                      </Flex>
                    )}
                  />
                  <Button type="submit" size="3" loading={loginEmail.isPending} style={{ width: "100%" }}>
                    {t("form.submit")}
                    <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                  </Button>
                </form>
              </Flex>
            </FormProvider>
          ) : (
            <FormProvider {...phoneForm}>
              <Flex asChild direction="column" gap="4">
                <form noValidate onSubmit={phoneForm.handleSubmit(onPhone)}>
                  <Controller
                    name="phone"
                    control={phoneForm.control}
                    render={({ field, fieldState }) => (
                      <Flex direction="column" gap="1">
                        <Text as="label" size="2" weight="medium" htmlFor="login-phone">
                          {t("form.phoneLabel")}
                        </Text>
                        <TextField.Root id="login-phone" size="3" type="tel" autoComplete="tel" placeholder={t("form.phonePlaceholder")} {...field} />
                        {fieldState.error ? (
                          <Text size="1" color="red">
                            {fieldState.error.message}
                          </Text>
                        ) : null}
                      </Flex>
                    )}
                  />
                  <Controller
                    name="password"
                    control={phoneForm.control}
                    render={({ field, fieldState }) => (
                      <Flex direction="column" gap="1">
                        <Text as="label" size="2" weight="medium" htmlFor="login-phone-password">
                          {t("form.passwordLabel")}
                        </Text>
                        <TextField.Root
                          id="login-phone-password"
                          size="3"
                          type={showPassword ? "text" : "password"}
                          autoComplete="current-password"
                          placeholder={t("form.passwordPlaceholder")}
                          {...field}
                        >
                          <TextField.Slot side="right">
                            <Button type="button" size="1" variant="soft" color="gray" onClick={() => setShowPassword((v) => !v)} aria-label={showPassword ? t("form.hidePassword") : t("form.showPassword")}>
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
                    name="rememberMe"
                    control={phoneForm.control}
                    render={({ field }) => (
                      <Flex asChild align="center" gap="2">
                        <Text as="label" size="2">
                          <Checkbox checked={!!field.value} onCheckedChange={(v) => field.onChange(v === true)} />
                          {t("form.rememberMe")}
                        </Text>
                      </Flex>
                    )}
                  />
                  <Button type="submit" size="3" loading={loginPhone.isPending} style={{ width: "100%" }}>
                    {t("form.submit")}
                    <AnimatedIcon icon={ArrowNarrowRightIcon} size={16} />
                  </Button>
                </form>
              </Flex>
            </FormProvider>
          )}

          <div className={styles.foot}>
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
          </div>
        </div>
      </div>
    </AuthChrome>
  );
}
