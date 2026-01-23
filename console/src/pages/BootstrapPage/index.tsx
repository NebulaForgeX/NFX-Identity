import type { BootstrapFormValues } from "@/elements/bootstrap";
import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider, useFormContext } from "react-hook-form";

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

import styles from "./styles.module.css";

const BootstrapFormContent = memo(() => {
  const { t } = useTranslation("BootstrapPage");
  const methods = useFormContext<BootstrapFormValues>();
  const { onSubmit, onSubmitError, isPending } = useSubmitBootstrap();

  return (
    <div className={styles.container}>
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
            <button
              type="button"
              className={styles.submitBtn}
              disabled={isPending}
              onClick={methods.handleSubmit(onSubmit, onSubmitError)}
            >
              {isPending ? t("initializing") : t("start_initialization")}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
});

BootstrapFormContent.displayName = "BootstrapFormContent";

/**
 * Bootstrap Page - 系统初始化页面
 * 当系统未初始化时显示此页面
 */
const BootstrapPage = memo(() => {
  const methods = useInitBootstrapForm();

  return (
    <FormProvider {...methods}>
      <BootstrapFormContent />
    </FormProvider>
  );
});

BootstrapPage.displayName = "BootstrapPage";

export default BootstrapPage;
