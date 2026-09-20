import type { CreateI18nResourcesResult, NameSpacesMap, Resources } from "nfx-ui/languages";

import enHooks from "./en/hooks.json";
import enLanguage from "./en/language.json";
import enAuthShell from "./en/pages/Account/AuthShell.json";
import enLogin from "./en/pages/Account/Login.json";
import enSelectProfile from "./en/pages/Account/SelectProfile.json";
import enSignup from "./en/pages/Account/Signup.json";
import enAssets from "./en/pages/Assets.json";
import enDesk from "./en/pages/Desk.json";
import enDirectory from "./en/pages/Directory.json";
import enProfileIdentity from "./en/pages/Profile/Identity.json";
import enProfileOverview from "./en/pages/Profile/Overview.json";
import enProfileSecurity from "./en/pages/Profile/Security.json";
import enUserProfileEdit from "./en/pages/User/Profile/Edit.json";
import enUserSetting from "./en/pages/User/Setting.json";
import frHooks from "./fr/hooks.json";
import frLanguage from "./fr/language.json";
import frAuthShell from "./fr/pages/Account/AuthShell.json";
import frLogin from "./fr/pages/Account/Login.json";
import frSelectProfile from "./fr/pages/Account/SelectProfile.json";
import frSignup from "./fr/pages/Account/Signup.json";
import frAssets from "./fr/pages/Assets.json";
import frDesk from "./fr/pages/Desk.json";
import frDirectory from "./fr/pages/Directory.json";
import frProfileIdentity from "./fr/pages/Profile/Identity.json";
import frProfileOverview from "./fr/pages/Profile/Overview.json";
import frProfileSecurity from "./fr/pages/Profile/Security.json";
import frUserProfileEdit from "./fr/pages/User/Profile/Edit.json";
import frUserSetting from "./fr/pages/User/Setting.json";
import zhHooks from "./zh/hooks.json";
import zhLanguage from "./zh/language.json";
import zhAuthShell from "./zh/pages/Account/AuthShell.json";
import zhLogin from "./zh/pages/Account/Login.json";
import zhSelectProfile from "./zh/pages/Account/SelectProfile.json";
import zhSignup from "./zh/pages/Account/Signup.json";
import zhAssets from "./zh/pages/Assets.json";
import zhDesk from "./zh/pages/Desk.json";
import zhDirectory from "./zh/pages/Directory.json";
import zhProfileIdentity from "./zh/pages/Profile/Identity.json";
import zhProfileOverview from "./zh/pages/Profile/Overview.json";
import zhProfileSecurity from "./zh/pages/Profile/Security.json";
import zhUserProfileEdit from "./zh/pages/User/Profile/Edit.json";
import zhUserSetting from "./zh/pages/User/Setting.json";

const PAGE = {
  AuthShell: "pages.Account.AuthShell",
  Login: "pages.Account.Login",
  Signup: "pages.Account.Signup",
  SelectProfile: "pages.Account.SelectProfile",
  UserSetting: "pages.User.Setting",
  UserProfileEdit: "pages.User.Profile.Edit",
  Desk: "pages.Desk",
  ProfileOverview: "pages.Profile.Overview",
  ProfileIdentity: "pages.Profile.Identity",
  ProfileSecurity: "pages.Profile.Security",
  Assets: "pages.Assets",
  Directory: "pages.Directory",
} as const;

const BUILTIN_I18N_NAMESPACES_MAP: NameSpacesMap = {
  language: "language",
  hooks: "hooks",
  ...PAGE,
};

const BUILTIN_I18N_NAMESPACES = Object.values(BUILTIN_I18N_NAMESPACES_MAP);

const BUILTIN_I18N_RESOURCES: Resources = {
  en: {
    language: enLanguage,
    hooks: enHooks,
    [PAGE.AuthShell]: enAuthShell,
    [PAGE.Login]: enLogin,
    [PAGE.Signup]: enSignup,
    [PAGE.SelectProfile]: enSelectProfile,
    [PAGE.UserSetting]: enUserSetting,
    [PAGE.UserProfileEdit]: enUserProfileEdit,
    [PAGE.Desk]: enDesk,
    [PAGE.ProfileOverview]: enProfileOverview,
    [PAGE.ProfileIdentity]: enProfileIdentity,
    [PAGE.ProfileSecurity]: enProfileSecurity,
    [PAGE.Assets]: enAssets,
    [PAGE.Directory]: enDirectory,
  },
  zh: {
    language: zhLanguage,
    hooks: zhHooks,
    [PAGE.AuthShell]: zhAuthShell,
    [PAGE.Login]: zhLogin,
    [PAGE.Signup]: zhSignup,
    [PAGE.SelectProfile]: zhSelectProfile,
    [PAGE.UserSetting]: zhUserSetting,
    [PAGE.UserProfileEdit]: zhUserProfileEdit,
    [PAGE.Desk]: zhDesk,
    [PAGE.ProfileOverview]: zhProfileOverview,
    [PAGE.ProfileIdentity]: zhProfileIdentity,
    [PAGE.ProfileSecurity]: zhProfileSecurity,
    [PAGE.Assets]: zhAssets,
    [PAGE.Directory]: zhDirectory,
  },
  fr: {
    language: frLanguage,
    hooks: frHooks,
    [PAGE.AuthShell]: frAuthShell,
    [PAGE.Login]: frLogin,
    [PAGE.Signup]: frSignup,
    [PAGE.SelectProfile]: frSelectProfile,
    [PAGE.UserSetting]: frUserSetting,
    [PAGE.UserProfileEdit]: frUserProfileEdit,
    [PAGE.Desk]: frDesk,
    [PAGE.ProfileOverview]: frProfileOverview,
    [PAGE.ProfileIdentity]: frProfileIdentity,
    [PAGE.ProfileSecurity]: frProfileSecurity,
    [PAGE.Assets]: frAssets,
    [PAGE.Directory]: frDirectory,
  },
};

export function getBuiltinI18nBundles(): CreateI18nResourcesResult {
  return {
    RESOURCES: BUILTIN_I18N_RESOURCES,
    NAME_SPACES_MAP: BUILTIN_I18N_NAMESPACES_MAP,
    NAME_SPACES: [...BUILTIN_I18N_NAMESPACES],
  };
}
