CREATE TABLE classes(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    member INT NOT NULL,
    major_id INT NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;