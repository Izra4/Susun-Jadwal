-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY id;

-- name: GetUsersByID :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users(email,name,nim,id_prodi)
    VALUES (?,?,?,?);

-- name: UpdateUser :exec
UPDATE users
    SET email = ?, name = ?, nim = ?, id_prodi = ?
    WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
    WHERE id = ?;