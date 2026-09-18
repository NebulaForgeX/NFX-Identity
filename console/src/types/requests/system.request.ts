export interface InitializeSystemStateRequest {
  version?: string;
  adminUsername: string;
  adminPassword: string;
  adminEmail: string;
  adminPhone?: string;
}

export interface ResetSystemStateRequest {
  metadata?: Record<string, unknown>;
}
