-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR NOT NULL UNIQUE,
  username VARCHAR NOT NULL UNIQUE,
  created_datetime TIMESTAMP NOT NULL,
  updated_datetime TIMESTAMP NOT NULL,
  profile_pic_url VARCHAR(255),
  admin BOOLEAN NOT NULL DEFAULT false
);
CREATE TABLE tokens (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id),
  expiration_datetime TIMESTAMP NOT NULL,
  token VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
DROP TABLE users;
-- +goose StatementEnd
