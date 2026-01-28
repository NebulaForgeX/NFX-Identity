import { useMemo } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useTranslation } from "react-i18next";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type EducationFormValues } from "../schemas/educationSchema";
import type { UserEducation } from "@/types";

export const useInitEducationForm = (education?: UserEducation) => {
  const { t } = useTranslation("elements.directory");

  // 使用 useMemo 缓存 schema，避免每次渲染都创建新的 schema
  const EducationFormSchema = useMemo(
    () =>
      z.object({
        school: z.string().trim().min(1, t("education.school.required")),
        degree: z.string().trim().optional(),
        major: z.string().trim().optional(),
        field: z.string().trim().optional(), // 映射到 fieldOfStudy
        startDate: z.string().optional(),
        endDate: z.string().optional(),
        isCurrent: z.boolean().optional(),
        description: z.string().trim().optional(),
        grade: z.string().trim().optional(),
        activities: z.string().trim().optional(),
        achievements: z.string().trim().optional(),
      }),
    [t]
  );

  // 缓存 resolver 以确保稳定性
  const resolver = useMemo(() => zodResolver(EducationFormSchema), [EducationFormSchema]);

  const form = useForm<EducationFormValues>({
    resolver,
    mode: "onChange",
    defaultValues: {
      school: education?.school || "",
      degree: education?.degree || "",
      major: education?.major || "",
      fieldOfStudy: education?.fieldOfStudy || "",
      startDate: education?.startDate ? new Date(education.startDate).toISOString().split("T")[0] : "",
      endDate: education?.endDate ? new Date(education.endDate).toISOString().split("T")[0] : "",
      isCurrent: education?.isCurrent || false,
      description: education?.description || "",
      grade: education?.grade || "",
      activities: education?.activities || "",
      achievements: education?.achievements || "",
    },
  });

  return form;
};
