-- name: CreateUser :execresult
INSERT INTO users (id, name, email, password, is_active, is_admin)
VALUES($1, $2, $3, $4, $5, $6);

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :execresult
UPDATE users
SET name = $1, email = $2, is_active = $3, is_admin = $4
WHERE id = $5;

-- name: UpdateUserPassword :execresult
UPDATE users
SET password = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

