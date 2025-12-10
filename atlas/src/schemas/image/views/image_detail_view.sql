-- Image Detail View
-- Purpose: Provides detailed view of an image with all its variants in a single query
-- Use Case: Image detail pages, admin views, API endpoints showing full image information
-- Returns: Multiple rows per image (one per variant) with complete image and variant metadata
--
-- Example query:
--   SELECT * FROM image.image_detail_view 
--   WHERE image_id = 'some-uuid'
--   ORDER BY variant_key;
--
-- Note: This view uses LEFT JOIN on variants, so images without variants will have one row
--       with NULL variant fields. Images with multiple variants will have multiple rows.
--       Use this view when you need to display all available variants of an image.

CREATE VIEW "image"."image_detail_view" AS
SELECT
  i."id" AS "image_id",
  i."type_id",
  t."key" AS "type_key",
  t."description" AS "type_description",
  i."user_id",
  i."source_domain",
  i."filename",
  i."original_filename",
  i."mime_type" AS "original_mime_type",
  i."size" AS "original_size",
  i."width" AS "original_width",
  i."height" AS "original_height",
  i."storage_path" AS "original_storage_path",
  i."url" AS "original_url",
  i."is_public",
  i."metadata" AS "image_metadata",
  v."id" AS "variant_id",
  v."variant_key",
  v."width" AS "variant_width",
  v."height" AS "variant_height",
  v."size" AS "variant_size",
  v."mime_type" AS "variant_mime_type",
  v."storage_path" AS "variant_storage_path",
  v."url" AS "variant_url",
  i."created_at" AS "image_created_at",
  i."updated_at" AS "image_updated_at"
FROM
  "image"."images" i
  LEFT JOIN "image"."image_types" t ON i."type_id" = t."id"
  LEFT JOIN "image"."image_variants" v ON v."image_id" = i."id"
WHERE
  i."deleted_at" IS NULL;

