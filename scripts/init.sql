-- Active: 1781056435421@@192.168.1.64@10104@nfxidentity_dev
-- scripts/init.sql
-- Seed one login-capable account with BOTH profile kinds.
--   email:    lyuchongkailyu@gmail.com (primary, verified)
--   password: Lucas127 (bcrypt via pgcrypto public.crypt/public.gen_salt)
--   forger profile:    display_name "Lucas Lyu", forger_roles = ALL forger roles
--   authority profile: display_name "Lucas Lyu", authority_roles = ALL authority roles
-- IDs use uuidv7() (PostgreSQL 18+, matches the app's uuid.NewV7()).
-- pgcrypto lives in the public schema; functions are schema-qualified so the
-- script works even when the session search_path excludes public.
-- Idempotent: skips if the email already exists (non-deleted).
-- Run manually against the target DB, e.g. psql "$DATABASE_URL" -f scripts/init.sql

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

DO $$
DECLARE
  v_email         text := 'lyuchongkailyu@gmail.com';
  v_password      text := 'Lucas127';
  v_display_name  text := 'Lucas Lyu';
  v_account_id    uuid := uuidv7();
  v_forger_id     uuid := uuidv7();
  v_authority_id  uuid := uuidv7();
BEGIN
  IF EXISTS (
    SELECT 1 FROM auth."Emails"
    WHERE LOWER("email") = LOWER(v_email) AND "deleted_at" IS NULL
  ) THEN
    RAISE NOTICE 'init.sql: account for % already exists, skipping', v_email;
    RETURN;
  END IF;

  INSERT INTO auth."Accounts" ("id", "account_status", "signup_platform")
  VALUES (v_account_id, 'active', 'nfxidentity');

  INSERT INTO auth."Identities" ("id", "account_id", "identity_provider", "provider_subject", "password_hash")
  VALUES (uuidv7(), v_account_id, 'password', v_email, public.crypt(v_password, public.gen_salt('bf')));

  INSERT INTO auth."Emails" ("id", "account_id", "email", "is_primary", "verified_at")
  VALUES (uuidv7(), v_account_id, v_email, TRUE, CURRENT_TIMESTAMP);

  INSERT INTO auth."ForgerProfiles" ("id", "account_id", "forger_roles", "profile_language", "display_name", "timezone")
  VALUES (
    v_forger_id,
    v_account_id,
    ARRAY['forger']::auth.forger_role[],
    'en',
    v_display_name,
    'UTC'
  );

  INSERT INTO auth."ForgerProfileSettings" ("id", "login_notification")
  VALUES (v_forger_id, TRUE);

  INSERT INTO auth."AuthorityProfiles" ("id", "account_id", "authority_roles", "profile_language", "display_name", "timezone")
  VALUES (
    v_authority_id,
    v_account_id,
    ARRAY['auditor', 'administrator', 'owner']::auth.authority_role[],
    'en',
    v_display_name,
    'UTC'
  );

  INSERT INTO auth."AuthorityProfileSettings" ("id", "login_notification")
  VALUES (v_authority_id, TRUE);

  RAISE NOTICE 'init.sql: seeded account % (forger %, authority %)', v_account_id, v_forger_id, v_authority_id;
END $$;
