import { AnimatedIcon, ArrowNarrowRightIcon, EyeIcon, EyeOffIcon } from "nfx-ui/icons";
import { useState } from "react";
import { Button, Checkbox, Flex, Heading, Link, Separator, Tabs, Text, TextField } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { useLoginWithEmail, useLoginWithPhone } from "nfx-ui/hooks";
import { LoginFormData, LoginWithPhoneFormData, useInitLoginForm, useInitLoginWithPhoneForm } from "nfx-ui/schemas";
import { Controller, FormProvider, SubmitHandler } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { routerEventEmitter } from "@/events/router";
import { ROUTES } from "@/navigations";
import AuthShell from "@/pages/Auth/shared/AuthShell";
import { safeOr } from "@/utils";

import styles from "./s.module.css";

export default function LoginPage() {
  const { t } = useTranslation("pages.Account.Login");
  const emailForm = useInitLoginForm();
  const phoneForm = useInitLoginWithPhoneForm();
  const loginEmail = useLoginWithEmail();
  const loginPhone = useLoginWithPhone();
  const [showPassword, setShowPassword] = useState(false);
  const [channel, setChannel] = useState<"email" | "phone">("email");

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
    <AuthShell brandEyebrow={t("brand.access", { name: APP_NAME })} brandTitle={t("brand.title")} heroFooter={t("heroFooter")}>
      <Flex direction="column" gap="5">
        <Flex direction="column" gap="1" className="js-auth-stagger">
          <Text as="p" size="1" weight="bold" className={styles.kicker}>
            {t("form.welcomeBack")}
          </Text>
          <Heading as="h2" size="6">
            {t("form.title", { name: APP_NAME })}
          </Heading>
          <Text as="p" size="2" color="gray">
            {t("form.subtitle")}
          </Text>
        </Flex>

        <Tabs.Root value={channel} onValueChange={(v) => setChannel(v as "email" | "phone")}>
          <Tabs.List>
            <Tabs.Trigger value="email">{t("form.channelEmail")}</Tabs.Trigger>
            <Tabs.Trigger value="phone">{t("form.channelPhone")}</Tabs.Trigger>
          </Tabs.List>
        </Tabs.Root>

        {channel === "email" ? (
          <FormProvider {...emailForm}>
            <Flex asChild direction="column" gap="4" className="js-auth-stagger">
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
            <Flex asChild direction="column" gap="4" className="js-auth-stagger">
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

        <Separator size="4" className="js-auth-stagger" />
        <Text as="p" size="2" align="center" color="gray" className="js-auth-stagger">
          {t("promo.newTo", { name: APP_NAME })}{" "}
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
        </Text>
      </Flex>
    </AuthShell>
  );
}
