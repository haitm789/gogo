-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.


-- https://viblo.asia/p/managing-database-migration-in-go-aWj5336b56m

-- goose -dir=migrations  mysql "pandog:pandog@tcp(gogo-mysql)/gogo?charset=utf8" create create_users sql 