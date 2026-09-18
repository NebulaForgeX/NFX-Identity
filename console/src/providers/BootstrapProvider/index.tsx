import type { ReactNode } from "react";

import { memo } from "react";
import { Button, Flex, Heading, Text, TextField } from "@radix-ui/themes";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { Controller, FormProvider, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { z } from "zod";
import { useUnifiedSuspenseQuery } from "nfx-ui/hooks";
import { hideModal, showError, showLoading, showSuccess } from "nfx-ui/stores";
import { getApiErrorMessage } from "nfx-ui/utils";

import { getSystemStateLatest, initializeSystemState } from "@/apis/system";
import { Suspense } from "@/components";
import { SYSTEM_SYSTEM_STATE_INIT } from "@/constants/system.query.key";
import AuthShell from "@/pages/LoginPage/AuthShell";
import type { SystemState } from "@/types/domain/system.domain";

interface BootstrapProviderProps {
  children: ReactNode;
}

type BootstrapFormValues = {
  version: string;
  adminUsername: string;
  adminPassword: string;
  adminPasswordConfirm: string;
  adminEmail: string;
  adminPhone: string;
};

const Field = memo(function Field({
  name,
  label,
  type = "text",
}: {
  name: keyof BootstrapFormValues;
  label: string;
  type?: "text" | "password" | "email";
}) {
  return (
    <Controller
      name={name}
      render={({ field, fieldState }) => (
        <Flex direction="column" gap="1">
          <Text as="label" size="2" weight="medium">
            {label}
          </Text>
          <TextField.Root size="3" type={type} color={fieldState.error ? "red" : undefined} {...field} />
          {fieldState.error?.message ? (
            <Text size="1" color="red">
              {fieldState.error.message}
            </Text>
          ) : null}
        </Flex>
      )}
    />
  );
});

const BootstrapContent = memo(({ children }: { children: ReactNode }) => {
  const { t } = useTranslation("BootstrapProvider");
  const systemState = useUnifiedSuspenseQuery(async () => getSystemStateLatest(), SYSTEM_SYSTEM_STATE_INIT, {});
  const methods = useForm<BootstrapFormValues>({
    resolver: zodResolver(
      z
        .object({
          version: z.string().trim().min(1, t("version_required")),
          adminUsername: z.string().trim().min(3, t("username_min")),
          adminPassword: z.string().min(8, t("password_min")),
          adminPasswordConfirm: z.string().min(8, t("password_confirm_required")),
          adminEmail: z.string().trim().email(t("email_invalid")).min(1, t("email_required")),
          adminPhone: z.string().trim(),
        })
        .refine((data) => data.adminPassword === data.adminPasswordConfirm, {
          message: t("password_mismatch"),
          path: ["adminPasswordConfirm"],
        }),
    ),
    defaultValues: {
      version: "1.0.0",
      adminUsername: "",
      adminPassword: "",
      adminPasswordConfirm: "",
      adminEmail: "",
      adminPhone: "",
    },
  });

  const init = useMutation({
    mutationFn: async (values: BootstrapFormValues) => {
      await initializeSystemState({
        version: values.version,
        adminUsername: values.adminUsername.trim(),
        adminPassword: values.adminPassword,
        adminEmail: values.adminEmail.trim(),
        adminPhone: values.adminPhone.trim() || undefined,
      });
    },
    onMutate: () => {
      showLoading({ message: t("initializing"), canClose: false });
    },
    onSuccess: () => {
      hideModal("loading");
      showSuccess({
        title: t("success_title"),
        message: t("success_message"),
        onClick: () => window.location.reload(),
      });
    },
    onError: (error: Error) => {
      hideModal("loading");
      showError(getApiErrorMessage(error, t("failed_title")), t("failed_title"));
    },
  });

  const state = systemState.data as SystemState | undefined;
  if (!state?.initialized) {
    return (
      <AuthShell brandEyebrow="NFX Identity" brandTitle={t("title")} heroFooter={t("description")}>
        <FormProvider {...methods}>
          <Flex direction="column" gap="5" className="js-auth-stagger">
            <Flex direction="column" gap="1">
              <Heading as="h2" size="6">
                {t("title")}
              </Heading>
              <Text as="p" size="2" color="gray">
                {t("description")}
              </Text>
            </Flex>
            <Field name="version" label={t("version_section")} />
            <Text size="2" weight="bold">
              {t("admin_section")}
            </Text>
            <Field name="adminUsername" label={t("username")} />
            <Field name="adminPassword" label={t("password")} type="password" />
            <Field name="adminPasswordConfirm" label={t("password_confirm")} type="password" />
            <Field name="adminEmail" label={t("email")} type="email" />
            <Text size="2" weight="bold">
              {t("optional_section")}
            </Text>
            <Field name="adminPhone" label={t("phone")} />
            <Button type="button" size="3" loading={init.isPending} onClick={methods.handleSubmit((values) => init.mutateAsync(values))}>
              {t("start_initialization")}
            </Button>
          </Flex>
        </FormProvider>
      </AuthShell>
    );
  }

  return <>{children}</>;
});
BootstrapContent.displayName = "BootstrapContent";

export function BootstrapProvider({ children }: BootstrapProviderProps) {
  const { t } = useTranslation("BootstrapProvider");
  return (
    <Suspense loadingText={t("checking_system_status")}>
      <BootstrapContent>{children}</BootstrapContent>
    </Suspense>
  );
}

export default BootstrapProvider;
