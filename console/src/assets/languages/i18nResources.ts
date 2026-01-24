import type { ValueOf } from "@/utils/types";

import en_AddEducationPage from "./en/AddEducationPage.json";
import en_AddOccupationPage from "./en/AddOccupationPage.json";
import en_BootstrapProvider from "./en/BootstrapProvider.json";
import en_elements_bootstrap from "./en/elements/bootstrap.json";
import en_elements_directory from "./en/elements/directory.json";
import en_EditEducationPage from "./en/EditEducationPage.json";
import en_EditOccupationPage from "./en/EditOccupationPage.json";
import en_EditPreferencePage from "./en/EditPreferencePage.json";
import en_EditProfilePage from "./en/EditProfilePage.json";
import en_LoginPage from "./en/LoginPage.json";
import en_ProfilePage from "./en/ProfilePage.json";
import en_components from "./en/components.json";
import fr_AddEducationPage from "./fr/AddEducationPage.json";
import fr_AddOccupationPage from "./fr/AddOccupationPage.json";
import fr_BootstrapProvider from "./fr/BootstrapProvider.json";
import fr_elements_bootstrap from "./fr/elements/bootstrap.json";
import fr_elements_directory from "./fr/elements/directory.json";
import fr_EditEducationPage from "./fr/EditEducationPage.json";
import fr_EditOccupationPage from "./fr/EditOccupationPage.json";
import fr_EditPreferencePage from "./fr/EditPreferencePage.json";
import fr_EditProfilePage from "./fr/EditProfilePage.json";
import fr_LoginPage from "./fr/LoginPage.json";
import fr_ProfilePage from "./fr/ProfilePage.json";
import fr_components from "./fr/components.json";
import zh_AddEducationPage from "./zh/AddEducationPage.json";
import zh_AddOccupationPage from "./zh/AddOccupationPage.json";
import zh_BootstrapProvider from "./zh/BootstrapProvider.json";
import zh_elements_bootstrap from "./zh/elements/bootstrap.json";
import zh_elements_directory from "./zh/elements/directory.json";
import zh_EditEducationPage from "./zh/EditEducationPage.json";
import zh_EditOccupationPage from "./zh/EditOccupationPage.json";
import zh_EditPreferencePage from "./zh/EditPreferencePage.json";
import zh_EditProfilePage from "./zh/EditProfilePage.json";
import zh_LoginPage from "./zh/LoginPage.json";
import zh_ProfilePage from "./zh/ProfilePage.json";
import zh_components from "./zh/components.json";

// 所有语言包内容
export const RESOURCES = {
  en: {
    BootstrapProvider: en_BootstrapProvider,
    "elements.bootstrap": en_elements_bootstrap,
    "elements.directory": en_elements_directory,
    LoginPage: en_LoginPage,
    ProfilePage: en_ProfilePage,
    AddEducationPage: en_AddEducationPage,
    AddOccupationPage: en_AddOccupationPage,
    EditEducationPage: en_EditEducationPage,
    EditOccupationPage: en_EditOccupationPage,
    EditPreferencePage: en_EditPreferencePage,
    EditProfilePage: en_EditProfilePage,
    components: en_components,
  },
  zh: {
    BootstrapProvider: zh_BootstrapProvider,
    "elements.bootstrap": zh_elements_bootstrap,
    "elements.directory": zh_elements_directory,
    LoginPage: zh_LoginPage,
    ProfilePage: zh_ProfilePage,
    AddEducationPage: zh_AddEducationPage,
    AddOccupationPage: zh_AddOccupationPage,
    EditEducationPage: zh_EditEducationPage,
    EditOccupationPage: zh_EditOccupationPage,
    EditPreferencePage: zh_EditPreferencePage,
    EditProfilePage: zh_EditProfilePage,
    components: zh_components,
  },
  fr: {
    BootstrapProvider: fr_BootstrapProvider,
    "elements.bootstrap": fr_elements_bootstrap,
    "elements.directory": fr_elements_directory,
    LoginPage: fr_LoginPage,
    ProfilePage: fr_ProfilePage,
    AddEducationPage: fr_AddEducationPage,
    AddOccupationPage: fr_AddOccupationPage,
    EditEducationPage: fr_EditEducationPage,
    EditOccupationPage: fr_EditOccupationPage,
    EditPreferencePage: fr_EditPreferencePage,
    EditProfilePage: fr_EditProfilePage,
    components: fr_components,
  },
};

// 所有命名空间
export const NAME_SPACES_MAP = {
  BootstrapProvider: "BootstrapProvider",
  "elements.bootstrap": "elements.bootstrap",
  "elements.directory": "elements.directory",
  LoginPage: "LoginPage",
  ProfilePage: "ProfilePage",
  AddEducationPage: "AddEducationPage",
  AddOccupationPage: "AddOccupationPage",
  EditEducationPage: "EditEducationPage",
  EditOccupationPage: "EditOccupationPage",
  EditPreferencePage: "EditPreferencePage",
  EditProfilePage: "EditProfilePage",
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
