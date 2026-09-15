CREATE SCHEMA IF NOT EXISTS "asset";
COMMENT ON SCHEMA "asset" IS 'Independent Asset service schema: Images, Files, Videos, Audios metadata. uploader_id is logical ref to Auth (no cross-schema FK). Stored in Stack MinIO.';
