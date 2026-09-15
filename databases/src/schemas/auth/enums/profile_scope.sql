CREATE TYPE "auth".profile_scope AS ENUM (
  'forger',
  'authority'
);
COMMENT ON TYPE "auth".profile_scope IS 'JWT profile_scope / login kind.';
