-- name: ListAllMajorSchedules :many
SELECT * FROM subject_schedules
ORDER BY
        CASE
            WHEN day = 'Senin' THEN 1
        WHEN day = 'Selasa' THEN 2
        WHEN day = 'Rabu' THEN 3
        WHEN day = 'Kamis' THEN 4
        WHEN day = 'Jum''at' THEN 5
        ELSE 6
END;

-- name: GetSchedulesById :one
SELECT * FROM subject_schedules
WHERE id = ?;

-- name: CreateNewSubjectSchedules :execresult
insert into subject_schedules(day,time,room,class_id)
    VALUES(?,?,?,?);

-- name: UpdateSchedule :exec
UPDATE subject_schedules
    SET day = ?, time = ?, room = ?, class_id = ?
    WHERE id = ?;

-- name: DeleteSchedule :exec
DELETE FROM subject_schedules
    WHERE id = ?;