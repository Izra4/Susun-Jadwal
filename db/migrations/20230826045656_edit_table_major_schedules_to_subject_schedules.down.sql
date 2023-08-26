ALTER TABLE subject_schedules
    DROP COLUMN createdAt,
    DROP COLUMN updatedAt,
    DROP COLUMN deletedAt;

ALTER TABLE subject_schedules
    RENAME TO major_schedules;
