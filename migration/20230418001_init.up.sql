CREATE TABLE "public"."accounting_provider" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "is_deleted" bool DEFAULT false
)
;

COMMENT ON COLUMN "public"."accounting_provider"."name" IS 'provider name';
COMMENT ON COLUMN "public"."accounting_provider"."is_deleted" IS 'is deleted';

INSERT INTO accounting_provider(name, is_deleted) values('Xero', false), ('MYOB', false);