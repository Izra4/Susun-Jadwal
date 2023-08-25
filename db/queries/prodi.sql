-- name: GetAllProdi :many
SELECT * FROM program_studies
ORDER BY id ASC;

-- name: GetProdiById :one
SELECT * FROM program_studies
WHERE id = ?;

-- name: CreateNewProdi :execresult
INSERT INTO program_studies(name)
    VALUES (?);

-- name: UpdateProdi :exec
UPDATE program_studies
    SET name = ?
    WHERE id = ?;

-- name: DeleteProdi :exec
DELETE FROM program_studies
    WHERE id = ?;