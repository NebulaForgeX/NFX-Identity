import { ShieldCheck } from "nfx-ui/icons";
import { useState } from "react";
import { Box, Button, Flex, Grid, Text, TextField } from "@radix-ui/themes";
import { LanguageEnum } from "nfx-ui/enums";
import { useChangePassword, useListEmails, useSendChangePasswordVerificationCode } from "nfx-ui/hooks";
import { usePreferenceStore } from "nfx-ui/stores";
import { isVerificationCodeComplete, normalizeVerificationCode } from "nfx-ui/utils";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { safeArray } from "@/utils";

import { LedgerSection } from "./Ledger";

export default function SecurityView() {
  const { t } = useTranslation("pages.Profile.Security");
  const currentLanguage = usePreferenceStore((s) => s.language);
  const changePassword = useChangePassword();
  const sendCode = useSendChangePasswordVerificationCode();
  const emails = useListEmails();
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [verificationCode, setVerificationCode] = useState("");
  const primaryEmail = safeArray(emails.data?.items).find((e) => e.isPrimary)?.email;
  const busy = changePassword.isPending || sendCode.isPending;
  const canSubmit = currentPassword.length > 0 && newPassword.length >= 8 && isVerificationCodeComplete(verificationCode) && !busy;

  return (
    <PageFrame>
      <PageHeader icon={ShieldCheck} title={t("title")} description={t("description")} />
      <LedgerSection title={t("sections.password.title")} description={t("sections.password.description")}>
        <Flex direction="column" gap="4">
          <Grid columns={{ initial: "1", sm: "2" }} gap="4">
            <Flex direction="column" gap="1">
              <Text size="1" color="gray">
                {t("labels.currentPassword")}
              </Text>
              <TextField.Root size="2" type="password" value={currentPassword} disabled={busy} onChange={(e) => setCurrentPassword(e.target.value)} />
            </Flex>
            <Flex direction="column" gap="1">
              <Text size="1" color="gray">
                {t("labels.newPassword")}
              </Text>
              <TextField.Root size="2" type="password" value={newPassword} disabled={busy} onChange={(e) => setNewPassword(e.target.value)} />
            </Flex>
          </Grid>
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">
              {t("labels.verificationCode")}
            </Text>
            <Text size="1" color="gray">
              {primaryEmail ? t("labels.passwordSendCodeHint", { email: primaryEmail }) : t("labels.passwordSendCodeHintNoEmail")}
            </Text>
            <Flex gap="2" wrap="wrap" align="center">
              <Box minWidth="160px" >
                <TextField.Root
                  size="2"
                  autoComplete="one-time-code"
                  value={verificationCode}
                  disabled={busy}
                  onChange={(e) => setVerificationCode(normalizeVerificationCode(e.target.value))}
                  placeholder={t("labels.verificationCodePlaceholder")}
                />
              </Box>
              <Button
                type="button"
                size="2"
                variant="outline"
                loading={sendCode.isPending}
                disabled={busy || !primaryEmail}
                onClick={() => void sendCode.mutateAsync({ lang: currentLanguage ?? LanguageEnum.EN })}
              >
                {t("actions.sendCode")}
              </Button>
            </Flex>
          </Flex>
        </Flex>
        <Flex justify="end">
          <Button
            size="2"
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
      </LedgerSection>
    </PageFrame>
  );
}
