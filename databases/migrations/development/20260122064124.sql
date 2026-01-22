-- Modify "user_credentials" table
ALTER TABLE "auth"."user_credentials" ALTER COLUMN "id" DROP DEFAULT, DROP COLUMN "user_id";
-- Modify "tenant_settings" table
ALTER TABLE "tenants"."tenant_settings" ALTER COLUMN "id" DROP DEFAULT, DROP COLUMN "tenant_id", ADD CONSTRAINT "tenant_settings_id_fkey" FOREIGN KEY ("id") REFERENCES "tenants"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "user_preferences" table
ALTER TABLE "directory"."user_preferences" ALTER COLUMN "id" DROP DEFAULT, DROP COLUMN "user_id", ADD CONSTRAINT "user_preferences_id_fkey" FOREIGN KEY ("id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "user_profiles" table
ALTER TABLE "directory"."user_profiles" ALTER COLUMN "id" DROP DEFAULT, DROP COLUMN "user_id", ADD CONSTRAINT "user_profiles_id_fkey" FOREIGN KEY ("id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
