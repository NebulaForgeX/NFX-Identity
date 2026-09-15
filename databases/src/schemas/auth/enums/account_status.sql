CREATE TYPE "auth".account_status AS ENUM (
  'active',
  'suspended',
  'deleted'
);
COMMENT ON TYPE "auth".account_status IS 'Account lifecycle.';
