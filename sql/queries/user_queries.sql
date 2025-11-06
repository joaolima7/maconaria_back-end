-- name: CreateUser :execresult
INSERT INTO users (id, name, email, password, cim, degree, is_active, is_admin)
VALUES(?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: GetUserByCIM :one
SELECT * FROM users WHERE cim = ?;

-- name: UpdateUser :execresult
UPDATE users
SET name = ?, email = ?, cim = ?, degree = ?, is_active = ?, is_admin = ?
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;