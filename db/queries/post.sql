-- -- name: GetPostByID :one
-- SELECT id, author_id, title, body, created_at, updated_at
-- FROM posts
-- WHERE id = $1;

-- -- name: ListPosts :many
-- SELECT id, author_id, title, body, created_at, updated_at
-- FROM posts
-- ORDER BY created_at DESC;

-- -- name: CreatePost :exec
-- INSERT INTO posts (id, author_id, title, body, created_at, updated_at)
-- VALUES ($1, $2, $3, $4, $5, $6);
