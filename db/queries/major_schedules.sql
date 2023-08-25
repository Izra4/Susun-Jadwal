-- name: ListAllMajorSchedules :many
SELECT * FROM major_schedules
ORDER BY day;

-- name: GetSchedulesByDay :many
SELECT * FROM major_schedules
WHERE day = ?;

-- name: GetSchedulesById :one
SELECT * FROM major_schedules
WHERE id = ?;