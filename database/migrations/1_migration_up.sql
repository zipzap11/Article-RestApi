-- +migrate Up
CREATE TABLE "authors" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);
-- +migrate StatementBegin
CREATE TABLE "articles" (
  "id" uuid PRIMARY KEY NOT NULL,
  "author_id" uuid NOT NULL,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamp NOT NULl
);
ALTER TABLE "articles" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");
CREATE INDEX idx_articles_id ON articles (id);
-- +migrate StatementEnd
