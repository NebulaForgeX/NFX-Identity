import { ShieldCheck } from "nfx-ui/icons";
import { useState } from "react";
import { Badge, Flex, Select, Table, Text, TextField } from "@radix-ui/themes";
import { AuthAuthorityRoleEnum, UI_ASSIGNABLE_AUTH_AUTHORITY_ROLES } from "nfx-ui/enums";
import { useGetPublicProfileCard, useListOwnerAuthorityProfiles, useListOwnerForgerProfiles, useUpdateAuthorityProfileRoles } from "nfx-ui/hooks";
import type { Profile } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { formatDateTime, safeArray, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "./Ledger";
import styles from "./Ledger/s.module.css";

export default function DirectoryView() {
  const { t } = useTranslation("pages.Directory");
  const [query, setQuery] = useState("");
  const [selectedId, setSelectedId] = useState("");
  const forgers = useListOwnerForgerProfiles(query);
  const authorities = useListOwnerAuthorityProfiles(query);
  const updateRoles = useUpdateAuthorityProfileRoles();
  const card = useGetPublicProfileCard(selectedId);
  const forgerItems = safeArray(forgers.data?.items);
  const authorityItems = safeArray(authorities.data?.items);
  const ownerError = forgers.error || authorities.error;

  return (
    <PageFrame>
      <PageHeader
        icon={ShieldCheck}
        title={t("title")}
        description={t("description")}
        actions={
          <div className={styles.search}>
            <TextField.Root variant="classic" placeholder={t("search")} value={query} onChange={(e) => setQuery(e.target.value)} />
          </div>
        }
      />
      {ownerError ? (
        <Text size="2" color="red" mb="3">
          {(ownerError as Error).message || t("ownerRequired")}
        </Text>
      ) : null}

      <LedgerSection title={t("forger.title")} description={t("forger.description")}>
        {forgerItems.length === 0 ? (
          <Text size="2" color="gray">
            {t("forger.empty")}
          </Text>
        ) : (
          <div className={styles.tableWrap}>
          <Table.Root variant="ghost" size="2">
            <Table.Header>
              <Table.Row>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.name")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.city")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.roles")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.created")}</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {forgerItems.map((item: Profile.Response.ForgerProfileItem) => (
                <Table.Row key={item.profileId} style={{ cursor: "pointer" }} onClick={() => setSelectedId(item.profileId)}>
                  <Table.Cell>{safeStringable(item.displayName) || item.profileId}</Table.Cell>
                  <Table.Cell>{safeStringable(item.city) || "—"}</Table.Cell>
                  <Table.Cell>
                    <Flex gap="1" wrap="wrap">
                      {safeArray(item.forgerRoles).map((role) => (
                        <Badge key={role} variant="outline" size="1">
                          {role}
                        </Badge>
                      ))}
                    </Flex>
                  </Table.Cell>
                  <Table.Cell>{item.createdAt ? formatDateTime(item.createdAt) : "—"}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table.Root>
          </div>
        )}
      </LedgerSection>

      <LedgerSection title={t("authority.title")} description={t("authority.description")}>
        {authorityItems.length === 0 ? (
          <Text size="2" color="gray">
            {t("authority.empty")}
          </Text>
        ) : (
          <div className={styles.tableWrap}>
          <Table.Root variant="ghost" size="2">
            <Table.Header>
              <Table.Row>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.name")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.roles")}</Table.ColumnHeaderCell>
                <Table.ColumnHeaderCell className={styles.head}>{t("columns.assign")}</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {authorityItems.map((item: Profile.Response.AuthorityProfileItem) => {
                const roles = safeArray(item.authorityRoles);
                const isOwner = roles.includes(AuthAuthorityRoleEnum.OWNER);
                const current = roles.find((role) => role !== AuthAuthorityRoleEnum.OWNER) ?? AuthAuthorityRoleEnum.ADMINISTRATOR;
                return (
                  <Table.Row key={item.profileId} style={{ cursor: "pointer" }} onClick={() => setSelectedId(item.profileId)}>
                    <Table.Cell>{safeStringable(item.displayName) || item.profileId}</Table.Cell>
                    <Table.Cell>{roles.join(", ")}</Table.Cell>
                    <Table.Cell>
                      {isOwner ? (
                        <Text size="2">owner</Text>
                      ) : (
                        <Select.Root
                          value={current}
                          onValueChange={(value) => {
                            updateRoles.mutate({
                              profileId: item.profileId,
                              authorityRoles: [value as AuthAuthorityRoleEnum],
                            });
                          }}
                        >
                          <Select.Trigger />
                          <Select.Content>
                            {UI_ASSIGNABLE_AUTH_AUTHORITY_ROLES.map((role) => (
                              <Select.Item key={role} value={role}>
                                {role}
                              </Select.Item>
                            ))}
                          </Select.Content>
                        </Select.Root>
                      )}
                    </Table.Cell>
                  </Table.Row>
                );
              })}
            </Table.Body>
          </Table.Root>
          </div>
        )}
      </LedgerSection>

      {selectedId ? (
        <LedgerSection title={t("card.title")} description={t("card.description")}>
          {card.isError ? (
            <Text size="2" color="red">
              {(card.error as Error).message}
            </Text>
          ) : card.data ? (
            <FieldList>
              <FieldRow label={t("columns.name")} value={safeStringable(card.data.displayName) || selectedId} />
              <FieldRow label={t("columns.city")} value={safeStringable(card.data.city) || "—"} />
              <FieldRow label={t("columns.country")} value={safeStringable(card.data.country) || "—"} />
              <FieldRow label="website" value={safeStringable(card.data.website) || "—"} />
              <FieldRow label="timezone" value={safeStringable(card.data.timezone) || "—"} />
              <FieldRow label={t("columns.created")} value={card.data.createdAt ? formatDateTime(card.data.createdAt) : "—"} />
            </FieldList>
          ) : (
            <Text size="2" color="gray">
              {t("card.loading")}
            </Text>
          )}
        </LedgerSection>
      ) : null}
    </PageFrame>
  );
}
