import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

const IsCurrentOccupationController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const { register } = useFormContext<OccupationFormValues>();

  return (
    <div style={{ display: "flex", alignItems: "center", gap: "0.5rem" }}>
      <input
        type="checkbox"
        id="isCurrentOccupation"
        {...register("isCurrent")}
      />
      <label htmlFor="isCurrentOccupation" style={{ margin: 0 }}>
        {t("occupation.isCurrent.label")}
      </label>
    </div>
  );
});

IsCurrentOccupationController.displayName = "IsCurrentOccupationController";

export default IsCurrentOccupationController;
