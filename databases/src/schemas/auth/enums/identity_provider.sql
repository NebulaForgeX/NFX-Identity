CREATE TYPE "auth".identity_provider AS ENUM (
  'password'
);
COMMENT ON TYPE "auth".identity_provider IS 'password=PASSWORD; OAuth values added later.';

-- Maintainer notes (below DDL):
-- 与域模型 IdentityProvider 对应；未来可 ALTER TYPE 增加 OAuth。
-- Maps to IdentityProvider; extend when adding OAuth.
