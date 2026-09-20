import { Settings, User } from "lucide-react";
import { Avatar, Button, Flex, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { LucideIcon } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { scopePaths } from "@/navigations";
import { buildImageUrl, resolveAccountDisplayName, resolveAccountInitial, safeNullable } from "@/utils";

import styles from "./s.module.css";

export default function UserTopBar() {
  const { t } = useTranslation("language");
  const { data, profile, kind } = useCurrentProfile();
  const paths = scopePaths(kind);
  const accountId = safeNullable(data?.account.id);
  const displayName = resolveAccountDisplayName(profile?.displayName, accountId);
  const initial = resolveAccountInitial(profile?.displayName, accountId);
  const avatarImageId = safeNullable(profile?.avatars?.[0]?.imageId);

  return (
    <Flex align="center" justify="between" gap="4" wrap="wrap" py="4" px="6" position="sticky" top="0" className={styles.bar}>
      <Flex align="center" gap="3" minWidth="0">
        <Avatar size="2" radius="none" src={avatarImageId ? buildImageUrl(avatarImageId) : undefined} fallback={initial} />
        <Flex direction="column" minWidth="0">
          <Text size="2" weight="bold" truncate>
            {displayName}
          </Text>
          <Text size="1" color="gray" truncate>
            {kind} · Identity
          </Text>
        </Flex>
      </Flex>

      <Flex align="center" gap="2" wrap="wrap">
        <Button size="2" variant="soft" color="gray" onClick={() => routerEventEmitter.navigate({ to: paths.overview })}>
          <LucideIcon icon={User} size={14} />
          {t("header.profile")}
        </Button>
        <Button size="2" variant="soft" color="gray" onClick={() => routerEventEmitter.navigate({ to: paths.settings })}>
          <LucideIcon icon={Settings} size={14} />
          {t("sidebar.settingsItem")}
        </Button>
      </Flex>
    </Flex>
  );
}
