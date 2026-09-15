import { memo } from "react";
import { Flex, Text, TextField } from "@radix-ui/themes";
import { Controller, useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";

import type { PhoneLoginFormValues } from "../../schemas/loginSchema";

const LoginPhoneController = memo(() => {
  const { t } = useTranslation("LoginPage");
  const { control, formState: { errors } } = useFormContext<PhoneLoginFormValues>();

  return (
    <Flex gap="2" width="100%">
      <Controller
        name="code"
        control={control}
        render={({ field, fieldState }) => (
          <Flex direction="column" gap="1" style={{ width: 96 }}>
            <Text as="label" size="2" weight="medium">
              +
            </Text>
            <TextField.Root size="3" type="tel" inputMode="numeric" placeholder="86" color={fieldState.error ? "red" : undefined} value={field.value} onChange={(e) => field.onChange(e.target.value.replace(/\D/g, ""))} />
            {errors.code ? (
              <Text size="1" color="red">
                {errors.code.message}
              </Text>
            ) : null}
          </Flex>
        )}
      />
      <Controller
        name="phone"
        control={control}
        render={({ field, fieldState }) => (
          <Flex direction="column" gap="1" flexGrow="1">
            <Text as="label" size="2" weight="medium">
              {t("phone")}
            </Text>
            <TextField.Root size="3" type="tel" autoComplete="tel" inputMode="numeric" placeholder={t("phone")} color={fieldState.error ? "red" : undefined} value={field.value} onChange={(e) => field.onChange(e.target.value.replace(/\D/g, ""))} />
            {errors.phone ? (
              <Text size="1" color="red">
                {errors.phone.message}
              </Text>
            ) : null}
          </Flex>
        )}
      />
    </Flex>
  );
});

LoginPhoneController.displayName = "LoginPhoneController";
export default LoginPhoneController;
