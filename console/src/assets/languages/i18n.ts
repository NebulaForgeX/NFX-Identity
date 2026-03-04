import { changeLanguage, type LanguageEnum } from "nfx-ui/languages";

/** 切换语言方法，与 Sjgz-Admin / NFX-Vault 一致使用 nfx-ui/languages。 */
export const ChangeLanguage = (lng: LanguageEnum) => {
  changeLanguage(lng);
};
