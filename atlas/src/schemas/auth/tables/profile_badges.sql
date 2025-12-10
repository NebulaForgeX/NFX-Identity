-- Profile-Badges many-to-many relationship table
-- Links profiles to badges they have earned
CREATE TABLE IF NOT EXISTS "auth"."profile_badges" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "profile_id" UUID NOT NULL REFERENCES "auth"."profiles"("id") ON DELETE CASCADE,
  "badge_id" UUID NOT NULL REFERENCES "auth"."badges"("id") ON DELETE CASCADE,
  "description" TEXT DEFAULT '',
  "level" INTEGER DEFAULT 1,
  "earned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("profile_id", "badge_id")
);

CREATE INDEX IF NOT EXISTS "idx_profile_badges_profile_id" ON "auth"."profile_badges"("profile_id");
CREATE INDEX IF NOT EXISTS "idx_profile_badges_badge_id" ON "auth"."profile_badges"("badge_id");
CREATE INDEX IF NOT EXISTS "idx_profile_badges_earned_at" ON "auth"."profile_badges"("earned_at");

