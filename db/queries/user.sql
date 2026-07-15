-- name: CheckEmailAndUsernameExisting :one
SELECT
    EXISTS (
        SELECT 1
        FROM users as u
        WHERE
            u.email = $1 and deleted_at IS NULL
    ) AS email_exists,
    EXISTS (
        SELECT 1
        FROM users as u
        WHERE
            u.username = $2 and deleted_at IS NULL
    ) AS username_exists;

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    username,
    full_name,
    password_hash,
    is_verified,
    created_at,
    updated_at
FROM users
WHERE
    email = $1 and deleted_at IS NULL
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO
    users (
        id,
        email,
        username,
        full_name,
        password_hash
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    );