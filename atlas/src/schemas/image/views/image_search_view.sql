-- Image Search View
-- Purpose: Provides images with aggregated tags array for tag-based search functionality
-- Use Case: Content discovery, AI-powered search, tag filtering, "find similar images"
-- Returns: One row per image with all tags aggregated into a single array column
--
-- Example queries:
--   -- Find all images tagged with 'cat'
--   SELECT * FROM image.image_search_view 
--   WHERE 'cat' = ANY(tags) AND is_public = TRUE
--   ORDER BY created_at DESC;
--
--   -- Find images with multiple tags
--   SELECT * FROM image.image_search_view 
--   WHERE tags && ARRAY['cat', 'nature']::TEXT[]
--   ORDER BY created_at DESC;
--
--   -- Search using array containment
--   SELECT * FROM image.image_search_view 
--   WHERE 'cat' = ANY(tags) OR 'dog' = ANY(tags);
--
-- Note: Tags are aggregated using ARRAY_AGG, so images without tags will have NULL or empty array
--       Use PostgreSQL array operators (ANY, &&) for efficient tag-based queries

CREATE VIEW "image"."image_search_view" AS
SELECT
  i."id",
  i."type_id",
  t."key" AS "type_key",
  i."user_id",
  i."source_domain",
  i."filename",
  i."url" AS "original_url",
  i."is_public",
  i."metadata",
  ARRAY_REMOVE(ARRAY_AGG(DISTINCT it."tag"), NULL) AS "tags",
  i."created_at",
  i."updated_at"
FROM
  "image"."images" i
  LEFT JOIN "image"."image_types" t ON i."type_id" = t."id"
  LEFT JOIN "image"."image_tags" it ON it."image_id" = i."id"
WHERE
  i."deleted_at" IS NULL
GROUP BY
  i."id", i."type_id", t."key", i."user_id", i."source_domain",
  i."filename", i."url", i."is_public", i."metadata", i."created_at", i."updated_at";

