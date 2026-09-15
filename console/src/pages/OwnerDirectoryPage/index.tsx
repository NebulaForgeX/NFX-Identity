import { memo, useState } from "react";
import { Button, Card, Flex, Table, TextField } from "@radix-ui/themes";
import { Shield } from "lucide-react";
import { useQuery } from "@tanstack/react-query";
import { useAuthRepository } from "nfx-ui/apis";
import { CardHeader, EmptyState, PageHeader } from "nfx-ui/components";
import { PageFrame } from "nfx-ui/layouts";
import type { Profile } from "nfx-ui/types";

const OwnerDirectoryPage = memo(() => {
  const auth = useAuthRepository();
  const [query, setQuery] = useState("");
  const forgers = useQuery({
    queryKey: ["owner-forger", query],
    queryFn: () => auth.ListOwnerForgerProfiles({ limit: 50, offset: 0, query }),
  });
  const authorities = useQuery({
    queryKey: ["owner-authority", query],
    queryFn: () => auth.ListOwnerAuthorityProfiles({ limit: 50, offset: 0, query }),
  });
  const forgerItems = forgers.data?.items ?? [];
  const authorityItems = authorities.data?.items ?? [];
  const ownerError = forgers.error || authorities.error;

  return (
    <PageFrame>
      <PageHeader
        icon={Shield}
        title="Owner directory"
        description="List Forger and Authority profiles on this login center."
        actions={<TextField.Root placeholder="Search profiles" value={query} onChange={(e) => setQuery(e.target.value)} />}
      />
      {ownerError ? <EmptyState icon={Shield} title="Owner role required to list all profiles." /> : null}
      <Flex direction="column" gap="4">
        <Card>
          <CardHeader icon={<Shield size={18} />} title="Forger" />
          {forgerItems.length === 0 ? (
            <EmptyState icon={Shield} title="No Forger profiles" />
          ) : (
            <Table.Root>
              <Table.Header>
                <Table.Row>
                  <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>City</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>Roles</Table.ColumnHeaderCell>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {forgerItems.map((item: Profile.Response.ForgerProfileItem) => (
                  <Table.Row key={item.profileId}>
                    <Table.Cell>{item.displayName}</Table.Cell>
                    <Table.Cell>{item.city}</Table.Cell>
                    <Table.Cell>{(item.forgerRoles ?? []).join(", ")}</Table.Cell>
                  </Table.Row>
                ))}
              </Table.Body>
            </Table.Root>
          )}
        </Card>
        <Card>
          <CardHeader icon={<Shield size={18} />} title="Authority" />
          {authorityItems.length === 0 ? (
            <EmptyState icon={Shield} title="No Authority profiles" />
          ) : (
            <Table.Root>
              <Table.Header>
                <Table.Row>
                  <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>Roles</Table.ColumnHeaderCell>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {authorityItems.map((item: Profile.Response.AuthorityProfileItem) => (
                  <Table.Row key={item.profileId}>
                    <Table.Cell>{item.displayName}</Table.Cell>
                    <Table.Cell>{(item.authorityRoles ?? []).join(", ")}</Table.Cell>
                  </Table.Row>
                ))}
              </Table.Body>
            </Table.Root>
          )}
        </Card>
        <Button variant="soft" onClick={() => { void forgers.refetch(); void authorities.refetch(); }}>
          Refresh
        </Button>
      </Flex>
    </PageFrame>
  );
});

OwnerDirectoryPage.displayName = "OwnerDirectoryPage";
export default OwnerDirectoryPage;
