import type { ValueOf } from "@/utils/types";

import en_BootstrapProvider from "./en/BootstrapProvider.json";
import en_elements_bootstrap from "./en/elements.bootstrap.json";
import en_LoginPage from "./en/LoginPage.json";
import en_components from "./en/components.json";
import fr_BootstrapProvider from "./fr/BootstrapProvider.json";
import fr_elements_bootstrap from "./fr/elements.bootstrap.json";
import fr_LoginPage from "./fr/LoginPage.json";
import fr_components from "./fr/components.json";
import zh_BootstrapProvider from "./zh/BootstrapProvider.json";
import zh_elements_bootstrap from "./zh/elements.bootstrap.json";
import zh_LoginPage from "./zh/LoginPage.json";
import zh_components from "./zh/components.json";

// 所有语言包内容
export const RESOURCES = {
  en: {
    BootstrapProvider: en_BootstrapProvider,
    "elements.bootstrap": en_elements_bootstrap,
    LoginPage: en_LoginPage,
    components: en_components,
  },
  zh: {
    BootstrapProvider: zh_BootstrapProvider,
    "elements.bootstrap": zh_elements_bootstrap,
    LoginPage: zh_LoginPage,
    components: zh_components,
  },
  fr: {
    BootstrapProvider: fr_BootstrapProvider,
    "elements.bootstrap": fr_elements_bootstrap,
    LoginPage: fr_LoginPage,
    components: fr_components,
  },
};

// 所有命名空间
export const NAME_SPACES_MAP = {
  BootstrapProvider: "BootstrapProvider",
  "elements.bootstrap": "elements.bootstrap",
  LoginPage: "LoginPage",
  components: "components",
};

export const NAME_SPACES = Object.values(NAME_SPACES_MAP);

// 所有语言类型
export const LANGUAGE = {
  EN: "en",
  ZH: "zh",
  FR: "fr",
} as const;

export type Language = ValueOf<typeof LANGUAGE>;
