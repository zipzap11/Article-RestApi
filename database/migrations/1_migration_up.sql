-- +migrate Up
CREATE TABLE "authors" (
  "id" uuid PRIMARY KEY,
  "name" varchar
);
-- +migrate StatementBegin
CREATE TABLE "articles" (
  "id" uuid PRIMARY KEY,
  "author_id" uuid,
  "title" varchar,
  "body" text,
  "created_at" timestamp
);
ALTER TABLE "articles" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");
-- +migrate StatementEnd
