import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserProfile } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

interface SkillsCardProps {
  userId: string;
}

const SkillsCard = memo(({ userId }: SkillsCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <SkillsCardContent userId={userId} />
    </Suspense>
  );
});

SkillsCard.displayName = "SkillsCard";

const SkillsCardContent = memo(({ userId }: SkillsCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: userProfile } = useUserProfile({ id: userId });

  const skills = userProfile?.skills as Record<string, number> | undefined;

  if (!skills || Object.keys(skills).length === 0) {
    return null;
  }

  // 将技能转换为数组并按分数排序
  const skillsArray = Object.entries(skills)
    .map(([name, score]) => ({
      name,
      score: typeof score === "number" ? score : Number(score) || 0,
    }))
    .sort((a, b) => b.score - a.score);

  return (
    <div className={styles.card}>
      <div className={styles.cardHeader}>
        <h3 className={styles.title}>{t("skills")}</h3>
        <button
          className={styles.editButton}
          onClick={() => navigate(ROUTES.EDIT_PROFILE)}
          title={t("edit")}
        >
          <Edit size={16} />
        </button>
      </div>
      <div className={styles.list}>
        {skillsArray.map((skill) => (
          <div key={skill.name} className={styles.item}>
            <div className={styles.skillInfo}>
              <span className={styles.skillName}>{skill.name}</span>
              <span className={styles.skillScore}>{skill.score}/10</span>
            </div>
            <div className={styles.progressBar}>
              <div
                className={styles.progressFill}
                style={{ width: `${(skill.score / 10) * 100}%` }}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
});

SkillsCardContent.displayName = "SkillsCardContent";

export default SkillsCard;
