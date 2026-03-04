import { memo, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

import { LanguageEnum, changeLanguage, useLanguageLabel } from "nfx-ui/languages";
import { useLanguageSync } from "@/hooks/useUserPreferenceSync";

import styles from "./styles.module.css";

interface LanguageSwitcherProps {
  status?: "primary" | "default";
}

const LANGUAGE_OPTIONS: LanguageEnum[] = [LanguageEnum.EN, LanguageEnum.ZH, LanguageEnum.FR];

const LanguageSwitcher = memo(({ status = "primary" }: LanguageSwitcherProps) => {
  const { i18n } = useTranslation();
  const { syncLanguage } = useLanguageSync();
  const { getLanguageDisplayName } = useLanguageLabel();
  const [isOpen, setIsOpen] = useState(false);
  const wrapperRef = useRef<HTMLDivElement>(null);

  const currentLanguage = (i18n.language as LanguageEnum) || LanguageEnum.ZH;

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleLanguageChange = (lng: LanguageEnum) => {
    changeLanguage(lng);
    syncLanguage(lng);
    setIsOpen(false);
  };

  return (
    <div className={styles.nbSelect} ref={wrapperRef}>
      <button
        className={`${styles.selectButton} ${styles[status]}`}
        onClick={() => setIsOpen(!isOpen)}
        aria-expanded={isOpen}
        aria-haspopup="listbox"
      >
        <span className={styles.buttonText}>{getLanguageDisplayName(currentLanguage)}</span>
        <svg
          className={`${styles.chevronIcon} ${isOpen ? styles.open : ""}`}
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
        >
          <path d="m6 9 6 6 6-6" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      </button>

      <div
        className={`${styles.optionsPanel} ${styles[status]} ${isOpen ? styles.open : styles.closed}`}
      >
        <ul className={styles.optionsList} role="listbox">
          {LANGUAGE_OPTIONS.map((lang) => (
            <li
              key={lang}
              className={`${styles.option} ${lang === currentLanguage ? styles.selected : ""}`}
              onClick={() => handleLanguageChange(lang)}
              role="option"
              aria-selected={lang === currentLanguage}
            >
              <span>{getLanguageDisplayName(lang)}</span>
              {lang === currentLanguage && (
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                  <path d="M20 6L9 17l-5-5" strokeLinecap="round" strokeLinejoin="round" />
                </svg>
              )}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
});

LanguageSwitcher.displayName = "LanguageSwitcher";

export default LanguageSwitcher;
