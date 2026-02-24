import type { Language } from "./i18nResources";

import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";

import { getErrorTranslations } from "@/apis/system.api";

import { NAME_SPACES, RESOURCES } from "./i18nResources";

// 所有语言和命名空间资源
const fallbackLng = "en";

/** 从后端加载错误码翻译并注入到 i18n 的 errors 命名空间（挂载目录更新即生效） */
export async function loadErrorTranslations(lng: string): Promise<void> {
  const supported = ["en", "zh", "fr"];
  if (!supported.includes(lng)) return;
  try {
    const json = await getErrorTranslations(lng);
    if (json && typeof json === "object") {
      i18n.addResourceBundle(lng, "errors", json, true, true);
    }
  } catch {
    // 网络或后端未就绪时忽略，getApiErrorMessage 会回退到 message / fallback
  }
}

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    compatibilityJSON: "v4",
    resources: RESOURCES,
    lng: "en", // 默认使用英文
    fallbackLng,
    ns: NAME_SPACES,
    defaultNS: NAME_SPACES[0],
    interpolation: {
      escapeValue: false,
    },
    keySeparator: ".",
    detection: {
      order: ["navigator", "htmlTag", "path", "subdomain"],
    },
  });

// 初始语言加载错误码翻译；切换语言时再加载对应语言
loadErrorTranslations(i18n.language || fallbackLng);
i18n.on("languageChanged", loadErrorTranslations);

export default i18n;

// 切换语言方法
export const ChangeLanguage = (lng: Language) => {
  i18n.changeLanguage(lng);
};
