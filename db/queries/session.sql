-- name: SaveSession :exec
INSERT INTO
    sessions (
        id,
        user_id,
        user_agent,
        refresh_token_hash,
        expired_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: GetAllSessions :many
SELECT s.id, s.user_id, s.refresh_token_hash, s.user_agent, s.ip_address, s.expired_at, s.revoked_at, s.created_at
FROM SESSIONS as s
WHERE
    s.user_id = $1
    AND s.deleted_at IS NULL
    AND s.revoked_at IS NULL;

-- name: GetSessionById :one
SELECT s.id, s.user_id, s.refresh_token_hash, s.user_agent, s.ip_address, s.expired_at, s.revoked_at, s.created_at
FROM SESSIONS as s
WHERE
    s.user_id = $1
    AND s.id = $2
    AND s.deleted_at IS NULL
    AND s.revoked_at IS NULL;

-- name: UpdateSession :exec
UPDATE SESSIONS as s SET revoked_at = now() WHERE s.id = $1;