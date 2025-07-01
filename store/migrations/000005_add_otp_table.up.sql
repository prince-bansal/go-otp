CREATE TABLE IF NOT EXISTS otps (
    id INT AUTO_INCREMENT,
    otp varchar(255),
    mobile_number VARCHAR(255),
    organisation_id INT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (organisation_id) REFERENCES organisations(id)
)