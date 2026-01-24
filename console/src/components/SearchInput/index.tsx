import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Search, X } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

interface SearchInputProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
}

const SearchInput = memo(({ value, onChange, placeholder }: SearchInputProps) => {
  const { t } = useTranslation("components");
  const defaultPlaceholder = placeholder ?? t("searchInput.placeholder");
  
  const handleClear = () => {
    onChange("");
  };

  return (
    <div className={styles.searchContainer}>
      <Search size={18} className={styles.searchIcon} />
      <input
        type="text"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={defaultPlaceholder}
        className={styles.searchInput}
      />
      {value && (
        <button onClick={handleClear} className={styles.clearBtn} aria-label={t("searchInput.clearSearch")}>
          <X size={16} />
        </button>
      )}
    </div>
  );
});

SearchInput.displayName = "SearchInput";

export default SearchInput;
