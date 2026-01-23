import type { BootstrapFormValues } from "@/elements/bootstrap";
import type { ReactNode } from "react";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider, useFormContext } from "react-hook-form";

import Suspense from "@/components/Suspense";
import { Button, LanguageSwitcher, ThemeSwitcher } from "@/components";
import {
  useInitBootstrapForm,
  useSubmitBootstrap,
  VersionController,
  AdminUsernameController,
  AdminPasswordController,
  AdminPasswordConfirmController,
  AdminEmailController,
  AdminPhoneController,
} from "@/elements/bootstrap";

import { useSystemInit } from "@/hooks/useSystem";

import styles from "./style.module.css";

interface BootstrapProviderProps {
  children: ReactNode;
}

/**
 * BootstrapFormContent - 系统初始化表单内容
 */
const BootstrapFormContent = memo(() => {
  const { t } = useTranslation("BootstrapProvider");
  const methods = useFormContext<BootstrapFormValues>();
  const { onSubmit, onSubmitError, isPending } = useSubmitBootstrap();

  return (
    <div className={styles.container}>
      {/* 左上角语言和主题切换按钮 */}
      <div className={styles.topControls}>
        <LanguageSwitcher status="default" />
        <ThemeSwitcher status="default" />
      </div>

      <div className={styles.formCard}>
        <h2 className={styles.formTitle}>{t("title")}</h2>
        <p className={styles.formDescription}>{t("description")}</p>

        <form
          onSubmit={(e) => {
            e.preventDefault();
          }}
          className={styles.form}
        >
          {/* 版本信息 */}
          <div className={styles.section}>
            <h3 className={styles.sectionTitle}>{t("version_section")}</h3>
            <VersionController />
          </div>

          {/* 管理员账户信息 */}
          <div className={styles.section}>
            <h3 className={styles.sectionTitle}>{t("admin_section")}</h3>
            <AdminUsernameController />
            <AdminPasswordController />
            <AdminPasswordConfirmController />
          </div>

          {/* 可选信息 */}
          <div className={styles.section}>
            <h3 className={styles.sectionTitle}>{t("optional_section")}</h3>
            <AdminEmailController />
            <AdminPhoneController />
          </div>

          <div className={styles.actions}>
            <Button
              type="button"
              variant="primary"
              size="medium"
              loading={isPending}
              disabled={isPending}
              onClick={methods.handleSubmit(onSubmit, onSubmitError)}
            >
              {isPending ? t("initializing") : t("start_initialization")}
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
});

BootstrapFormContent.displayName = "BootstrapFormContent";

/**
 * BootstrapContent - 系统初始化内容（必须在 Suspense 内，useSystemInit 为 suspense 模式会 throw）
 * 检查系统是否已初始化，如果未初始化则显示初始化表单
 */
const BootstrapContent = memo(({ children }: { children: ReactNode }) => {
  const systemState = useSystemInit();
  const methods = useInitBootstrapForm();

  if (!systemState.data.initialized) {
    return (
      <FormProvider {...methods}>
        <BootstrapFormContent />
      </FormProvider>
    );
  }

  return <>{children}</>;
});
BootstrapContent.displayName = "BootstrapContent";

/**
 * BootstrapProvider - 系统初始化 Provider
 * 检查系统是否已初始化，如果未初始化则显示初始化表单
 */
export function BootstrapProvider({ children }: BootstrapProviderProps) {
  return (
    <Suspense
      loadingType="truck"
      loadingText="检查系统状态..."
      loadingSize="medium"
    >
      <BootstrapContent>{children}</BootstrapContent>
    </Suspense>
  );
}
