-- name: CreateUser :execresult
INSERT INTO users (id, name, email, password, cim, degree, is_active, is_admin, is_regular)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllUsers :many
SELECT id, name, email, password, cim, degree, is_active, is_admin, is_regular, created_at, updated_at FROM users;

-- name: GetUserByID :one
SELECT id, name, email, password, cim, degree, is_active, is_admin, is_regular, created_at, updated_at FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, name, email, password, cim, degree, is_active, is_admin, is_regular, created_at, updated_at FROM users WHERE email = ?;

-- name: GetUserByCIM :one
SELECT id, name, email, password, cim, degree, is_active, is_admin, is_regular, created_at, updated_at FROM users WHERE cim = ?;

-- name: UpdateUser :execresult
UPDATE users
SET name = ?, email = ?, cim = ?, degree = ?, is_active = ?, is_admin = ?, is_regular = ?
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;