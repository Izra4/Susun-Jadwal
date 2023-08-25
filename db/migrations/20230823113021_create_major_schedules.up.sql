CREATE TABLE major_schedules (
    id INT NOT NULL AUTO_INCREMENT,
    day VARCHAR(10) NOT NULL,
    time VARCHAR(50) NOT NULL,
    room VARCHAR(10) NOT NULL,
    class_id INT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_class_id
        FOREIGN KEY (class_id) REFERENCES classes(id)
) ENGINE = InnoDB;
