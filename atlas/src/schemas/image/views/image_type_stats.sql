-- Image Type Statistics View
-- Purpose: Provides aggregated statistics for each image type (counts, public/private breakdown)
-- Use Case: Admin dashboards, analytics, reports, monitoring image usage by type
-- Returns: One row per image type with aggregated statistics
--
-- Example queries:
--   -- Get all type statistics
--   SELECT * FROM image.image_type_stats ORDER BY total_images DESC;
--
--   -- Find types with most private images
--   SELECT type_key, private_images FROM image.image_type_stats 
--   ORDER BY private_images DESC;
--
-- Statistics included:
--   - total_images: Total number of images of this type (excluding deleted)
--   - public_images: Count of public images
--   - private_images: Count of private images
--   - first_created_at: When the first image of this type was uploaded
--   - last_created_at: When the most recent image of this type was uploaded

CREATE VIEW "image"."image_type_stats" AS
SELECT
  t."id",
  t."key" AS "type_key",
  t."description",
  t."max_width",
  t."max_height",
  t."aspect_ratio",
  t."is_system",
  COUNT(i."id") AS "total_images",
  COUNT(*) FILTER (WHERE i."is_public") AS "public_images",
  COUNT(*) FILTER (WHERE NOT i."is_public") AS "private_images",
  MIN(i."created_at") AS "first_created_at",
  MAX(i."created_at") AS "last_created_at"
FROM
  "image"."image_types" t
  LEFT JOIN "image"."images" i ON i."type_id" = t."id" AND i."deleted_at" IS NULL
GROUP BY
  t."id", t."key", t."description", t."max_width", t."max_height",
  t."aspect_ratio", t."is_system";

