CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    role VARCHAR(50) DEFAULT 'customer'
);

INSERT INTO users (username, email, password_hash, first_name, last_name, role)
VALUES ('admin', 'alihassankhan285@gmail.com', '$2a$12$BG4eccDYQWqii3gbW/KP9evZ/3XrUUko7OfdKb73ia.AWAVJp7.LS', 'Admin', 'User', 'admin');