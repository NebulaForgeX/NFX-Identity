import type { ReactNode } from "react";

import { memo } from "react";
import { FormProvider } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { Button, LanguageSwitcher, Suspense, ThemeSwitcher } from "@/components";
import {
  AdminEmailController,
  AdminPasswordConfirmController,
  AdminPasswordController,
  AdminPhoneController,
  AdminUsernameController,
  useInitBootstrapForm,
  useSubmitBootstrap,
  VersionController,
} from "@/elements/bootstrap";
import { useSystemInit } from "@/hooks/useSystem";

import styles from "./style.module.css";

interface BootstrapProviderProps {
  children: ReactNode;
}

/**
 * BootstrapContent - 系统初始化内容（必须在 Suspense 内，useSystemInit 为 suspense 模式会 throw）
 * 检查系统是否已初始化，如果未初始化则显示初始化表单
 */
const BootstrapContent = memo(({ children }: { children: ReactNode }) => {
  const { t } = useTranslation("BootstrapProvider");
  const systemState = useSystemInit();
  const methods = useInitBootstrapForm();
  const { onSubmit, onSubmitError, isPending } = useSubmitBootstrap();

  console.log("🔍 System state:", systemState.data);

  if (!systemState.data.initialized) {
    return (
      <div className={styles.container}>
        {/* 左上角语言和主题切换按钮 */}
        <div className={styles.topControls}>
          <LanguageSwitcher status="default" />
          <ThemeSwitcher status="default" />
        </div>
        <FormProvider {...methods}>
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
                  disabled={isPending}
                  onClick={methods.handleSubmit(onSubmit, onSubmitError)}
                >
                  {t("start_initialization")}
                </Button>
              </div>
            </form>
          </div>
        </FormProvider>
      </div>
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
  const { t } = useTranslation("BootstrapProvider");
  return (
    <Suspense loadingType="truck" loadingText={t("checking_system_status")} loadingSize="medium">
      <BootstrapContent>{children}</BootstrapContent>
    </Suspense>
  );
}

export default BootstrapProvider;
