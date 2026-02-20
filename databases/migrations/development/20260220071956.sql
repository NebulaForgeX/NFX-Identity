-- Add new schema named "access"
CREATE SCHEMA "access";
-- Set comment to schema: "access"
COMMENT ON SCHEMA "access" IS 'Authorization and access control schema';
-- Add new schema named "audit"
CREATE SCHEMA "audit";
-- Set comment to schema: "audit"
COMMENT ON SCHEMA "audit" IS 'Enterprise audit and compliance schema';
-- Add new schema named "auth"
CREATE SCHEMA "auth";
-- Set comment to schema: "auth"
COMMENT ON SCHEMA "auth" IS 'Authentication and authorization schema';
-- Add new schema named "clients"
CREATE SCHEMA "clients";
-- Set comment to schema: "clients"
COMMENT ON SCHEMA "clients" IS 'Client credentials and application management schema';
-- Add new schema named "directory"
CREATE SCHEMA "directory";
-- Set comment to schema: "directory"
COMMENT ON SCHEMA "directory" IS 'User directory and profile schema';
-- Add new schema named "image"
CREATE SCHEMA "image";
-- Set comment to schema: "image"
COMMENT ON SCHEMA "image" IS 'Image schema';
-- Add new schema named "system"
CREATE SCHEMA "system";
-- Set comment to schema: "system"
COMMENT ON SCHEMA "system" IS 'System-level state and administration schema';
-- Add new schema named "tenants"
CREATE SCHEMA "tenants";
-- Set comment to schema: "tenants"
COMMENT ON SCHEMA "tenants" IS 'Multi-tenant management schema';
-- Create extension "pgcrypto"
CREATE EXTENSION "pgcrypto" WITH SCHEMA "public" VERSION "1.3";
-- Create extension "btree_gist"
CREATE EXTENSION "btree_gist" WITH SCHEMA "public" VERSION "1.8";
-- Create enum type "environment"
CREATE TYPE "clients"."environment" AS ENUM ('production', 'staging', 'development', 'test');
-- Create enum type "failure_code"
CREATE TYPE "auth"."failure_code" AS ENUM ('bad_password', 'user_not_found', 'locked', 'mfa_required', 'mfa_failed', 'account_disabled', 'credential_expired', 'rate_limited', 'ip_blocked', 'device_not_trusted', 'other');
-- Create enum type "scope_type"
CREATE TYPE "access"."scope_type" AS ENUM ('TENANT', 'APP', 'GLOBAL');
-- Create enum type "subject_type"
CREATE TYPE "access"."subject_type" AS ENUM ('USER', 'CLIENT');
-- Create enum type "grant_type"
CREATE TYPE "access"."grant_type" AS ENUM ('ROLE', 'PERMISSION');
-- Create enum type "grant_effect"
CREATE TYPE "access"."grant_effect" AS ENUM ('ALLOW', 'DENY');
-- Create enum type "credential_status"
CREATE TYPE "clients"."credential_status" AS ENUM ('active', 'expired', 'revoked', 'rotating');
-- Create "actions" table
CREATE TABLE "access"."actions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" character varying(255) NOT NULL,
  "service" character varying(255) NOT NULL,
  "status" character varying(50) NOT NULL DEFAULT 'active',
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "actions_key_key" UNIQUE ("key")
);
-- Create index "idx_actions_deleted_at" to table: "actions"
CREATE INDEX "idx_actions_deleted_at" ON "access"."actions" ("deleted_at");
-- Create index "idx_actions_is_system" to table: "actions"
CREATE INDEX "idx_actions_is_system" ON "access"."actions" ("is_system");
-- Create index "idx_actions_key" to table: "actions"
CREATE INDEX "idx_actions_key" ON "access"."actions" ("key");
-- Create index "idx_actions_service" to table: "actions"
CREATE INDEX "idx_actions_service" ON "access"."actions" ("service");
-- Create index "idx_actions_status" to table: "actions"
CREATE INDEX "idx_actions_status" ON "access"."actions" ("status");
-- Create "grants" table
CREATE TABLE "access"."grants" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "subject_type" "access"."subject_type" NOT NULL,
  "subject_id" uuid NOT NULL,
  "grant_type" "access"."grant_type" NOT NULL,
  "grant_ref_id" uuid NOT NULL,
  "tenant_id" uuid NULL,
  "app_id" uuid NULL,
  "resource_type" character varying(100) NULL,
  "resource_id" uuid NULL,
  "effect" "access"."grant_effect" NOT NULL DEFAULT 'ALLOW',
  "expires_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_grants_app_id" to table: "grants"
CREATE INDEX "idx_grants_app_id" ON "access"."grants" ("app_id");
-- Create index "idx_grants_effect" to table: "grants"
CREATE INDEX "idx_grants_effect" ON "access"."grants" ("effect");
-- Create index "idx_grants_expires_at" to table: "grants"
CREATE INDEX "idx_grants_expires_at" ON "access"."grants" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_grants_grant_type" to table: "grants"
CREATE INDEX "idx_grants_grant_type" ON "access"."grants" ("grant_type", "grant_ref_id");
-- Create index "idx_grants_resource" to table: "grants"
CREATE INDEX "idx_grants_resource" ON "access"."grants" ("resource_type", "resource_id");
-- Create index "idx_grants_revoked_at" to table: "grants"
CREATE INDEX "idx_grants_revoked_at" ON "access"."grants" ("revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_grants_subject" to table: "grants"
CREATE INDEX "idx_grants_subject" ON "access"."grants" ("subject_type", "subject_id");
-- Create index "idx_grants_subject_tenant" to table: "grants"
CREATE INDEX "idx_grants_subject_tenant" ON "access"."grants" ("subject_type", "subject_id", "tenant_id");
-- Create index "idx_grants_tenant_id" to table: "grants"
CREATE INDEX "idx_grants_tenant_id" ON "access"."grants" ("tenant_id");
-- Create "system_state" table
CREATE TABLE "system"."system_state" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "initialized" boolean NOT NULL DEFAULT false,
  "initialized_at" timestamp NULL,
  "initialization_version" character varying(50) NULL,
  "last_reset_at" timestamp NULL,
  "last_reset_by" uuid NULL,
  "reset_count" integer NOT NULL DEFAULT 0,
  "metadata" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_system_state_created_at" to table: "system_state"
CREATE INDEX "idx_system_state_created_at" ON "system"."system_state" ("created_at" DESC);
-- Create index "idx_system_state_initialized" to table: "system_state"
CREATE INDEX "idx_system_state_initialized" ON "system"."system_state" ("initialized");
-- Set comment to table: "system_state"
COMMENT ON TABLE "system"."system_state" IS 'System initialization state: tracks if system has been bootstrapped. Query logic: SELECT initialized FROM system_state ORDER BY created_at DESC LIMIT 1. If no record exists OR initialized = false, system is not initialized.';
-- Set comment to column: "id" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."id" IS 'UUID primary key (not fixed, allows multiple records for reset/re-initialization)';
-- Set comment to column: "initialized" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialized" IS 'Whether system has been initialized (checked on service startup). Always check latest record by created_at DESC.';
-- Set comment to column: "initialized_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialized_at" IS 'Timestamp when system was initialized via /bootstrap/initialize';
-- Set comment to column: "initialization_version" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialization_version" IS 'Version of initialization schema/data for migration tracking';
-- Set comment to column: "last_reset_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."last_reset_at" IS 'Timestamp when system was last reset';
-- Set comment to column: "last_reset_by" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."last_reset_by" IS 'User ID who reset the system. Even if database is cleared after reset, can be traced via log files for accountability';
-- Set comment to column: "created_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."created_at" IS 'Record creation time. Used to determine latest state (ORDER BY created_at DESC LIMIT 1)';
-- Create enum type "verification_status"
CREATE TYPE "tenants"."verification_status" AS ENUM ('PENDING', 'VERIFIED', 'FAILED', 'EXPIRED');
-- Create enum type "verification_method"
CREATE TYPE "tenants"."verification_method" AS ENUM ('DNS', 'TXT', 'HTML', 'FILE');
-- Create enum type "group_type"
CREATE TYPE "tenants"."group_type" AS ENUM ('department', 'team', 'group', 'other');
-- Create enum type "tenant_app_status"
CREATE TYPE "tenants"."tenant_app_status" AS ENUM ('ACTIVE', 'DISABLED', 'SUSPENDED');
-- Create enum type "actor_type"
CREATE TYPE "audit"."actor_type" AS ENUM ('user', 'service', 'system', 'admin');
-- Create enum type "result_type"
CREATE TYPE "audit"."result_type" AS ENUM ('success', 'failure', 'deny', 'error');
-- Create enum type "risk_level"
CREATE TYPE "audit"."risk_level" AS ENUM ('low', 'medium', 'high', 'critical');
-- Create enum type "data_classification"
CREATE TYPE "audit"."data_classification" AS ENUM ('public', 'internal', 'confidential', 'restricted');
-- Create enum type "retention_action"
CREATE TYPE "audit"."retention_action" AS ENUM ('archive', 'delete', 'export');
-- Create "actor_snapshots" table
CREATE TABLE "audit"."actor_snapshots" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "actor_type" "audit"."actor_type" NOT NULL,
  "actor_id" uuid NOT NULL,
  "display_name" character varying(255) NULL,
  "email" character varying(255) NULL,
  "client_name" character varying(255) NULL,
  "tenant_id" uuid NULL,
  "snapshot_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "snapshot_data" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "actor_snapshots_actor_type_actor_id_snapshot_at_key" UNIQUE ("actor_type", "actor_id", "snapshot_at")
);
-- Create index "idx_actor_snapshots_actor" to table: "actor_snapshots"
CREATE INDEX "idx_actor_snapshots_actor" ON "audit"."actor_snapshots" ("actor_type", "actor_id", "snapshot_at" DESC);
-- Create index "idx_actor_snapshots_snapshot_at" to table: "actor_snapshots"
CREATE INDEX "idx_actor_snapshots_snapshot_at" ON "audit"."actor_snapshots" ("snapshot_at");
-- Create index "idx_actor_snapshots_tenant_id" to table: "actor_snapshots"
CREATE INDEX "idx_actor_snapshots_tenant_id" ON "audit"."actor_snapshots" ("tenant_id");
-- Create "event_retention_policies" table
CREATE TABLE "audit"."event_retention_policies" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "policy_name" character varying(255) NOT NULL,
  "tenant_id" uuid NULL,
  "action_pattern" character varying(255) NULL,
  "data_classification" "audit"."data_classification" NULL,
  "risk_level" "audit"."risk_level" NULL,
  "retention_days" integer NOT NULL,
  "retention_action" "audit"."retention_action" NOT NULL DEFAULT 'archive',
  "archive_location" text NULL,
  "status" character varying(50) NOT NULL DEFAULT 'active',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "event_retention_policies_policy_name_key" UNIQUE ("policy_name")
);
-- Create index "idx_event_retention_policies_policy_name" to table: "event_retention_policies"
CREATE INDEX "idx_event_retention_policies_policy_name" ON "audit"."event_retention_policies" ("policy_name");
-- Create index "idx_event_retention_policies_status" to table: "event_retention_policies"
CREATE INDEX "idx_event_retention_policies_status" ON "audit"."event_retention_policies" ("status");
-- Create index "idx_event_retention_policies_tenant_id" to table: "event_retention_policies"
CREATE INDEX "idx_event_retention_policies_tenant_id" ON "audit"."event_retention_policies" ("tenant_id");
-- Create "event_search_index" table
CREATE TABLE "audit"."event_search_index" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "event_id" character varying(255) NOT NULL,
  "tenant_id" uuid NULL,
  "app_id" uuid NULL,
  "actor_type" "audit"."actor_type" NOT NULL,
  "actor_id" uuid NOT NULL,
  "action" character varying(255) NOT NULL,
  "target_type" character varying(100) NULL,
  "target_id" uuid NULL,
  "result" "audit"."result_type" NOT NULL,
  "occurred_at" timestamp NOT NULL,
  "ip" inet NULL,
  "risk_level" "audit"."risk_level" NOT NULL DEFAULT 'low',
  "data_classification" "audit"."data_classification" NOT NULL DEFAULT 'internal',
  "tags" text[] NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "event_search_index_event_id_key" UNIQUE ("event_id")
);
-- Create index "idx_event_search_index_action" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_action" ON "audit"."event_search_index" ("action", "occurred_at");
-- Create index "idx_event_search_index_actor" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_actor" ON "audit"."event_search_index" ("actor_type", "actor_id", "occurred_at");
-- Create index "idx_event_search_index_event_id" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_event_id" ON "audit"."event_search_index" ("event_id");
-- Create index "idx_event_search_index_result" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_result" ON "audit"."event_search_index" ("result", "occurred_at");
-- Create index "idx_event_search_index_risk_level" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_risk_level" ON "audit"."event_search_index" ("risk_level", "occurred_at");
-- Create index "idx_event_search_index_tags" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_tags" ON "audit"."event_search_index" USING gin ("tags");
-- Create index "idx_event_search_index_target" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_target" ON "audit"."event_search_index" ("target_type", "target_id");
-- Create index "idx_event_search_index_tenant_occurred" to table: "event_search_index"
CREATE INDEX "idx_event_search_index_tenant_occurred" ON "audit"."event_search_index" ("tenant_id", "occurred_at");
-- Create "events" table
CREATE TABLE "audit"."events" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "event_id" character varying(255) NOT NULL,
  "occurred_at" timestamp NOT NULL,
  "received_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "tenant_id" uuid NULL,
  "app_id" uuid NULL,
  "actor_type" "audit"."actor_type" NOT NULL,
  "actor_id" uuid NOT NULL,
  "actor_tenant_member_id" uuid NULL,
  "action" character varying(255) NOT NULL,
  "target_type" character varying(100) NULL,
  "target_id" uuid NULL,
  "result" "audit"."result_type" NOT NULL,
  "failure_reason_code" character varying(100) NULL,
  "http_method" character varying(10) NULL,
  "http_path" character varying(500) NULL,
  "http_status" integer NULL,
  "request_id" character varying(255) NULL,
  "trace_id" character varying(255) NULL,
  "ip" inet NULL,
  "user_agent" text NULL,
  "geo_country" character varying(10) NULL,
  "risk_level" "audit"."risk_level" NOT NULL DEFAULT 'low',
  "data_classification" "audit"."data_classification" NOT NULL DEFAULT 'internal',
  "prev_hash" character varying(64) NULL,
  "event_hash" character varying(64) NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "events_event_id_key" UNIQUE ("event_id")
);
-- Create index "idx_events_action" to table: "events"
CREATE INDEX "idx_events_action" ON "audit"."events" ("action");
-- Create index "idx_events_action_occurred" to table: "events"
CREATE INDEX "idx_events_action_occurred" ON "audit"."events" ("action", "occurred_at");
-- Create index "idx_events_actor" to table: "events"
CREATE INDEX "idx_events_actor" ON "audit"."events" ("actor_type", "actor_id");
-- Create index "idx_events_actor_occurred" to table: "events"
CREATE INDEX "idx_events_actor_occurred" ON "audit"."events" ("actor_type", "actor_id", "occurred_at");
-- Create index "idx_events_app_id" to table: "events"
CREATE INDEX "idx_events_app_id" ON "audit"."events" ("app_id");
-- Create index "idx_events_data_classification" to table: "events"
CREATE INDEX "idx_events_data_classification" ON "audit"."events" ("data_classification");
-- Create index "idx_events_event_id" to table: "events"
CREATE INDEX "idx_events_event_id" ON "audit"."events" ("event_id");
-- Create index "idx_events_occurred_at" to table: "events"
CREATE INDEX "idx_events_occurred_at" ON "audit"."events" ("occurred_at");
-- Create index "idx_events_request_id" to table: "events"
CREATE INDEX "idx_events_request_id" ON "audit"."events" ("request_id");
-- Create index "idx_events_result" to table: "events"
CREATE INDEX "idx_events_result" ON "audit"."events" ("result");
-- Create index "idx_events_risk_level" to table: "events"
CREATE INDEX "idx_events_risk_level" ON "audit"."events" ("risk_level");
-- Create index "idx_events_target" to table: "events"
CREATE INDEX "idx_events_target" ON "audit"."events" ("target_type", "target_id");
-- Create index "idx_events_tenant_action" to table: "events"
CREATE INDEX "idx_events_tenant_action" ON "audit"."events" ("tenant_id", "action", "occurred_at");
-- Create index "idx_events_tenant_id" to table: "events"
CREATE INDEX "idx_events_tenant_id" ON "audit"."events" ("tenant_id");
-- Create index "idx_events_tenant_occurred" to table: "events"
CREATE INDEX "idx_events_tenant_occurred" ON "audit"."events" ("tenant_id", "occurred_at");
-- Create index "idx_events_trace_id" to table: "events"
CREATE INDEX "idx_events_trace_id" ON "audit"."events" ("trace_id");
-- Create "hash_chain_checkpoints" table
CREATE TABLE "audit"."hash_chain_checkpoints" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "checkpoint_id" character varying(255) NOT NULL,
  "tenant_id" uuid NULL,
  "partition_date" date NOT NULL,
  "checkpoint_hash" character varying(64) NOT NULL,
  "prev_checkpoint_hash" character varying(64) NULL,
  "event_count" integer NOT NULL,
  "first_event_id" character varying(255) NULL,
  "last_event_id" character varying(255) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" character varying(255) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "hash_chain_checkpoints_checkpoint_id_key" UNIQUE ("checkpoint_id")
);
-- Create index "idx_hash_chain_checkpoints_checkpoint_id" to table: "hash_chain_checkpoints"
CREATE INDEX "idx_hash_chain_checkpoints_checkpoint_id" ON "audit"."hash_chain_checkpoints" ("checkpoint_id");
-- Create index "idx_hash_chain_checkpoints_partition_date" to table: "hash_chain_checkpoints"
CREATE INDEX "idx_hash_chain_checkpoints_partition_date" ON "audit"."hash_chain_checkpoints" ("partition_date" DESC);
-- Create index "idx_hash_chain_checkpoints_prev_hash" to table: "hash_chain_checkpoints"
CREATE INDEX "idx_hash_chain_checkpoints_prev_hash" ON "audit"."hash_chain_checkpoints" ("prev_checkpoint_hash");
-- Create index "idx_hash_chain_checkpoints_tenant_id" to table: "hash_chain_checkpoints"
CREATE INDEX "idx_hash_chain_checkpoints_tenant_id" ON "audit"."hash_chain_checkpoints" ("tenant_id");
-- Create enum type "credential_type"
CREATE TYPE "auth"."credential_type" AS ENUM ('password', 'passkey', 'oauth_link', 'saml', 'ldap');
-- Create enum type "credential_status"
CREATE TYPE "auth"."credential_status" AS ENUM ('active', 'disabled', 'expired');
-- Create enum type "user_status"
CREATE TYPE "directory"."user_status" AS ENUM ('pending', 'active', 'deactive');
-- Create enum type "lock_reason"
CREATE TYPE "auth"."lock_reason" AS ENUM ('too_many_attempts', 'admin_lock', 'risk_detected', 'suspicious_activity', 'compliance', 'other');
-- Create enum type "reset_delivery"
CREATE TYPE "auth"."reset_delivery" AS ENUM ('email', 'sms');
-- Create enum type "reset_status"
CREATE TYPE "auth"."reset_status" AS ENUM ('issued', 'used', 'expired', 'revoked');
-- Create enum type "mfa_type"
CREATE TYPE "auth"."mfa_type" AS ENUM ('totp', 'sms', 'email', 'webauthn', 'backup_code');
-- Create enum type "session_revoke_reason"
CREATE TYPE "auth"."session_revoke_reason" AS ENUM ('user_logout', 'admin_revoke', 'password_changed', 'device_changed', 'account_locked', 'suspicious_activity', 'session_expired', 'other');
-- Create "account_lockouts" table
CREATE TABLE "auth"."account_lockouts" (
  "user_id" uuid NOT NULL,
  "locked_until" timestamp NULL,
  "lock_reason" "auth"."lock_reason" NOT NULL,
  "locked_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "locked_by" character varying(255) NULL,
  "actor_id" uuid NULL,
  "unlocked_at" timestamp NULL,
  "unlocked_by" character varying(255) NULL,
  "unlock_actor_id" uuid NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("user_id")
);
-- Create index "idx_account_lockouts_locked_at" to table: "account_lockouts"
CREATE INDEX "idx_account_lockouts_locked_at" ON "auth"."account_lockouts" ("locked_at");
-- Create index "idx_account_lockouts_locked_until" to table: "account_lockouts"
CREATE INDEX "idx_account_lockouts_locked_until" ON "auth"."account_lockouts" ("locked_until") WHERE (locked_until IS NOT NULL);
-- Create "login_attempts" table
CREATE TABLE "auth"."login_attempts" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "identifier" character varying(255) NOT NULL,
  "user_id" uuid NULL,
  "ip" inet NULL,
  "ua_hash" character varying(64) NULL,
  "device_fingerprint" character varying(255) NULL,
  "success" boolean NOT NULL DEFAULT false,
  "failure_code" "auth"."failure_code" NULL,
  "mfa_required" boolean NOT NULL DEFAULT false,
  "mfa_verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_login_attempts_created_at" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_created_at" ON "auth"."login_attempts" ("created_at");
-- Create index "idx_login_attempts_identifier" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_identifier" ON "auth"."login_attempts" ("identifier", "created_at");
-- Create index "idx_login_attempts_ip" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_ip" ON "auth"."login_attempts" ("ip", "created_at");
-- Create index "idx_login_attempts_success" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_success" ON "auth"."login_attempts" ("success", "created_at");
-- Create index "idx_login_attempts_user_id" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_user_id" ON "auth"."login_attempts" ("user_id", "created_at");
-- Create "mfa_factors" table
CREATE TABLE "auth"."mfa_factors" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "factor_id" character varying(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "type" "auth"."mfa_type" NOT NULL,
  "secret_encrypted" text NULL,
  "phone" character varying(20) NULL,
  "email" character varying(255) NULL,
  "name" character varying(255) NULL,
  "enabled" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "last_used_at" timestamp NULL,
  "recovery_codes_hash" text NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "mfa_factors_factor_id_key" UNIQUE ("factor_id")
);
-- Create index "idx_mfa_factors_deleted_at" to table: "mfa_factors"
CREATE INDEX "idx_mfa_factors_deleted_at" ON "auth"."mfa_factors" ("deleted_at");
-- Create index "idx_mfa_factors_enabled" to table: "mfa_factors"
CREATE INDEX "idx_mfa_factors_enabled" ON "auth"."mfa_factors" ("user_id", "enabled") WHERE (enabled = true);
-- Create index "idx_mfa_factors_factor_id" to table: "mfa_factors"
CREATE INDEX "idx_mfa_factors_factor_id" ON "auth"."mfa_factors" ("factor_id");
-- Create index "idx_mfa_factors_type" to table: "mfa_factors"
CREATE INDEX "idx_mfa_factors_type" ON "auth"."mfa_factors" ("type");
-- Create index "idx_mfa_factors_user_id" to table: "mfa_factors"
CREATE INDEX "idx_mfa_factors_user_id" ON "auth"."mfa_factors" ("user_id");
-- Create "password_history" table
CREATE TABLE "auth"."password_history" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "password_hash" character varying(255) NOT NULL,
  "hash_alg" character varying(50) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_password_history_created_at" to table: "password_history"
CREATE INDEX "idx_password_history_created_at" ON "auth"."password_history" ("created_at");
-- Create index "idx_password_history_user_id" to table: "password_history"
CREATE INDEX "idx_password_history_user_id" ON "auth"."password_history" ("user_id", "created_at" DESC);
-- Create "password_resets" table
CREATE TABLE "auth"."password_resets" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "reset_id" character varying(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "delivery" "auth"."reset_delivery" NOT NULL,
  "code_hash" character varying(255) NOT NULL,
  "expires_at" timestamp NOT NULL,
  "used_at" timestamp NULL,
  "requested_ip" inet NULL,
  "ua_hash" character varying(64) NULL,
  "attempt_count" integer NOT NULL DEFAULT 0,
  "status" "auth"."reset_status" NOT NULL DEFAULT 'issued',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "password_resets_reset_id_key" UNIQUE ("reset_id")
);
-- Create index "idx_password_resets_created_at" to table: "password_resets"
CREATE INDEX "idx_password_resets_created_at" ON "auth"."password_resets" ("created_at");
-- Create index "idx_password_resets_expires_at" to table: "password_resets"
CREATE INDEX "idx_password_resets_expires_at" ON "auth"."password_resets" ("expires_at");
-- Create index "idx_password_resets_reset_id" to table: "password_resets"
CREATE INDEX "idx_password_resets_reset_id" ON "auth"."password_resets" ("reset_id");
-- Create index "idx_password_resets_status" to table: "password_resets"
CREATE INDEX "idx_password_resets_status" ON "auth"."password_resets" ("status");
-- Create index "idx_password_resets_user_id" to table: "password_resets"
CREATE INDEX "idx_password_resets_user_id" ON "auth"."password_resets" ("user_id");
-- Create enum type "revoke_reason"
CREATE TYPE "auth"."revoke_reason" AS ENUM ('user_logout', 'admin_revoke', 'password_changed', 'rotation', 'account_locked', 'device_changed', 'suspicious_activity', 'other');
-- Create "refresh_tokens" table
CREATE TABLE "auth"."refresh_tokens" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "token_id" character varying(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "app_id" uuid NULL,
  "client_id" character varying(255) NULL,
  "session_id" uuid NULL,
  "issued_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" timestamp NOT NULL,
  "revoked_at" timestamp NULL,
  "revoke_reason" "auth"."revoke_reason" NULL,
  "rotated_from" uuid NULL,
  "device_id" character varying(255) NULL,
  "ip" inet NULL,
  "ua_hash" character varying(64) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "refresh_tokens_token_id_key" UNIQUE ("token_id"),
  CONSTRAINT "refresh_tokens_rotated_from_fkey" FOREIGN KEY ("rotated_from") REFERENCES "auth"."refresh_tokens" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_refresh_tokens_app_id" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_app_id" ON "auth"."refresh_tokens" ("app_id");
-- Create index "idx_refresh_tokens_expires_at" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_expires_at" ON "auth"."refresh_tokens" ("expires_at");
-- Create index "idx_refresh_tokens_revoked_at" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_revoked_at" ON "auth"."refresh_tokens" ("revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_refresh_tokens_session_id" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_session_id" ON "auth"."refresh_tokens" ("session_id");
-- Create index "idx_refresh_tokens_token_id" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_token_id" ON "auth"."refresh_tokens" ("token_id");
-- Create index "idx_refresh_tokens_user_id" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_user_id" ON "auth"."refresh_tokens" ("user_id");
-- Create "sessions" table
CREATE TABLE "auth"."sessions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "session_id" character varying(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "app_id" uuid NULL,
  "client_id" character varying(255) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "last_seen_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" timestamp NOT NULL,
  "ip" inet NULL,
  "ua_hash" character varying(64) NULL,
  "device_id" character varying(255) NULL,
  "device_fingerprint" character varying(255) NULL,
  "device_name" character varying(255) NULL,
  "revoked_at" timestamp NULL,
  "revoke_reason" "auth"."session_revoke_reason" NULL,
  "revoked_by" character varying(255) NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "sessions_session_id_key" UNIQUE ("session_id")
);
-- Create index "idx_sessions_app_id" to table: "sessions"
CREATE INDEX "idx_sessions_app_id" ON "auth"."sessions" ("app_id");
-- Create index "idx_sessions_expires_at" to table: "sessions"
CREATE INDEX "idx_sessions_expires_at" ON "auth"."sessions" ("expires_at");
-- Create index "idx_sessions_last_seen_at" to table: "sessions"
CREATE INDEX "idx_sessions_last_seen_at" ON "auth"."sessions" ("last_seen_at");
-- Create index "idx_sessions_revoked_at" to table: "sessions"
CREATE INDEX "idx_sessions_revoked_at" ON "auth"."sessions" ("revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_sessions_session_id" to table: "sessions"
CREATE INDEX "idx_sessions_session_id" ON "auth"."sessions" ("session_id");
-- Create index "idx_sessions_user_active" to table: "sessions"
CREATE INDEX "idx_sessions_user_active" ON "auth"."sessions" ("user_id", "revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_sessions_user_id" to table: "sessions"
CREATE INDEX "idx_sessions_user_id" ON "auth"."sessions" ("user_id");
-- Create "trusted_devices" table
CREATE TABLE "auth"."trusted_devices" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "device_id" character varying(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "device_fingerprint_hash" character varying(255) NOT NULL,
  "device_name" character varying(255) NULL,
  "trusted_until" timestamp NOT NULL,
  "last_used_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "ip" inet NULL,
  "ua_hash" character varying(64) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "trusted_devices_user_id_device_id_key" UNIQUE ("user_id", "device_id")
);
-- Create index "idx_trusted_devices_device_id" to table: "trusted_devices"
CREATE INDEX "idx_trusted_devices_device_id" ON "auth"."trusted_devices" ("device_id");
-- Create index "idx_trusted_devices_trusted_until" to table: "trusted_devices"
CREATE INDEX "idx_trusted_devices_trusted_until" ON "auth"."trusted_devices" ("trusted_until");
-- Create index "idx_trusted_devices_user_id" to table: "trusted_devices"
CREATE INDEX "idx_trusted_devices_user_id" ON "auth"."trusted_devices" ("user_id");
-- Create "user_credentials" table
CREATE TABLE "auth"."user_credentials" (
  "id" uuid NOT NULL,
  "credential_type" "auth"."credential_type" NOT NULL DEFAULT 'password',
  "password_hash" character varying(255) NULL,
  "hash_alg" character varying(50) NULL,
  "hash_params" jsonb NULL DEFAULT '{}',
  "password_updated_at" timestamp NULL,
  "last_success_login_at" timestamp NULL,
  "status" "auth"."credential_status" NOT NULL DEFAULT 'active',
  "must_change_password" boolean NOT NULL DEFAULT false,
  "version" integer NOT NULL DEFAULT 1,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_credentials_deleted_at" to table: "user_credentials"
CREATE INDEX "idx_user_credentials_deleted_at" ON "auth"."user_credentials" ("deleted_at");
-- Create index "idx_user_credentials_status" to table: "user_credentials"
CREATE INDEX "idx_user_credentials_status" ON "auth"."user_credentials" ("status");
-- Create index "idx_user_credentials_type" to table: "user_credentials"
CREATE INDEX "idx_user_credentials_type" ON "auth"."user_credentials" ("credential_type");
-- Create enum type "app_type"
CREATE TYPE "clients"."app_type" AS ENUM ('server', 'service', 'internal', 'partner', 'third_party');
-- Create enum type "app_status"
CREATE TYPE "clients"."app_status" AS ENUM ('active', 'disabled', 'suspended', 'pending');
-- Create enum type "invitation_status"
CREATE TYPE "tenants"."invitation_status" AS ENUM ('PENDING', 'ACCEPTED', 'EXPIRED', 'REVOKED');
-- Create enum type "allowlist_status"
CREATE TYPE "clients"."allowlist_status" AS ENUM ('active', 'disabled', 'revoked');
-- Create enum type "rate_limit_type"
CREATE TYPE "clients"."rate_limit_type" AS ENUM ('requests_per_second', 'requests_per_minute', 'requests_per_hour', 'requests_per_day');
-- Create enum type "member_source"
CREATE TYPE "tenants"."member_source" AS ENUM ('MANUAL', 'INVITE', 'SCIM', 'SSO', 'HR_SYNC', 'IMPORT');
-- Create enum type "api_key_status"
CREATE TYPE "clients"."api_key_status" AS ENUM ('active', 'revoked', 'expired');
-- Create enum type "member_status"
CREATE TYPE "tenants"."member_status" AS ENUM ('INVITED', 'ACTIVE', 'SUSPENDED', 'REMOVED');
-- Create enum type "tenant_status"
CREATE TYPE "tenants"."tenant_status" AS ENUM ('ACTIVE', 'SUSPENDED', 'CLOSED', 'PENDING');
-- Create "permissions" table
CREATE TABLE "access"."permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "permissions_key_key" UNIQUE ("key")
);
-- Create index "idx_permissions_deleted_at" to table: "permissions"
CREATE INDEX "idx_permissions_deleted_at" ON "access"."permissions" ("deleted_at");
-- Create index "idx_permissions_is_system" to table: "permissions"
CREATE INDEX "idx_permissions_is_system" ON "access"."permissions" ("is_system");
-- Create index "idx_permissions_key" to table: "permissions"
CREATE INDEX "idx_permissions_key" ON "access"."permissions" ("key");
-- Create "action_requirements" table
CREATE TABLE "access"."action_requirements" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "action_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "group_id" integer NOT NULL DEFAULT 1,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "action_requirements_action_id_permission_id_key" UNIQUE ("action_id", "permission_id"),
  CONSTRAINT "action_requirements_action_id_fkey" FOREIGN KEY ("action_id") REFERENCES "access"."actions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "action_requirements_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "access"."permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_action_requirements_action_group" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_action_group" ON "access"."action_requirements" ("action_id", "group_id");
-- Create index "idx_action_requirements_action_id" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_action_id" ON "access"."action_requirements" ("action_id");
-- Create index "idx_action_requirements_permission_id" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_permission_id" ON "access"."action_requirements" ("permission_id");
-- Create "apps" table
CREATE TABLE "clients"."apps" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "app_id" character varying(255) NOT NULL,
  "tenant_id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "type" "clients"."app_type" NOT NULL DEFAULT 'server',
  "status" "clients"."app_status" NOT NULL DEFAULT 'pending',
  "environment" "clients"."environment" NOT NULL DEFAULT 'development',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "updated_by" uuid NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "apps_app_id_key" UNIQUE ("app_id")
);
-- Create index "idx_apps_app_id" to table: "apps"
CREATE INDEX "idx_apps_app_id" ON "clients"."apps" ("app_id");
-- Create index "idx_apps_deleted_at" to table: "apps"
CREATE INDEX "idx_apps_deleted_at" ON "clients"."apps" ("deleted_at");
-- Create index "idx_apps_environment" to table: "apps"
CREATE INDEX "idx_apps_environment" ON "clients"."apps" ("environment");
-- Create index "idx_apps_status" to table: "apps"
CREATE INDEX "idx_apps_status" ON "clients"."apps" ("status");
-- Create index "idx_apps_tenant_environment" to table: "apps"
CREATE INDEX "idx_apps_tenant_environment" ON "clients"."apps" ("tenant_id", "environment");
-- Create index "idx_apps_tenant_id" to table: "apps"
CREATE INDEX "idx_apps_tenant_id" ON "clients"."apps" ("tenant_id");
-- Create "api_keys" table
CREATE TABLE "clients"."api_keys" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key_id" character varying(255) NOT NULL,
  "app_id" uuid NOT NULL,
  "key_hash" character varying(255) NOT NULL,
  "hash_alg" character varying(50) NOT NULL,
  "name" character varying(255) NOT NULL,
  "status" "clients"."api_key_status" NOT NULL DEFAULT 'active',
  "expires_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  "last_used_at" timestamp NULL,
  "created_by" uuid NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "api_keys_key_id_key" UNIQUE ("key_id"),
  CONSTRAINT "api_keys_app_id_fkey" FOREIGN KEY ("app_id") REFERENCES "clients"."apps" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_api_keys_app_id" to table: "api_keys"
CREATE INDEX "idx_api_keys_app_id" ON "clients"."api_keys" ("app_id");
-- Create index "idx_api_keys_app_status" to table: "api_keys"
CREATE INDEX "idx_api_keys_app_status" ON "clients"."api_keys" ("app_id", "status") WHERE (status = 'active'::clients.api_key_status);
-- Create index "idx_api_keys_expires_at" to table: "api_keys"
CREATE INDEX "idx_api_keys_expires_at" ON "clients"."api_keys" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_api_keys_key_id" to table: "api_keys"
CREATE INDEX "idx_api_keys_key_id" ON "clients"."api_keys" ("key_id");
-- Create index "idx_api_keys_last_used_at" to table: "api_keys"
CREATE INDEX "idx_api_keys_last_used_at" ON "clients"."api_keys" ("last_used_at");
-- Create index "idx_api_keys_status" to table: "api_keys"
CREATE INDEX "idx_api_keys_status" ON "clients"."api_keys" ("status");
-- Create "client_credentials" table
CREATE TABLE "clients"."client_credentials" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "app_id" uuid NOT NULL,
  "client_id" character varying(255) NOT NULL,
  "secret_hash" character varying(255) NOT NULL,
  "hash_alg" character varying(50) NOT NULL,
  "status" "clients"."credential_status" NOT NULL DEFAULT 'active',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "rotated_at" timestamp NULL,
  "expires_at" timestamp NULL,
  "last_used_at" timestamp NULL,
  "created_by" uuid NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "client_credentials_client_id_key" UNIQUE ("client_id"),
  CONSTRAINT "client_credentials_app_id_fkey" FOREIGN KEY ("app_id") REFERENCES "clients"."apps" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_client_credentials_app_id" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_app_id" ON "clients"."client_credentials" ("app_id");
-- Create index "idx_client_credentials_app_status" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_app_status" ON "clients"."client_credentials" ("app_id", "status") WHERE (status = 'active'::clients.credential_status);
-- Create index "idx_client_credentials_client_id" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_client_id" ON "clients"."client_credentials" ("client_id");
-- Create index "idx_client_credentials_expires_at" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_expires_at" ON "clients"."client_credentials" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_client_credentials_last_used_at" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_last_used_at" ON "clients"."client_credentials" ("last_used_at");
-- Create index "idx_client_credentials_status" to table: "client_credentials"
CREATE INDEX "idx_client_credentials_status" ON "clients"."client_credentials" ("status");
-- Create "client_scopes" table
CREATE TABLE "clients"."client_scopes" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "app_id" uuid NOT NULL,
  "scope" character varying(255) NOT NULL,
  "granted_by" uuid NULL,
  "granted_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "client_scopes_app_id_scope_key" UNIQUE ("app_id", "scope"),
  CONSTRAINT "client_scopes_app_id_fkey" FOREIGN KEY ("app_id") REFERENCES "clients"."apps" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_client_scopes_app_id" to table: "client_scopes"
CREATE INDEX "idx_client_scopes_app_id" ON "clients"."client_scopes" ("app_id");
-- Create index "idx_client_scopes_app_scope" to table: "client_scopes"
CREATE INDEX "idx_client_scopes_app_scope" ON "clients"."client_scopes" ("app_id", "scope");
-- Create index "idx_client_scopes_expires_at" to table: "client_scopes"
CREATE INDEX "idx_client_scopes_expires_at" ON "clients"."client_scopes" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_client_scopes_revoked_at" to table: "client_scopes"
CREATE INDEX "idx_client_scopes_revoked_at" ON "clients"."client_scopes" ("revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_client_scopes_scope" to table: "client_scopes"
CREATE INDEX "idx_client_scopes_scope" ON "clients"."client_scopes" ("scope");
-- Create "tenants" table
CREATE TABLE "tenants"."tenants" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tenant_id" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "display_name" character varying(255) NULL,
  "status" "tenants"."tenant_status" NOT NULL DEFAULT 'PENDING',
  "primary_domain" character varying(255) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "tenants_tenant_id_key" UNIQUE ("tenant_id")
);
-- Create index "idx_tenants_deleted_at" to table: "tenants"
CREATE INDEX "idx_tenants_deleted_at" ON "tenants"."tenants" ("deleted_at");
-- Create index "idx_tenants_name" to table: "tenants"
CREATE INDEX "idx_tenants_name" ON "tenants"."tenants" ("name");
-- Create index "idx_tenants_primary_domain" to table: "tenants"
CREATE INDEX "idx_tenants_primary_domain" ON "tenants"."tenants" ("primary_domain") WHERE (primary_domain IS NOT NULL);
-- Create index "idx_tenants_status" to table: "tenants"
CREATE INDEX "idx_tenants_status" ON "tenants"."tenants" ("status");
-- Create index "idx_tenants_tenant_id" to table: "tenants"
CREATE INDEX "idx_tenants_tenant_id" ON "tenants"."tenants" ("tenant_id");
-- Create "domain_verifications" table
CREATE TABLE "tenants"."domain_verifications" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tenant_id" uuid NOT NULL,
  "domain" character varying(255) NOT NULL,
  "verification_method" "tenants"."verification_method" NOT NULL DEFAULT 'DNS',
  "verification_token" character varying(255) NULL,
  "status" "tenants"."verification_status" NOT NULL DEFAULT 'PENDING',
  "verified_at" timestamp NULL,
  "expires_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "domain_verifications_tenant_id_domain_key" UNIQUE ("tenant_id", "domain"),
  CONSTRAINT "domain_verifications_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_domain_verifications_domain" to table: "domain_verifications"
CREATE INDEX "idx_domain_verifications_domain" ON "tenants"."domain_verifications" ("domain");
-- Create index "idx_domain_verifications_status" to table: "domain_verifications"
CREATE INDEX "idx_domain_verifications_status" ON "tenants"."domain_verifications" ("status");
-- Create index "idx_domain_verifications_tenant_id" to table: "domain_verifications"
CREATE INDEX "idx_domain_verifications_tenant_id" ON "tenants"."domain_verifications" ("tenant_id");
-- Create "groups" table
CREATE TABLE "tenants"."groups" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "group_id" character varying(255) NOT NULL,
  "tenant_id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "type" "tenants"."group_type" NOT NULL DEFAULT 'group',
  "parent_group_id" uuid NULL,
  "description" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "deleted_at" timestamp NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "groups_group_id_key" UNIQUE ("group_id"),
  CONSTRAINT "groups_parent_group_id_fkey" FOREIGN KEY ("parent_group_id") REFERENCES "tenants"."groups" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "groups_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_groups_deleted_at" to table: "groups"
CREATE INDEX "idx_groups_deleted_at" ON "tenants"."groups" ("deleted_at");
-- Create index "idx_groups_group_id" to table: "groups"
CREATE INDEX "idx_groups_group_id" ON "tenants"."groups" ("group_id");
-- Create index "idx_groups_parent_group_id" to table: "groups"
CREATE INDEX "idx_groups_parent_group_id" ON "tenants"."groups" ("parent_group_id");
-- Create index "idx_groups_tenant_id" to table: "groups"
CREATE INDEX "idx_groups_tenant_id" ON "tenants"."groups" ("tenant_id");
-- Create index "idx_groups_type" to table: "groups"
CREATE INDEX "idx_groups_type" ON "tenants"."groups" ("type");
-- Create "image_types" table
CREATE TABLE "image"."image_types" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" text NOT NULL,
  "description" text NULL,
  "max_width" integer NULL,
  "max_height" integer NULL,
  "aspect_ratio" text NULL,
  "is_system" boolean NULL DEFAULT false,
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_types_key_key" UNIQUE ("key")
);
-- Create index "idx_image_types_is_system" to table: "image_types"
CREATE INDEX "idx_image_types_is_system" ON "image"."image_types" ("is_system");
-- Create index "idx_image_types_key" to table: "image_types"
CREATE INDEX "idx_image_types_key" ON "image"."image_types" ("key");
-- Set comment to table: "image_types"
COMMENT ON TABLE "image"."image_types" IS 'Image type definitions for different use cases and business domains';
-- Set comment to column: "key" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."key" IS 'Type key identifier: avatar, background, product_cover, product_gallery, post_image, banner, badge_icon, etc.';
-- Set comment to column: "description" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."description" IS 'Human-readable description of what this image type is used for';
-- Set comment to column: "max_width" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."max_width" IS 'Maximum allowed width in pixels (NULL means no limit)';
-- Set comment to column: "max_height" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."max_height" IS 'Maximum allowed height in pixels (NULL means no limit)';
-- Set comment to column: "aspect_ratio" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."aspect_ratio" IS 'Preferred aspect ratio hint: 1:1, 16:9, 4:3, etc. Used for frontend auto-cropping';
-- Set comment to column: "is_system" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."is_system" IS 'TRUE for system-built types (cannot be deleted), FALSE for custom types';
-- Create "images" table
CREATE TABLE "image"."images" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "type_id" uuid NULL,
  "user_id" uuid NULL,
  "tenant_id" uuid NULL,
  "app_id" uuid NULL,
  "source_domain" text NULL,
  "filename" text NOT NULL,
  "original_filename" text NOT NULL,
  "mime_type" text NOT NULL,
  "size" bigint NOT NULL,
  "width" integer NULL,
  "height" integer NULL,
  "storage_path" text NOT NULL,
  "url" text NULL,
  "is_public" boolean NOT NULL DEFAULT false,
  "metadata" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "images_type_id_fkey" FOREIGN KEY ("type_id") REFERENCES "image"."image_types" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "idx_images_app_id" to table: "images"
CREATE INDEX "idx_images_app_id" ON "image"."images" ("app_id");
-- Create index "idx_images_created_at" to table: "images"
CREATE INDEX "idx_images_created_at" ON "image"."images" ("created_at");
-- Create index "idx_images_deleted_at" to table: "images"
CREATE INDEX "idx_images_deleted_at" ON "image"."images" ("deleted_at");
-- Create index "idx_images_filename" to table: "images"
CREATE INDEX "idx_images_filename" ON "image"."images" ("filename");
-- Create index "idx_images_is_public" to table: "images"
CREATE INDEX "idx_images_is_public" ON "image"."images" ("is_public");
-- Create index "idx_images_mime_type" to table: "images"
CREATE INDEX "idx_images_mime_type" ON "image"."images" ("mime_type");
-- Create index "idx_images_source_domain" to table: "images"
CREATE INDEX "idx_images_source_domain" ON "image"."images" ("source_domain");
-- Create index "idx_images_tenant_app" to table: "images"
CREATE INDEX "idx_images_tenant_app" ON "image"."images" ("tenant_id", "app_id") WHERE (deleted_at IS NULL);
-- Create index "idx_images_tenant_id" to table: "images"
CREATE INDEX "idx_images_tenant_id" ON "image"."images" ("tenant_id");
-- Create index "idx_images_tenant_user" to table: "images"
CREATE INDEX "idx_images_tenant_user" ON "image"."images" ("tenant_id", "user_id") WHERE (deleted_at IS NULL);
-- Create index "idx_images_type_id" to table: "images"
CREATE INDEX "idx_images_type_id" ON "image"."images" ("type_id");
-- Create index "idx_images_type_public" to table: "images"
CREATE INDEX "idx_images_type_public" ON "image"."images" ("type_id", "is_public") WHERE (deleted_at IS NULL);
-- Create index "idx_images_user_id" to table: "images"
CREATE INDEX "idx_images_user_id" ON "image"."images" ("user_id");
-- Create index "idx_images_user_public" to table: "images"
CREATE INDEX "idx_images_user_public" ON "image"."images" ("user_id", "is_public") WHERE (deleted_at IS NULL);
-- Set comment to table: "images"
COMMENT ON TABLE "image"."images" IS 'Main table storing original image metadata for all uploaded images';
-- Set comment to column: "type_id" on table: "images"
COMMENT ON COLUMN "image"."images"."type_id" IS 'Reference to image_types table defining the image category/usage';
-- Set comment to column: "user_id" on table: "images"
COMMENT ON COLUMN "image"."images"."user_id" IS 'UUID of user who uploaded the image (from directory.users, application-level consistency)';
-- Set comment to column: "tenant_id" on table: "images"
COMMENT ON COLUMN "image"."images"."tenant_id" IS 'Tenant isolation: which tenant this image belongs to (from tenants.tenants.id, application-level consistency)';
-- Set comment to column: "app_id" on table: "images"
COMMENT ON COLUMN "image"."images"."app_id" IS 'App isolation: which app uploaded this image (from clients.apps.id, application-level consistency)';
-- Set comment to column: "source_domain" on table: "images"
COMMENT ON COLUMN "image"."images"."source_domain" IS 'Source service identifier: auth, product, post, cms, etc.';
-- Set comment to column: "filename" on table: "images"
COMMENT ON COLUMN "image"."images"."filename" IS 'Current filename after processing (may differ from original_filename)';
-- Set comment to column: "original_filename" on table: "images"
COMMENT ON COLUMN "image"."images"."original_filename" IS 'Original filename as uploaded by user';
-- Set comment to column: "mime_type" on table: "images"
COMMENT ON COLUMN "image"."images"."mime_type" IS 'MIME type: image/jpeg, image/png, image/webp, image/gif, etc.';
-- Set comment to column: "size" on table: "images"
COMMENT ON COLUMN "image"."images"."size" IS 'File size in bytes';
-- Set comment to column: "width" on table: "images"
COMMENT ON COLUMN "image"."images"."width" IS 'Image width in pixels';
-- Set comment to column: "height" on table: "images"
COMMENT ON COLUMN "image"."images"."height" IS 'Image height in pixels';
-- Set comment to column: "storage_path" on table: "images"
COMMENT ON COLUMN "image"."images"."storage_path" IS 'Storage backend path (filesystem path or S3 object key)';
-- Set comment to column: "url" on table: "images"
COMMENT ON COLUMN "image"."images"."url" IS 'Public accessible URL (CDN URL or direct storage URL)';
-- Set comment to column: "is_public" on table: "images"
COMMENT ON COLUMN "image"."images"."is_public" IS 'TRUE = publicly accessible, FALSE = requires authentication';
-- Set comment to column: "metadata" on table: "images"
COMMENT ON COLUMN "image"."images"."metadata" IS 'Extended JSON metadata: EXIF, color profiles, AI labels, etc.';
-- Set comment to column: "deleted_at" on table: "images"
COMMENT ON COLUMN "image"."images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
-- Create "image_tags" table
CREATE TABLE "image"."image_tags" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "image_id" uuid NOT NULL,
  "tag" text NOT NULL,
  "confidence" double precision NULL,
  "created_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_tags_image_id_tag_key" UNIQUE ("image_id", "tag"),
  CONSTRAINT "image_tags_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"."images" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_image_tags_confidence" to table: "image_tags"
CREATE INDEX "idx_image_tags_confidence" ON "image"."image_tags" ("confidence") WHERE (confidence IS NOT NULL);
-- Create index "idx_image_tags_image_id" to table: "image_tags"
CREATE INDEX "idx_image_tags_image_id" ON "image"."image_tags" ("image_id");
-- Create index "idx_image_tags_tag" to table: "image_tags"
CREATE INDEX "idx_image_tags_tag" ON "image"."image_tags" ("tag");
-- Create index "idx_image_tags_tag_confidence" to table: "image_tags"
CREATE INDEX "idx_image_tags_tag_confidence" ON "image"."image_tags" ("tag", "confidence" DESC);
-- Set comment to table: "image_tags"
COMMENT ON TABLE "image"."image_tags" IS 'Image tags for content-based search: AI labels and user-defined tags';
-- Set comment to column: "image_id" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."image_id" IS 'Reference to image in images table';
-- Set comment to column: "tag" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."tag" IS 'Tag text (normalized, e.g., lowercase). Examples: cat, shoes, selfie, nature';
-- Set comment to column: "confidence" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."confidence" IS 'Confidence score 0.0-1.0 for AI-generated tags, NULL for user tags';
-- Set comment to column: "created_at" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."created_at" IS 'When this tag was added to the image';
-- Create "image_variants" table
CREATE TABLE "image"."image_variants" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "image_id" uuid NOT NULL,
  "variant_key" text NOT NULL,
  "width" integer NULL,
  "height" integer NULL,
  "size" bigint NULL,
  "mime_type" text NULL,
  "storage_path" text NOT NULL,
  "url" text NULL,
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_variants_image_id_variant_key_key" UNIQUE ("image_id", "variant_key"),
  CONSTRAINT "image_variants_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"."images" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_image_variants_image_id" to table: "image_variants"
CREATE INDEX "idx_image_variants_image_id" ON "image"."image_variants" ("image_id");
-- Create index "idx_image_variants_url" to table: "image_variants"
CREATE INDEX "idx_image_variants_url" ON "image"."image_variants" ("url") WHERE (url IS NOT NULL);
-- Create index "idx_image_variants_variant_key" to table: "image_variants"
CREATE INDEX "idx_image_variants_variant_key" ON "image"."image_variants" ("variant_key");
-- Set comment to table: "image_variants"
COMMENT ON TABLE "image"."image_variants" IS 'Stores derived versions of images: thumbnails, resized versions, format conversions';
-- Set comment to column: "image_id" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."image_id" IS 'Reference to parent image in images table';
-- Set comment to column: "variant_key" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."variant_key" IS 'Variant type: thumbnail (150x150), small (300px), medium (720px), large (1080px), webp (converted), etc.';
-- Set comment to column: "width" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."width" IS 'Variant width in pixels';
-- Set comment to column: "height" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."height" IS 'Variant height in pixels';
-- Set comment to column: "size" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."size" IS 'Variant file size in bytes';
-- Set comment to column: "mime_type" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."mime_type" IS 'Variant MIME type (may differ from original, e.g., webp conversion)';
-- Set comment to column: "storage_path" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."storage_path" IS 'Storage backend path for this variant';
-- Set comment to column: "url" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."url" IS 'Public accessible URL for this variant';
-- Create "invitations" table
CREATE TABLE "tenants"."invitations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "invite_id" character varying(255) NOT NULL,
  "tenant_id" uuid NOT NULL,
  "email" character varying(255) NOT NULL,
  "token_hash" character varying(255) NOT NULL,
  "expires_at" timestamp NOT NULL,
  "status" "tenants"."invitation_status" NOT NULL DEFAULT 'PENDING',
  "invited_by" uuid NOT NULL,
  "invited_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "accepted_by_user_id" uuid NULL,
  "accepted_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoked_at" timestamp NULL,
  "revoke_reason" text NULL,
  "role_ids" uuid[] NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "invitations_invite_id_key" UNIQUE ("invite_id"),
  CONSTRAINT "invitations_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_invitations_email" to table: "invitations"
CREATE INDEX "idx_invitations_email" ON "tenants"."invitations" ("email");
-- Create index "idx_invitations_expires_at" to table: "invitations"
CREATE INDEX "idx_invitations_expires_at" ON "tenants"."invitations" ("expires_at");
-- Create index "idx_invitations_invite_id" to table: "invitations"
CREATE INDEX "idx_invitations_invite_id" ON "tenants"."invitations" ("invite_id");
-- Create index "idx_invitations_role_ids" to table: "invitations"
CREATE INDEX "idx_invitations_role_ids" ON "tenants"."invitations" USING gin ("role_ids");
-- Create index "idx_invitations_status" to table: "invitations"
CREATE INDEX "idx_invitations_status" ON "tenants"."invitations" ("status");
-- Create index "idx_invitations_tenant_id" to table: "invitations"
CREATE INDEX "idx_invitations_tenant_id" ON "tenants"."invitations" ("tenant_id");
-- Create index "idx_invitations_tenant_status" to table: "invitations"
CREATE INDEX "idx_invitations_tenant_status" ON "tenants"."invitations" ("tenant_id", "status");
-- Create "ip_allowlist" table
CREATE TABLE "clients"."ip_allowlist" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "rule_id" character varying(255) NOT NULL,
  "app_id" uuid NOT NULL,
  "cidr" cidr NOT NULL,
  "description" text NULL,
  "status" "clients"."allowlist_status" NOT NULL DEFAULT 'active',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" uuid NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "ip_allowlist_rule_id_key" UNIQUE ("rule_id"),
  CONSTRAINT "ip_allowlist_app_id_fkey" FOREIGN KEY ("app_id") REFERENCES "clients"."apps" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_ip_allowlist_app_id" to table: "ip_allowlist"
CREATE INDEX "idx_ip_allowlist_app_id" ON "clients"."ip_allowlist" ("app_id");
-- Create index "idx_ip_allowlist_app_status" to table: "ip_allowlist"
CREATE INDEX "idx_ip_allowlist_app_status" ON "clients"."ip_allowlist" ("app_id", "status") WHERE (status = 'active'::clients.allowlist_status);
-- Create index "idx_ip_allowlist_rule_id" to table: "ip_allowlist"
CREATE INDEX "idx_ip_allowlist_rule_id" ON "clients"."ip_allowlist" ("rule_id");
-- Create index "idx_ip_allowlist_status" to table: "ip_allowlist"
CREATE INDEX "idx_ip_allowlist_status" ON "clients"."ip_allowlist" ("status");
-- Create "members" table
CREATE TABLE "tenants"."members" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "member_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tenant_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "status" "tenants"."member_status" NOT NULL DEFAULT 'INVITED',
  "source" "tenants"."member_source" NOT NULL DEFAULT 'MANUAL',
  "joined_at" timestamp NULL,
  "left_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "external_ref" character varying(255) NULL,
  "metadata" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "members_member_id_key" UNIQUE ("member_id"),
  CONSTRAINT "members_tenant_id_user_id_key" UNIQUE ("tenant_id", "user_id"),
  CONSTRAINT "members_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_members_member_id" to table: "members"
CREATE INDEX "idx_members_member_id" ON "tenants"."members" ("member_id");
-- Create index "idx_members_source" to table: "members"
CREATE INDEX "idx_members_source" ON "tenants"."members" ("source");
-- Create index "idx_members_status" to table: "members"
CREATE INDEX "idx_members_status" ON "tenants"."members" ("status");
-- Create index "idx_members_tenant_id" to table: "members"
CREATE INDEX "idx_members_tenant_id" ON "tenants"."members" ("tenant_id");
-- Create index "idx_members_tenant_status" to table: "members"
CREATE INDEX "idx_members_tenant_status" ON "tenants"."members" ("tenant_id", "status");
-- Create index "idx_members_user_id" to table: "members"
CREATE INDEX "idx_members_user_id" ON "tenants"."members" ("user_id");
-- Create "member_app_roles" table
CREATE TABLE "tenants"."member_app_roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "member_id" uuid NOT NULL,
  "app_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  "assigned_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" uuid NULL,
  "expires_at" timestamp NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "member_app_roles_member_id_app_id_role_id_key" UNIQUE ("member_id", "app_id", "role_id"),
  CONSTRAINT "member_app_roles_member_id_fkey" FOREIGN KEY ("member_id") REFERENCES "tenants"."members" ("member_id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_member_app_roles_app_id" to table: "member_app_roles"
CREATE INDEX "idx_member_app_roles_app_id" ON "tenants"."member_app_roles" ("app_id");
-- Create index "idx_member_app_roles_expires_at" to table: "member_app_roles"
CREATE INDEX "idx_member_app_roles_expires_at" ON "tenants"."member_app_roles" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_member_app_roles_member_id" to table: "member_app_roles"
CREATE INDEX "idx_member_app_roles_member_id" ON "tenants"."member_app_roles" ("member_id");
-- Create index "idx_member_app_roles_role_id" to table: "member_app_roles"
CREATE INDEX "idx_member_app_roles_role_id" ON "tenants"."member_app_roles" ("role_id");
-- Create "member_groups" table
CREATE TABLE "tenants"."member_groups" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "member_id" uuid NOT NULL,
  "group_id" uuid NOT NULL,
  "assigned_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" uuid NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "member_groups_member_id_group_id_key" UNIQUE ("member_id", "group_id"),
  CONSTRAINT "member_groups_group_id_fkey" FOREIGN KEY ("group_id") REFERENCES "tenants"."groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "member_groups_member_id_fkey" FOREIGN KEY ("member_id") REFERENCES "tenants"."members" ("member_id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_member_groups_group_id" to table: "member_groups"
CREATE INDEX "idx_member_groups_group_id" ON "tenants"."member_groups" ("group_id");
-- Create index "idx_member_groups_member_id" to table: "member_groups"
CREATE INDEX "idx_member_groups_member_id" ON "tenants"."member_groups" ("member_id");
-- Create "member_roles" table
CREATE TABLE "tenants"."member_roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tenant_id" uuid NOT NULL,
  "member_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  "assigned_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" uuid NULL,
  "expires_at" timestamp NULL,
  "scope" character varying(100) NULL,
  "revoked_at" timestamp NULL,
  "revoked_by" uuid NULL,
  "revoke_reason" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "member_roles_tenant_id_member_id_role_id_key" UNIQUE ("tenant_id", "member_id", "role_id"),
  CONSTRAINT "member_roles_member_id_fkey" FOREIGN KEY ("member_id") REFERENCES "tenants"."members" ("member_id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "member_roles_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_member_roles_expires_at" to table: "member_roles"
CREATE INDEX "idx_member_roles_expires_at" ON "tenants"."member_roles" ("expires_at") WHERE (expires_at IS NOT NULL);
-- Create index "idx_member_roles_member_id" to table: "member_roles"
CREATE INDEX "idx_member_roles_member_id" ON "tenants"."member_roles" ("member_id");
-- Create index "idx_member_roles_revoked_at" to table: "member_roles"
CREATE INDEX "idx_member_roles_revoked_at" ON "tenants"."member_roles" ("revoked_at") WHERE (revoked_at IS NULL);
-- Create index "idx_member_roles_role_id" to table: "member_roles"
CREATE INDEX "idx_member_roles_role_id" ON "tenants"."member_roles" ("role_id");
-- Create index "idx_member_roles_tenant_id" to table: "member_roles"
CREATE INDEX "idx_member_roles_tenant_id" ON "tenants"."member_roles" ("tenant_id");
-- Create "rate_limits" table
CREATE TABLE "clients"."rate_limits" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "app_id" uuid NOT NULL,
  "limit_type" "clients"."rate_limit_type" NOT NULL DEFAULT 'requests_per_minute',
  "limit_value" integer NOT NULL,
  "window_seconds" integer NOT NULL,
  "description" text NULL,
  "status" character varying(50) NOT NULL DEFAULT 'active',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "rate_limits_app_id_limit_type_key" UNIQUE ("app_id", "limit_type"),
  CONSTRAINT "rate_limits_app_id_fkey" FOREIGN KEY ("app_id") REFERENCES "clients"."apps" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_rate_limits_app_id" to table: "rate_limits"
CREATE INDEX "idx_rate_limits_app_id" ON "clients"."rate_limits" ("app_id");
-- Create index "idx_rate_limits_status" to table: "rate_limits"
CREATE INDEX "idx_rate_limits_status" ON "clients"."rate_limits" ("status");
-- Create "roles" table
CREATE TABLE "access"."roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "scope_type" "access"."scope_type" NOT NULL DEFAULT 'TENANT',
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "roles_key_key" UNIQUE ("key")
);
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "access"."roles" ("deleted_at");
-- Create index "idx_roles_is_system" to table: "roles"
CREATE INDEX "idx_roles_is_system" ON "access"."roles" ("is_system");
-- Create index "idx_roles_key" to table: "roles"
CREATE INDEX "idx_roles_key" ON "access"."roles" ("key");
-- Create index "idx_roles_scope_type" to table: "roles"
CREATE INDEX "idx_roles_scope_type" ON "access"."roles" ("scope_type");
-- Create "role_permissions" table
CREATE TABLE "access"."role_permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "role_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "role_permissions_role_id_permission_id_key" UNIQUE ("role_id", "permission_id"),
  CONSTRAINT "role_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "access"."permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "role_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "access"."roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_role_permissions_permission_id" to table: "role_permissions"
CREATE INDEX "idx_role_permissions_permission_id" ON "access"."role_permissions" ("permission_id");
-- Create index "idx_role_permissions_role_id" to table: "role_permissions"
CREATE INDEX "idx_role_permissions_role_id" ON "access"."role_permissions" ("role_id");
-- Create index "idx_role_permissions_role_permission" to table: "role_permissions"
CREATE INDEX "idx_role_permissions_role_permission" ON "access"."role_permissions" ("role_id", "permission_id");
-- Create "scopes" table
CREATE TABLE "access"."scopes" (
  "scope" character varying(255) NOT NULL,
  "description" text NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("scope")
);
-- Create index "idx_scopes_deleted_at" to table: "scopes"
CREATE INDEX "idx_scopes_deleted_at" ON "access"."scopes" ("deleted_at");
-- Create index "idx_scopes_is_system" to table: "scopes"
CREATE INDEX "idx_scopes_is_system" ON "access"."scopes" ("is_system");
-- Create "scope_permissions" table
CREATE TABLE "access"."scope_permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "scope" character varying(255) NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "scope_permissions_scope_permission_id_key" UNIQUE ("scope", "permission_id"),
  CONSTRAINT "scope_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "access"."permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "scope_permissions_scope_fkey" FOREIGN KEY ("scope") REFERENCES "access"."scopes" ("scope") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_scope_permissions_permission_id" to table: "scope_permissions"
CREATE INDEX "idx_scope_permissions_permission_id" ON "access"."scope_permissions" ("permission_id");
-- Create index "idx_scope_permissions_scope" to table: "scope_permissions"
CREATE INDEX "idx_scope_permissions_scope" ON "access"."scope_permissions" ("scope");
-- Create index "idx_scope_permissions_scope_permission" to table: "scope_permissions"
CREATE INDEX "idx_scope_permissions_scope_permission" ON "access"."scope_permissions" ("scope", "permission_id");
-- Create "tenant_apps" table
CREATE TABLE "tenants"."tenant_apps" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tenant_id" uuid NOT NULL,
  "app_id" uuid NOT NULL,
  "status" "tenants"."tenant_app_status" NOT NULL DEFAULT 'ACTIVE',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "settings" jsonb NULL DEFAULT '{}',
  PRIMARY KEY ("id"),
  CONSTRAINT "tenant_apps_tenant_id_app_id_key" UNIQUE ("tenant_id", "app_id"),
  CONSTRAINT "tenant_apps_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_tenant_apps_app_id" to table: "tenant_apps"
CREATE INDEX "idx_tenant_apps_app_id" ON "tenants"."tenant_apps" ("app_id");
-- Create index "idx_tenant_apps_status" to table: "tenant_apps"
CREATE INDEX "idx_tenant_apps_status" ON "tenants"."tenant_apps" ("status");
-- Create index "idx_tenant_apps_tenant_id" to table: "tenant_apps"
CREATE INDEX "idx_tenant_apps_tenant_id" ON "tenants"."tenant_apps" ("tenant_id");
-- Create "tenant_settings" table
CREATE TABLE "tenants"."tenant_settings" (
  "id" uuid NOT NULL,
  "enforce_mfa" boolean NOT NULL DEFAULT false,
  "allowed_email_domains" text[] NULL,
  "session_ttl_minutes" integer NULL,
  "password_policy" jsonb NULL DEFAULT '{}',
  "login_policy" jsonb NULL DEFAULT '{}',
  "mfa_policy" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "tenant_settings_id_fkey" FOREIGN KEY ("id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "users" table
CREATE TABLE "directory"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" character varying(50) NOT NULL,
  "status" "directory"."user_status" NOT NULL DEFAULT 'pending',
  "is_verified" boolean NOT NULL DEFAULT false,
  "last_login_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_username_key" UNIQUE ("username")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "directory"."users" ("deleted_at");
-- Create index "idx_users_last_login_at" to table: "users"
CREATE INDEX "idx_users_last_login_at" ON "directory"."users" ("last_login_at");
-- Create index "idx_users_status" to table: "users"
CREATE INDEX "idx_users_status" ON "directory"."users" ("status");
-- Create index "idx_users_username" to table: "users"
CREATE INDEX "idx_users_username" ON "directory"."users" ("username");
-- Create "user_avatars" table
CREATE TABLE "directory"."user_avatars" (
  "user_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("user_id"),
  CONSTRAINT "user_avatars_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_avatars_image_id" to table: "user_avatars"
CREATE INDEX "idx_user_avatars_image_id" ON "directory"."user_avatars" ("image_id");
-- Set comment to table: "user_avatars"
COMMENT ON TABLE "directory"."user_avatars" IS 'User avatar table - stores current avatar for each user (one-to-one relationship)';
-- Set comment to column: "user_id" on table: "user_avatars"
COMMENT ON COLUMN "directory"."user_avatars"."user_id" IS 'Reference to directory.users.id (primary key, one-to-one)';
-- Set comment to column: "image_id" on table: "user_avatars"
COMMENT ON COLUMN "directory"."user_avatars"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key)';
-- Create "badges" table
CREATE TABLE "directory"."badges" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "icon_url" character varying(255) NULL,
  "color" character varying(20) NULL,
  "category" character varying(50) NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "badges_name_key" UNIQUE ("name")
);
-- Create index "idx_badges_category" to table: "badges"
CREATE INDEX "idx_badges_category" ON "directory"."badges" ("category");
-- Create index "idx_badges_deleted_at" to table: "badges"
CREATE INDEX "idx_badges_deleted_at" ON "directory"."badges" ("deleted_at");
-- Create index "idx_badges_name" to table: "badges"
CREATE INDEX "idx_badges_name" ON "directory"."badges" ("name");
-- Create "user_badges" table
CREATE TABLE "directory"."user_badges" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "badge_id" uuid NOT NULL,
  "description" text NULL DEFAULT '',
  "level" integer NULL DEFAULT 1,
  "earned_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_badges_user_id_badge_id_key" UNIQUE ("user_id", "badge_id"),
  CONSTRAINT "user_badges_badge_id_fkey" FOREIGN KEY ("badge_id") REFERENCES "directory"."badges" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "user_badges_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_badges_badge_id" to table: "user_badges"
CREATE INDEX "idx_user_badges_badge_id" ON "directory"."user_badges" ("badge_id");
-- Create index "idx_user_badges_earned_at" to table: "user_badges"
CREATE INDEX "idx_user_badges_earned_at" ON "directory"."user_badges" ("earned_at");
-- Create index "idx_user_badges_user_id" to table: "user_badges"
CREATE INDEX "idx_user_badges_user_id" ON "directory"."user_badges" ("user_id");
-- Create "user_educations" table
CREATE TABLE "directory"."user_educations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "school" character varying(255) NOT NULL,
  "degree" character varying(100) NULL,
  "major" character varying(255) NULL,
  "field_of_study" character varying(255) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "grade" character varying(50) NULL,
  "activities" text NULL,
  "achievements" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_educations_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_educations_degree" to table: "user_educations"
CREATE INDEX "idx_user_educations_degree" ON "directory"."user_educations" ("degree");
-- Create index "idx_user_educations_deleted_at" to table: "user_educations"
CREATE INDEX "idx_user_educations_deleted_at" ON "directory"."user_educations" ("deleted_at");
-- Create index "idx_user_educations_school" to table: "user_educations"
CREATE INDEX "idx_user_educations_school" ON "directory"."user_educations" ("school");
-- Create index "idx_user_educations_user_id" to table: "user_educations"
CREATE INDEX "idx_user_educations_user_id" ON "directory"."user_educations" ("user_id");
-- Create "user_emails" table
CREATE TABLE "directory"."user_emails" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "email" character varying(255) NOT NULL,
  "is_primary" boolean NOT NULL DEFAULT false,
  "is_verified" boolean NOT NULL DEFAULT false,
  "verified_at" timestamp NULL,
  "verification_token" character varying(255) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_emails_email_key" UNIQUE ("email"),
  CONSTRAINT "user_emails_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_emails_deleted_at" to table: "user_emails"
CREATE INDEX "idx_user_emails_deleted_at" ON "directory"."user_emails" ("deleted_at");
-- Create index "idx_user_emails_email" to table: "user_emails"
CREATE INDEX "idx_user_emails_email" ON "directory"."user_emails" ("email");
-- Create index "idx_user_emails_is_primary" to table: "user_emails"
CREATE INDEX "idx_user_emails_is_primary" ON "directory"."user_emails" ("user_id", "is_primary") WHERE (is_primary = true);
-- Create index "idx_user_emails_is_verified" to table: "user_emails"
CREATE INDEX "idx_user_emails_is_verified" ON "directory"."user_emails" ("is_verified");
-- Create index "idx_user_emails_user_id" to table: "user_emails"
CREATE INDEX "idx_user_emails_user_id" ON "directory"."user_emails" ("user_id");
-- Create "user_images" table
CREATE TABLE "directory"."user_images" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "display_order" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_images_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_images_created_at" to table: "user_images"
CREATE INDEX "idx_user_images_created_at" ON "directory"."user_images" ("created_at");
-- Create index "idx_user_images_deleted_at" to table: "user_images"
CREATE INDEX "idx_user_images_deleted_at" ON "directory"."user_images" ("deleted_at");
-- Create index "idx_user_images_display_order" to table: "user_images"
CREATE INDEX "idx_user_images_display_order" ON "directory"."user_images" ("user_id", "display_order") WHERE (deleted_at IS NULL);
-- Create index "idx_user_images_image_id" to table: "user_images"
CREATE INDEX "idx_user_images_image_id" ON "directory"."user_images" ("image_id");
-- Create index "idx_user_images_user_id" to table: "user_images"
CREATE INDEX "idx_user_images_user_id" ON "directory"."user_images" ("user_id");
-- Create index "idx_user_images_user_image_unique" to table: "user_images"
CREATE UNIQUE INDEX "idx_user_images_user_image_unique" ON "directory"."user_images" ("user_id", "image_id") WHERE (deleted_at IS NULL);
-- Create index "idx_user_images_user_order" to table: "user_images"
CREATE INDEX "idx_user_images_user_order" ON "directory"."user_images" ("user_id", "display_order") WHERE (deleted_at IS NULL);
-- Set comment to table: "user_images"
COMMENT ON TABLE "directory"."user_images" IS 'User images table - stores all user images with display order. Image type is stored in image.images.type_id';
-- Set comment to column: "user_id" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."user_id" IS 'Reference to directory.users.id';
-- Set comment to column: "image_id" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key). Image type can be queried via image.images.type_id';
-- Set comment to column: "display_order" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."display_order" IS 'Display order: 0 = current/active image, higher numbers = later in sequence';
-- Set comment to column: "deleted_at" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
-- Create "user_occupations" table
CREATE TABLE "directory"."user_occupations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "company" character varying(255) NOT NULL,
  "position" character varying(255) NOT NULL,
  "department" character varying(255) NULL,
  "industry" character varying(100) NULL,
  "location" character varying(255) NULL,
  "employment_type" character varying(50) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "responsibilities" text NULL,
  "achievements" text NULL,
  "skills_used" text[] NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_occupations_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_occupations_company" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_company" ON "directory"."user_occupations" ("company");
-- Create index "idx_user_occupations_deleted_at" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_deleted_at" ON "directory"."user_occupations" ("deleted_at");
-- Create index "idx_user_occupations_industry" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_industry" ON "directory"."user_occupations" ("industry");
-- Create index "idx_user_occupations_is_current" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_is_current" ON "directory"."user_occupations" ("is_current");
-- Create index "idx_user_occupations_position" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_position" ON "directory"."user_occupations" ("position");
-- Create index "idx_user_occupations_user_id" to table: "user_occupations"
CREATE INDEX "idx_user_occupations_user_id" ON "directory"."user_occupations" ("user_id");
-- Create "user_phones" table
CREATE TABLE "directory"."user_phones" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "phone" character varying(20) NOT NULL,
  "country_code" character varying(10) NULL,
  "is_primary" boolean NOT NULL DEFAULT false,
  "is_verified" boolean NOT NULL DEFAULT false,
  "verified_at" timestamp NULL,
  "verification_code" character varying(10) NULL,
  "verification_expires_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_phones_phone_key" UNIQUE ("phone"),
  CONSTRAINT "user_phones_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_phones_deleted_at" to table: "user_phones"
CREATE INDEX "idx_user_phones_deleted_at" ON "directory"."user_phones" ("deleted_at");
-- Create index "idx_user_phones_is_primary" to table: "user_phones"
CREATE INDEX "idx_user_phones_is_primary" ON "directory"."user_phones" ("user_id", "is_primary") WHERE (is_primary = true);
-- Create index "idx_user_phones_is_verified" to table: "user_phones"
CREATE INDEX "idx_user_phones_is_verified" ON "directory"."user_phones" ("is_verified");
-- Create index "idx_user_phones_phone" to table: "user_phones"
CREATE INDEX "idx_user_phones_phone" ON "directory"."user_phones" ("phone");
-- Create index "idx_user_phones_user_id" to table: "user_phones"
CREATE INDEX "idx_user_phones_user_id" ON "directory"."user_phones" ("user_id");
-- Create "user_preferences" table
CREATE TABLE "directory"."user_preferences" (
  "id" uuid NOT NULL,
  "theme" character varying(50) NULL DEFAULT 'light',
  "language" character varying(10) NULL DEFAULT 'en',
  "timezone" character varying(50) NULL DEFAULT 'Etc/UTC',
  "notifications" jsonb NULL DEFAULT '{}',
  "privacy" jsonb NULL DEFAULT '{}',
  "display" jsonb NULL DEFAULT '{}',
  "other" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_preferences_id_fkey" FOREIGN KEY ("id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_preferences_deleted_at" to table: "user_preferences"
CREATE INDEX "idx_user_preferences_deleted_at" ON "directory"."user_preferences" ("deleted_at");
-- Create "user_profiles" table
CREATE TABLE "directory"."user_profiles" (
  "id" uuid NOT NULL,
  "role" character varying(100) NULL,
  "first_name" character varying(100) NULL,
  "last_name" character varying(100) NULL,
  "nickname" character varying(50) NULL,
  "display_name" character varying(100) NULL,
  "bio" text NULL,
  "birthday" date NULL,
  "age" integer NULL,
  "gender" character varying(20) NULL,
  "location" character varying(255) NULL,
  "website" character varying(255) NULL,
  "github" character varying(255) NULL,
  "social_links" jsonb NULL DEFAULT '{}',
  "skills" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_profiles_nickname_key" UNIQUE ("nickname"),
  CONSTRAINT "user_profiles_id_fkey" FOREIGN KEY ("id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_profiles_birthday" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_birthday" ON "directory"."user_profiles" ("birthday");
-- Create index "idx_user_profiles_deleted_at" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_deleted_at" ON "directory"."user_profiles" ("deleted_at");
-- Create index "idx_user_profiles_display_name" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_display_name" ON "directory"."user_profiles" ("display_name");
-- Create index "idx_user_profiles_first_and_last_name" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_first_and_last_name" ON "directory"."user_profiles" ("first_name", "last_name");
-- Create index "idx_user_profiles_gender" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_gender" ON "directory"."user_profiles" ("gender");
-- Create index "idx_user_profiles_github" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_github" ON "directory"."user_profiles" ("github");
-- Create index "idx_user_profiles_location" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_location" ON "directory"."user_profiles" ("location");
-- Create index "idx_user_profiles_nickname" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_nickname" ON "directory"."user_profiles" ("nickname");
-- Create index "idx_user_profiles_role" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_role" ON "directory"."user_profiles" ("role");
-- Create index "idx_user_profiles_website" to table: "user_profiles"
CREATE INDEX "idx_user_profiles_website" ON "directory"."user_profiles" ("website");
