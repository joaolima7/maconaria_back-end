-- name: CreatePost :execresult
INSERT INTO posts(id, title, description, date, image, user_id, post_type)
VALUES($1, $2, $3, $4, $5, $6, $7);

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetAllPosts :many
SELECT * FROM posts ORDER BY date DESC;

-- name: UpdatePost :execresult
UPDATE posts
SET title = $1, description = $2, date = $3, image = $4, post_type = $5
WHERE id = $6;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: GetAllPostsByUserID :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY date DESC;

