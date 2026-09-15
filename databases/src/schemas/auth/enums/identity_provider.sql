CREATE TYPE "auth".identity_provider AS ENUM (
  'password',
  'github'
);
COMMENT ON TYPE "auth".identity_provider IS 'password=email/phone credential; github=OAuth.';
