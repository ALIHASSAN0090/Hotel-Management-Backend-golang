CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    number_of_persons INT NOT NULL,
    dine_in_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);