-- name: CreateWordKey :execresult
INSERT INTO wordkeys (id, wordkey, active)
VALUES (?, ?, ?);

-- name: GetWordKeyByID :one
SELECT id, wordkey, active, created_at
FROM wordkeys
WHERE id = ?;

-- name: GetWordKeyByActive :one
SELECT id, wordkey, active, created_at
FROM wordkeys
WHERE active = true
LIMIT 1;

-- name: GetAllWordKeys :many
SELECT id, wordkey, active, created_at
FROM wordkeys
ORDER BY created_at DESC;

-- name: UpdateWordKey :execresult
UPDATE wordkeys
SET wordkey = ?, active = ?
WHERE id = ?;

-- name: DeactivateAllWordKeys :exec
UPDATE wordkeys
SET active = false
WHERE active = true;

-- name: DeleteWordKey :exec
DELETE FROM wordkeys WHERE id = ?;
