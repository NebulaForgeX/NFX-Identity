-- Create enum type "permission_category"
CREATE TYPE "permission"."permission_category" AS ENUM ('system', 'user');
-- Modify "permissions" table
ALTER TABLE "permission"."permissions" ALTER COLUMN "category" TYPE "permission"."permission_category" USING "category"::"permission"."permission_category", ALTER COLUMN "category" SET NOT NULL;
