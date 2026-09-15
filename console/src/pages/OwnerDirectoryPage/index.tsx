import { Button, Flex, Heading, Table, Text, TextField } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { useAuthRepository } from "nfx-ui/apis";
import type { Profile } from "nfx-ui/types";

export default function OwnerDirectoryPage() {
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

  return (
    <Flex direction="column" gap="5" p="5">
      <Heading>Owner directory</Heading>
      <TextField.Root placeholder="Search profiles" value={query} onChange={(e) => setQuery(e.target.value)} />
      <Heading size="4">Forger</Heading>
      <Table.Root>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>City</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Roles</Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {(forgers.data?.items ?? []).map((item: Profile.Response.ForgerProfileItem) => (
            <Table.Row key={item.profileId}>
              <Table.Cell>{item.displayName}</Table.Cell>
              <Table.Cell>{item.city}</Table.Cell>
              <Table.Cell>{(item.forgerRoles ?? []).join(", ")}</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
      <Heading size="4">Authority</Heading>
      <Table.Root>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Roles</Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {(authorities.data?.items ?? []).map((item: Profile.Response.AuthorityProfileItem) => (
            <Table.Row key={item.profileId}>
              <Table.Cell>{item.displayName}</Table.Cell>
              <Table.Cell>{(item.authorityRoles ?? []).join(", ")}</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
      <Text size="2" color="gray">
        {forgers.error || authorities.error ? "Owner role required to list all profiles." : null}
      </Text>
      <Button variant="ghost" onClick={() => { void query; }}>
        Refresh
      </Button>
    </Flex>
  );
}
