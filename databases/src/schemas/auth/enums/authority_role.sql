CREATE TYPE "auth".authority_role AS ENUM (
  'auditor',
  'administrator',
  'owner'
);
COMMENT ON TYPE "auth".authority_role IS 'Authority profile roles; membership only, not hierarchical.';
