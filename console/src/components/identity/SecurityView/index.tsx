import { useState } from "react";
import { Box, Button, Flex, Grid, Text, TextField } from "@radix-ui/themes";
import { LanguageEnum } from "nfx-ui/enums";
import { useChangePassword, useListEmails, useSendChangePasswordVerificationCode } from "nfx-ui/hooks";
import { AnimatedIcon, CheckedIcon, ShieldCheck } from "nfx-ui/icons";
import { usePreferenceStore } from "nfx-ui/stores";
import { isVerificationCodeComplete, normalizeVerificationCode } from "nfx-ui/utils";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { safeArray } from "@/utils";

import { LedgerSection } from "../Ledger";
import styles from "./s.module.css";

function Step({ done, label }: { done: boolean; label: string }) {
  return (
    <Flex align="center" gap="3">
      <Box className={styles.step} data-done={done ? "true" : "false"}>
        <Flex align="center" justify="center" width="100%" height="100%">
          {done ? <AnimatedIcon icon={CheckedIcon} size={12} /> : null}
        </Flex>
      </Box>
      <Text size="2" color={done ? undefined : "gray"}>
        {label}
      </Text>
    </Flex>
  );
}

export default function SecurityView() {
  const { t } = useTranslation("pages.Profile.Security");
  const currentLanguage = usePreferenceStore((s) => s.language);
  const changePassword = useChangePassword({ successMsg: t("toasts.changePasswordSuccess") });
  const sendCode = useSendChangePasswordVerificationCode({ successMsg: t("toasts.sendVerificationCodeSuccess") });
  const emails = useListEmails();
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [verificationCode, setVerificationCode] = useState("");
  const primaryEmail = safeArray(emails.data?.items).find((e) => e.isPrimary)?.email;
  const busy = changePassword.isPending || sendCode.isPending;
  const codeReady = isVerificationCodeComplete(verificationCode);
  const canSubmit = currentPassword.length > 0 && newPassword.length >= 8 && codeReady && !busy;

  return (
    <PageFrame>
      <PageHeader icon={ShieldCheck} title={t("title")} description={t("description")} />
      <Grid columns={{ initial: "1", lg: "minmax(0, 1.5fr) minmax(0, 1fr)" }} gap="5" align="start">
        <LedgerSection title={t("sections.password.title")} description={t("sections.password.description")}>
          <Flex direction="column" gap="5">
            <Grid columns={{ initial: "1", sm: "2" }} gap="4">
              <Flex direction="column" gap="2">
                <Text as="label" size="2" weight="medium">
                  {t("labels.currentPassword")}
                </Text>
                <TextField.Root size="3" type="password" value={currentPassword} disabled={busy} onChange={(e) => setCurrentPassword(e.target.value)} />
              </Flex>
              <Flex direction="column" gap="2">
                <Text as="label" size="2" weight="medium">
                  {t("labels.newPassword")}
                </Text>
                <TextField.Root size="3" type="password" value={newPassword} disabled={busy} onChange={(e) => setNewPassword(e.target.value)} />
              </Flex>
            </Grid>
            <Flex direction="column" gap="2">
              <Text as="label" size="2" weight="medium">
                {t("labels.verificationCode")}
              </Text>
              <Text size="1" color="gray">
                {primaryEmail ? t("labels.passwordSendCodeHint", { email: primaryEmail }) : t("labels.passwordSendCodeHintNoEmail")}
              </Text>
              <Flex gap="2" wrap="wrap" align="center">
                <Box className={styles.codeField}>
                  <TextField.Root
                    size="3"
                    autoComplete="one-time-code"
                    value={verificationCode}
                    disabled={busy}
                    onChange={(e) => setVerificationCode(normalizeVerificationCode(e.target.value))}
                    placeholder={t("labels.verificationCodePlaceholder")}
                  />
                </Box>
                <Button
                  type="button"
                  size="3"
                  variant="outline"
                  loading={sendCode.isPending}
                  disabled={busy || !primaryEmail}
                  onClick={() => void sendCode.mutateAsync({ lang: currentLanguage ?? LanguageEnum.EN })}
                >
                  {t("actions.sendCode")}
                </Button>
              </Flex>
            </Flex>
            <Flex justify="end">
              <Button
                size="3"
                loading={changePassword.isPending}
                disabled={!canSubmit}
                onClick={() =>
                  changePassword
                    .mutateAsync({
                      currentPassword,
                      newPassword,
                      verificationCode: normalizeVerificationCode(verificationCode),
                    })
                    .then(() => {
                      setCurrentPassword("");
                      setNewPassword("");
                      setVerificationCode("");
                    })
                }
              >
                {t("actions.updatePassword")}
              </Button>
            </Flex>
          </Flex>
        </LedgerSection>
        <LedgerSection sticky title={t("sections.checklist.title")} description={t("sections.checklist.description")}>
          <Flex direction="column" gap="3">
            <Step done={Boolean(primaryEmail)} label={t("sections.checklist.email")} />
            <Step done={currentPassword.length > 0} label={t("sections.checklist.current")} />
            <Step done={newPassword.length >= 8} label={t("sections.checklist.length")} />
            <Step done={codeReady} label={t("sections.checklist.code")} />
          </Flex>
        </LedgerSection>
      </Grid>
    </PageFrame>
  );
}
