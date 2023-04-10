DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

CREATE TABLE "public"."users" ("id" serial NOT NULL,"username" VARCHAR(64) NOT NULL,"first_name" text,"last_name" text,"email" text,"phone" text, PRIMARY KEY ("id"));
CREATE UNIQUE INDEX "users_username_idx" ON "public"."users" USING BTREE ("username");

ALTER SEQUENCE users_id_seq RESTART WITH 1;
UPDATE "public"."users" SET id=nextval('users_id_seq');