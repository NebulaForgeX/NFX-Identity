import type { DataResponse } from "nfx-ui/types";
import { publicClient } from "nfx-ui/apis";

import type { InitializeSystemStateRequest } from "@/types/requests/system.request";
import type { SystemState } from "@/types/domain/system.domain";

export async function getSystemStateLatest(): Promise<SystemState> {
  const { data } = await publicClient.get<DataResponse<SystemState>>("/system/system-state/latest");
  return data.data;
}

export async function initializeSystemState(params: InitializeSystemStateRequest): Promise<void> {
  await publicClient.post("/system/system-state/initialize", params, { timeout: 240_000 });
}
