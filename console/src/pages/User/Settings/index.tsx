import { useState } from "react";
import { Button, Card, Flex, Heading, Text, TextField } from "@radix-ui/themes";
import { Settings2 } from "lucide-react";
import { useMutation } from "@tanstack/react-query";
import { useTranslation } from "react-i18next";
import { useAuthRepository } from "nfx-ui/apis";
import { LanguageEnum } from "nfx-ui/enums";
import { PageFrame } from "nfx-ui/layouts";
import { PageHeader, ThemeSettings } from "nfx-ui/components";

export default function SettingsPage() {
  const { t } = useTranslation("EditPreferencePage");
  const auth = useAuthRepository();
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [code, setCode] = useState("");

  const sendCode = useMutation({
    mutationFn: async () => auth.SendChangePasswordVerificationCode({ lang: LanguageEnum.ZH }),
  });
  const change = useMutation({
    mutationFn: async () => auth.ChangePassword({ currentPassword, newPassword, verificationCode: code }),
  });

  const error = (sendCode.error || change.error) as Error | null;

  return (
    <PageFrame>
      <PageHeader icon={Settings2} title={t("title")} description={t("subtitle")} />
      <Flex direction="column" gap="6" width="100%">
        <ThemeSettings />
        <Card size="3">
          <Flex direction="column" gap="3">
            <Heading size="4">{t("changePassword")}</Heading>
            <Text size="2" color="gray">
              {t("changePasswordHint")}
            </Text>
            <TextField.Root type="password" value={currentPassword} onChange={(e) => setCurrentPassword(e.target.value)} placeholder={t("currentPassword")} />
            <TextField.Root type="password" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} placeholder={t("newPassword")} />
            <Flex gap="2">
              <TextField.Root value={code} onChange={(e) => setCode(e.target.value)} placeholder={t("verificationCode")} />
              <Button type="button" variant="soft" onClick={() => sendCode.mutate()} loading={sendCode.isPending}>
                {t("sendCode")}
              </Button>
            </Flex>
            {error ? (
              <Text size="2" color="red">
                {error.message}
              </Text>
            ) : null}
            {change.isSuccess ? (
              <Text size="2" color="green">
                {t("passwordUpdated")}
              </Text>
            ) : null}
            <Button onClick={() => change.mutate()} loading={change.isPending} disabled={!currentPassword || !newPassword || !code}>
              {t("submit")}
            </Button>
          </Flex>
        </Card>
      </Flex>
    </PageFrame>
  );
}
