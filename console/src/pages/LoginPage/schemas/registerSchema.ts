import { z } from "zod";

// 创建邮箱注册 Schema 工厂函数
export const createRegisterSchema = (t: (key: string) => string) => {
  return z.object({
    email: z.email(t("validation.emailInvalid")),
    verificationCode: z
      .string()
      .trim()
      .min(1, t("validation.verificationCodeRequired"))
      .length(6, t("validation.verificationCodeLength")),
    password: z
      .string()
      .min(1, t("validation.passwordRequired"))
      .min(6, t("validation.passwordMinLength")),
  });
};

// 基础 Schema（用于类型推断，不包含翻译）
const BaseRegisterSchema = z.object({
  email: z.email(),
  verificationCode: z.string().trim().min(1).length(6),
  password: z.string().min(1).min(6),
});

export type RegisterFormValues = z.infer<typeof BaseRegisterSchema>;
