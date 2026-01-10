// System Domain Types - 基于 NFX-ID Backend

export interface SystemState {
  id: string;
  version: string;
  initialized: boolean;
  initializedAt?: string;
  initializedBy?: string;
  metadata?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}
