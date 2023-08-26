-- name: GetAllKrs :many
SELECT * FROM krs
ORDER BY id ASC;

-- name: GetKrsByID :one
SELECT * FROM krs
WHERE id = ?;

-- name: GetKrsByIDUser :many
SELECT * FROM krs
WHERE userID = ?;

-- name: AddKrs :execresult
INSERT INTO krs(totals, userID)
    VALUES (?,?);

-- name: UpdateKrs :exec
UPDATE krs
    SET totals = ?, userID = ?
    WHERE id = ?;

-- name: DeleteKrs :exec
DELETE FROM krs
    WHERE id = ?;