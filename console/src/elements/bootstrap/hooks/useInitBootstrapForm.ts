import { zodResolver } from "@hookform/resolvers/zod";
import { useTranslation } from "react-i18next";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type BootstrapFormValues } from "../schemas/bootstrapSchema";

export const useInitBootstrapForm = () => {
  const { t } = useTranslation("elements.bootstrap");

  // 动态创建 schema，使用翻译
  const BootstrapFormSchema = z
    .object({
      Version: z.string().trim().min(1, t("version.required")).default("1.0.0"),
      AdminUsername: z
        .string()
        .trim()
        .min(1, t("admin_username.required"))
        .min(3, t("admin_username.min_length")),
      AdminPassword: z.string().min(8, t("admin_password.min_length")),
      AdminPasswordConfirm: z.string().min(8, t("admin_password_confirm.required")),
      AdminEmail: z
        .string()
        .trim()
        .refine((val) => !val || z.string().email().safeParse(val).success, {
          message: t("admin_email.invalid"),
        })
        .optional(),
      AdminPhone: z.string().trim().optional(),
      AdminCountryCode: z.string().trim().optional(),
    })
    .refine((data) => data.AdminPassword === data.AdminPasswordConfirm, {
      message: t("admin_password_confirm.mismatch"),
      path: ["AdminPasswordConfirm"],
    });

  const form = useForm<BootstrapFormValues>({
    resolver: zodResolver(BootstrapFormSchema),
    mode: "onChange",
    defaultValues: {
      Version: "1.0.0",
      AdminUsername: "",
      AdminPassword: "",
      AdminPasswordConfirm: "",
      AdminEmail: "",
      AdminPhone: "",
      AdminCountryCode: "",
    },
  });

  return form;
};
