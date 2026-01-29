import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Button, Input, IconButton, Suspense } from "@/components";
import { ArrowLeft } from "@/assets/icons/lucide";
import { ROUTES } from "@/types/navigation";

import { CreateActionForm, ActionByKeyContent } from "./components";
import styles from "./styles.module.css";

const AddActionPage = memo(() => {
  const { t } = useTranslation("AddActionPage");
  const navigate = useNavigate();
  const [lookupKey, setLookupKey] = useState("");
  const [lookupActionKey, setLookupActionKey] = useState("");

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

      <CreateActionForm />

      <div className={styles.card}>
        <h3 className={styles.cardTitle}>{t("lookupActionByKey")}</h3>
        <div className={styles.formRow}>
          <Input
            label={t("key")}
            value={lookupKey}
            onChange={(e) => setLookupKey(e.target.value)}
            placeholder="e.g. directory.users.read"
            fullWidth
          />
          <Button
            variant="secondary"
            onClick={() => setLookupActionKey(lookupKey.trim())}
            disabled={!lookupKey.trim()}
          >
            {t("lookup")}
          </Button>
        </div>
        {lookupActionKey && (
          <Suspense
            loadingType="ecg"
            loadingText={t("loading")}
            loadingSize="small"
            loadingContainerClassName={styles.loading}
          >
            <ActionByKeyContent actionKey={lookupActionKey} />
          </Suspense>
        )}
      </div>
    </div>
  );
});

AddActionPage.displayName = "AddActionPage";

export default AddActionPage;
