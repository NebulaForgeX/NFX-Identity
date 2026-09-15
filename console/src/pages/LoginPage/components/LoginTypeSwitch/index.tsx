import { memo } from "react";
import { SegmentedControl } from "@radix-ui/themes";
import { useTranslation } from "react-i18next";

interface LoginTypeSwitchProps {
  loginType: "email" | "phone";
  onLoginTypeChange: (type: "email" | "phone") => void;
}

const LoginTypeSwitch = memo(({ loginType, onLoginTypeChange }: LoginTypeSwitchProps) => {
  const { t } = useTranslation("LoginPage");

  return (
    <SegmentedControl.Root size="2" value={loginType} onValueChange={(value) => onLoginTypeChange(value as "email" | "phone")}>
      <SegmentedControl.Item value="email">{t("loginType.email")}</SegmentedControl.Item>
      <SegmentedControl.Item value="phone">{t("loginType.phone")}</SegmentedControl.Item>
    </SegmentedControl.Root>
  );
});

LoginTypeSwitch.displayName = "LoginTypeSwitch";
export default LoginTypeSwitch;
