ALTER TABLE classes
    ADD createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER id,
    ADD updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER createdAt,
    ADD deletedAt TIMESTAMP DEFAULT NULL AFTER updatedAt;

ALTER TABLE classes
    CHANGE major_id subject_id INT NOT NULL;

ALTER TABLE classes
    ADD CONSTRAINT fk_subject_id FOREIGN KEY (subject_id) REFERENCES subjects(id);
