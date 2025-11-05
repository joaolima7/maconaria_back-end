-- name: CreateAcacia :execresult
INSERT INTO acacias (id, name, terms, is_president, deceased, image_data)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAcaciaByID :one
SELECT id, name, terms, is_president, deceased, image_data, created_at, updated_at
FROM acacias
WHERE id = ?;

-- name: GetAcaciaByName :one
SELECT id, name, terms, is_president, deceased, image_data, created_at, updated_at
FROM acacias
WHERE name = ?;

-- name: GetAllAcacias :many
SELECT id, name, terms, is_president, deceased, image_data, created_at, updated_at
FROM acacias
ORDER BY name ASC;

-- name: UpdateAcacia :execresult
UPDATE acacias
SET name = ?, terms = ?, is_president = ?, deceased = ?, image_data = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteAcacia :exec
DELETE FROM acacias WHERE id = ?;