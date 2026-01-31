import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const IsCurrentEducationController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const { register } = useFormContext<EducationFormValues>();

  return (
    <div style={{ display: "flex", alignItems: "center", gap: "0.5rem" }}>
      <input
        type="checkbox"
        id="isCurrent"
        {...register("isCurrent")}
      />
      <label htmlFor="isCurrent" style={{ margin: 0 }}>
        {t("education.isCurrent.label")}
      </label>
    </div>
  );
});

IsCurrentEducationController.displayName = "IsCurrentEducationController";

export default IsCurrentEducationController;
