-- Badges table for user achievements and recognition
-- Similar structure to roles table
CREATE TABLE IF NOT EXISTS "directory"."badges" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" VARCHAR(50) NOT NULL UNIQUE,
  "description" TEXT,
  "icon_url" VARCHAR(255),
  "color" VARCHAR(20),
  "category" VARCHAR(50), -- e.g., "achievement", "skill", "community", "special"
  "is_system" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_badges_name" ON "directory"."badges"("name");
CREATE INDEX IF NOT EXISTS "idx_badges_category" ON "directory"."badges"("category");
CREATE INDEX IF NOT EXISTS "idx_badges_deleted_at" ON "directory"."badges"("deleted_at");

