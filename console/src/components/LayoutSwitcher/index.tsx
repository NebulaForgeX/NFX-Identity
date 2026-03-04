import { memo, useEffect, useRef, useState } from "react";

import { useLayoutLabel } from "nfx-ui/languages";
import { LAYOUT_MODE_VALUES, LayoutModeEnum } from "nfx-ui/layouts";
import { useLayout } from "@/providers";
import { useLayoutSync } from "@/hooks/useUserPreferenceSync";

import styles from "./styles.module.css";

interface LayoutSwitcherProps {
  status?: "primary" | "default";
}

const LayoutSwitcher = memo(({ status = "primary" }: LayoutSwitcherProps) => {
  const { layoutMode, setLayoutMode } = useLayout();
  const { syncLayout } = useLayoutSync();
  const { getLayoutDisplayName } = useLayoutLabel();
  const [isOpen, setIsOpen] = useState(false);
  const wrapperRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleChange = (mode: LayoutModeEnum) => {
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
        <span className={styles.buttonText}>{getLayoutDisplayName(layoutMode)}</span>
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
          {LAYOUT_MODE_VALUES.map((mode) => (
            <li
              key={mode}
              className={`${styles.option} ${mode === layoutMode ? styles.selected : ""}`}
              onClick={() => handleChange(mode)}
              role="option"
              aria-selected={mode === layoutMode}
            >
              <span>{getLayoutDisplayName(mode)}</span>
              {mode === layoutMode && (
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
