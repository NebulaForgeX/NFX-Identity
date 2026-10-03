import { AnimatedIcon, MagnifierIcon, ShieldCheck, UsersIcon } from "nfx-ui/icons";
import { useState } from "react";
import { Badge, Box, Button, Callout, Flex, Grid, Select, Table, Text, TextField } from "@radix-ui/themes";
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
  const updateRoles = useUpdateAuthorityProfileRoles({ successMsg: t("toasts.updateRolesSuccess") });
  const card = useGetPublicProfileCard(selectedId);
  const forgerItems = safeArray(forgers.data?.items);
  const authorityItems = safeArray(authorities.data?.items);
  const ownerError = forgers.error || authorities.error;

  return (
    <PageFrame>
      <PageHeader icon={ShieldCheck} title={t("title")} description={t("description")} />
      <Box className={styles.search} data-reveal="">
        <TextField.Root size="3" placeholder={t("search")} value={query} onChange={(e) => setQuery(e.target.value)}>
          <TextField.Slot>
            <AnimatedIcon icon={MagnifierIcon} size={16} />
          </TextField.Slot>
        </TextField.Root>
      </Box>
      {ownerError ? (
        <Callout.Root color="red" variant="surface" data-reveal="">
          <Callout.Text>{(ownerError as Error).message || t("ownerRequired")}</Callout.Text>
        </Callout.Root>
      ) : null}

      <Grid columns={{ initial: "1", lg: "minmax(0, 1.6fr) minmax(0, 0.8fr)" }} gap="5" align="start">
        <Flex direction="column" gap="5" minWidth="0">
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
              <Table.Root variant="surface" size="2">
                  <Table.Header>
                    <Table.Row>
                      <Table.ColumnHeaderCell>{t("columns.name")}</Table.ColumnHeaderCell>
                      <Table.ColumnHeaderCell>{t("columns.city")}</Table.ColumnHeaderCell>
                      <Table.ColumnHeaderCell>{t("columns.roles")}</Table.ColumnHeaderCell>
                      <Table.ColumnHeaderCell>{t("columns.created")}</Table.ColumnHeaderCell>
                    </Table.Row>
                  </Table.Header>
                  <Table.Body>
                    {forgerItems.map((item: Profile.Response.ForgerProfileItem) => (
                      <Table.Row
                        key={item.profileId}
                        className={styles.row}
                        data-selected={selectedId === item.profileId ? "true" : undefined}
                        onClick={() => setSelectedId(item.profileId)}
                      >
                        <Table.RowHeaderCell>
                          <Text weight="medium">{safeStringable(item.displayName) || item.profileId}</Text>
                        </Table.RowHeaderCell>
                        <Table.Cell>{safeStringable(item.city) || "—"}</Table.Cell>
                        <Table.Cell>
                          <Flex gap="1" wrap="wrap">
                            {safeArray(item.forgerRoles).map((role) => (
                              <Badge key={role} variant="outline" color="gray" size="1">
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
              <Table.Root variant="surface" size="2">
                  <Table.Header>
                    <Table.Row>
                      <Table.ColumnHeaderCell>{t("columns.name")}</Table.ColumnHeaderCell>
                      <Table.ColumnHeaderCell>{t("columns.roles")}</Table.ColumnHeaderCell>
                      <Table.ColumnHeaderCell>{t("columns.assign")}</Table.ColumnHeaderCell>
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
                          className={styles.row}
                          data-selected={selectedId === item.profileId ? "true" : undefined}
                          onClick={() => setSelectedId(item.profileId)}
                        >
                          <Table.RowHeaderCell>
                            <Text weight="medium">{safeStringable(item.displayName) || item.profileId}</Text>
                          </Table.RowHeaderCell>
                          <Table.Cell>
                            <Flex gap="1" wrap="wrap">
                              {roles.map((role) => (
                                <Badge key={role} variant="outline" color="gray" size="1">
                                  {role}
                                </Badge>
                              ))}
                            </Flex>
                          </Table.Cell>
                          <Table.Cell onClick={(event) => event.stopPropagation()}>
                            {isOwner ? (
                              <Badge size="1" variant="surface">
                                {t("labels.owner")}
                              </Badge>
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
                                <Select.Trigger variant="ghost" />
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
            )}
          </LedgerSection>
        </Flex>

        <LedgerSection sticky title={t("card.title")} description={t("card.description")}>
          {!selectedId ? (
            <EmptyState icon={UsersIcon} title={t("card.pick")} />
          ) : card.isError ? (
            <Text size="2" color="red">
              {(card.error as Error).message}
            </Text>
          ) : card.data ? (
            <FieldList>
              <FieldRow label={t("columns.name")} value={safeStringable(card.data.displayName) || selectedId} />
              <FieldRow label={t("columns.city")} value={safeStringable(card.data.city) || "—"} />
              <FieldRow label={t("columns.country")} value={safeStringable(card.data.country) || "—"} />
              <FieldRow label={t("columns.website")} value={safeStringable(card.data.website) || "—"} />
              <FieldRow label={t("columns.timezone")} value={safeStringable(card.data.timezone) || "—"} />
              <FieldRow label={t("columns.created")} value={card.data.createdAt ? formatDateTime(card.data.createdAt) : "—"} />
            </FieldList>
          ) : (
            <Text size="2" color="gray">
              {t("card.loading")}
            </Text>
          )}
        </LedgerSection>
      </Grid>
    </PageFrame>
  );
}
