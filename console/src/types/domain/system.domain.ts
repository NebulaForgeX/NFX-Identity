export interface SystemState {
  id: string;
  initialized: boolean;
  initializedAt?: string;
  initializationVersion?: string;
  lastResetAt?: string;
  lastResetBy?: string;
  resetCount?: number;
  metadata?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}
