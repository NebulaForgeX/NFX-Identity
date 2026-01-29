import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { IconButton } from "@/components";
import { ArrowLeft } from "@/assets/icons/lucide";
import { ROUTES } from "@/types/navigation";

import CreateRoleForm from "./components/CreateRoleForm";
import styles from "./styles.module.css";

const AddRolePage = memo(() => {
  const { t } = useTranslation("AddRolePage");
  const navigate = useNavigate();

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <IconButton
          variant="ghost"
          leftIcon={<ArrowLeft size={20} />}
          onClick={() => navigate(ROUTES.PERMISSION_MANAGEMENT)}
          className={styles.backButton}
        >
          {t("back")}
        </IconButton>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <CreateRoleForm />
    </div>
  );
});

AddRolePage.displayName = "AddRolePage";

export default AddRolePage;
