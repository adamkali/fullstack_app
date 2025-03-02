-- name: FindTokenByToken :one
SELECT *
    FROM tokens 
    WHERE token = $1
    AND now() > expiration_datetime;

-- name: FindTokenByUserId :one
SELECT *
    FROM tokens 
    WHERE user_id = $1
    AND now() > expiration_datetime;

-- name: UpdateTokenByUserId :exec
UPDATE tokens 
    SET token = $1, expiration_datetime = $2
    WHERE user_id = $3;

-- name: CreateToken :one
INSERT INTO tokens (
    user_id, expiration_datetime, token
) VALUES ( $1, $2, $3 )
RETURNING *;

