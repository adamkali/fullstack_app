-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: FindUserByUsername :one
SELECT *
    FROM users 
    WHERE username = $1;

-- name: FindUserByEmail :one
SELECT *
    FROM users 
    WHERE email = $1;

-- name: FindBCryptHashByUsername :one
SELECT b_crypt_hash
    FROM users
    Where username = $1;

-- name: FindBCryptHashByEmail :one
SELECT b_crypt_hash
    FROM users
    Where email = $1;

-- name: FindUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (
  email, username, created_datetime, updated_datetime, profile_pic_url, admin, b_crypt_hash 
) VALUES ($1, $2, now(), now(), NULL, false, $3)
RETURNING *;

-- name: CreateUserAdmin :one
INSERT INTO users (
  email, username, created_datetime, updated_datetime, profile_pic_url, admin, b_crypt_hash
) VALUES ($1, $2, now(), now(), NULL, true, $3)
RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

