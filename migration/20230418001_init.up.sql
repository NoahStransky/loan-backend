CREATE TABLE "public"."accounting_provider" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "is_deleted" bool DEFAULT false
)
;

COMMENT ON COLUMN "public"."accounting_provider"."name" IS 'provider name';
COMMENT ON COLUMN "public"."accounting_provider"."is_deleted" IS 'is deleted';

INSERT INTO accounting_provider values('Xero', false), ('MYOB', false);