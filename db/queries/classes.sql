-- name: AddNewClass :execresult
INSERT INTO classes(name,member,major_id)
    VALUES (?,?,?);


-- name: ListClass :many
SELECT * FROM classes
ORDER BY name;

-- name: GetClassById :one
SELECT * FROM classes
WHERE id = ? LIMIT 1;

-- name: GetClassByName :many
SELECT * FROM classes
WHERE name = ?;

-- name: GetClassByMajorId :many
SELECT * FROM classes
WHERE major_id = ?;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = ?;

-- name: GetClassNameById :one
SELECT name FROM classes
WHERE id = ?;