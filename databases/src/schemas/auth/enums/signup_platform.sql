CREATE TYPE "auth".signup_platform AS ENUM (
  'nfxidentity',
  'nfxnews',
  'nfxstorages',
  'nfxedge'
);
COMMENT ON TYPE "auth".signup_platform IS 'Product that first created the account; written only at signup.';
