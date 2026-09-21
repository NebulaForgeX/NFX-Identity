CREATE TYPE "auth".profile_scope AS ENUM (
  'community',
  'authority'
);
COMMENT ON TYPE "auth".profile_scope IS 'JWT profile_scope / login kind: community | authority. forger is a community role, not a kind.';
