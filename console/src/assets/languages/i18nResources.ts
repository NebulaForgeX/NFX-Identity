import type { ValueOf } from "@/utils/types";

import en_BootstrapPage from "./en/BootstrapPage.json";
import en_elements_bootstrap from "./en/elements.bootstrap.json";
import fr_BootstrapPage from "./fr/BootstrapPage.json";
import fr_elements_bootstrap from "./fr/elements.bootstrap.json";
import zh_BootstrapPage from "./zh/BootstrapPage.json";
import zh_elements_bootstrap from "./zh/elements.bootstrap.json";

// 所有语言包内容
export const RESOURCES = {
  en: {
    BootstrapPage: en_BootstrapPage,
    "elements.bootstrap": en_elements_bootstrap,
  },
  zh: {
    BootstrapPage: zh_BootstrapPage,
    "elements.bootstrap": zh_elements_bootstrap,
  },
  fr: {
    BootstrapPage: fr_BootstrapPage,
    "elements.bootstrap": fr_elements_bootstrap,
  },
};

// 所有命名空间
export const NAME_SPACES_MAP = {
  BootstrapPage: "BootstrapPage",
  "elements.bootstrap": "elements.bootstrap",
};

export const NAME_SPACES = Object.values(NAME_SPACES_MAP);

// 所有语言类型
export const LANGUAGE = {
  EN: "en",
  ZH: "zh",
  FR: "fr",
} as const;

export type Language = ValueOf<typeof LANGUAGE>;
