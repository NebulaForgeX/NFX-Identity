-- Image List Item View
-- Purpose: Provides a flattened view of images with type information and primary thumbnail URL
-- Use Case: API endpoints for listing images, admin dashboards, gallery views
-- Returns: One row per image with type details and a single thumbnail variant URL
-- 
-- Example query:
--   SELECT * FROM image.image_list_item 
--   WHERE type_key = 'avatar' AND is_public = TRUE
--   ORDER BY created_at DESC 
--   LIMIT 20 OFFSET 0;
--
-- Performance: Optimized for list queries with thumbnail URL included in single query

CREATE VIEW "image"."image_list_item" AS
SELECT
  i."id",
  i."type_id",
  t."key" AS "type_key",
  t."description" AS "type_description",
  i."user_id",
  i."source_domain",
  i."filename",
  i."original_filename",
  i."mime_type",
  i."size" AS "original_size",
  i."width" AS "original_width",
  i."height" AS "original_height",
  i."storage_path",
  i."url" AS "original_url",
  i."is_public",
  v."url" AS "thumbnail_url",
  v."width" AS "thumbnail_width",
  v."height" AS "thumbnail_height",
  i."metadata",
  i."created_at",
  i."updated_at"
FROM
  "image"."images" i
  LEFT JOIN "image"."image_types" t ON i."type_id" = t."id"
  LEFT JOIN LATERAL (
    SELECT
      iv."url",
      iv."width",
      iv."height"
    FROM
      "image"."image_variants" iv
    WHERE
      iv."image_id" = i."id"
      AND iv."variant_key" = 'thumbnail'
    ORDER BY
      iv."created_at" DESC
    LIMIT 1
  ) v ON TRUE
WHERE
  i."deleted_at" IS NULL;
