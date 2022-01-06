-- +migrate Down
DROP TABLE "authors"
-- +migrate StatementBegin
DROP TABLE "articles"
-- +migrate StatementEnd