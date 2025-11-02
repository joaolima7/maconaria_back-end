-- name: CreatePost :execresult
INSERT INTO posts (
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetPostByID :one
SELECT
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id,
  created_at, updated_at
FROM posts
WHERE id = ?;

-- name: GetAllPosts :many
SELECT
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id,
  created_at, updated_at
FROM posts
ORDER BY created_at DESC;

-- name: GetFeaturedPosts :many
SELECT
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id,
  created_at, updated_at
FROM posts
WHERE is_featured = 1
ORDER BY created_at DESC;

-- name: GetPostsByType :many
SELECT
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id,
  created_at, updated_at
FROM posts
WHERE post_type = ?
ORDER BY created_at DESC;

-- name: UpdatePost :execresult
UPDATE posts
SET
  title = ?, category = ?, small_description = ?, complete_description = ?,
  date = ?, time = ?, location = ?, is_featured = ?, post_type = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- name: GetAllPostsByUserID :many
SELECT
  id, title, category, small_description, complete_description,
  date, time, location, is_featured, post_type, user_id,
  created_at, updated_at
FROM posts
WHERE user_id = ?
ORDER BY created_at DESC;