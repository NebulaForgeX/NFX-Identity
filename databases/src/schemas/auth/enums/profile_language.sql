CREATE TYPE "auth".profile_language AS ENUM (
  'en',
  'zh',
  'fr'
);
COMMENT ON TYPE "auth".profile_language IS 'Preferred UI locale.';
