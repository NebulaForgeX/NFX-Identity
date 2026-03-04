import type { QueryKey } from "@tanstack/react-query";
import { createItemKey, createKey, createListKey } from "nfx-ui/constants";
import { DOMAIN_SYSTEM, DOMAIN_SYSTEM_STATE } from "./domain.key";

// ========== SystemState ==========
export const SYSTEM_SYSTEM_STATE_LIST = createListKey(DOMAIN_SYSTEM, DOMAIN_SYSTEM_STATE);
export const SYSTEM_SYSTEM_STATE = createItemKey(DOMAIN_SYSTEM, DOMAIN_SYSTEM_STATE);
export const SYSTEM_SYSTEM_STATE_LATEST: QueryKey = createKey(DOMAIN_SYSTEM, "item", DOMAIN_SYSTEM_STATE, "latest");
export const SYSTEM_SYSTEM_STATE_INIT: QueryKey = createKey(DOMAIN_SYSTEM, "item", DOMAIN_SYSTEM_STATE, "init", "public");

