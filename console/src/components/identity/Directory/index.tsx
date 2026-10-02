import { ShieldCheck, UsersIcon } from "nfx-ui/icons";
import { useState } from "react";
import { Badge, Box, Button, Container, Flex, Grid, Section, Select, Table, Text, TextField } from "@radix-ui/themes";
import { AuthAuthorityRoleEnum, UI_ASSIGNABLE_AUTH_AUTHORITY_ROLES } from "nfx-ui/enums";
import { useCurrentProfile, useGetPublicProfileCard, useListOwnerAuthorityProfiles, useListOwnerForgerProfiles, useUpdateAuthorityProfileRoles } from "nfx-ui/hooks";
import type { Profile } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

import { EmptyState, PageHeader } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { scopePaths } from "@/navigations";
import { formatDateTime, safeArray, safeStringable } from "@/utils";

import { FieldList, FieldRow, LedgerSection } from "../Ledger";
import styles from "./s.module.css";

export default function DirectoryView() {
  const { t } = useTranslation("pages.Directory");
  const { kind } = useCurrentProfile();
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
      <PageHeader icon={ShieldCheck} title={t("title")} description={t("description")} />
      <Section size="1" py="4">
        <TextField.Root className={styles.search} variant="classic" placeholder={t("search")} value={query} onChange={(e) => setQuery(e.target.value)} />
      </Section>
      {ownerError ? (
        <Text size="2">
          {(ownerError as Error).message || t("ownerRequired")}
        </Text>
      ) : null}

      <Grid columns="1" gap="6" className={styles.board}>
        <Flex direction="column" gap="6" minWidth="0">
          <LedgerSection title={t("forger.title")} description={t("forger.description")}>
            {forgerItems.length === 0 ? (
              <EmptyState
                icon={UsersIcon}
                title={t("forger.empty")}
                action={
                  <Button
                    size="2"
                    onClick={() => (query ? setQuery("") : routerEventEmitter.navigate({ to: scopePaths(kind).identity }))}
                  >
                    {query ? t("search") : t("actions.createForger")}
                  </Button>
                }
              />
            ) : (
              <Box className={styles.table}>
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
                      <Table.Row
                        key={item.profileId}
                        className={selectedId === item.profileId ? styles.rowSelected : styles.rowHit}
                        onClick={() => setSelectedId(item.profileId)}
                      >
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
              </Box>
            )}
          </LedgerSection>

          <LedgerSection title={t("authority.title")} description={t("authority.description")}>
            {authorityItems.length === 0 ? (
              <EmptyState
                icon={ShieldCheck}
                title={t("authority.empty")}
                action={
                  <Button
                    size="2"
                    onClick={() => (query ? setQuery("") : routerEventEmitter.navigate({ to: scopePaths(kind).identity }))}
                  >
                    {query ? t("search") : t("actions.createAuthority")}
                  </Button>
                }
              />
            ) : (
              <Box className={styles.table}>
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
                        <Table.Row
                          key={item.profileId}
                          className={selectedId === item.profileId ? styles.rowSelected : styles.rowHit}
                          onClick={() => setSelectedId(item.profileId)}
                        >
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
              </Box>
            )}
          </LedgerSection>
        </Flex>

        <Box position="sticky" top="4" className={styles.detail}>
          {selectedId ? (
            <Section size="1" py="4">
              <Container size="2" px="4">
                <LedgerSection title={t("card.title")} description={t("card.description")}>
                  {card.isError ? (
                    <Text size="2">
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
              </Container>
            </Section>
          ) : (
            <EmptyState icon={ShieldCheck} title={t("card.title")} description={t("card.description")} />
          )}
        </Box>
      </Grid>
    </PageFrame>
  );
}
