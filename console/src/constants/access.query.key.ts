import type { QueryKey } from "@tanstack/react-query";
import { createItemKey, createKey, createListKey } from "nfx-ui/constants";
import { DOMAIN_ACCESS, DOMAIN_ACCESS_ROLE, DOMAIN_ACCESS_PERMISSION, DOMAIN_ACCESS_SCOPE, DOMAIN_ACCESS_GRANT, DOMAIN_ACCESS_ROLE_PERMISSION, DOMAIN_ACCESS_SCOPE_PERMISSION, DOMAIN_ACCESS_ACTION, DOMAIN_ACCESS_ACTION_REQUIREMENT } from "./domain.key";

export const ACCESS_ROLE_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ROLE);
export const ACCESS_ROLE = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ROLE);

export const ACCESS_PERMISSION_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_PERMISSION);
export const ACCESS_PERMISSION = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_PERMISSION);

export const ACCESS_SCOPE_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_SCOPE);
export const ACCESS_SCOPE = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_SCOPE);

export const ACCESS_GRANT_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_GRANT);
export const ACCESS_GRANT = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_GRANT);
export const ACCESS_GRANTS_BY_SUBJECT = (subjectType: string, subjectId: string, tenantId?: string): QueryKey =>
  tenantId != null
    ? createKey(DOMAIN_ACCESS, "list", DOMAIN_ACCESS_GRANT, "by-subject", subjectType, subjectId, tenantId)
    : createKey(DOMAIN_ACCESS, "list", DOMAIN_ACCESS_GRANT, "by-subject", subjectType, subjectId);

export const ACCESS_ROLE_PERMISSION_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ROLE_PERMISSION);
export const ACCESS_ROLE_PERMISSION = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ROLE_PERMISSION);
export const ACCESS_ROLE_PERMISSIONS_BY_ROLE = (roleId: string): QueryKey =>
  createKey(DOMAIN_ACCESS, "list", DOMAIN_ACCESS_ROLE_PERMISSION, "by-role", roleId);

export const ACCESS_ACTION_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ACTION);
export const ACCESS_ACTION = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ACTION);

export const ACCESS_ACTION_REQUIREMENT_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ACTION_REQUIREMENT);
export const ACCESS_ACTION_REQUIREMENT = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_ACTION_REQUIREMENT);
export const ACCESS_ACTION_REQUIREMENTS_BY_PERMISSION = (permissionId: string): QueryKey =>
  createKey(DOMAIN_ACCESS, "list", DOMAIN_ACCESS_ACTION_REQUIREMENT, "by-permission", permissionId);

export const ACCESS_SCOPE_PERMISSION_LIST = createListKey(DOMAIN_ACCESS, DOMAIN_ACCESS_SCOPE_PERMISSION);
export const ACCESS_SCOPE_PERMISSION = createItemKey(DOMAIN_ACCESS, DOMAIN_ACCESS_SCOPE_PERMISSION);

