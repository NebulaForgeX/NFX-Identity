-- Images Table
-- Purpose: Stores original image metadata and file information
-- This is the main table for all uploaded images across all services
-- Note: user_id has no foreign key constraint to maintain service independence (microservices architecture)

CREATE TABLE "image"."images" (
  -- Primary key
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  
  -- Image type reference (e.g., avatar, product_cover)
  -- NULL allowed for legacy images or untyped uploads
  "type_id" UUID REFERENCES "image"."image_types"("id") ON DELETE SET NULL,
  
  -- User who uploaded the image (UUID from directory.users)
  -- No foreign key to avoid cross-service dependencies (application-level consistency)
  "user_id" UUID,
  
  -- Tenant isolation: which tenant this image belongs to
  -- References tenants.tenants.id (application-level consistency)
  "tenant_id" UUID,
  
  "application_id" UUID,
  
  -- Source domain/service that uploaded this image (e.g., 'auth', 'product', 'post', 'cms')
  -- Helps with analytics and service-specific queries
  "source_domain" TEXT,
  
  -- Current filename (may differ from original after processing)
  "filename" TEXT NOT NULL,
  
  -- Original filename as uploaded by user
  "original_filename" TEXT NOT NULL,
  
  -- MIME type (e.g., 'image/jpeg', 'image/png', 'image/webp')
  "mime_type" TEXT NOT NULL,
  
  -- File size in bytes
  "size" BIGINT NOT NULL,
  
  -- Image dimensions in pixels
  "width" INT,
  "height" INT,
  
  -- Storage path (local filesystem or S3 object key)
  -- This is the canonical path used by storage backend
  "storage_path" TEXT NOT NULL,
  
  -- Public URL for accessing the image (if applicable)
  -- May be CDN URL or direct storage URL
  "url" TEXT,
  
  -- Privacy flag: TRUE = publicly accessible, FALSE = private/requires auth
  "is_public" BOOLEAN NOT NULL DEFAULT FALSE,
  
  -- Extended metadata as JSON (EXIF data, AI labels, color profiles, etc.)
  -- Structure: {"exif": {...}, "colors": [...], "ai_labels": [...]}
  "metadata" JSONB DEFAULT '{}'::jsonb,
  
  -- Timestamps
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW(),
  
  -- Soft delete timestamp (NULL = active, NOT NULL = deleted)
  "deleted_at" TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX "idx_images_type_id" ON "image"."images"("type_id");
CREATE INDEX "idx_images_user_id" ON "image"."images"("user_id");
CREATE INDEX "idx_images_tenant_id" ON "image"."images"("tenant_id");
CREATE INDEX "idx_images_application_id" ON "image"."images"("application_id");
CREATE INDEX "idx_images_is_public" ON "image"."images"("is_public");
CREATE INDEX "idx_images_deleted_at" ON "image"."images"("deleted_at");
CREATE INDEX "idx_images_source_domain" ON "image"."images"("source_domain");
CREATE INDEX "idx_images_created_at" ON "image"."images"("created_at");
CREATE INDEX "idx_images_filename" ON "image"."images"("filename");
CREATE INDEX "idx_images_mime_type" ON "image"."images"("mime_type");

-- Composite indexes for common query patterns
CREATE INDEX "idx_images_user_public" ON "image"."images"("user_id", "is_public") WHERE "deleted_at" IS NULL;
CREATE INDEX "idx_images_type_public" ON "image"."images"("type_id", "is_public") WHERE "deleted_at" IS NULL;
CREATE INDEX "idx_images_tenant_user" ON "image"."images"("tenant_id", "user_id") WHERE "deleted_at" IS NULL;
CREATE INDEX "idx_images_tenant_application" ON "image"."images"("tenant_id", "application_id") WHERE "deleted_at" IS NULL;

-- Table and column comments
COMMENT ON TABLE "image"."images" IS 'Main table storing original image metadata for all uploaded images';
COMMENT ON COLUMN "image"."images"."type_id" IS 'Reference to image_types table defining the image category/usage';
COMMENT ON COLUMN "image"."images"."user_id" IS 'UUID of user who uploaded the image (from directory.users, application-level consistency)';
COMMENT ON COLUMN "image"."images"."tenant_id" IS 'Tenant isolation: which tenant this image belongs to (from tenants.tenants.id, application-level consistency)';
COMMENT ON COLUMN "image"."images"."application_id" IS 'Which application uploaded this image (clients.applications.id, application-level)';
COMMENT ON COLUMN "image"."images"."source_domain" IS 'Source service identifier: auth, product, post, cms, etc.';
COMMENT ON COLUMN "image"."images"."filename" IS 'Current filename after processing (may differ from original_filename)';
COMMENT ON COLUMN "image"."images"."original_filename" IS 'Original filename as uploaded by user';
COMMENT ON COLUMN "image"."images"."mime_type" IS 'MIME type: image/jpeg, image/png, image/webp, image/gif, etc.';
COMMENT ON COLUMN "image"."images"."size" IS 'File size in bytes';
COMMENT ON COLUMN "image"."images"."width" IS 'Image width in pixels';
COMMENT ON COLUMN "image"."images"."height" IS 'Image height in pixels';
COMMENT ON COLUMN "image"."images"."storage_path" IS 'Storage backend path (filesystem path or S3 object key)';
COMMENT ON COLUMN "image"."images"."url" IS 'Public accessible URL (CDN URL or direct storage URL)';
COMMENT ON COLUMN "image"."images"."is_public" IS 'TRUE = publicly accessible, FALSE = requires authentication';
COMMENT ON COLUMN "image"."images"."metadata" IS 'Extended JSON metadata: EXIF, color profiles, AI labels, etc.';
COMMENT ON COLUMN "image"."images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
