-- name: GetAllSubjects :many
SELECT * FROM subjects
ORDER BY name ASC;

-- name: GetSubjectById :one
SELECT * FROM subjects
    WHERE id = ?;

-- name: CreateNewSubject :execresult
INSERT INTO subjects (name,curriculum,sks,id_prodi)
    VALUES(?,?,?,?);

-- name: DeleteSubject :exec
DELETE FROM subjects
    WHERE id = ?;

-- name: UpdateSubject :exec
UPDATE subjects
    SET name = ?, curriculum = ?, sks = ?, id_prodi = ?
    WHERE id = ?;