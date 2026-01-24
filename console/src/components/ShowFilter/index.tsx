import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Eye, EyeOff, Filter } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

export interface ShowFilterValue {
  enabled: boolean;
  value: boolean | null; // null = 全部, true = 显示, false = 隐藏
}

interface ShowFilterProps {
  value: ShowFilterValue;
  onChange: (value: ShowFilterValue) => void;
}

const ShowFilter = memo(({ value, onChange }: ShowFilterProps) => {
  const { t } = useTranslation("components");
  
  const handleToggleEnabled = () => {
    onChange({ ...value, enabled: !value.enabled });
  };

  const handleSelectShow = (show: boolean | null) => {
    onChange({ ...value, value: show });
  };

  return (
    <div className={styles.container}>
      {/* Toggle 开关 */}
      <div className={styles.toggleContainer}>
        <button
          type="button"
          className={`${styles.toggleButton} ${value.enabled ? styles.enabled : ""}`}
          onClick={handleToggleEnabled}
          aria-label={value.enabled ? t("showFilter.disableFilter") : t("showFilter.enableFilter")}
        >
          <Filter size={16} />
          <span>{value.enabled ? t("showFilter.filterEnabled") : t("showFilter.filterDisabled")}</span>
        </button>
      </div>

      {/* 过滤选项 */}
      {value.enabled && (
        <div className={styles.optionsContainer}>
          <button
            type="button"
            className={`${styles.option} ${value.value === null ? styles.active : ""}`}
            onClick={() => handleSelectShow(null)}
          >
            {t("showFilter.all")}
          </button>
          <button
            type="button"
            className={`${styles.option} ${value.value === true ? styles.active : ""}`}
            onClick={() => handleSelectShow(true)}
          >
            <Eye size={16} />
            <span>{t("showFilter.show")}</span>
          </button>
          <button
            type="button"
            className={`${styles.option} ${value.value === false ? styles.active : ""}`}
            onClick={() => handleSelectShow(false)}
          >
            <EyeOff size={16} />
            <span>{t("showFilter.hide")}</span>
          </button>
        </div>
      )}
    </div>
  );
});

ShowFilter.displayName = "ShowFilter";

export default ShowFilter;
