import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext, useFieldArray, Controller } from "react-hook-form";

import { Button, Input, Slider } from "@/components";
import { Plus, Trash2 } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

const SkillsController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    control,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  const { fields, append, remove } = useFieldArray({
    control,
    name: "skills",
  });

  const handleAdd = () => {
    append({ name: "", score: 0 });
  };

  const handleRemove = (index: number) => {
    remove(index);
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <label className={styles.label}>{t("profile.skills.label")}</label>
        <Button type="button" variant="secondary" onClick={handleAdd} className={styles.addButton}>
          <Plus size={16} />
          {t("keyValueEditor.add")}
        </Button>
      </div>
      {errors.skills && <div className={styles.error}>{errors.skills.message as string}</div>}
      <div className={styles.skillsList}>
        {fields.map((field, index) => (
          <div key={field.id} className={styles.skillItem}>
            <Controller
              control={control}
              name={`skills.${index}.name` as const}
              render={({ field: nameField }) => (
                <Input
                  placeholder={t("profile.skills.namePlaceholder")}
                  value={nameField.value || ""}
                  onChange={nameField.onChange}
                  onBlur={nameField.onBlur}
                  className={styles.nameInput}
                  fullWidth
                />
              )}
            />
            <div className={styles.sliderWrapper}>
              <Controller
                control={control}
                name={`skills.${index}.score` as const}
                render={({ field: scoreField }) => (
                  <Slider
                    value={typeof scoreField.value === "number" ? scoreField.value : 0}
                    onChange={(value) => {
                      scoreField.onChange(value);
                      scoreField.onBlur();
                    }}
                    min={0}
                    max={10}
                    step={1}
                    showValue={true}
                    className={styles.slider}
                  />
                )}
              />
            </div>
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
    </div>
  );
});

SkillsController.displayName = "SkillsController";

export default SkillsController;
