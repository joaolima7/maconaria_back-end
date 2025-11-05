-- name: CreateWorker :execresult
INSERT INTO workers (
  id, number, name, registration, birth_date,
  initiation_date, elevation_date, exaltation_date, affiliation_date, installation_date,
  emeritus_mason_date, provect_mason_date, image_data, deceased
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetWorkerByID :one
SELECT
  id, number, name, registration, birth_date,
  initiation_date, elevation_date, exaltation_date, affiliation_date, installation_date,
  emeritus_mason_date, provect_mason_date, image_data, deceased,
  created_at, updated_at
FROM workers
WHERE id = ?;

-- name: GetWorkerByNumber :one
SELECT
  id, number, name, registration, birth_date,
  initiation_date, elevation_date, exaltation_date, affiliation_date, installation_date,
  emeritus_mason_date, provect_mason_date, image_data, deceased,
  created_at, updated_at
FROM workers
WHERE number = ?;

-- name: GetWorkerByRegistration :one
SELECT
  id, number, name, registration, birth_date,
  initiation_date, elevation_date, exaltation_date, affiliation_date, installation_date,
  emeritus_mason_date, provect_mason_date, image_data, deceased,
  created_at, updated_at
FROM workers
WHERE registration = ?;

-- name: GetAllWorkers :many
SELECT
  id, number, name, registration, birth_date,
  initiation_date, elevation_date, exaltation_date, affiliation_date, installation_date,
  emeritus_mason_date, provect_mason_date, image_data, deceased,
  created_at, updated_at
FROM workers
ORDER BY number ASC;

-- name: UpdateWorker :execresult
UPDATE workers
SET
  number = ?, name = ?, registration = ?, birth_date = ?,
  initiation_date = ?, elevation_date = ?, exaltation_date = ?, affiliation_date = ?,
  installation_date = ?, emeritus_mason_date = ?, provect_mason_date = ?,
  image_data = ?, deceased = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteWorker :exec
DELETE FROM workers WHERE id = ?;