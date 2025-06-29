CREATE TABLE IF NOT EXISTS api_keys (
    id INT AUTO_INCREMENT,
    api_key VARCHAR(255) UNIQUE,
    organisation_id INT NOT NULL,
    expiry TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (organisation_id) REFERENCES organisations(id)
)