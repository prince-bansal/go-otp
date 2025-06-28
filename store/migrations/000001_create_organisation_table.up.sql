CREATE TABLE IF NOT EXISTS organisations (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    description text,
    created_at timestamp,
    updated_at timestamp
)

