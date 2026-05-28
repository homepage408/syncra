-- name: GetUserByEmail :one
SELECT id, email, password_hash
FROM users
WHERE email = $1
LIMIT 1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id;
