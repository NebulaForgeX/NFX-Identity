import { LayoutDashboardIcon } from "nfx-ui/icons";
import { Badge, Button, Flex, Text } from "@radix-ui/themes";
import { useCurrentProfile } from "nfx-ui/hooks";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import type { ScopePaths } from "@/navigations";
import { formatDateTime, safeArray, safeNullable, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "./Ledger";
import Masthead, { profileRoles } from "./Masthead";

export default function DeskView({ paths }: { paths: ScopePaths }) {
  const { t } = useTranslation("pages.Desk");
  const { data, profile, kind } = useCurrentProfile();
  const roles = profileRoles(kind, data);
  const emails = safeArray(data?.emails);
  const phones = safeArray(data?.phones);
  const identities = safeArray(data?.identities);

  return (
    <PageFrame>
      <PageHeader icon={LayoutDashboardIcon} title={t("title")} description={t("description")} />
      <Masthead
        kind={kind}
        data={data}
        profile={profile}
        action={
          <Button size="2" onClick={() => routerEventEmitter.navigate({ to: paths.overview })}>
            {t("openRegistry")}
          </Button>
        }
      />
      <LedgerSection title={t("session.title")} description={t("session.description")}>
        <FieldList>
          <FieldRow label={t("session.accountId")} value={safeNullable(data?.account.id) || t("empty")} />
          <FieldRow label={t("session.status")} value={safeStringable(data?.account.accountStatus) || t("empty")} />
          <FieldRow label={t("session.platform")} value={safeStringable(data?.account.signupPlatform) || t("empty")} />
          <FieldRow label={t("session.kind")} value={kind} />
          <FieldRow label={t("session.profileId")} value={safeNullable(profile?.profileId) || t("empty")} />
          <FieldRow label={t("session.roles")}>
            <Flex gap="2" wrap="wrap">
              {roles.length ? roles.map((role) => <Badge key={role} variant="outline" size="1">{role}</Badge>) : <Text size="2">{t("empty")}</Text>}
            </Flex>
          </FieldRow>
          <FieldRow label={t("session.emails")} value={String(emails.length)} />
          <FieldRow label={t("session.phones")} value={String(phones.length)} />
          <FieldRow label={t("session.links")} value={String(identities.length)} />
          <FieldRow label={t("session.loginNotification")} value={profile?.settings?.loginNotification ? t("on") : t("off")} />
          <FieldRow label={t("session.created")} value={data?.account.createdAt ? formatDateTime(data.account.createdAt) : t("empty")} />
        </FieldList>
      </LedgerSection>
    </PageFrame>
  );
}
