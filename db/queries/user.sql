-- name: CheckEmailAndUsernameExisting :one
SELECT
    EXISTS (
        SELECT 1
        FROM users as u
        WHERE
            u.email = $1
    ) AS email_exists,
    EXISTS (
        SELECT 1
        FROM users as u
        WHERE
            u.username = $2
    ) AS username_exists;

-- name: GetUserByEmail :one

SELECT
    id,
    email,
    password_hash,
    is_verified,
    created_at,
    updated_at
FROM users
WHERE
    email = $1
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO
    users (
        id,
        email,
        username,
        full_name,
        password_hash,
        is_verified,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    );