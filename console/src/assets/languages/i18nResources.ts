import type { ValueOf } from "nfx-ui/types";

import en_components from "./en/components.json";
import en_EditPreferencePage from "./en/EditPreferencePage.json";
import en_ImagesPage from "./en/ImagesPage.json";
import en_LoginPage from "./en/LoginPage.json";
import en_ProfilePage from "./en/ProfilePage.json";
import fr_components from "./fr/components.json";
import fr_EditPreferencePage from "./fr/EditPreferencePage.json";
import fr_ImagesPage from "./fr/ImagesPage.json";
import fr_LoginPage from "./fr/LoginPage.json";
import fr_ProfilePage from "./fr/ProfilePage.json";
import zh_components from "./zh/components.json";
import zh_EditPreferencePage from "./zh/EditPreferencePage.json";
import zh_ImagesPage from "./zh/ImagesPage.json";
import zh_LoginPage from "./zh/LoginPage.json";
import zh_ProfilePage from "./zh/ProfilePage.json";

export const RESOURCES = {
  en: {
    LoginPage: en_LoginPage,
    ProfilePage: en_ProfilePage,
    ImagesPage: en_ImagesPage,
    EditPreferencePage: en_EditPreferencePage,
    components: en_components,
  },
  zh: {
    LoginPage: zh_LoginPage,
    ProfilePage: zh_ProfilePage,
    ImagesPage: zh_ImagesPage,
    EditPreferencePage: zh_EditPreferencePage,
    components: zh_components,
  },
  fr: {
    LoginPage: fr_LoginPage,
    ProfilePage: fr_ProfilePage,
    ImagesPage: fr_ImagesPage,
    EditPreferencePage: fr_EditPreferencePage,
    components: fr_components,
  },
};

export const NAME_SPACES_MAP = {
  LoginPage: "LoginPage",
  ProfilePage: "ProfilePage",
  ImagesPage: "ImagesPage",
  EditPreferencePage: "EditPreferencePage",
  components: "components",
};

export const NAME_SPACES = Object.values(NAME_SPACES_MAP);

export function getBuiltinI18nBundles() {
  return { RESOURCES, NAME_SPACES_MAP, NAME_SPACES };
}

export const LANGUAGE = {
  EN: "en",
  ZH: "zh",
  FR: "fr",
} as const;

export type Language = ValueOf<typeof LANGUAGE>;
