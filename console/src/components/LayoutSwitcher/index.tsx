import { memo, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

import { useLayout } from "@/providers/LayoutProvider/useLayout";
import { useLayoutSync } from "@/hooks/useUserPreferenceSync";

import styles from "./styles.module.css";

interface LayoutSwitcherProps {
  status?: "primary" | "default";
}

const LayoutSwitcher = memo(({ status = "primary" }: LayoutSwitcherProps) => {
  const { t } = useTranslation("components");
  const { layoutMode, setLayoutMode } = useLayout();
  const { syncLayout } = useLayoutSync();
  const [isOpen, setIsOpen] = useState(false);
  const wrapperRef = useRef<HTMLDivElement>(null);

  const getDisplayName = (mode: "show" | "hide"): string => {
    return mode === "show" ? t("layoutSwitcher.showSidebar") : t("layoutSwitcher.hideSidebar");
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleChange = (mode: "show" | "hide") => {
    setLayoutMode(mode);
    syncLayout(mode);
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
        <span className={styles.buttonText}>{getDisplayName(layoutMode)}</span>
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
          {(["show", "hide"] as const).map((m) => (
            <li
              key={m}
              className={`${styles.option} ${m === layoutMode ? styles.selected : ""}`}
              onClick={() => handleChange(m)}
              role="option"
              aria-selected={m === layoutMode}
            >
              <span>{getDisplayName(m)}</span>
              {m === layoutMode && (
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

LayoutSwitcher.displayName = "LayoutSwitcher";

export default LayoutSwitcher;
