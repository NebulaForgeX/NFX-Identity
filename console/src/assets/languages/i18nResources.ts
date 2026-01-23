import type { ValueOf } from "@/utils/types";

import en_BootstrapProvider from "./en/BootstrapProvider.json";
import en_elements_bootstrap from "./en/elements.bootstrap.json";
import fr_BootstrapProvider from "./fr/BootstrapProvider.json";
import fr_elements_bootstrap from "./fr/elements.bootstrap.json";
import zh_BootstrapProvider from "./zh/BootstrapProvider.json";
import zh_elements_bootstrap from "./zh/elements.bootstrap.json";

// 所有语言包内容
export const RESOURCES = {
  en: {
    BootstrapProvider: en_BootstrapProvider,
    "elements.bootstrap": en_elements_bootstrap,
  },
  zh: {
    BootstrapProvider: zh_BootstrapProvider,
    "elements.bootstrap": zh_elements_bootstrap,
  },
  fr: {
    BootstrapProvider: fr_BootstrapProvider,
    "elements.bootstrap": fr_elements_bootstrap,
  },
};

// 所有命名空间
export const NAME_SPACES_MAP = {
  BootstrapProvider: "BootstrapProvider",
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
