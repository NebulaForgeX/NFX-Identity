-- User-Badges many-to-many relationship table
-- Links users to badges they have earned
CREATE TABLE IF NOT EXISTS "directory"."user_badges" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "badge_id" UUID NOT NULL REFERENCES "directory"."badges"("id") ON DELETE CASCADE,
  "description" TEXT DEFAULT '',
  "level" INTEGER DEFAULT 1,
  "earned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("user_id", "badge_id")
);

CREATE INDEX IF NOT EXISTS "idx_user_badges_user_id" ON "directory"."user_badges"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_badges_badge_id" ON "directory"."user_badges"("badge_id");
CREATE INDEX IF NOT EXISTS "idx_user_badges_earned_at" ON "directory"."user_badges"("earned_at");

