import { memo, useState } from "react";
import { useTranslation } from "react-i18next";

import { Button, Input } from "@/components";
import { useCreateAction } from "@/hooks/useAccess";

import styles from "../styles.module.css";

const CreateActionForm = memo(() => {
  const { t } = useTranslation("AddActionPage");
  const [key, setKey] = useState("");
  const [name, setName] = useState("");
  const [service, setService] = useState("");
  const [status, setStatus] = useState("active");
  const [description, setDescription] = useState("");

  const createAction = useCreateAction();

  const handleCreate = () => {
    if (!key.trim() || !name.trim() || !service.trim()) return;
    createAction.mutate(
      {
        key: key.trim(),
        name: name.trim(),
        service: service.trim(),
        status: status.trim() || "active",
        description: description.trim() || undefined,
        isSystem: false,
      },
      {
        onSuccess: () => {
          setKey("");
          setName("");
          setService("");
          setStatus("active");
          setDescription("");
        },
      },
    );
  };

  return (
    <div className={styles.card}>
      <h3 className={styles.cardTitle}>{t("createAction")}</h3>
      <div className={styles.formRow}>
        <Input
          label={t("key")}
          value={key}
          onChange={(e) => setKey(e.target.value)}
          placeholder="e.g. directory.users.read"
          fullWidth
        />
        <Input
          label={t("name")}
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Users Read"
          fullWidth
        />
      </div>
      <div className={styles.formRow}>
        <Input
          label={t("service")}
          value={service}
          onChange={(e) => setService(e.target.value)}
          placeholder="e.g. directory"
          fullWidth
        />
        <Input
          label={t("status")}
          value={status}
          onChange={(e) => setStatus(e.target.value)}
          placeholder="active"
          fullWidth
        />
      </div>
      <Input
        label={t("description")}
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        fullWidth
      />
      <div className={styles.submitWrap}>
        <Button onClick={handleCreate} loading={createAction.isPending}>
          {t("createAction")}
        </Button>
      </div>
    </div>
  );
});
CreateActionForm.displayName = "CreateActionForm";

export default CreateActionForm;
