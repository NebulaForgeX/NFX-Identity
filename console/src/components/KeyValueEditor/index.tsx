import { memo, useState, useEffect } from "react";
import { useTranslation } from "react-i18next";

import { Button, Input } from "@/components";
import { Plus, Trash2 } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

export interface KeyValuePair {
  key: string;
  value: string | string[];
}

interface KeyValueEditorProps {
  label: string;
  pairs: KeyValuePair[];
  onChange: (pairs: KeyValuePair[]) => void;
  valueType?: "string" | "array";
  keyPlaceholder?: string;
  valuePlaceholder?: string;
  error?: string;
}

const KeyValueEditor = memo(
  ({ label, pairs, onChange, valueType = "string", keyPlaceholder, valuePlaceholder, error }: KeyValueEditorProps) => {
    const { t } = useTranslation("elements.directory");
    const [localPairs, setLocalPairs] = useState<KeyValuePair[]>(pairs);

    // Sync with external pairs changes
    useEffect(() => {
      setLocalPairs(pairs);
    }, [pairs]);

    const handleAdd = () => {
      const newPairs = [...localPairs, { key: "", value: valueType === "array" ? [] : "" }];
      setLocalPairs(newPairs);
      onChange(newPairs);
    };

    const handleRemove = (index: number) => {
      const newPairs = localPairs.filter((_, i) => i !== index);
      setLocalPairs(newPairs);
      onChange(newPairs);
    };

    const handleKeyChange = (index: number, key: string) => {
      const newPairs = [...localPairs];
      newPairs[index] = { ...newPairs[index], key };
      setLocalPairs(newPairs);
      onChange(newPairs);
    };

    const handleValueChange = (index: number, value: string) => {
      const newPairs = [...localPairs];
      if (valueType === "array") {
        // Split by comma and trim
        const arrayValue = value
          .split(",")
          .map((v) => v.trim())
          .filter((v) => v.length > 0);
        newPairs[index] = { ...newPairs[index], value: arrayValue };
      } else {
        newPairs[index] = { ...newPairs[index], value };
      }
      setLocalPairs(newPairs);
      onChange(newPairs);
    };

    const getDisplayValue = (value: string | string[]): string => {
      if (Array.isArray(value)) {
        return value.join(", ");
      }
      return value;
    };

    return (
      <div className={styles.container}>
        <label className={styles.label}>{label}</label>
        {error && <div className={styles.error}>{error}</div>}
        <div className={styles.pairs}>
          {localPairs.map((pair, index) => (
            <div key={index} className={styles.pair}>
              <Input
                placeholder={keyPlaceholder || t("keyValueEditor.key.placeholder")}
                value={pair.key}
                onChange={(e) => handleKeyChange(index, e.target.value)}
                className={styles.keyInput}
              />
              <Input
                placeholder={
                  valuePlaceholder ||
                  (valueType === "array"
                    ? t("keyValueEditor.value.arrayPlaceholder")
                    : t("keyValueEditor.value.placeholder"))
                }
                value={getDisplayValue(pair.value)}
                onChange={(e) => handleValueChange(index, e.target.value)}
                className={styles.valueInput}
              />
              <Button
                type="button"
                variant="ghost"
                onClick={() => handleRemove(index)}
                className={styles.removeButton}
                aria-label={t("keyValueEditor.remove")}
              >
                <Trash2 size={16} />
              </Button>
            </div>
          ))}
        </div>
        <Button type="button" variant="secondary" onClick={handleAdd} className={styles.addButton}>
          <Plus size={16} />
          {t("keyValueEditor.add")}
        </Button>
      </div>
    );
  },
);

KeyValueEditor.displayName = "KeyValueEditor";

export default KeyValueEditor;
