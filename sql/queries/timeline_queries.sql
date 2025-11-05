-- name: CreateTimeline :execresult
INSERT INTO timelines (id, period, pdf_data, is_highlight)
VALUES (?, ?, ?, ?);

-- name: GetTimelineByID :one
SELECT id, period, pdf_data, is_highlight, created_at, updated_at
FROM timelines
WHERE id = ?;

-- name: GetTimelineByPeriod :one
SELECT id, period, pdf_data, is_highlight, created_at, updated_at
FROM timelines
WHERE period = ?;

-- name: GetAllTimelines :many
SELECT id, period, pdf_data, is_highlight, created_at, updated_at
FROM timelines
ORDER BY period ASC;

-- name: UpdateTimeline :execresult
UPDATE timelines
SET period = ?, pdf_data = ?, is_highlight = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteTimeline :exec
DELETE FROM timelines WHERE id = ?;