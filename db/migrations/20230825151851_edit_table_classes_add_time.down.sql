# ALTER TABLE classes
#     DROP FOREIGN KEY fk_subject_id;
#
ALTER TABLE classes
    CHANGE subject_id major_id INT NOT NULL;

ALTER TABLE classes
    DROP COLUMN createdAt,
    DROP COLUMN updatedAt,
    DROP COLUMN deletedAt;
