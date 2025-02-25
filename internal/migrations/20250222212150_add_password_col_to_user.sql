-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN b_crypt_hash VARCHAR NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN b_crypt_hash;
-- +goose StatementEnd
