-- name: CreateLibrary :execresult
INSERT INTO libraries (id, title, small_description, degree, file_url, cover_url, link)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetLibraryByID :one
SELECT id, title, small_description, degree, file_url, cover_url, link, created_at, updated_at
FROM libraries
WHERE id = ?;

-- name: GetLibraryByTitle :one
SELECT id, title, small_description, degree, file_url, cover_url, link, created_at, updated_at
FROM libraries
WHERE title = ?;

-- name: GetAllLibraries :many
SELECT id, title, small_description, degree, file_url, cover_url, link, created_at, updated_at
FROM libraries
ORDER BY created_at DESC;

-- name: GetLibrariesByDegree :many
SELECT id, title, small_description, degree, file_url, cover_url, link, created_at, updated_at
FROM libraries
WHERE degree = ?
ORDER BY created_at DESC;

-- name: UpdateLibrary :execresult
UPDATE libraries
SET title = ?, small_description = ?, degree = ?, file_url = ?, cover_url = ?, link = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteLibrary :exec
DELETE FROM libraries WHERE id = ?;