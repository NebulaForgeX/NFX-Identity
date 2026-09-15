CREATE TYPE "auth".forger_role AS ENUM (
  'forger'
);
COMMENT ON TYPE "auth".forger_role IS 'Forger profile roles; membership only, not hierarchical.';
