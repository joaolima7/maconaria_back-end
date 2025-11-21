-- name: CreatePostImage :execresult
INSERT INTO post_images (id, post_id, image_url)
VALUES (?, ?, ?);

-- name: GetPostImages :many
SELECT id, post_id, image_url, created_at
FROM post_images
WHERE post_id = ?
ORDER BY created_at;

-- name: GetPostImageByID :one
SELECT id, post_id, image_url, created_at
FROM post_images
WHERE id = ?;

-- name: DeletePostImagesByPostID :exec
DELETE FROM post_images WHERE post_id = ?;

-- name: DeletePostImageByID :exec
DELETE FROM post_images WHERE id = ?;