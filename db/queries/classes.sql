-- name: AddNewClass :execresult
INSERT INTO classes(name,member,subject_id)
    VALUES (?,?,?);

-- name: ListClass :many
SELECT * FROM classes
ORDER BY name;

-- name: GetClassById :one
SELECT * FROM classes
WHERE id = ?;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = ?;

-- name: UpdateClass :exec
UPDATE classes
    SET name = ?, member = ?, subject_id = ?
    WHERE id = ?;