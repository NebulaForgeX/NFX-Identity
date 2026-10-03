import { Flex, RadioCards, Text } from "@radix-ui/themes";
import { LanguageEnum } from "nfx-ui/enums";
import { useSyncPreference } from "nfx-ui/hooks";
import { usePreferenceStore } from "nfx-ui/stores";
import { useTranslation } from "react-i18next";

import { LedgerSection } from "../Ledger";

const LANG_CODE: Record<LanguageEnum, string> = {
  [LanguageEnum.EN]: "EN",
  [LanguageEnum.ZH]: "ZH",
  [LanguageEnum.FR]: "FR",
};

const LANG_NAME_KEY: Record<LanguageEnum, string> = {
  [LanguageEnum.EN]: "labels.langEn",
  [LanguageEnum.ZH]: "labels.langZh",
  [LanguageEnum.FR]: "labels.langFr",
};

export default function LanguageSettings() {
  const { t } = useTranslation("pages.User.Setting");
  const currentLanguage = usePreferenceStore((s) => s.language);
  const { syncPreference } = useSyncPreference();

  return (
    <LedgerSection title={t("sections.language.title")} description={t("sections.language.description")}>
      <RadioCards.Root size="2" columns="3" gap="3" value={currentLanguage} onValueChange={(v) => syncPreference({ language: v as LanguageEnum })}>
        {(Object.values(LanguageEnum) as LanguageEnum[]).map((lang) => (
          <RadioCards.Item key={lang} value={lang}>
            <Flex direction="column" align="center" gap="1" width="100%">
              <Text size="4" weight="bold">
                {LANG_CODE[lang]}
              </Text>
              <Text size="1" color="gray">
                {t(LANG_NAME_KEY[lang])}
              </Text>
            </Flex>
          </RadioCards.Item>
        ))}
      </RadioCards.Root>
    </LedgerSection>
  );
}
